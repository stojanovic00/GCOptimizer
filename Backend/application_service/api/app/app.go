package app

import (
	"application_service/api/handler"
	"application_service/config"
	"application_service/core/domain"
	"application_service/core/service"
	db_client "application_service/persistence/client"
	"application_service/persistence/repo"
	application_pb "common/proto/application/generated"
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"log"
	"net"
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
		&domain.ApparatusAnnouncement{},
		&domain.Competition{},
		&domain.ContestantApplication{},
		&domain.Contestant{},
		&domain.DelegationMemberApplication{},
		&domain.DelegationMember{},
		&domain.DelegationMemberPosition{},
		&domain.DelegationMemberProposition{},
		&domain.Judge{},
		&domain.JudgeApplication{},
		&domain.SportsOrganisation{},
		&domain.TeamComposition{},
	)

	if err != nil {
		return err
	}
	log.Printf("Connected and updated pg database")

	soRepo := repo.NewSportsOrganisationRepoPg(pgClient)
	soService := service.NewSportsOrganisationService(soRepo)
	rpcHandler := handler.NewHandlerRpc(soService)

	a.startGrpcServer(rpcHandler)
	return nil
}

func (a *App) initPGClient() *gorm.DB {
	client, err := db_client.GetPostgresClient(
		a.Config.ApplicationDbHost, a.Config.ApplicationDbUser,
		a.Config.ApplicationDbPass, a.Config.ApplicationDbName,
		a.Config.ApplicationDbPort)
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
	application_pb.RegisterApplicationServiceServer(grpcServer, handlerRpc)
	fmt.Printf("Listening on port: %s\n", a.Config.Port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
