package main

import (
	"fmt"
	"log"

	"github.com/Go11Group/Javokhir-A/homework/lesson30/transactions/db/postgres"
	"github.com/Go11Group/Javokhir-A/homework/lesson30/transactions/internal/app/models"
	"github.com/Go11Group/Javokhir-A/homework/lesson30/transactions/repositories"
	"github.com/Go11Group/Javokhir-A/homework/lesson30/transactions/test"
)

const (
	user     = "postgres"
	port     = 5432
	host     = "localhost"
	password = "1702"
	dbName   = "shopping"
)

var dns = fmt.Sprintf("host=%s port=%d password=%s database=%s user=%s sslmode=disable", host, port, password, dbName, user)

func main() {
	err := postgres.ConnectionDB(dns)
	if err != nil {
		log.Fatal(err)
	}

	db := postgres.DB
	// err = db.AutoMigrate(&models.Product{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = db.AutoMigrate(&models.User{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	users := []*models.User{}

	if users, err = test.ReadingUsersJson(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", users[0])

	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)
	fmt.Println("connected", userRepo, productRepo)

}
