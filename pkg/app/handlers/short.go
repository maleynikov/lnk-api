package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"gitlab.maleynikov.me/url-short/api/pkg/app/models"
	"gitlab.maleynikov.me/url-short/api/pkg/util"
)

type ShortRequest struct {
	URL string `json:"url"`
}

type ShortResponse struct {
	Code string `json:"code"`
}

func (b *ShortResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (b *ShortRequest) Bind(r *http.Request) error {
	if b.URL == "" {
		return errors.New("missing required URL field")
	}
	return nil
}

func (h *Handler) Short(w http.ResponseWriter, r *http.Request) {
	data := &ShortRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	url := models.Url{
		Code:  util.OID(data.URL),
		Value: data.URL,
	}
	h.db.Create(&url)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, &ShortResponse{Code: url.Code})
}
