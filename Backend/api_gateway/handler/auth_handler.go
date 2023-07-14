package handler

import (
	auth_pb "common/proto/auth/generated"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type AuthHandler struct {
	client auth_pb.AuthServiceClient
}

func NewAuthHandler(client auth_pb.AuthServiceClient) *AuthHandler {
	return &AuthHandler{client: client}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var acc auth_pb.Account
	err := ctx.ShouldBindJSON(&acc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.client.Create(context.TODO(), &acc)
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.AlreadyExists:
				ctx.JSON(http.StatusConflict, grpcError.Message())
				return

			}
		}

		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var acc auth_pb.Account
	err := ctx.ShouldBindJSON(&acc)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	response, err := h.client.Login(context.TODO(), &acc)
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, grpcError.Message())
				return
			case codes.InvalidArgument:
				ctx.JSON(http.StatusBadRequest, grpcError.Message())
				return
			case codes.Unknown:
				ctx.JSON(http.StatusBadRequest, grpcError.Message())
				return
			}
		}

		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}
