package client

import (
	application_pb "common/proto/application/generated"
	auth_pb "common/proto/auth/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func getConnection(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
}

func NewAuthClient(address string) auth_pb.AuthServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Authorization service: %v", err)
	}
	return auth_pb.NewAuthServiceClient(conn)
}
func NewApplicationClient(address string) application_pb.ApplicationServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Authorization service: %v", err)
	}
	return application_pb.NewApplicationServiceClient(conn)
}
