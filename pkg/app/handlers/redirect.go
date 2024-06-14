package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	oid := chi.URLParam(r, "oid")
	w.Write([]byte("Here is Open ID: " + oid))
}
