package handler

import (
	"auth_service/api/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ScheduleHandler struct {
}

func (h *ScheduleHandler) Test(ctx *gin.Context) {
	userInfo, err := middleware.ParseUserInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, userInfo)
}
