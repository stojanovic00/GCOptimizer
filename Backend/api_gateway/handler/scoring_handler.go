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
	var id scoring_pb.IdMessage

	err := ctx.ShouldBindJSON(&id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parsing error"})
		return
	}

	_, err = h.client.StartCompetition(context.Background(), &id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}
func (h *ScoringHandler) GetApparatusesWithoutPanel(ctx *gin.Context) {
	var compId scoring_pb.IdMessage

	err := ctx.ShouldBindJSON(&compId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parsing error"})
		return
	}

	result, err := h.client.GetApparatusesWithoutPanel(context.Background(), &compId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, result)
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
	var request scoring_pb.AssignJudgeRequest
	panelId := ctx.Param("id")
	request.PanelId = panelId

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parsing error"})
		return
	}

	result, err := h.client.AssignJudge(context.Background(), &request)
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

	ctx.JSON(http.StatusOK, result)
}
func (h *ScoringHandler) AssignScoreCalculationMethod(ctx *gin.Context) {
	var request scoring_pb.AssignScoreCalculationRequest
	panelId := ctx.Param("id")
	request.PanelId = panelId

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parsing error"})
		return
	}

	_, err = h.client.AssignScoreCalculation(context.Background(), &request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}
