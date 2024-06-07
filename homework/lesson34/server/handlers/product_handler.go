package handlers

import "net/http"

func (h *Handlers) GetProductByID(w http.ResponseWriter, r *http.Request) {
	// Your handler logic here
	w.Write([]byte("GetProductByID handler"))
}
