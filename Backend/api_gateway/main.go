package main

import (
	"api_gateway/app"
	"api_gateway/config"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	app := &app.App{}

	var err error
	app.Config, err = config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	app.PublicRouter, app.PrivateRouter, err = app.CreateRoutersAndSetRoutes()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	publicAddress := fmt.Sprintf("%s:%s", app.Config.Host, app.Config.PublicPort)
	publicServer := &http.Server{
		Handler:           app.PublicRouter,
		Addr:              publicAddress,
		WriteTimeout:      15 * time.Second,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 100 * time.Millisecond,
		MaxHeaderBytes:    2048,
	}

	go func() {
		if err := publicServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	privateAddress := fmt.Sprintf("%s:%s", app.Config.Host, app.Config.PrivatePort)
	privateServer := &http.Server{
		Handler:           app.PrivateRouter,
		Addr:              privateAddress,
		WriteTimeout:      15 * time.Second,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 100 * time.Millisecond,
		MaxHeaderBytes:    2048,
	}

	go func() {
		if err := privateServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	//GRACEFUL SHUTDOWN
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	timeoutTime := 2
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutTime)*time.Second)
	defer cancel()

	err = publicServer.Shutdown(ctx)
	if err != nil {
		log.Fatal("Public server shutdown:", err)
	}

	err = privateServer.Shutdown(ctx)
	if err != nil {
		log.Fatal("Private server shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Printf("timeout of %d seconds.", timeoutTime)
	}
	log.Println("Server exiting")
}
