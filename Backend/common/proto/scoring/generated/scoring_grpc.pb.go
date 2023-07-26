// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: scoring.proto

package scoring_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ScoringServiceClient is the client API for ScoringService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ScoringServiceClient interface {
	StartCompetition(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*EmptyMessage, error)
}

type scoringServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewScoringServiceClient(cc grpc.ClientConnInterface) ScoringServiceClient {
	return &scoringServiceClient{cc}
}

func (c *scoringServiceClient) StartCompetition(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, "/scoring_pb.ScoringService/StartCompetition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScoringServiceServer is the server API for ScoringService service.
// All implementations must embed UnimplementedScoringServiceServer
// for forward compatibility
type ScoringServiceServer interface {
	StartCompetition(context.Context, *IdMessage) (*EmptyMessage, error)
	mustEmbedUnimplementedScoringServiceServer()
}

// UnimplementedScoringServiceServer must be embedded to have forward compatible implementations.
type UnimplementedScoringServiceServer struct {
}

func (UnimplementedScoringServiceServer) StartCompetition(context.Context, *IdMessage) (*EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartCompetition not implemented")
}
func (UnimplementedScoringServiceServer) mustEmbedUnimplementedScoringServiceServer() {}

// UnsafeScoringServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ScoringServiceServer will
// result in compilation errors.
type UnsafeScoringServiceServer interface {
	mustEmbedUnimplementedScoringServiceServer()
}

func RegisterScoringServiceServer(s grpc.ServiceRegistrar, srv ScoringServiceServer) {
	s.RegisterService(&ScoringService_ServiceDesc, srv)
}

func _ScoringService_StartCompetition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoringServiceServer).StartCompetition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scoring_pb.ScoringService/StartCompetition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoringServiceServer).StartCompetition(ctx, req.(*IdMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// ScoringService_ServiceDesc is the grpc.ServiceDesc for ScoringService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ScoringService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scoring_pb.ScoringService",
	HandlerType: (*ScoringServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartCompetition",
			Handler:    _ScoringService_StartCompetition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scoring.proto",
}