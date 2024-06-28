package service

import (
	"context"
	"database/sql"
	"fmt"

	pb "github.com/Go11Group/Javokhir-A/homework/lesson46_/Weather/genproto/weather"
	rp "github.com/Go11Group/Javokhir-A/homework/lesson46_/Weather/repositories"
)

type WeatherService interface {
	GetCurrentWeather(ctx context.Context, in *pb.WeatherRequest) (*pb.WeatherResponse, error)
	GetWeatherForecast(ctx context.Context, in *pb.ForecastRequest) (*pb.ForecastResponse, error)
	ReportWeatherCondition(ctx context.Context, in *pb.WeatherCondition) (*pb.WeatherReportResponse, error)
}

type weatherService struct {
	pb.UnimplementedWeatherServiceServer
	Repositories *rp.Repositories
}

func NewWeatherService(db *sql.DB) *weatherService {
	return &weatherService{
		Repositories: rp.NewRepository(db),
	}
}

func (w *weatherService) GetCurrentWeather(ctx context.Context, in *pb.WeatherRequest) (*pb.WeatherResponse, error) {

	wCon, err := w.Repositories.Weather.GetWeather(ctx, in.Location)
	if err != nil {
		return nil,
			fmt.Errorf("could not get current weather: %v", err)
	}

	con := &pb.WeatherCondition{}

	con.Humidity = float32(wCon.Humidity)
	con.Pressure = float32(wCon.Pressure)
	con.Temperature = float32(wCon.Temperature)
	con.WindSpeed = float32(wCon.WindSpeed)
	con.Location = in.Location

	return &pb.WeatherResponse{Condition: con}, nil
}

func (w *weatherService) GetWeatherForecast(ctx context.Context, in *pb.ForecastRequest) (*pb.ForecastResponse, error) {
	// Placeholder implementation
	return &pb.ForecastResponse{
		Forecast: []*pb.WeatherCondition{
			{
				Temperature: 26.0,
				Humidity:    55,
				WindSpeed:   12.0,
			},
			{
				Temperature: 24.0,
				Humidity:    65,
				WindSpeed:   8.0,
			},
		},
	}, nil
}

func (w *weatherService) ReportWeatherCondition(ctx context.Context, in *pb.WeatherCondition) (*pb.WeatherReportResponse, error) {
	wCon := &rp.WeatherCondition{
		Location:    in.Location,
		Temperature: float64(in.Temperature),
		Humidity:    float64(in.Humidity),
		Pressure:    float64(in.Pressure),
		WindSpeed:   float64(in.WindSpeed),
	}

	err := w.Repositories.Weather.ReportWeatherCondition(ctx, wCon)
	if err != nil {
		return &pb.WeatherReportResponse{
			Status: "could not report weather condition",
		}, fmt.Errorf("could not report weather condition: %v", err)
	}

	return &pb.WeatherReportResponse{
		Status: "Condition reported successfully",
	}, nil
}
