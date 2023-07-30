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

	scoringServiceAddress := fmt.Sprintf("%s:%s", a.Config.ScoringServiceHost, a.Config.ScoringServicePort)
	scoringClient := grpc_client.NewScoringClient(scoringServiceAddress)
	scoringHandler := handler.NewScoringHandler(scoringClient)

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
	privateRouter.Use(corsMiddleware)

	// ROUTES
	privateRouter.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Endpoint doesn't exist"})
	})

	//Schedule
	schGroup := privateRouter.Group("scheduling")
	schGroup.Use(middleware.ValidateAndExtractToken())
	schGroup.POST("schedule", middleware.Authorize("Schedule_crud"), schedulingHandler.GenerateSchedule)
	schGroup.GET("schedule/:id", middleware.Authorize("Schedule_crud"), schedulingHandler.GetByCompetitionId)

	//Live scoring
	scoGroup := privateRouter.Group("scoring")
	scoGroup.Use(middleware.ValidateAndExtractToken())
	scoGroup.POST("competition/:id", middleware.Authorize("LiveSchedule_cru"), scoringHandler.StartCompetition)

	scoGroup.GET("judge", middleware.Authorize("LiveJudge_r"), scoringHandler.GetLoggedJudgeInfo)
	scoGroup.GET("competition/:id/contestant", middleware.Authorize("LiveContestant_r"), scoringHandler.GetCurrentApparatusContestants)
	scoGroup.GET("competition/:id/contestant/current", middleware.Authorize("LiveContestant_r"), scoringHandler.GetNextCurrentApparatusContestant)

	scoGroup.POST("competition/:id/temp-score", middleware.Authorize("Score_c"), scoringHandler.SubmitTempScore)
	scoGroup.GET("competition/:id/temp-score", middleware.Authorize("Score_r"), scoringHandler.GetContestantsTempScores)
	scoGroup.GET("competition/:id/score/can-calculate", middleware.Authorize("Score_r"), scoringHandler.CanCalculateScore)
	scoGroup.GET("competition/:id/score/calculate", middleware.Authorize("Score_r"), scoringHandler.CalculateScore)
	scoGroup.POST("competition/:id/score", middleware.Authorize("Score_c"), scoringHandler.SubmitScore)
	scoGroup.GET("competition/:id/score", middleware.Authorize("Score_c"), scoringHandler.GetScore)

	scoGroup.POST("competition/:id/rotation/finish", middleware.Authorize("LiveSchedule_cru"), scoringHandler.FinishRotation)
	scoGroup.POST("competition/:id/session/finish", middleware.Authorize("LiveSchedule_cru"), scoringHandler.FinishSession)
	scoGroup.POST("competition/:id/finish", middleware.Authorize("Scoreboard_c"), scoringHandler.FinishCompetition)

	scoGroup.GET("competition/:id/rotation/finish-check", middleware.Authorize("LiveSchedule_cru"), scoringHandler.IsRotationFinished)
	scoGroup.GET("competition/:id/session/finish-check", middleware.Authorize("LiveSchedule_cru"), scoringHandler.IsSessionFinished)
	scoGroup.GET("competition/:id/finish-check", middleware.Authorize("LiveSchedule_cru"), scoringHandler.IsCompetitionFinished)

	//Judging panel
	jpGroup := scoGroup.Group("judging-panel")
	jpGroup.GET("competition/:id/unassigned", middleware.Authorize("JudgingPanel_crud"), scoringHandler.GetApparatusesWithoutPanel)
	jpGroup.POST("", middleware.Authorize("JudgingPanel_crud"), scoringHandler.CreateJudgingPanelsForApparatus)
	jpGroup.POST("/:id/judge", middleware.Authorize("JudgingPanel_crud"), scoringHandler.AssignJudge)
	jpGroup.GET("judge/competition/:id", middleware.Authorize("JudgingPanel_crud"), scoringHandler.GetAssignedJudges)
	jpGroup.POST("/:id/score-calc-method", middleware.Authorize("JudgingPanel_crud"), scoringHandler.AssignScoreCalculationMethod)

	return publicRouter, privateRouter, nil
}
