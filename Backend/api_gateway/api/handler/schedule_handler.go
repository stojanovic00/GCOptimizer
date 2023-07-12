package handler

import (
	"github.com/gin-gonic/gin"
)

type ScheduleHandler struct {
}

func (h *ScheduleHandler) Test(ctx *gin.Context) {
	ctx.String(200, "%s", "Success!")
}
