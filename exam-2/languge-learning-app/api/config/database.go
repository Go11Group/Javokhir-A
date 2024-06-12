package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDbConnection(dns string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, fmt.Errorf("opening db failed: " + err.Error())
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("connecting to database failed: " + err.Error())
	}

	fmt.Println("Pong")

	return db, nil
}
