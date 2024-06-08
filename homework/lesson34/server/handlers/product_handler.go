package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handlers) GetProductByID(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query()
	id, err := strconv.Atoi(val.Get("id"))

	w.Header().Set("content-type", "application/json")

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	prodcut, err := h.repos.GetProductByID(uint(id))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = json.NewEncoder(w).Encode(prodcut)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
}

func (h *Handlers) CreateProduct(w http.ResponseWriter, r *http.Request) {

}

func (h *Handlers) DeleteProduct(w http.ResponseWriter, r *http.Request) {

}
