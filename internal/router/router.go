package router

import (
	"encoding/json"
	"go-dbsqlc/internal/app"
	"go-dbsqlc/internal/middleware"
	"net/http"

	"github.com/justinas/alice"
)

type Response struct {
	Message string `json:"message"`
}

func NewRouter(app *app.App) http.Handler {
	mux := http.NewServeMux()

	// User Route
	mux.HandleFunc("POST /users", app.Handler.User.Create)
	mux.HandleFunc("GET /users", app.Handler.User.List)
	mux.HandleFunc("GET /users/{id}", app.Handler.User.GetById)
	mux.HandleFunc("PUT /users/{id}", app.Handler.User.Update)
	mux.HandleFunc("DELETE /users/{id}", app.Handler.User.Delete)
	mux.HandleFunc("POST /users/{id}/avatar", app.Handler.User.UploadAvatar)

	// Product Route
	mux.HandleFunc("GET /products/{id}", app.Handler.Product.GetById)
	mux.HandleFunc("POST /products", app.Handler.Product.Create)
	mux.HandleFunc("GET /products", app.Handler.Product.List)
	// NotFound Route
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		response := Response{
			Message: "Route Not Found",
		}
		json.NewEncoder(w).Encode(response)
	})

	// Manage Midlleware -> alice library
	chain := alice.New(
		middleware.Logger,
		middleware.ApiKeyMiddleware(app.Config.ApiKey),
	)

	return chain.Then(mux)
}
