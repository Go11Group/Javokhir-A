package api

import (
	h "exam/api/handler"
	"exam/storage"
	"net/http"
)

type Option struct {
	Storage storage.IStorage
}

func New(option Option) *http.ServeMux {
	router := http.NewServeMux()

	handler := h.New(&h.HandlerConfig{
		Storage: option.Storage,
	})

	router.HandleFunc("POST /author", handler.CreateAuthor)
	router.HandleFunc("POST /book", handler.CreateBook)

	router.HandleFunc("GET /author", handler.GetAuthors)
	router.HandleFunc("GET /book", handler.GetBooks)

	router.HandleFunc("GET /author/{id}", handler.GetAuthorById)
	router.HandleFunc("GET /book/{id}", handler.GetBookById)

	router.HandleFunc("GET /authorName/{name}", handler.GetAuthorByName)
	router.HandleFunc("GET /bookTitle/{title}", handler.GetBookByTitle)

	router.HandleFunc("PUT /author/{id}", handler.UpdateAuthorByID)
	router.HandleFunc("PUT /book/{id}", handler.UpdateBookById)

	router.HandleFunc("DELETE /author/{id}", handler.DeleteAuthorByID)
	router.HandleFunc("DELETE /book/{id}", handler.DeleteBookByID)

	return router
}






// postman     8081      8080