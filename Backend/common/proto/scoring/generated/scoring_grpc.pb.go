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
	// Judging panel
	GetApparatusesWithoutPanel(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*ApparatusList, error)
	CreateJudgingPanelsForApparatus(ctx context.Context, in *CreateJudgingPanelsForApparatusRequest, opts ...grpc.CallOption) (*CreateJudgingPanelsForApparatusResponse, error)
	AssignJudge(ctx context.Context, in *AssignJudgeRequest, opts ...grpc.CallOption) (*EmptyMessage, error)
	GetAssignedJudges(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*JudgeList, error)
	AssignScoreCalculation(ctx context.Context, in *AssignScoreCalculationRequest, opts ...grpc.CallOption) (*EmptyMessage, error)
	// Live scoring
	// Judge and contestant info
	GetLoggedJudgeInfo(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*JudgeJudgingInfo, error)
	GetCurrentApparatusContestants(ctx context.Context, in *GetByApparatusRequest, opts ...grpc.CallOption) (*ContestantList, error)
	GetNextCurrentApparatusContestant(ctx context.Context, in *GetByApparatusRequest, opts ...grpc.CallOption) (*Contestant, error)
	// Scores
	SubmitTempScore(ctx context.Context, in *TempScore, opts ...grpc.CallOption) (*EmptyMessage, error)
	GetContestantsTempScores(ctx context.Context, in *ScoreRequest, opts ...grpc.CallOption) (*TempScoreList, error)
	CalculateScore(ctx context.Context, in *ScoreRequest, opts ...grpc.CallOption) (*Score, error)
	SubmitScore(ctx context.Context, in *Score, opts ...grpc.CallOption) (*EmptyMessage, error)
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

