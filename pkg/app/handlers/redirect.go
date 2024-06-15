package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gitlab.maleynikov.me/url-short/api/pkg/app/models"
)

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	oid := chi.URLParam(r, "oid")
	url := models.Url{}
	result := h.db.Where("code = ?", oid).First(&url)
	if result.Error != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	stat := models.Statistic{
		Url: url,
		IP:  r.RemoteAddr,
	}
	h.db.Create(&stat)
	http.Redirect(w, r, url.Value, http.StatusMovedPermanently)
}
