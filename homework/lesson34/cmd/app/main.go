package main

import (
	"fmt"
	"log"

	"github.com/Go11Group/Javokhir-A/homework/lesson34/db/postgres"
	"github.com/Go11Group/Javokhir-A/homework/lesson34/server"
)

const (
	user     = "postgres"
	port     = 5432
	host     = "localhost"
	password = "1702"
	dbName   = "shopping"
)

var dns = fmt.Sprintf("host=%s port=%d password=%s dbname=%s user=%s sslmode=disable", host, port, password, dbName, user)

func main() {
	err := postgres.ConnectionDB(dns)
	if err != nil {
		log.Fatal(err)
	}
	defer postgres.CloseDb()

	db := postgres.DB
	// uniRepo := repositories.NewUniverseRepository(db)
	// userRepo := repositories.NewUserRepository(db)

	ser := server.NewServer(db)

	ser.Start()

}
