package app

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
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
	router := gin.Default()

	SetupRoutes(router, s.db)

	if err := router.Run(addr); err != nil {
		log.Fatal("Failed to start server: " + err.Error())
		return
	}
}
