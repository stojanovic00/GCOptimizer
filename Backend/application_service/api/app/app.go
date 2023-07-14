package app

import (
	"application_service/config"
	"application_service/core/domain"
	db_client "application_service/persistence/client"
	"gorm.io/gorm"
	"log"
)

type App struct {
	Config config.Config
}

func (a *App) Run() error {
	//DB INIT
	log.Printf("Connecting to database...")
	pgClient := a.initPGClient()
	// POMAZE BOG
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
		&domain.SportsOrganisation{},
		&domain.TeamComposition{},
	)

	// POMAZE BOG
	if err != nil {
		return err
	}
	log.Printf("Connected and updated pg database")

	//a.startGrpcServer(accountHandler)
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

//func (a *App) startGrpcServer(accountHandler *handler.AccountHandler) {
//	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", a.Config.Port))
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//
//	grpcServer := grpc.NewServer()
//	auth_pb.RegisterAuthServiceServer(grpcServer, accountHandler)
//	fmt.Printf("Listening on port: %s\n", a.Config.Port)
//
//	if err := grpcServer.Serve(listener); err != nil {
//		log.Fatalf("failed to serve: %s", err)
//	}
//}
