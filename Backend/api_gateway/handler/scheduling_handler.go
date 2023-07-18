package handler

import (
	scheduling_pb "common/proto/scheduling/generated"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SchedulingHandler struct {
	client scheduling_pb.SchedulingServiceClient
}

func NewSchedulingHandler(client scheduling_pb.SchedulingServiceClient) *SchedulingHandler {
	return &SchedulingHandler{client: client}
}

func (h *SchedulingHandler) Test(ctx *gin.Context) {
	var message scheduling_pb.TestMessage

	err := ctx.ShouldBindJSON(&message)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parsing error"})
		return
	}

	resp, err := h.client.Test(context.Background(), &scheduling_pb.TestMessage{Message: message.Message})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp.Response)
}
