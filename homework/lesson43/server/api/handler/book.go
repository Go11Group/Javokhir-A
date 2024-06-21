package handler

import (
	"encoding/json"
	m "exam/models"
	"net/http"
	"time"
)

type BookTest struct {
	AuthorName      string `json:"author_name"`
	Title           string `json:"title"`
	PublicationDate string `json:"publication_date"`
	ISBN            string `json:"isbn"`
	Description     string `json:"description"`
}

func (h *handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook BookTest
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !IsDateValid(newBook.PublicationDate) {
		http.Error(w, "Pulication date is wrong!", http.StatusBadRequest)
		return
	}

	parsedDate, err := time.Parse("2006-01-02", newBook.PublicationDate)
	if err != nil {
		http.Error(w, "Pulication date is wrong!", http.StatusBadRequest)
		return
	}

	result := m.Book{
		AuthorName:      newBook.AuthorName,
		Title:           newBook.Title,
		PublicationDate: parsedDate,
		ISBN:            newBook.ISBN,
		Description:     newBook.Description,
	}

	book, name, err := h.storage.Book().CreateBook(r.Context(), result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	book.AuthorName = name

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (h *handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.storage.Book().GetBooks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *handler) GetBookById(w http.ResponseWriter, r *http.Request) {
	id := GetIDFromURL(r.URL.Path)
	if id == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	book, err := h.storage.Book().GetBookById(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (h *handler) UpdateBookById(w http.ResponseWriter, r *http.Request) {
	id := GetIDFromURL(r.URL.Path)
	if id == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedBook BookTest

	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, "No book found with the given ID", http.StatusBadRequest)
		return
	}

	parsedDate, err := time.Parse("2006-01-02", updatedBook.PublicationDate)
	if err != nil {
		http.Error(w, "Publication date is wrong!", http.StatusBadRequest)
		return
	}

	book := m.Book{
		Title:           updatedBook.Title,
		AuthorName:      updatedBook.AuthorName,
		PublicationDate: parsedDate,
		ISBN:            updatedBook.ISBN,
		Description:     updatedBook.Description,
	}

	updatedBookResult, name, err := h.storage.Book().UpdateBookById(r.Context(), id, book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updatedBookResult.AuthorName = name

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(updatedBookResult); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *handler) DeleteBookByID(w http.ResponseWriter, r *http.Request) {
	id := GetIDFromURL(r.URL.Path)
	if id == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err := h.storage.Book().DeleteBookByID(r.Context(), id)
	if err != nil {
		http.Error(w, "No book found with the given ID", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Book is deleted!!! ID: " + id)
}

func (h *handler) GetBookByTitle(w http.ResponseWriter, r *http.Request) {
	title := GetIDFromURL(r.URL.Path)
	if title == "" {
		http.Error(w, "Invalid Title", http.StatusBadRequest)
		return
	}

	book, err := h.storage.Book().GetBookByTitle(r.Context(), title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
