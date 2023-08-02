package app

import (
	"api_gateway/config"
	"log"
)
import "github.com/gin-gonic/gin"

type App struct {
	Config        config.Config
	PublicRouter  *gin.Engine
	PrivateRouter *gin.Engine
}

func NewApp() (*App, error) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	app := &App{
		Config: config,
	}

	err = app.CreateRoutersAndSetRoutes()
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return app, nil
}
