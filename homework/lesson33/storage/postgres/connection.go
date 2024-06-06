package postgress

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dns = "host=localhost port=5432 user=postgres database=testing sslmode=disable password=1702"
)

func NewDBConnection() (*gorm.DB, error) {

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed while connecting to database: %v", err)
	}

	return db, nil
}
