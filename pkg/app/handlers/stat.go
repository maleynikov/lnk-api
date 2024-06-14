package handlers

import "net/http"

func StatHandler(w http.ResponseWriter, r *http.Request) {
	oid := r.URL.Query().Get("oid")
	if oid == "" {
		w.Write([]byte("All stat"))
		return
	}
	w.Write([]byte("Stat OID: " + oid))
}
