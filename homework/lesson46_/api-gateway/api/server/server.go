package server

import (
	"log"

	"github.com/Go11Group/Javokhir-A/homework/lesson46_/api-gateway/api/handlers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Handlers handlers.Handler
}

func NewServer(addr string) *Server {
	wHandler, err := handlers.NewWeatherHandler(addr)
	if err != nil {
		log.Fatal(err)
	}

	return &Server{
		Handlers: *wHandler,
	}
}

func (s *Server) StartServer() {
	router := gin.Default()

	s.SetupRouter(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
