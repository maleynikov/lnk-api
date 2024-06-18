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

func (b *ShortRequest) Bind(r *http.Request) error {
	if b.URL == "" {
		return errors.New("missing required URL field")
	}
	return nil
}

func (h *Handler) Short(w http.ResponseWriter, r *http.Request) {
	data := &ShortRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, Err(err))
		return
	}
	url := models.Url{
		Code:  util.OID(data.URL),
		Value: data.URL,
	}
	h.db.Create(&url)

	render.Status(r, http.StatusCreated)
	render.Render(w, r, Ok(struct {
		Code string `json:"code"`
	}{Code: url.Code}))
}
