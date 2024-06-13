package config

import (
	"database/sql"

	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db *sql.DB
}

func InitServer(db *sql.DB) *Server {
	return &Server{
		db: db,
	}
}

func (s Server) Start(addr string) {
	router := gin.Default()

	routes.SetupRoutes(router, s.db)

	router.Run(addr)
}
