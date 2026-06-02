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

	mux.HandleFunc("POST /users", app.UserHandler.Create)
	mux.HandleFunc("GET /users", app.UserHandler.List)
	mux.HandleFunc("GET /users/{id}", app.UserHandler.GetById)
	mux.HandleFunc("PUT /users/{id}", app.UserHandler.Update)
	mux.HandleFunc("DELETE /users/{id}", app.UserHandler.Delete)
	mux.HandleFunc("POST /users/{id}/avatar", app.UserHandler.UploadAvatar)

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
