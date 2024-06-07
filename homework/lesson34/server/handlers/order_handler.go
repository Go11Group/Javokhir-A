package handlers

import (
	"net/http"
)

func (h *Handlers) GetOrderByID(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Get order by Id"))
}
