package repositories

import "database/sql"

type Repositories struct {
	db      *sql.DB
	Weather *WeatherRepository
}

func NewRepository(db *sql.DB) *Repositories {
	return &Repositories{
		db:      db,
		Weather: NewWeatherRepository(db),
	}
}
