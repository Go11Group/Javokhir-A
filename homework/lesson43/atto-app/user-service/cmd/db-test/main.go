package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Go11Group/Javokhir-A/homework/lesson43/atto-app/user-service/repository"
	"github.com/Go11Group/Javokhir-A/homework/lesson43/atto-app/user-service/service"
	"github.com/Go11Group/Javokhir-A/homework/lesson43/atto-app/user-service/utils"
	_ "github.com/lib/pq"
)

func main() {
	dns := utils.CreateUrlForDatabseConnection("atto")
	db, err := sql.Open("postgres", dns)
	if err != nil {
		log.Fatal(err)
	}
	user := repository.CreateUser{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
		Phone:     "+1234567890", // Replace with a valid phone number if needed
		Age:       25,
		Password:  "secretpassword", // This is a plain text password, replace with a strong password for actual use
	}

	userSer := service.NewUserService(db)
	id, err := userSer.Create(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}
