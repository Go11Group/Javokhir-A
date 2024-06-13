package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	db *sql.DB
}

func NewServer(db *sql.DB) *Server {
	return &Server{
		db: db,
	}
}

func (s *Server) Start(addr string) {
	router := mux.NewRouter()

	SetupRoutes(router, s.db)

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal("Failed to start server:" + err.Error())
		return
	}

}
