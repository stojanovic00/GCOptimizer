package app

import (
	"auth_service/api/handler"
	"auth_service/config"
	"auth_service/core/domain"
	"auth_service/core/service"
	db_client "auth_service/persistence/client"
	"auth_service/persistence/repo"
	auth_pb "common/proto/auth/generated"
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
	err := pgClient.AutoMigrate(&domain.Account{}, &domain.Role{}, &domain.Permission{})
	if err != nil {
		return err
	}
	log.Printf("Connected and updated pg database")

	accountRepo := repo.NewAccountRepositoryPg(pgClient)
	mailService := service.NewMailService(a.Config.MailClientMail, a.Config.MailClientPassword)
	accountService := service.NewAccountService(accountRepo, mailService)
	accountHandler := handler.NewAccountHandler(accountService)

	a.startGrpcServer(accountHandler)
	return nil
}

func (a *App) initPGClient() *gorm.DB {
	client, err := db_client.GetPostgresClient(
		a.Config.AuthDbHost, a.Config.AuthDbUser,
		a.Config.AuthDbPass, a.Config.AuthDbName,
		a.Config.AuthDbPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
func (a *App) startGrpcServer(accountHandler *handler.AccountHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", a.Config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	auth_pb.RegisterAuthServiceServer(grpcServer, accountHandler)
	fmt.Printf("Listening on port: %s\n", a.Config.Port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
