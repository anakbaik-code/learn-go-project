package app

import (
	"go-dbsqlc/internal/handler"
	"go-dbsqlc/internal/repository"
	"go-dbsqlc/internal/service"

	"github.com/google/wire"
)

// User Set
var UserSet = wire.NewSet(
	repository.NewUserRepository,
	service.NewUserService,
	handler.NewUserHandler,
)

// Product Set
var ProductSet = wire.NewSet(
	repository.NewProductRepository,
	service.NewProductService,
	handler.NewProductHandler,
)
