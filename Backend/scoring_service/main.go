package main

import (
	"log"
	"scoring_service/api/app"
	"scoring_service/config"
)

func main() {
	app := &app.App{}

	var err error
	app.Config, err = config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	err = app.Run()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
