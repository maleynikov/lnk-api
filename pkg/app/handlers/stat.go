package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

type Data []struct {
	Value     string `json:"url"`
	IP        string `json:"ip"`
	CreatedAt string `json:"created_at"`
}

func (h *Handler) Stat(w http.ResponseWriter, r *http.Request) {

	query := h.db.Table("statistics s").
		Joins("left join urls on s.url_id = urls.id").
		Joins("left join users on s.user_id = users.id")

	oid := r.URL.Query().Get("oid")
	if oid != "" {
		query = query.Where("urls.code = ?", oid)
	}

	data := Data{}
	query.Scan(&data)

	render.Status(r, http.StatusOK)
	render.Render(w, r, Ok(data))
}
