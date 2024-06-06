package postgres

import (
	"database/sql"
	"fmt"
)

const (
	user     = "postgres"
	host     = "localhost"
	port     = 5432
	dbName   = "lesson29practice"
	password = "1702"
)

func NewConnection() (*sql.DB, error) {
	conStr := fmt.Sprintf("user=%s host=%s database=%s port=%d password=%s sslmode=disable",
		user, host, dbName, port, password)

	db, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, fmt.Errorf("failed while connecting %w", err)
	}

	return db, nil
}
