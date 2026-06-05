package app

import (
	"go-dbsqlc/internal/config"
	handler "go-dbsqlc/internal/handler"
	
)

type App struct {
	Config  *config.Config
	Handler *handler.Handlers
}

func NewApp(config *config.Config,handler *handler.Handlers) *App{
	return &App{
		Config: config,
		Handler: handler,
	}
}
