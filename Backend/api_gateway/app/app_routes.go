package app

import (
	grpc_client "api_gateway/client"
	handler "api_gateway/handler"
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

	applicationServiceAddress := fmt.Sprintf("%s:%s", a.Config.ApplicationServiceHost, a.Config.ApplicationServicePort)
	applicationClient := grpc_client.NewApplicationClient(applicationServiceAddress)
	applicationHandler := handler.NewApplicationHandler(applicationClient, authClient)

	schedulingServiceAddress := fmt.Sprintf("%s:%s", a.Config.SchedulingServiceHost, a.Config.SchedulingServicePort)
	schedulingClient := grpc_client.NewSchedulingClient(schedulingServiceAddress)
	schedulingHandler := handler.NewSchedulingHandler(schedulingClient)

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

	//Unauth
	soGroup := applicationGroupPublic.Group("/sports-organisation")
	soGroup.POST("", applicationHandler.RegisterSportsOrganisation)

	//Auth
	soGroup.Use(middleware.ValidateAndExtractToken())
	soGroup.GET("", middleware.Authorize("SportsOrg_rud"), applicationHandler.GetLoggedSportsOrganisation)
	soGroup.GET("judge", middleware.Authorize("Judge_crud"), applicationHandler.GetSportOrganisationJudges)
	soGroup.POST("judge", middleware.Authorize("Judge_crud"), applicationHandler.RegisterJudge)
	soGroup.GET("contestant", middleware.Authorize("Contestant_crud"), applicationHandler.GetSportOrganisationContestants)
	soGroup.POST("contestant", middleware.Authorize("Contestant_crud"), applicationHandler.RegisterContestant)

	compGroup := applicationGroupPublic.Group("/competition")
	//Unauth
	compGroup.GET("", applicationHandler.GetAllCompetitions)
	compGroup.GET("/:id", applicationHandler.GetCompetitionById)
	//Auth
	compGroup.Use(middleware.ValidateAndExtractToken())
	compGroup.POST("", middleware.Authorize("Competition_cud"), applicationHandler.CreateCompetition)
	compGroup.POST("/:id/age-category", middleware.Authorize("Competition_cud"), applicationHandler.AddAgeCategory)
	compGroup.POST("/:id/delegation-member-prop", middleware.Authorize("Competition_cud"), applicationHandler.AddDelegationMemberProposition)

	compGroup.POST("/:id/app/judge", middleware.Authorize("Application_crud"), applicationHandler.CreateJudgeApplication)
	compGroup.GET("/:id/app/judge", middleware.Authorize("Application_crud"), applicationHandler.GetAllJudgeApplications)
	compGroup.POST("/:id/app/contestant", middleware.Authorize("Application_crud"), applicationHandler.CreateContestantApplication)
	compGroup.GET("/:id/app/contestant", middleware.Authorize("Application_crud"), applicationHandler.GetAllContestantApplications)

	//PRIVATE
	privateRouter := gin.Default()

	// ROUTES
	privateRouter.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Endpoint doesn't exist"})
	})

	schGroup := privateRouter.Group("scheduling")
	schGroup.Use(middleware.ValidateAndExtractToken())
	schGroup.POST("test", middleware.Authorize("Application_crud"), schedulingHandler.Test)

	return publicRouter, privateRouter, nil
}
