package middleware

import (
	grpc_client "auth_service/api/client"
	"auth_service/core/service"
	auth_pb "common/proto/auth/generated"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"net/http"
	"os"
	"strings"
)

//NOTE!!!
//This middleware should be used in service which uses auth service e.g. api gateway
//NOTE!!!

var (
	authServiceHost = os.Getenv("AUTH_SERVICE_HOST")
	authServicePort = os.Getenv("AUTH_SERVICE_PORT")
)

func ValidateAndExtractToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tokenString string
		bearerToken := ctx.Request.Header.Get("Authorization")
		if bearerToken == "" {
			// Maybe it is web socket opening request(it can't have Authorize header) so check its dedicated auth header
			bearerToken = ctx.Request.Header.Get("Sec-Websocket-Protocol")
			if bearerToken == "" {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": "No authentication header provided"})
				ctx.Abort()
				return
			}
			//It doesn't have "Bearer " part so no need for parsing
			tokenString = bearerToken
		} else {
			//Removes "Bearer "
			tokenString = strings.Split(bearerToken, " ")[1]
		}

		claims, err := service.VerifyToken(tokenString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": err.Error()})
			ctx.Abort()
			return
		}

		if len(ctx.Keys) == 0 {
			ctx.Keys = make(map[string]interface{})
		}

		//Packing relevant claims to "inter-service token"
		ctx.Keys["user-info"] = &auth_pb.UserInfo{
			Email: claims.Email,
			Role:  claims.Role,
		}

		ctx.Next()
	}
}

func Authorize(permission string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authServiceAddress := fmt.Sprintf("%s:%s", authServiceHost, authServicePort)
		authClient := grpc_client.NewAuthClient(authServiceAddress)
		ctxWithInfo, err := GetGrpcContextWithUserInfo(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			ctx.Abort()
			return
		}

		hasPermission, err := authClient.HasPermission(ctxWithInfo, &auth_pb.HasPermissionRequest{Permission: permission})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"errors": err.Error()})
			ctx.Abort()
			return
		}

		if !hasPermission.Value {
			ctx.JSON(http.StatusUnauthorized, gin.H{"errors": "You don't have permission for this action"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func ParseUserInfo(ctx *gin.Context) (*auth_pb.UserInfo, error) {
	userInfoInterf := ctx.Keys["user-info"]
	if userInfoInterf == nil {
		return nil, errors.New("User info not present in  context")
	}
	return userInfoInterf.(*auth_pb.UserInfo), nil
}

func GetGrpcContextWithUserInfo(ctx *gin.Context) (context.Context, error) {
	userInfo, err := ParseUserInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return nil, err
	}
	return metadata.NewOutgoingContext(context.Background(), metadata.Pairs("user-info", userInfo.String())), nil
}
