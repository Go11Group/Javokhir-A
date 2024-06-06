package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Go11Group/Javokhir-A/at_lesson/lesson29Practice/models"
	"github.com/Go11Group/Javokhir-A/at_lesson/lesson29Practice/storage/postgres"
	_ "github.com/lib/pq"
)

const (
	usersJsonPath = "storage/jsons/users.json"
)

func main() {
	db, err := postgres.NewConnection()

	userRepo := postgres.NewUserRepo(db)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connected")
	}
	defer db.Close()

	uJsonFile, err := os.Open(usersJsonPath)

	if err != nil {
		log.Println(err)
	}

	users := []models.User{}

	decoder := json.NewDecoder(uJsonFile)
	err = decoder.Decode(&users)

	if err != nil {
		log.Println(err)
	}

	for _, user := range users {
		err := userRepo.Create(user)
		if err != nil {
			log.Fatal(err)
		}
	}

}
