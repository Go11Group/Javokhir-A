package handlers

import (
	"context"
	"fmt"

	wr "github.com/Go11Group/Javokhir-A/homework/lesson46_/api-gateway/genproto/grpc"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewWeatherHandler(grpcAddr string) (*Handler, error) {
	conn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	return &Handler{
		weatherClient: wr.NewWeatherServiceClient(conn),
	}, nil
}

type Handler struct {
	weatherClient wr.WeatherServiceClient // gRPC weatherClient
}

func (h *Handler) GetWeather(c *gin.Context) {
	city := c.Param("city")
	weatherRequest := &wr.WeatherRequest{Location: city}

	weatherResponse, err := h.weatherClient.GetCurrentWeather(context.Background(), weatherRequest)
	if err != nil {
		c.JSON(500, gin.H{"error": "No weather information found"})
		return
	}

	if weatherResponse.Condition != nil {
		c.JSON(200, gin.H{"weather": weatherResponse.Condition})
		return
	}

	c.JSON(500, gin.H{"error": "No weather information found"})
}

func (h *Handler) ReportWeatherCondition(c *gin.Context) {
	city := c.Param("city")
	weatherRequest := &wr.WeatherRequest{Location: city}

	weatherResponse, err := h.weatherClient.GetCurrentWeather(context.Background(), weatherRequest)
	if err != nil {
		c.JSON(500, gin.H{"error": "No weather information found"})
		return
	}

	if weatherResponse.Condition != nil {
		c.JSON(200, gin.H{"weather": weatherResponse.Condition})
		return
	}

	c.JSON(500, gin.H{"error": "No weather information found"})
}
