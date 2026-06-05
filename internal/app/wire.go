//go:build wireinject

package app

import (
	"database/sql"
	"go-dbsqlc/db"
	"go-dbsqlc/internal/config"
	"go-dbsqlc/internal/handler"


	"github.com/google/wire"
)

func InitializeApp(cfg *config.Config) (*App, func(), error) {
    wire.Build(
        // DB provides
        config.NewMySQL,
        db.New,
        wire.Bind(new(db.DBTX), new(*sql.DB)),

        // App Provides
        UserSet,
        ProductSet,
        
        handler.NewHandlers,
        wire.Struct(new(handler.HandlersParam), "*"),       
        NewApp,
    )
    return nil, nil, nil 
}
