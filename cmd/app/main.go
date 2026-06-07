package main

import (
	"go-dbsqlc/internal/app"
	"go-dbsqlc/internal/config"
	"go-dbsqlc/internal/logger"
	"go-dbsqlc/internal/router"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// load config
	cfg := config.LoadConfig()

	// logger init
	logger.InitLogger(cfg.LogLevel)
	slog.Info("config loaded ...")

	// init app
	app, cleanup, err := app.InitializeApp(cfg)
	if err != nil {
		slog.Error("failed initialize application")
		os.Exit(1)
	}
	defer cleanup()
	slog.Info("dependencies initialized")

	// init router
	r := router.NewRouter(app)
	slog.Info("router initialized")

	addr := ":" + cfg.AppPort
	slog.Info("server started", "addr", addr)

	// start server
	if err := http.ListenAndServe(addr, r); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
