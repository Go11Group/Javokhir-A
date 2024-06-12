package main

import (
	"log"

	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/config"
)

func main() {
	dns := "host=localhost user=postgres database=learning_language_app password=1702 port=5432 sslmode=disable"
	db, err := config.NewDbConnection(dns)
	if err != nil {
		log.Fatal(err)
	}

	server := config.InitServer(db)
	server.Start(":8080")
}
