package handler

import (
	scoring_pb "common/proto/scoring/generated"
	"context"
	"github.com/google/uuid"
	"scoring_service/api/mapper"
	"scoring_service/core/domain"
	"scoring_service/core/service"
)

type HandlerRpc struct {
	scoring_pb.UnimplementedScoringServiceServer
	compService *service.ScheduleService
	jpService   *service.JudgePanelService
}

func NewHandlerRpc(compService *service.ScheduleService, jpService *service.JudgePanelService) *HandlerRpc {
	return &HandlerRpc{compService: compService, jpService: jpService}
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
func (h *HandlerRpc) GetApparatusesWithoutPanel(ctx context.Context, id *scoring_pb.IdMessage) (*scoring_pb.ApparatusList, error) {
	compId, _ := uuid.Parse(id.Id)
	apparatuses, err := h.jpService.GetApparatusesWithoutPanel(compId)
	if err != nil {
		return &scoring_pb.ApparatusList{}, nil
	}

	return &scoring_pb.ApparatusList{Apparatuses: mapper.ApparatusListDomToPb(apparatuses)}, nil
}

func (h *HandlerRpc) CreateJudgingPanelsForApparatus(ctx context.Context, request *scoring_pb.CreateJudgingPanelsForApparatusRequest) (*scoring_pb.CreateJudgingPanelsForApparatusResponse, error) {
	compId, _ := uuid.Parse(request.CompetitionId)
	apparatus := domain.Apparatus(request.Apparatus.Number())

	dPanId, ePanId, err := h.jpService.CreateJudgingPanelsForApparatus(apparatus, compId)
	if err != nil {
		return &scoring_pb.CreateJudgingPanelsForApparatusResponse{}, nil
	}

	return &scoring_pb.CreateJudgingPanelsForApparatusResponse{
		DPanelId: dPanId.String(),
		EPanelId: ePanId.String(),
	}, nil
}
func (h *HandlerRpc) AssignJudge(ctx context.Context, request *scoring_pb.AssignJudgeRequest) (*scoring_pb.EmptyMessage, error) {
	judge := mapper.JudgePbToDom(request.Judge)
	panelId, _ := uuid.Parse(request.PanelId)

	err := h.jpService.AssignJudge(judge, panelId)
	if err != nil {
		return &scoring_pb.EmptyMessage{}, err
	}

	return &scoring_pb.EmptyMessage{}, nil
}

func (h *HandlerRpc) GetAssignedJudges(ctx context.Context, id *scoring_pb.IdMessage) (*scoring_pb.JudgeList, error) {
	compId, _ := uuid.Parse(id.Id)

	judges, err := h.jpService.GetAssignedJudges(compId)
	if err != nil {
		return &scoring_pb.JudgeList{}, err
	}

	return &scoring_pb.JudgeList{Judges: mapper.JudgeListDomToPb(judges)}, err
}
func (h *HandlerRpc) AssignScoreCalculation(ctx context.Context, request *scoring_pb.AssignScoreCalculationRequest) (*scoring_pb.EmptyMessage, error) {
	method := mapper.ScoreCalcMethodPbToDom(request.Method)
	panelId, _ := uuid.Parse(request.PanelId)

	err := h.jpService.AssignScoreCalculationMethod(method, panelId)
	if err != nil {
		return &scoring_pb.EmptyMessage{}, err
	}

	return &scoring_pb.EmptyMessage{}, nil
}
