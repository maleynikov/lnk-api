package main

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"gitlab.maleynikov.me/url-short/api/pkg/util"
)

type ShortPayload struct {
	URL string `json:"url"`
}

type ShortRequest struct {
	*ShortPayload
}

func (b *ShortRequest) Bind(r *http.Request) error {
	if b.ShortPayload == nil {
		return errors.New("missing required ShortPayload fields")
	}
	return nil
}

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	StatusText string `json:"status"`
	AppCode    int64  `json:"code,omitempty"`
	ErrorText  string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

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
	r.Get("/stat", statHandler)

	return r
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("URL Short API!"))
}

func shortHandler(w http.ResponseWriter, r *http.Request) {
	data := &ShortRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	w.Write([]byte("URL OID: " + util.OID(data.URL)))
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	oid := chi.URLParam(r, "oid")
	w.Write([]byte("Here is Open ID: " + oid))
}

func statHandler(w http.ResponseWriter, r *http.Request) {
	oid := r.URL.Query().Get("oid")
	if oid == "" {
		w.Write([]byte("All stat"))
		return
	}
	w.Write([]byte("Stat OID: " + oid))
}
