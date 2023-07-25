package app

import (
	"gorm.io/gorm"
	"log"
	"scoring_service/config"
	"scoring_service/core/domain"
	db_client "scoring_service/persistence/client"
)

type App struct {
	Config config.Config
}

func (a *App) Run() error {
	//DB INIT
	log.Printf("Connecting to database...")
	pgClient := a.initPGClient()
	err := pgClient.AutoMigrate(
		&domain.Address{},
		&domain.AgeCategory{},
		&domain.Competition{},
		&domain.Contestant{},
		&domain.Judge{},
		&domain.Panel{},
		&domain.Schedule{},
		&domain.ScheduleSlot{},
		&domain.ScoreCalculationMethod{},
		&domain.Score{},
		&domain.Session{},
		&domain.SportsOrganization{},
		&domain.TeamComposition{},
		&domain.TempScore{},
	)

	if err != nil {
		return err
	}
	log.Printf("Connected and updated pg database")

	//soRepo := repo.NewSportsOrganisationRepoPg(pgClient)
	//soService := service.NewSportsOrganisationService(soRepo)
	//rpcHandler := handler.NewHandlerRpc(soService, dmService, compService, appService)

	//a.startGrpcServer(rpcHandler)
	return nil
}

func (a *App) initPGClient() *gorm.DB {
	client, err := db_client.GetPostgresClient(
		a.Config.ScoringDbHost, a.Config.ScoringDbUser,
		a.Config.ScoringDbPass, a.Config.ScoringDbName,
		a.Config.ScoringDbPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

//func (a *App) startGrpcServer(handlerRpc *handler.HandlerRpc) {
//	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", a.Config.Port))
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//
//	grpcServer := grpc.NewServer()
//	application_pb.RegisterApplicationServiceServer(grpcServer, handlerRpc)
//	fmt.Printf("Listening on port: %s\n", a.Config.Port)
//
//	if err := grpcServer.Serve(listener); err != nil {
//		log.Fatalf("failed to serve: %s", err)
//	}
//}
