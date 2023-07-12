package client

import (
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
