# Structure Folder Clean Architecture Go

├── cmd
│   └── app
│       └── main.go
├── db
│   ├── db.go
│   ├── models.go
│   ├── products.sql.go
│   ├── query
│   │   ├── products.sql
│   │   └── users.sql
│   ├── schema
│   │   ├── products.sql
│   │   └── users.sql
│   └── users.sql.go
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   ├── app.go
│   │   ├── wire_gen.go
│   │   ├── wire.go
│   │   └── wire_set.go
│   ├── config
│   │   ├── config.go
│   │   └── database.go
│   ├── domain
│   │   ├── product.go
│   │   └── user.go
│   ├── handler
│   │   ├── handler.go
│   │   ├── product_handler.go
│   │   └── user_handler.go
│   ├── logger
│   │   └── logger.go
│   ├── middleware
│   │   ├── api_key.go
│   │   └── logging.go
│   ├── repository
│   │   ├── product_repo.go
│   │   └── user_repo.go
│   ├── router
│   │   └── router.go
│   ├── service
│   │   ├── product_service.go
│   │   └── user_service.go
│   └── validator
│       ├── product_validator.go
│       └── user_validator.go
├── note.txt
├── sqlc.yml
├── storage
│   └── logs
│       └── app.log
└── uploads
    └── avatar1.jpg
