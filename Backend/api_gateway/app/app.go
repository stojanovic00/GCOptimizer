package app

import (
	"api_gateway/config"
)
import "github.com/gin-gonic/gin"

type App struct {
	Config        config.Config
	PublicRouter  *gin.Engine
	PrivateRouter *gin.Engine
}
