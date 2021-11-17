package main

import (
	"backend-task/internal/handlers"
	"backend-task/internal/repository"
	"backend-task/internal/server"
	"backend-task/pgk/config"
	"backend-task/pgk/logger"
	"context"
	"log"
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

	repo := repository.New(cfg, logger)
	handler := handlers.New(repo, cfg.Ttl, logger)
	server := server.New(repo, handler)

	go func() {
		if err := server.Run(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
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
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
