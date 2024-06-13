package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func NewDatabase(dns string) error {

	db, err := sql.Open("postgres", dns)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}

	log.Println("Pong!")

	Db = db
	return nil
}
