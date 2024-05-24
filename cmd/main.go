package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/hyper-dot/books-server/internal/handler"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/user", handler.UserHandler())

	fmt.Println("Listening to port 8080 !!")
	http.ListenAndServe(":8080", r)
}
