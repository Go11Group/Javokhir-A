package app

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

// Server struct holds the database connection.
type Server struct {
	db *sql.DB
}

// NewServer is a constructor function that returns a new Server instance with the provided database connection.
func NewServer(db *sql.DB) *Server {
	return &Server{
		db: db,
	}
}

// Start method starts the HTTP server with the given address.
func (s *Server) Start(addr string) {
	// Create a new router using gin.
	router := gin.Default()

	// Setup the routes for the application.
	SetupRoutes(router, s.db)

	// Start the HTTP server.
	if err := router.Run(addr); err != nil {
		log.Fatal("Failed to start server: " + err.Error())
		return
	}
}
