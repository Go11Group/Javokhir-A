package main

import (
	"log"

	"github.com/Go11Group/Javokhir-A/at_lesson/GormPractice/models"
	"github.com/Go11Group/Javokhir-A/at_lesson/GormPractice/storage/postgres"
)

func main() {
	db, err := postgres.NewConnection()

	if err != nil {
		log.Fatal(err)
	}

	userRepo := postgres.CreateNewUserRepo(db)

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}

	user1 := models.User{
		FirstName: "Javokhir",
		LastName:  "Abdusamatov",
		Email:     "abdusamatovjavohir@gmai.com",
		Password:  "1234",
		Age:       22,
		Filed:     "Programming",
		Geneder:   "Male",
		IsEmploee: true,
	}
	user1.ID = 2
	err = userRepo.Create(user1)
	if err != nil {
		log.Println("user1", err)
	}
	err = userRepo.Delete(&user1)

	if err != nil {
		log.Println(err)
	}
}
