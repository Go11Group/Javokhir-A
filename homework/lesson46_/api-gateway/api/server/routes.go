package server

import "github.com/gin-gonic/gin"

func (s *Server) SetupRouter(router *gin.Engine) {

	w := router.Group("/weather")
	{
		w.GET("/:city", s.Handlers.GetWeather)
		w.POST("/", s.Handlers.ReportWeatherCondition)
		w.GET("/weather")
	}

}
