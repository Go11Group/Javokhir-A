package server

import (
	"github.com/Go11Group/Javokhir-A/homework/lesson43/atto-app/api-gateway/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{addr: addr}
}

func (s *Server) Start() error {
	router := gin.Default()

	routes.SetupRoutes(router)

	return router.Run(s.addr)
}
