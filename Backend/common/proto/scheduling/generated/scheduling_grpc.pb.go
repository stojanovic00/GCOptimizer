// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: scheduling.proto

package scheduling_pb

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

// SchedulingServiceClient is the client API for SchedulingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SchedulingServiceClient interface {
	Test(ctx context.Context, in *TestMessage, opts ...grpc.CallOption) (*TestResponse, error)
}

type schedulingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSchedulingServiceClient(cc grpc.ClientConnInterface) SchedulingServiceClient {
	return &schedulingServiceClient{cc}
}

func (c *schedulingServiceClient) Test(ctx context.Context, in *TestMessage, opts ...grpc.CallOption) (*TestResponse, error) {
	out := new(TestResponse)
	err := c.cc.Invoke(ctx, "/scheduling_pb.SchedulingService/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchedulingServiceServer is the server API for SchedulingService service.
// All implementations must embed UnimplementedSchedulingServiceServer
// for forward compatibility
type SchedulingServiceServer interface {
	Test(context.Context, *TestMessage) (*TestResponse, error)
	mustEmbedUnimplementedSchedulingServiceServer()
}

// UnimplementedSchedulingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSchedulingServiceServer struct {
}

func (UnimplementedSchedulingServiceServer) Test(context.Context, *TestMessage) (*TestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (UnimplementedSchedulingServiceServer) mustEmbedUnimplementedSchedulingServiceServer() {}

// UnsafeSchedulingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SchedulingServiceServer will
// result in compilation errors.
type UnsafeSchedulingServiceServer interface {
	mustEmbedUnimplementedSchedulingServiceServer()
}

func RegisterSchedulingServiceServer(s grpc.ServiceRegistrar, srv SchedulingServiceServer) {
	s.RegisterService(&SchedulingService_ServiceDesc, srv)
}

func _SchedulingService_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulingServiceServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduling_pb.SchedulingService/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulingServiceServer).Test(ctx, req.(*TestMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// SchedulingService_ServiceDesc is the grpc.ServiceDesc for SchedulingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SchedulingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scheduling_pb.SchedulingService",
	HandlerType: (*SchedulingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _SchedulingService_Test_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scheduling.proto",
}
