package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/beeploop/our-srts/internal/config"
	"github.com/beeploop/our-srts/internal/infrastructure/http"
	"github.com/beeploop/our-srts/internal/infrastructure/persistence"
	"github.com/beeploop/our-srts/internal/infrastructure/storage"
	"github.com/beeploop/our-srts/internal/server"
	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.Load()

	logfile, err := os.OpenFile(cfg.LOG_FILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %s\n", err.Error())
	}
	defer logfile.Close()

	logger := slog.New(slog.NewJSONHandler(logfile, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	db, err := persistence.NewMysql(mysql.Config{
		User:                 cfg.DB_USER,
		Passwd:               cfg.DB_PASS,
		Net:                  cfg.DB_NET,
		Addr:                 cfg.DB_HOST + ":" + cfg.DB_PORT,
		DBName:               cfg.DB_NAME,
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatalf("could not start db: %s\n", err.Error())
	}

	fs := storage.NewDiskStorage(cfg.UPLOAD_DIR)

	handler := http.NewRouter(cfg, db, fs)

	srv := server.NewServer(cfg, handler.Echo)

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
