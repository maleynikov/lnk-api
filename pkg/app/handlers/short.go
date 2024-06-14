package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"gitlab.maleynikov.me/url-short/api/pkg/util"
)

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

func ShortHandler(w http.ResponseWriter, r *http.Request) {
	data := &ShortRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	w.Write([]byte("URL OID: " + util.OID(data.URL)))
}
