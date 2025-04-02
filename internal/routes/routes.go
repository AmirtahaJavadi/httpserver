package routes

import (
	"fmt"
	"net/http"

	"github.com/amirtahajavadi/httpserver/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", HealthChecker)

	return r
}

func HealthChecker(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "status is available\n")
}
