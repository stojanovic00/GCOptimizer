package client

import (
	application_pb "common/proto/application/generated"
	auth_pb "common/proto/auth/generated"
	scheduling_pb "common/proto/scheduling/generated"
	scoring_pb "common/proto/scoring/generated"
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
		log.Fatalf("Failed to start gRPC connection to Application service: %v", err)
	}
	return application_pb.NewApplicationServiceClient(conn)
}

func NewSchedulingClient(address string) scheduling_pb.SchedulingServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to Scheduling service: %v", err)
	}
	return scheduling_pb.NewSchedulingServiceClient(conn)
}
func NewScoringClient(address string) scoring_pb.ScoringServiceClient {
	conn, err := getConnection(address)
	if err != nil {
		log.Fatalf("Failed to start gRPC connection to scoring service: %v", err)
	}
	return scoring_pb.NewScoringServiceClient(conn)
}
