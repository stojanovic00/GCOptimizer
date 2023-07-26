package handler

import (
	scoring_pb "common/proto/scoring/generated"
	"context"
	"github.com/google/uuid"
	"scoring_service/core/service"
)

type HandlerRpc struct {
	scoring_pb.UnimplementedScoringServiceServer
	compService *service.ScheduleService
}

func NewHandlerRpc(compService *service.ScheduleService) *HandlerRpc {
	return &HandlerRpc{compService: compService}
}

func (h *HandlerRpc) StartCompetition(ctx context.Context, id *scoring_pb.IdMessage) (*scoring_pb.EmptyMessage, error) {
	compId, err := uuid.Parse(id.Id)
	if err != nil {
		return &scoring_pb.EmptyMessage{}, err
	}

	err = h.compService.StartCompetition(compId)
	if err != nil {
		return &scoring_pb.EmptyMessage{}, err
	}

	return &scoring_pb.EmptyMessage{}, nil
}
