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
		Email:     "abdusamatovjavohir@gmail.com", // Corrected email
		Password:  "1234",
		Age:       22,
		Field:     "Programming",
		Gender:    "Male",
		IsEmploee: true,
	}
	user1.ID = 1
	// Uncomment to create user first before deleting
	// if err := userRepo.Create(user1); err != nil {
	// 	log.Println(err)
	// }

	if err := userRepo.DeleteUser(&user1); err != nil {
		log.Println(err)
	}
}
