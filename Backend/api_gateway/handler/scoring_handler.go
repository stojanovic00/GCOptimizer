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
