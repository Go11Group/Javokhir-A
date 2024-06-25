package server

import (
	"database/sql"
	"log"

	"github.com/Go11Group/Javokhir-A/homework/lesson43/atto-app/user-service/handlers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	handler *handlers.Handler
}

func NewServer(db *sql.DB) *Server {
	return &Server{handler: handlers.NewHandler(db)}
}

func (s *Server) Start(addr string) {
	r := gin.Default()
	s.SetupRoutes(r)

	if addr == "" {
		addr = ":8080"
	}

	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}

}
