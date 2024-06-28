package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Go11Group/Javokhir-A/homework/lesson46_/Weather/config"
	_ "github.com/lib/pq"
)

func NewDatabaseConnection() *sql.DB {
	dbConfig := config.Configs{}.Database

	file, err := os.Open("resources/database_config.json")

	if err != nil {
		log.Fatal(err)
	}

	err = json.NewDecoder(file).Decode(&dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	dns := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Name)
	db, err := sql.Open("postgres", dns)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
