package main

import (
	"HahaTeam/internal/config"
	"HahaTeam/internal/http-server/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	cfg := config.MustLoad()
	fmt.Printf("%#v\n", cfg)

	log := setUpLogger(cfg.Env)
	log.Info("starting work program", slog.String("env", cfg.Env))
	log.Debug("debug log enabled", slog.String("env", cfg.Env))

	/*strorage, err := postgres.NewStorage(cfg.StoragePath)
	if err != nil {
	log.Error("failed to init storage", err)
	os.Exit(1)
	}
	_ = storage*/

	router := gin.Default()
	router.POST("/register", handlers.CreateNewUser(nil))
	router.POST("/login", handlers.AuthenticationUser(nil))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}
	log.Error("server stopped")

}

func setUpLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)

	}

	return log
}
