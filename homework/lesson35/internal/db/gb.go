package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	host     = "localhost"
	port     = 5432
	password = "1702"
	user     = "postgres"
	dbname   = "leetcode"
)

func NewDBConnection() {
	dns := fmt.Sprintf("host=%s port=%d user=%s database=%s password=%s sslmode=disable",
		host, port, user, dbname, password)

	db, err := sql.Open("postgres", dns)

	if err != nil {
		log.Println("Couldn't open database connection:" + err.Error())
	}
	if err := db.Ping(); err != nil {
		log.Println("Connection failed:" + err.Error())
	}

	DB = db

	fmt.Println("Connection is set.")
}

func CloseDb() {
	if err := DB.Close(); err != nil {
		log.Fatal(err)
	}
}
