package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectDB(dsn string) error {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed while connecting %w", err)
	}

	DB = db
	DB.DB()
	return nil
}

func CloseDB() {
	if DB != nil {
		dbPostgers, err := DB.DB()
		if err != nil {
			fmt.Printf("Error getting uderlying database connection: %v\n", err)
			return
		}

		if err := dbPostgers.Close(); err != nil {
			fmt.Printf("Erorr closing database connection: %v\n", err)
			return
		}
	}
}
