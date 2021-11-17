package main

import (
	"backend-task/internal/handlers"
	"backend-task/internal/repository"
	"backend-task/internal/server"
	"backend-task/pgk/config"
	"backend-task/pgk/logger"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := logger.New()    // Logger init
	cfg, err := config.Init() // Config init
	if err != nil {
		logger.Error(err.Error())
	}

	repo := repository.New(cfg.Ttl, logger)
	handler := handlers.New(repo, cfg.Workers, logger)
	server := server.New(cfg.ServerPort, handler)

	go func() {
		if err := server.Run(); err != nil && err != http.ErrServerClosed {
			logger.Fatal(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Stop(ctx); err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("Server Exited Properly")
}
