package handler

import (
	"encoding/json"
	m "exam/models"
	"net/http"
	"time"
)

type AuthorTest struct {
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
	Biography string `json:"biography"`
}

func (h *handler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var newAuthor AuthorTest
	if err := json.NewDecoder(r.Body).Decode(&newAuthor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !IsDateValid(newAuthor.BirthDate) {
		http.Error(w, "Date of birth is wrong!", http.StatusBadRequest)
		return
	}

	parsedDate, err := time.Parse("2006-01-02", newAuthor.BirthDate)
	if err != nil {
		http.Error(w, "Date of birth is wrong!", http.StatusBadRequest)
		return
	}

	result := m.Author{
		Name:      newAuthor.Name,
		BirthDate: parsedDate,
		Biography: newAuthor.Biography,
	}

	id, err := h.storage.Author().CreateAuthor(r.Context(), result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Author is created!!!  ID: " + id)
}

func (h *handler) GetAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := h.storage.Author().GetAuthors(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

func (h *handler) GetAuthorById(w http.ResponseWriter, r *http.Request) {
	id := GetIDFromURL(r.URL.Path)
	if id == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	author, err := h.storage.Author().GetAuthorById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}

func (h *handler) UpdateAuthorByID(w http.ResponseWriter, r *http.Request) {
	id := GetIDFromURL(r.URL.Path)

	if id == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var newAuthor AuthorTest

	err := json.NewDecoder(r.Body).Decode(&newAuthor)
	if err != nil {
		http.Error(w, "Failed to decode body", http.StatusBadRequest)
		return
	}

	if !IsDateValid(newAuthor.BirthDate) {
		http.Error(w, "Date of birth is wrong!", http.StatusBadRequest)
		return
	}

	parsedDate, err := time.Parse("2006-01-02", newAuthor.BirthDate)
	if err != nil {
		http.Error(w, "Date of birth is wrong!", http.StatusBadRequest)
		return
	}

	result := m.Author{
		Name:      newAuthor.Name,
		BirthDate: parsedDate,
		Biography: newAuthor.Biography,
	}

	author, err := h.storage.Author().UpdateAuthorByID(r.Context(), id, result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(author); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *handler) DeleteAuthorByID(w http.ResponseWriter, r *http.Request) {
	id := GetIDFromURL(r.URL.Path)

	if id == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err := h.storage.Author().DeleteAuthorByID(r.Context(), id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Author is deleted!!!")
}

func (h *handler) GetAuthorByName(w http.ResponseWriter, r *http.Request) {
	name := GetIDFromURL(r.URL.Path)
	if name == "" {
		http.Error(w, "Invalid Name", http.StatusBadRequest)
		return
	}

	author, err := h.storage.Author().GetAuthorByName(r.Context(), name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}
