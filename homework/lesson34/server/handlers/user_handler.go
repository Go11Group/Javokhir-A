package handlers

import (
	"net/http"
)

func (h *Handlers) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Your handler logic here
	w.Write([]byte("GetUserByID handler"))
}
