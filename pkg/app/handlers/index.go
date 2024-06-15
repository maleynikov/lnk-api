package handlers

import (
	"net/http"

	"github.com/go-chi/render"
	"gorm.io/gorm"
)

type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("URL Short API!"))
}

type ErrResponse struct {
	Err        error  `json:"-"`
	StatusText string `json:"status"`
	ErrorText  string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, 400)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:        err,
		StatusText: "Invalid request",
		ErrorText:  err.Error(),
	}
}
