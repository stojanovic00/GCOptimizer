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
	scService   *service.ScoringService
}

func NewHandlerRpc(compService *service.ScheduleService, jpService *service.JudgePanelService, scService *service.ScoringService) *HandlerRpc {
	return &HandlerRpc{compService: compService, jpService: jpService, scService: scService}
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

func (h *HandlerRpc) GetLoggedJudgeInfo(ctx context.Context, _ *scoring_pb.EmptyMessage) (*scoring_pb.JudgeJudgingInfo, error) {
	userInfo, err := ParseUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	info, err := h.scService.GetJudgeJudgingInfo(userInfo.Email)
	if err != nil {
		return nil, err
	}

	return mapper.JudgeJudgingInfoDomToPb(info), nil
}
func (h *HandlerRpc) GetCurrentApparatusContestants(ctx context.Context, request *scoring_pb.GetByApparatusRequest) (*scoring_pb.ContestantList, error) {
	compId, _ := uuid.Parse(request.CompetitionId)
	apparatus := domain.Apparatus(request.Apparatus)

	contestants, err := h.scService.GetCurrentApparatusContestants(compId, apparatus)
	if err != nil {
		return nil, err
	}

	return &scoring_pb.ContestantList{Contestants: mapper.ContestantCompetingListDomToPbSorted(contestants, apparatus)}, nil
}

func (h *HandlerRpc) GetNextCurrentApparatusContestant(ctx context.Context, request *scoring_pb.GetByApparatusRequest) (*scoring_pb.Contestant, error) {
	compId, _ := uuid.Parse(request.CompetitionId)
	apparatus := domain.Apparatus(request.Apparatus)

	contestant, err := h.scService.GetNextCurrentApparatusContestant(compId, apparatus)
	if err != nil {
		return nil, err
	}

	return mapper.ContestantDomToPb(contestant), nil
}

func (h *HandlerRpc) SubmitTempScore(ctx context.Context, tempScore *scoring_pb.TempScore) (*scoring_pb.EmptyMessage, error) {
	err := h.scService.SubmitTempScore(mapper.TempScorePbToDom(tempScore))
	if err != nil {
		return nil, err
	}

	return &scoring_pb.EmptyMessage{}, nil
}

func (h *HandlerRpc) GetContestantsTempScores(ctx context.Context, request *scoring_pb.ScoreRequest) (*scoring_pb.TempScoreList, error) {
	competitionId, _ := uuid.Parse(request.CompetitionId)
	contestantId, _ := uuid.Parse(request.ContestantId)
	apparatus := domain.Apparatus(request.Apparatus)

	tempScores, err := h.scService.GetContestantsTempScores(competitionId, contestantId, apparatus)
	if err != nil {
		return nil, err
	}

	return &scoring_pb.TempScoreList{TempScores: mapper.TempScoreListDomToPb(tempScores)}, nil
}

func (h *HandlerRpc) CanCalculateScore(ctx context.Context, request *scoring_pb.ScoreRequest) (*scoring_pb.BoolMessage, error) {
	competitionId, _ := uuid.Parse(request.CompetitionId)
	contestantId, _ := uuid.Parse(request.ContestantId)
	apparatus := domain.Apparatus(request.Apparatus)

	canCalculate, err := h.scService.CanCalculateScore(competitionId, contestantId, apparatus)
	if err != nil {
		return nil, err
	}

	return &scoring_pb.BoolMessage{IsTrue: canCalculate}, nil
}

func (h *HandlerRpc) CalculateScore(ctx context.Context, request *scoring_pb.ScoreRequest) (*scoring_pb.Score, error) {
	competitionId, _ := uuid.Parse(request.CompetitionId)
	contestantId, _ := uuid.Parse(request.ContestantId)
	apparatus := domain.Apparatus(request.Apparatus)

	score, err := h.scService.CalculateScore(competitionId, contestantId, apparatus)
	if err != nil {
		return nil, err
	}

	return mapper.ScoreDomToPb(score), nil
}
func (h *HandlerRpc) SubmitScore(ctx context.Context, score *scoring_pb.Score) (*scoring_pb.EmptyMessage, error) {
	err := h.scService.SubmitScore(mapper.ScorePbToDom(score))
	if err != nil {
		return nil, err
	}

	return &scoring_pb.EmptyMessage{}, nil
}

func (h *HandlerRpc) GetScore(ctx context.Context, request *scoring_pb.ScoreRequest) (*scoring_pb.Score, error) {
	competitionId, _ := uuid.Parse(request.CompetitionId)
	contestantId, _ := uuid.Parse(request.ContestantId)
	apparatus := domain.Apparatus(request.Apparatus)

	score, err := h.scService.GetScore(competitionId, contestantId, apparatus)
	if err != nil {
		return nil, err
	}

	return mapper.ScoreDomToPb(score), nil
}

func (h *HandlerRpc) FinishRotation(ctx context.Context, request *scoring_pb.IdMessage) (*scoring_pb.EmptyMessage, error) {
	competitionId, _ := uuid.Parse(request.Id)
	err := h.scService.FinishRotation(competitionId)
	if err != nil {
		return nil, err
	}

	return &scoring_pb.EmptyMessage{}, nil
}

func (h *HandlerRpc) FinishSession(ctx context.Context, request *scoring_pb.IdMessage) (*scoring_pb.EmptyMessage, error) {
	competitionId, _ := uuid.Parse(request.Id)
	err := h.scService.FinishSession(competitionId)
	if err != nil {
		return nil, err
	}

	return &scoring_pb.EmptyMessage{}, nil
}
func (h *HandlerRpc) FinishCompetition(ctx context.Context, request *scoring_pb.IdMessage) (*scoring_pb.EmptyMessage, error) {
	competitionId, _ := uuid.Parse(request.Id)
	err := h.jpService.DeleteAllJudgeAccounts(competitionId)
	if err != nil {
		return nil, err
	}
	//Generate scoreboards
	err = h.scService.GenerateScoreboards(competitionId)
	if err != nil {
		return nil, err
	}

	return &scoring_pb.EmptyMessage{}, nil
}
func (h *HandlerRpc) IsRotationFinished(ctx context.Context, request *scoring_pb.IdMessage) (*scoring_pb.BoolMessage, error) {
	competitionId, _ := uuid.Parse(request.Id)
	finished, err := h.scService.IsRotationFinished(competitionId)
	if err != nil {
		return nil, err
	}

	return &scoring_pb.BoolMessage{IsTrue: finished}, nil
}
func (h *HandlerRpc) IsSessionFinished(ctx context.Context, request *scoring_pb.IdMessage) (*scoring_pb.BoolMessage, error) {
	competitionId, _ := uuid.Parse(request.Id)
	finished, err := h.scService.IsSessionFinished(competitionId)
	if err != nil {
		return nil, err
	}

	return &scoring_pb.BoolMessage{IsTrue: finished}, nil
}
func (h *HandlerRpc) IsCompetitionFinished(ctx context.Context, request *scoring_pb.IdMessage) (*scoring_pb.BoolMessage, error) {
	competitionId, _ := uuid.Parse(request.Id)
	finished, err := h.scService.IsCompetitionFinished(competitionId)
	if err != nil {
		return nil, err
	}

	return &scoring_pb.BoolMessage{IsTrue: finished}, nil
}
func (h *HandlerRpc) GetCurrentSessionInfo(ctx context.Context, request *scoring_pb.IdMessage) (*scoring_pb.CurrentSessionInfo, error) {
	competitionId, _ := uuid.Parse(request.Id)
	info, err := h.scService.GetCurrentSessionInfo(competitionId)
	if err != nil {
		return nil, err
	}

	return mapper.CurrentSessionInfoDomToPb(info), nil
}
func (h *HandlerRpc) GetScoreboards(ctx context.Context, request *scoring_pb.IdMessage) (*scoring_pb.ScoreBoardBundle, error) {
	competitionId, _ := uuid.Parse(request.Id)

	allAround, err := h.scService.GetAllAroundScoreBoards(competitionId)
	if err != nil {
		return nil, err
	}
	team, err := h.scService.GetTeamScoreBoards(competitionId)
	if err != nil {
		return nil, err
	}

	return &scoring_pb.ScoreBoardBundle{
		AllAroundScoreboards: mapper.AllAroundScoreBoardListDomToPb(allAround),
		TeamScoreboards:      mapper.TeamScoreBoardListDomToPb(team),
	}, nil
}
