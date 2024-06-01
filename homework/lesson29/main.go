package main

import (
	"fmt"
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
	fmt.Println(user1)

	// Uncomment to create user first before deleting
	if err := userRepo.Create(user1); err != nil {
		log.Println(err)
	}

	if users, err := userRepo.GetAllUsers(); err != nil {
		log.Println(err)
	} else {
		fmt.Println(users)
	}

	// if err := userRepo.Delete(&user1); err != nil {
	// 	log.Println(err)
	// }
}
