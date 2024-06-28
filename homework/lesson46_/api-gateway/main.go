// api/handlers/handler.go

package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	gw "github.com/Go11Group/Javokhir-A/homework/lesson46_/api-gateway/genproto/grpc" // Import generated protobuf package
	"google.golang.org/grpc"
)

type WeatherHandler struct {
	client gw.WeatherServiceClient // gRPC client
}

func NewWeatherHandler(grpcAddr string) (*WeatherHandler, error) {
	conn, err := grpc.NewClient(grpcAddr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	return &WeatherHandler{
		client: gw.NewWeatherServiceClient(conn),
	}, nil
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	// Parse request parameters
	city := r.URL.Query().Get("city")

	// Call gRPC service
	resp, err := h.client.GetWeather(context.Background(), &gw.WeatherRequest{City: city})
	if err != nil {
		log.Printf("gRPC error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Return response as JSON
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Printf("JSON marshal error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
