package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Go11Group/Javokhir-A/homework/lesson43/atto-app/billing-service/utils"
	_ "github.com/lib/pq"
)

func NewDbConneciton(dbName, driverName string) (*sql.DB, error) {
	dns := utils.CreateUrlForDatabseConnection(dbName)
	db, err := sql.Open(driverName, dns)
	if err != nil {
		return nil, fmt.Errorf("opening db connection failed: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	return db, nil
}