func (c *scoringServiceClient) GetApparatusesWithoutPanel(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*ApparatusList, error) {
	out := new(ApparatusList)
	err := c.cc.Invoke(ctx, "/scoring_pb.ScoringService/GetApparatusesWithoutPanel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoringServiceClient) CreateJudgingPanelsForApparatus(ctx context.Context, in *CreateJudgingPanelsForApparatusRequest, opts ...grpc.CallOption) (*CreateJudgingPanelsForApparatusResponse, error) {
	out := new(CreateJudgingPanelsForApparatusResponse)
	err := c.cc.Invoke(ctx, "/scoring_pb.ScoringService/CreateJudgingPanelsForApparatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoringServiceClient) AssignJudge(ctx context.Context, in *AssignJudgeRequest, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, "/scoring_pb.ScoringService/AssignJudge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoringServiceClient) GetAssignedJudges(ctx context.Context, in *IdMessage, opts ...grpc.CallOption) (*JudgeList, error) {
	out := new(JudgeList)
	err := c.cc.Invoke(ctx, "/scoring_pb.ScoringService/GetAssignedJudges", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoringServiceClient) AssignScoreCalculation(ctx context.Context, in *AssignScoreCalculationRequest, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, "/scoring_pb.ScoringService/AssignScoreCalculation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoringServiceClient) GetLoggedJudgeInfo(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (*JudgeJudgingInfo, error) {
	out := new(JudgeJudgingInfo)
	err := c.cc.Invoke(ctx, "/scoring_pb.ScoringService/GetLoggedJudgeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoringServiceClient) GetCurrentApparatusContestants(ctx context.Context, in *GetByApparatusRequest, opts ...grpc.CallOption) (*ContestantList, error) {
	out := new(ContestantList)
	err := c.cc.Invoke(ctx, "/scoring_pb.ScoringService/GetCurrentApparatusContestants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoringServiceClient) GetNextCurrentApparatusContestant(ctx context.Context, in *GetByApparatusRequest, opts ...grpc.CallOption) (*Contestant, error) {
	out := new(Contestant)
	err := c.cc.Invoke(ctx, "/scoring_pb.ScoringService/GetNextCurrentApparatusContestant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoringServiceClient) SubmitTempScore(ctx context.Context, in *TempScore, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, "/scoring_pb.ScoringService/SubmitTempScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoringServiceClient) GetContestantsTempScores(ctx context.Context, in *ScoreRequest, opts ...grpc.CallOption) (*TempScoreList, error) {
	out := new(TempScoreList)
	err := c.cc.Invoke(ctx, "/scoring_pb.ScoringService/GetContestantsTempScores", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoringServiceClient) CalculateScore(ctx context.Context, in *ScoreRequest, opts ...grpc.CallOption) (*Score, error) {
	out := new(Score)
	err := c.cc.Invoke(ctx, "/scoring_pb.ScoringService/CalculateScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scoringServiceClient) SubmitScore(ctx context.Context, in *Score, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, "/scoring_pb.ScoringService/SubmitScore", in, out, opts...)
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
	// Judging panel
	GetApparatusesWithoutPanel(context.Context, *IdMessage) (*ApparatusList, error)
	CreateJudgingPanelsForApparatus(context.Context, *CreateJudgingPanelsForApparatusRequest) (*CreateJudgingPanelsForApparatusResponse, error)
	AssignJudge(context.Context, *AssignJudgeRequest) (*EmptyMessage, error)
	GetAssignedJudges(context.Context, *IdMessage) (*JudgeList, error)
	AssignScoreCalculation(context.Context, *AssignScoreCalculationRequest) (*EmptyMessage, error)
	// Live scoring
	// Judge and contestant info
	GetLoggedJudgeInfo(context.Context, *EmptyMessage) (*JudgeJudgingInfo, error)
	GetCurrentApparatusContestants(context.Context, *GetByApparatusRequest) (*ContestantList, error)
	GetNextCurrentApparatusContestant(context.Context, *GetByApparatusRequest) (*Contestant, error)
	// Scores
	SubmitTempScore(context.Context, *TempScore) (*EmptyMessage, error)
	GetContestantsTempScores(context.Context, *ScoreRequest) (*TempScoreList, error)
	CalculateScore(context.Context, *ScoreRequest) (*Score, error)
	SubmitScore(context.Context, *Score) (*EmptyMessage, error)
	mustEmbedUnimplementedScoringServiceServer()
}

// UnimplementedScoringServiceServer must be embedded to have forward compatible implementations.
type UnimplementedScoringServiceServer struct {
}

func (UnimplementedScoringServiceServer) StartCompetition(context.Context, *IdMessage) (*EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartCompetition not implemented")
}
func (UnimplementedScoringServiceServer) GetApparatusesWithoutPanel(context.Context, *IdMessage) (*ApparatusList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApparatusesWithoutPanel not implemented")
}
func (UnimplementedScoringServiceServer) CreateJudgingPanelsForApparatus(context.Context, *CreateJudgingPanelsForApparatusRequest) (*CreateJudgingPanelsForApparatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJudgingPanelsForApparatus not implemented")
}
func (UnimplementedScoringServiceServer) AssignJudge(context.Context, *AssignJudgeRequest) (*EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignJudge not implemented")
}
func (UnimplementedScoringServiceServer) GetAssignedJudges(context.Context, *IdMessage) (*JudgeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssignedJudges not implemented")
}
func (UnimplementedScoringServiceServer) AssignScoreCalculation(context.Context, *AssignScoreCalculationRequest) (*EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignScoreCalculation not implemented")
}
func (UnimplementedScoringServiceServer) GetLoggedJudgeInfo(context.Context, *EmptyMessage) (*JudgeJudgingInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLoggedJudgeInfo not implemented")
}
func (UnimplementedScoringServiceServer) GetCurrentApparatusContestants(context.Context, *GetByApparatusRequest) (*ContestantList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentApparatusContestants not implemented")
}
func (UnimplementedScoringServiceServer) GetNextCurrentApparatusContestant(context.Context, *GetByApparatusRequest) (*Contestant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNextCurrentApparatusContestant not implemented")
}
func (UnimplementedScoringServiceServer) SubmitTempScore(context.Context, *TempScore) (*EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTempScore not implemented")
}
func (UnimplementedScoringServiceServer) GetContestantsTempScores(context.Context, *ScoreRequest) (*TempScoreList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContestantsTempScores not implemented")
}
func (UnimplementedScoringServiceServer) CalculateScore(context.Context, *ScoreRequest) (*Score, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateScore not implemented")
}
func (UnimplementedScoringServiceServer) SubmitScore(context.Context, *Score) (*EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitScore not implemented")
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

func _ScoringService_GetApparatusesWithoutPanel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoringServiceServer).GetApparatusesWithoutPanel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scoring_pb.ScoringService/GetApparatusesWithoutPanel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoringServiceServer).GetApparatusesWithoutPanel(ctx, req.(*IdMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoringService_CreateJudgingPanelsForApparatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateJudgingPanelsForApparatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoringServiceServer).CreateJudgingPanelsForApparatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scoring_pb.ScoringService/CreateJudgingPanelsForApparatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoringServiceServer).CreateJudgingPanelsForApparatus(ctx, req.(*CreateJudgingPanelsForApparatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoringService_AssignJudge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignJudgeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoringServiceServer).AssignJudge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scoring_pb.ScoringService/AssignJudge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoringServiceServer).AssignJudge(ctx, req.(*AssignJudgeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoringService_GetAssignedJudges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoringServiceServer).GetAssignedJudges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scoring_pb.ScoringService/GetAssignedJudges",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoringServiceServer).GetAssignedJudges(ctx, req.(*IdMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoringService_AssignScoreCalculation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignScoreCalculationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoringServiceServer).AssignScoreCalculation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scoring_pb.ScoringService/AssignScoreCalculation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoringServiceServer).AssignScoreCalculation(ctx, req.(*AssignScoreCalculationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoringService_GetLoggedJudgeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoringServiceServer).GetLoggedJudgeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scoring_pb.ScoringService/GetLoggedJudgeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoringServiceServer).GetLoggedJudgeInfo(ctx, req.(*EmptyMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoringService_GetCurrentApparatusContestants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByApparatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoringServiceServer).GetCurrentApparatusContestants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scoring_pb.ScoringService/GetCurrentApparatusContestants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoringServiceServer).GetCurrentApparatusContestants(ctx, req.(*GetByApparatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoringService_GetNextCurrentApparatusContestant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByApparatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoringServiceServer).GetNextCurrentApparatusContestant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scoring_pb.ScoringService/GetNextCurrentApparatusContestant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoringServiceServer).GetNextCurrentApparatusContestant(ctx, req.(*GetByApparatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoringService_SubmitTempScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TempScore)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoringServiceServer).SubmitTempScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scoring_pb.ScoringService/SubmitTempScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoringServiceServer).SubmitTempScore(ctx, req.(*TempScore))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoringService_GetContestantsTempScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoringServiceServer).GetContestantsTempScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scoring_pb.ScoringService/GetContestantsTempScores",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoringServiceServer).GetContestantsTempScores(ctx, req.(*ScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoringService_CalculateScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoringServiceServer).CalculateScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scoring_pb.ScoringService/CalculateScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoringServiceServer).CalculateScore(ctx, req.(*ScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ScoringService_SubmitScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Score)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScoringServiceServer).SubmitScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scoring_pb.ScoringService/SubmitScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScoringServiceServer).SubmitScore(ctx, req.(*Score))
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
		{
			MethodName: "GetApparatusesWithoutPanel",
			Handler:    _ScoringService_GetApparatusesWithoutPanel_Handler,
		},
		{
			MethodName: "CreateJudgingPanelsForApparatus",
			Handler:    _ScoringService_CreateJudgingPanelsForApparatus_Handler,
		},
		{
			MethodName: "AssignJudge",
			Handler:    _ScoringService_AssignJudge_Handler,
		},
		{
			MethodName: "GetAssignedJudges",
			Handler:    _ScoringService_GetAssignedJudges_Handler,
		},
		{
			MethodName: "AssignScoreCalculation",
			Handler:    _ScoringService_AssignScoreCalculation_Handler,
		},
		{
			MethodName: "GetLoggedJudgeInfo",
			Handler:    _ScoringService_GetLoggedJudgeInfo_Handler,
		},
		{
			MethodName: "GetCurrentApparatusContestants",
			Handler:    _ScoringService_GetCurrentApparatusContestants_Handler,
		},
		{
			MethodName: "GetNextCurrentApparatusContestant",
			Handler:    _ScoringService_GetNextCurrentApparatusContestant_Handler,
		},
		{
			MethodName: "SubmitTempScore",
			Handler:    _ScoringService_SubmitTempScore_Handler,
		},
		{
			MethodName: "GetContestantsTempScores",
			Handler:    _ScoringService_GetContestantsTempScores_Handler,
		},
		{
			MethodName: "CalculateScore",
			Handler:    _ScoringService_CalculateScore_Handler,
		},
		{
			MethodName: "SubmitScore",
			Handler:    _ScoringService_SubmitScore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scoring.proto",
}
