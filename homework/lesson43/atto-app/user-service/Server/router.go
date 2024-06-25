package server

import (
	"github.com/gin-gonic/gin"
)

func (s *Server) SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		users := api.Group("/user")
		{
			users.POST("/", s.handler.CreateUser)
			users.GET("/:id", s.handler.GetUserByID)
		}
	}
}
