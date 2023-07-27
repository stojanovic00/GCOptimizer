package app

import (
	scoring_pb "common/proto/scoring/generated"
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
	"scoring_service/api/client"
	"scoring_service/api/handler"
	"scoring_service/config"
	"scoring_service/core/domain"
	"scoring_service/core/service"
	db_client "scoring_service/persistence/client"
	"scoring_service/persistence/repo"
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

	schedulingServiceAddress := fmt.Sprintf("%s:%s", a.Config.SchedulingServiceHost, a.Config.SchedulingServicePort)
	schClient := client.NewSchedulingClient(schedulingServiceAddress)

	applicationServiceAddress := fmt.Sprintf("%s:%s", a.Config.ApplicationServiceHost, a.Config.ApplicationServicePort)
	appClient := client.NewApplicationClient(applicationServiceAddress)

	authServiceAddress := fmt.Sprintf("%s:%s", a.Config.AuthServiceHost, a.Config.AuthServicePort)
	authClient := client.NewAuthClient(authServiceAddress)

	schRepo := repo.NewScheduleRepoPg(pgClient)
	schService := service.NewScheduleService(appClient, schClient, schRepo)

	jpRepo := repo.NewJudgePanelRepoPg(pgClient)
	jpService := service.NewJudgePanelService(jpRepo, authClient)

	rpcHandler := handler.NewHandlerRpc(schService, jpService)

	a.startGrpcServer(rpcHandler)
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

func (a *App) startGrpcServer(handlerRpc *handler.HandlerRpc) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", a.Config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	scoring_pb.RegisterScoringServiceServer(grpcServer, handlerRpc)
	fmt.Printf("Listening on port: %s\n", a.Config.Port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
