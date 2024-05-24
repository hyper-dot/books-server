package handler

import (
	"github.com/go-chi/chi/v5"

	"github.com/hyper-dot/books-server/internal/service"
)

func UserHandler() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", service.AddNewUser)

	return r
}
