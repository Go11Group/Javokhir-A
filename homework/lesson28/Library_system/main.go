package main

import (
	"fmt"
	"librarySystem/storage/postgres"
	"log"
)

func main() {
	db, err := postgres.NewConnetion()
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Connected Successfully")
	}
	bookRepo := postgres.NewBookRepo(db)

	// books, err := bookRepo.GetAllBooks()
	// if err != nil {
	// 	log.Println(err)
	// }

	// fmt.Println(books)
	// newBook := models.Book{}
	// newBook.Id = uuid.NewString()
	// newBook.Title = "NewTitle"
	// newBook.Author = "NewBook"
	// newBook.Genre = "Fiction"
	// newBook.PublishedYear = 2022
	// err = bookRepo.Create(newBook)
	// if err != nil {
	// 	log.Println(err)
	// }
	// err = bookRepo.Delete(models.Book{Id: "74a46f19-3570-4837-9e9e-f60bcfca4329"})
	// if err != nil {
	// 	log.Println(err)
	// } else {
	// 	fmt.Println("deleted")
	// }
	fmt.Println(bookRepo.GetById("bf5b085d-7d66-47cc-a607-f5e4a19da759"))
}
