package app

import (
	grpc_client "api_gateway/client"
	handler2 "api_gateway/handler"
	"auth_service/api/middleware"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (a *App) CreateRoutersAndSetRoutes() (*gin.Engine, *gin.Engine, error) {
	//DEPENDENCIES
	authServiceAddress := fmt.Sprintf("%s:%s", a.Config.AuthServiceHost, a.Config.AuthServicePort)
	authClient := grpc_client.NewAuthClient(authServiceAddress)
	authHandler := handler2.NewAuthHandler(authClient)

	applicationServiceAddress := fmt.Sprintf("%s:%s", a.Config.ApplicationServiceHost, a.Config.ApplicationServicePort)
	applicationClient := grpc_client.NewApplicationClient(applicationServiceAddress)
	applicationHandler := handler2.NewApplicationHandler(applicationClient, authClient)

	scheduleHandler := handler2.ScheduleHandler{}

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

	//AUTH
	authGroup := publicRouter.Group("/auth")
	authGroup.POST("/register", authHandler.Register)
	authGroup.POST("/login", authHandler.Login)

	//APPLICATION
	applicationGroupPublic := publicRouter.Group("application")
	soGroup := applicationGroupPublic.Group("/sports-organisation")
	soGroup.POST("", applicationHandler.RegisterSportsOrganisation)

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