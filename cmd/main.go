package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/beeploop/our-srts/internal/config"
	"github.com/beeploop/our-srts/internal/infrastructure/http"
	"github.com/beeploop/our-srts/internal/server"
)

func main() {
	cfg := config.Load()

	handler := http.SetupRouter()

	srv := server.NewServer(cfg, handler)

	go func() {
		log.Printf("starting server on %s\n", cfg.PORT)

		if err := srv.Start(); err != nil {
			log.Fatalf("could not start server: %s\n", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Printf("shutting down server...")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("server shutdown failed: %s\n", err.Error())
	}

	log.Printf("server exited")
}
