package main

import (
	app "auth_service/api/app"
	"auth_service/config"
	"log"
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
