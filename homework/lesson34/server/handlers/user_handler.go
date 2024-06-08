package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Go11Group/Javokhir-A/homework/lesson34/internal/app/models"
	"github.com/Go11Group/Javokhir-A/homework/lesson34/repositories"
)

func (h *Handlers) GetUserByID(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query()
	id, err := strconv.Atoi(val.Get("id"))

	w.Header().Set("content-type", "application/json")

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	user, err := h.repos.GetUserByID(uint(id))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

}

func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repos.CreateUser(&user); err != nil {
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id, err := strconv.Atoi(values.Get("id"))
	if err != nil {
		http.Error(w, "Invalid user ID:"+err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repos.DeleterUser(id); err != nil {
		http.Error(w, "Failed while deteling: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User successfully deleted"))

}

func (h *Handlers) UpdateUserById(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	id, err := strconv.Atoi(values.Get("id"))
	if err != nil {
		http.Error(w, ":"+err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.repos.GetUserByID(uint(id))

	if err != nil {
		http.Error(w, "Invalid user ID:"+err.Error(), http.StatusBadRequest)
		return
	}
	var updatedUser models.User

	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Cannot decode user json:"+err.Error(), http.StatusBadRequest)
		return
	}
	updatedUser.ID = user.ID

	if err := h.repos.UpdateUser(&updatedUser); err != nil {
		http.Error(w, "Failed to update user:"+err.Error(), http.StatusNotModified)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User successfully updated"))
}

func (h *Handlers) GetAllUsersByFilter(w http.ResponseWriter, r *http.Request) {
	filter := repositories.Filter{}
	w.Header().Set("content-type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&filter); err != nil {
		http.Error(w, "Failed to decode filter: "+err.Error(), http.StatusBadRequest)
		return
	}
	var users []models.User

	if err := h.repos.FetchAll(&users, filter); err != nil {
		http.Error(w, "Fetching users has failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Cannot fetch users", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
