package server

import (
	"log"
	"net/http"

	"github.com/Go11Group/Javokhir-A/homework/lesson34/repositories"
	"gorm.io/gorm"
)

type Server struct {
	db           *gorm.DB
	Repositories *repositories.UniRepo
}

func NewServer(db *gorm.DB) *Server {
	return &Server{db: db}
}

func (s *Server) Start() {
	router := http.NewServeMux()
	s.Repositories = repositories.NewUniRepo(s.db)

	SetupRoutes(router, s.db)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
