package routes

import (
	"go-backend-server-template/internal/app"

	"github.com/go-chi/chi/v5"
)

func SetUpRoutes(a app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health")
}
