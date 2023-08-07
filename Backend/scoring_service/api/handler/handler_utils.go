package handler

import (
	auth_pb "common/proto/auth/generated"
	"context"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func ParseUserInfo(ctx context.Context) (*auth_pb.UserInfo, error) {
	userInfo := &auth_pb.UserInfo{}
	metadata, ok := metadata.FromIncomingContext(ctx)
	if ok {
		userInfoStr := metadata.Get("user-info")
		if len(userInfoStr) > 0 {
			err := proto.UnmarshalText(userInfoStr[0], userInfo)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "Failed unmarshalling context")
			}
		}
	}
	return userInfo, nil
}
