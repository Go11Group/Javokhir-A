package main

import (
	"log"

	server "github.com/Go11Group/Javokhir-A/homework/lesson43/atto-app/user-service/Server"
	"github.com/Go11Group/Javokhir-A/homework/lesson43/atto-app/user-service/storage/postgres"
)

func main() {
	db, err := postgres.NewDbConneciton("atto", "postgres")
	if err != nil {
		log.Fatal(err)
	}

	ser := server.NewServer(db)

	ser.Start(":8085")
}
