package handler

import (
	scoring_pb "common/proto/scoring/generated"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ScoringHandler struct {
	client scoring_pb.ScoringServiceClient
}

func NewScoringHandler(client scoring_pb.ScoringServiceClient) *ScoringHandler {
	return &ScoringHandler{client: client}
}

func (h *ScoringHandler) StartCompetition(ctx *gin.Context) {
	compId := ctx.Param("id")

	_, err := h.client.StartCompetition(context.Background(), &scoring_pb.IdMessage{Id: compId})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}
func (h *ScoringHandler) GetApparatusesWithoutPanel(ctx *gin.Context) {
	compId := ctx.Param("id")

	result, err := h.client.GetApparatusesWithoutPanel(context.Background(), &scoring_pb.IdMessage{Id: compId})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, result.Apparatuses)
}

func (h *ScoringHandler) CreateJudgingPanelsForApparatus(ctx *gin.Context) {
	var request scoring_pb.CreateJudgingPanelsForApparatusRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parsing error"})
		return
	}

	result, err := h.client.CreateJudgingPanelsForApparatus(context.Background(), &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, result)
}
func (h *ScoringHandler) AssignJudge(ctx *gin.Context) {
	var judge scoring_pb.Judge
	panelId := ctx.Param("id")

	err := ctx.ShouldBindJSON(&judge)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parsing error"})
		return
	}

	result, err := h.client.AssignJudge(context.Background(), &scoring_pb.AssignJudgeRequest{
		Judge:   &judge,
		PanelId: panelId,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, result)
}
func (h *ScoringHandler) GetAssignedJudges(ctx *gin.Context) {
	competitionId := ctx.Param("id")

	result, err := h.client.GetAssignedJudges(context.Background(), &scoring_pb.IdMessage{Id: competitionId})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, result.Judges)
}
func (h *ScoringHandler) AssignScoreCalculationMethod(ctx *gin.Context) {
	var method scoring_pb.ScoreCalculationMethod
	panelId := ctx.Param("id")

	err := ctx.ShouldBindJSON(&method)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parsing error"})
		return
	}

	_, err = h.client.AssignScoreCalculation(context.Background(), &scoring_pb.AssignScoreCalculationRequest{
		Method:  &method,
		PanelId: panelId,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}
