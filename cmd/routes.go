package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Define routes
	r.Get("/", indexHandler)
	r.Post("/short", shortHandler)
	r.Get("/r/{oid}", redirectHandler)

	return r
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello API!"))
}

func shortHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Short URL"))
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	oid := chi.URLParam(r, "oid")
	w.Write([]byte("Here is Open ID: " + oid))
}
