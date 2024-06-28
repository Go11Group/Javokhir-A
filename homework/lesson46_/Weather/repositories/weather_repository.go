package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type WeatherRepository struct {
	Db *sql.DB
}

func NewWeatherRepository(db *sql.DB) *WeatherRepository {
	return &WeatherRepository{
		Db: db,
	}
}

func (w *WeatherRepository) ReportWeatherCondition(ctx context.Context, wCon *WeatherCondition) error {
	newID := uuid.NewString()

	query := `
		INSERT INTO weather_report (
			id,
		    location,
            temperature,
            humidity,
            pressure,
            wind_speed,
			reported_date
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	stmt, err := w.Db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("preparing query failed: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		newID,
		wCon.Location,
		wCon.Temperature,
		wCon.Humidity,
		wCon.Pressure,
		wCon.WindSpeed,
		time.Now(),
	)
	if err != nil {
		return fmt.Errorf("executing query failed: %v", err)
	}
	return nil
}

func (w *WeatherRepository) GetWeather(ctx context.Context, location string) (*WeatherCondition, error) {
	wCon := WeatherCondition{Location: location}

	query := `
		SELECT 
            temperature,
            humidity,
            pressure,
            wind_speed,
			reported_date
		FROM weather_report
		WHERE location = $1
	`

	err := w.Db.QueryRowContext(ctx, query, location).Scan(
		&wCon.Temperature,
		&wCon.Humidity,
		&wCon.Pressure,
		&wCon.WindSpeed,
		&wCon.ReportedDate,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no weather data found for location: %v", location)
		}
		return nil, fmt.Errorf("querying weather data failed: %v", err)
	}
	return &wCon, nil
}

func (w *WeatherRepository) GetWeatherForecast(ctx context.Context, location string) ([]*WeatherCondition, error) {
	query := `
		SELECT 
            location,
            temperature,
            humidity,
            pressure,
            wind_speed,
			reported_date
		FROM weather_report
	`

	rows, err := w.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("querying weather forecast failed: %v", err)
	}
	defer rows.Close()

	var forecasts []*WeatherCondition
	for rows.Next() {
		var wCon WeatherCondition
		if err := rows.Scan(
			&wCon.Location,
			&wCon.Temperature,
			&wCon.Humidity,
			&wCon.Pressure,
			&wCon.WindSpeed,
			&wCon.ReportedDate,
		); err != nil {
			return nil, fmt.Errorf("scanning row failed: %v", err)
		}
		forecasts = append(forecasts, &wCon)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration failed: %v", err)
	}

	return forecasts, nil
}

type WeatherCondition struct {
	Location     string    `json:"location"`
	Temperature  float64   `json:"temperature"`
	Humidity     float64   `json:"humidity"`
	Pressure     float64   `json:"pressure"`
	WindSpeed    float64   `json:"wind_speed"`
	ReportedDate time.Time `json:"reported_date"`
}
