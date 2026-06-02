package main

import (
	"go-dbsqlc/internal/app"
	"go-dbsqlc/internal/config"
	"go-dbsqlc/internal/logger"
	"go-dbsqlc/internal/router"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	// logger init
	logfile := logger.InitLogger()
	defer logfile.Close()

	// load config
	cfg := config.LoadConfig()
	slog.Info("config loaded ...")

	// init app
	application := app.NewApp(cfg)
	slog.Info("dependencies initialized")

	// init router
	r := router.NewRouter(application)
	slog.Info("router initialized")

	addr := ":" + cfg.AppPort
	slog.Info("server running on", "addr", addr)

	// start server
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal("server failed : ", err)
	}
}
