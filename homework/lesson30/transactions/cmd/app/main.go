package main

import (
	"fmt"
	"log"

	"github.com/Go11Group/Javokhir-A/homework/lesson30/transactions/db/postgres"
	"github.com/Go11Group/Javokhir-A/homework/lesson30/transactions/repositories"
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
	uniRepo := repositories.NewUniverseRepository(db)
	userRepo := repositories.NewUserRepository(db)

	updateUser, err := userRepo.GetUserByID(1005)
	if err != nil {
		log.Fatal(err)
	}

	updateUser.UserName = "java"
	if err := uniRepo.Update(updateUser); err != nil {
		log.Fatal(err)
	}

	// newOrder := &models.Order{
	// 	UserId:    75,
	// 	ProductId: 65,
	// }

	// if err := uniRepo.Create(&newOrder); err != nil {
	// 	log.Fatal(err)
	// }

	// newOrder.ID = 3
	// if err := uniRepo.Update(newOrder); err != nil {
	// 	log.Fatal(err)
	// }

	// var users []models.User
	// if err := uniRepo.FetchAll(&users); err != nil {
	// 	log.Fatal(err)
	// }

	// for _, user := range users {
	// 	fmt.Println(user)
	// }

	// var products []models.Product
	// if err := uniRepo.FetchAll(&products); err != nil {
	// 	log.Fatal(err)
	// }

	// for _, product := range products {
	// 	fmt.Println(product)
	// }

}
