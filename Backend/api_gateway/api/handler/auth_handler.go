package handler

import (
	"github.com/gin-gonic/gin"
)

type AuthHadler struct {
}

func (h *AuthHadler) Login(ctx *gin.Context) {
	ctx.String(201, "%s", "Success!")
}
