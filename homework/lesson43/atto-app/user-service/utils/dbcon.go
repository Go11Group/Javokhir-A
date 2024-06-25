package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Go11Group/Javokhir-A/homework/lesson43/atto-app/user-service/config"
	_ "github.com/lib/pq"
)

func CreateUrlForDatabseConnection(dbname string) string {
	dbConfig := config.DatabaseConfig{}
	dbConfig.DBname = dbname

	file, err := os.Open("/home/javokhir/go/src/github.com/Go11Group/Javokhir-A/homework/lesson43/atto-app/user-service/resources/db_config.json")
	if err != nil {
		log.Fatal("opening json file failed: " + err.Error())
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&dbConfig); err != nil {
		log.Fatal("Decoding database file failed: " + err.Error())
	}

	dns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbname)

	db, err := sql.Open("postgres", dns)
	if err != nil {
		log.Fatal("opening postgres db failed: " + err.Error())
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("failed to connect to database: " + err.Error())
	}
	return dns

}

func WriteDnsToMakeFile(path, url string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("opening Makefile failed: " + err.Error())
	}
	defer file.Close()

	query := "\nDATABASE_URL := " + url + "\n"

	_, err = file.WriteString(query)
	if err != nil {
		log.Fatal("writing to Makefile failed: " + err.Error())
	}
}
