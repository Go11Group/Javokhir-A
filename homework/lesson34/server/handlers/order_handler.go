package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handlers) GetOrderByID(w http.ResponseWriter, r *http.Request) {

	val := r.URL.Query()
	id, err := strconv.Atoi(val.Get("id"))

	w.Header().Set("content-type", "application/json")

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	order, err := h.repos.GetOrderById(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = json.NewEncoder(w).Encode(order)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

func (h *Handlers) CreateOrder(w http.ResponseWriter, r *http.Request) {

}
func (h *Handlers) DeleteOrder(w http.ResponseWriter, r *http.Request) {

}
