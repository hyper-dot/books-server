package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// LISTEN AND SERVE
	fmt.Println("Listening to port 8080")
	http.ListenAndServe(":8080", r)
}
