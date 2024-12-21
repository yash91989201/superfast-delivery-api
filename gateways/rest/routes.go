package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func registerRoutes(handler *handler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return r
}
