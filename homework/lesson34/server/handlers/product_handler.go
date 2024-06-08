package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Go11Group/Javokhir-A/homework/lesson34/internal/app/models"
	"github.com/Go11Group/Javokhir-A/homework/lesson34/repositories"
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
	var product models.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repos.CreateProduct(&product); err != nil {
		http.Error(w, "Failed to create product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(product); err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id, err := strconv.Atoi(values.Get("id"))
	if err != nil {
		http.Error(w, "Invalid product ID:"+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repos.DeleteProduct(id); err != nil {
		http.Error(w, "Failed while deleting: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product successfully deleted"))
}
func (h *Handlers) UpdateProductByID(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id, err := strconv.Atoi(values.Get("id"))
	if err != nil {
		http.Error(w, ":"+err.Error(), http.StatusBadRequest)
		return
	}
	product, err := h.repos.GetProductByID(uint(id))

	if err != nil {
		http.Error(w, "Invalid product ID:"+err.Error(), http.StatusBadRequest)
		return
	}
	var updatedProduct models.Product

	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		http.Error(w, "Cannot decode product json:"+err.Error(), http.StatusBadRequest)
		return
	}
	updatedProduct.ID = product.ID

	if err := h.repos.UpdateProduct(&updatedProduct); err != nil {
		http.Error(w, "Failed to update product:"+err.Error(), http.StatusNotModified)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Product successfully updated"))
}

func (h *Handlers) GetAllProductsByFilter(w http.ResponseWriter, r *http.Request) {
	filter := repositories.Filter{}
	w.Header().Set("content-type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&filter); err != nil {
		http.Error(w, "Failed to decode filter: "+err.Error(), http.StatusBadRequest)
		return
	}
	var products []models.Product

	if err := h.repos.FetchAll(&products, filter); err != nil {
		http.Error(w, "Fetching products has failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "Cannot fetch products", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
