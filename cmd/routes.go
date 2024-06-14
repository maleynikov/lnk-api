package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"gitlab.maleynikov.me/url-short/api/pkg/app/handlers"
)

func routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Define routes
	r.Get("/", handlers.IndexHandler)
	r.Post("/short", handlers.ShortHandler)
	r.Get("/r/{oid}", handlers.RedirectHandler)
	r.Get("/stat", handlers.StatHandler)

	return r
}
