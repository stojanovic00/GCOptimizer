package app

import (
	grpc_client "api_gateway/api/client"
	"api_gateway/api/handler"
	"auth_service/api/middleware"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (a *App) CreateRoutersAndSetRoutes() (*gin.Engine, *gin.Engine, error) {
	//DEPENDENCIES
	authServiceAddress := fmt.Sprintf("%s:%s", a.Config.AuthServiceHost, a.Config.AuthServicePort)
	authClient := grpc_client.NewAuthClient(authServiceAddress)

	authHandler := handler.NewAuthHandler(authClient)
	scheduleHandler := handler.ScheduleHandler{}

	// MIDDLEWARE
	corsMiddleware := cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	})

	// PUBLIC
	publicRouter := gin.Default()
	publicRouter.Use(corsMiddleware)

	// ROUTES
	publicRouter.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Endpoint doesn't exist"})
	})

	authGroup := publicRouter.Group("/auth")
	authGroup.POST("/register", authHandler.Register)
	authGroup.POST("/login", authHandler.Login)

	//PRIVATE
	privateRouter := gin.Default()

	// ROUTES
	privateRouter.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Endpoint doesn't exist"})
	})

	testGroup := privateRouter.Group("/test", middleware.ValidateAndExtractToken())
	testGroup.GET("", middleware.Authorize("JudgesRead"), scheduleHandler.Test)
	return publicRouter, privateRouter, nil
}
