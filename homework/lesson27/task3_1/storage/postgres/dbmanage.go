package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DbConnection struct {
	*sql.DB
}

const (
	user     = "postgres"
	database = "library_system"
	password = "1702"
	port     = 5432
	host     = "localhost"
)

func NewConnetion() (*sql.DB, error) {
	driverName := "postgres"
	dataSourceName := fmt.Sprintf("user=%s host=%s port=%d password=%s databse=%s",
		user, host, port, password, database)
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	err = db.Ping()

	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return db, nil
}
