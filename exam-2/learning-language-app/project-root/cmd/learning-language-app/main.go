package main

import (
	"log"

	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/configs"
	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/storage"
)

func main() {
	dns := "user=postgres port=5432 password=1702 database=learning_language_app host=localhost sslmode=disable"
	err := storage.NewDatabase(dns)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	db := storage.Db
	ser := configs.NewServer(db)
	ser.Start(":8060")
}
