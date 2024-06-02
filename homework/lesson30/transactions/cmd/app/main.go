package main

import (
	"fmt"
	"log"

	"github.com/Go11Group/Javokhir-A/homework/lesson30/transactions/db/postgres"
	"github.com/Go11Group/Javokhir-A/homework/lesson30/transactions/internal/app/models"
	"github.com/Go11Group/Javokhir-A/homework/lesson30/transactions/repositories"
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
	defer postgres.CloseDb()

	db := postgres.DB
	// err = db.AutoMigrate(&models.Product{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = db.Autd

	// users := []*models.User{}
	// products := []*models.Product{}

	// if users, err = test.ReadingUsersJson(); err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%v", users[0])

	// userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)

	// fmt.Println("connected", userRepo, productRepo)

	// for _, user := range users {
	// 	fmt.Printf("%v\n", user)
	// }

	// for _, user := range users {
	// 	if err := userRepo.CreateUser(user); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// if products, err = test.ReadingProductJson(); err != nil {
	// 	log.Fatal(err)
	// }

	// for _, product := range products {
	// 	if err := productRepo.CreateProduct(product); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// if err := userRepo.DeleterUser(5); err != nil {
	// log.Fatal(err)
	// }
	// if err := productRepo.DeleteProduct(8); err != nil {
	// 	log.Fatal(err)
	// }

	product := &models.Product{}
	product.ID = 10
	product.Price = 522.2
	product.Description = "Testing update"

	// if err := productRepo.UpdateProduct(product); err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println(productRepo.GetProductByID(product.ID))
	// fmt.Println(userRepo.GetUserByID(7))
}
