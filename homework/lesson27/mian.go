package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	user   = "posgres"
	dbname = "univercity"
	pass   = "1702"
	host   = "localhost"
	port   = 5432
)

func main() {
	connectionString := fmt.Sprintf("host=%s port=%d database=%s user=%s pass=%s sslmode=disable",
		host, port, dbname, user, pass)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	query := ``
	db.Query()
}
