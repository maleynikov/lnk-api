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

type Status string

const (
	StatusSuccess Status = "success"
	StatusError   Status = "error"
)

type Response struct {
	Status  `json:"status"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func (rd *Response) Render(w http.ResponseWriter, r *http.Request) error {
	switch rd.Status {
	case StatusError:
		render.Status(r, 500)
	default:
		render.Status(r, 200)
	}
	return nil
}

func Err(err error) render.Renderer {
	return &Response{
		Status:  StatusError,
		Message: err.Error(),
	}
}

func Ok(data any) render.Renderer {
	return &Response{
		Status: StatusSuccess,
		Data:   data,
	}
}
