package handler

import (
	"auth_service/api/middleware"
	scheduling_pb "common/proto/scheduling/generated"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SchedulingHandler struct {
	client scheduling_pb.SchedulingServiceClient
}

func NewSchedulingHandler(client scheduling_pb.SchedulingServiceClient) *SchedulingHandler {
	return &SchedulingHandler{client: client}
}

func (h *SchedulingHandler) GenerateSchedule(ctx *gin.Context) {
	var params scheduling_pb.SchedulingParameters

	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Parsing error"})
		return
	}

	userInfoContext, err := middleware.GetGrpcContextWithUserInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No user info provided"})
		return
	}
	scheduleDto, err := h.client.GenerateSchedule(userInfoContext, &params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, scheduleDto)
}
