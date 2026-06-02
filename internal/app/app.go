package app

import (
	"go-dbsqlc/db"
	"go-dbsqlc/internal/config"
	handler "go-dbsqlc/internal/handler"
	"go-dbsqlc/internal/repository"
	"go-dbsqlc/internal/service"
)

type App struct {
	Config *config.Config
	UserHandler *handler.UserHandler
}

func NewApp(cfg *config.Config) *App {
	sqlDB := config.NewMySQL(cfg)
	db := db.New(sqlDB)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	return &App{
		Config: cfg,
		UserHandler: userHandler,
	}
}
