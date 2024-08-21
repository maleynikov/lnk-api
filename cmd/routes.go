package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"gitlab.maleynikov.me/url-short/api/pkg/app/handlers"
)

func (s *Server) routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// Define routes
	h := handlers.NewHandler(s.storage.DB)

	r.Get("/", h.Index)
	r.Post("/cut", h.Short)
	r.Get("/r:{oid}", h.Redirect)
	r.Get("/stat", h.Stat)
	r.Options("/", func(w http.ResponseWriter, r *http.Request) {})

	return r
}
