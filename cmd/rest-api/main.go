package main

import (
	config2 "RestApi/config"
	"RestApi/internal/http/server"
	"RestApi/pkg/logger"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func gracefulShutdown(srv *http.Server, log *logger.Logger) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown")
	}
	log.Info("Server exiting")
}

func main() {
	config, err := config2.LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}
	log := logger.NewLogger()

	mux := http.NewServeMux()
	server.RegRoutes(mux, log)

	srv := &http.Server{
		Addr:    ":" + config.ServerPort,
		Handler: mux,
	}

	go func() {
		log.Info("Starting server on port " + config.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed to start")
		}
	}()
	gracefulShutdown(srv, log)
}
