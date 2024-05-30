package main

import (
	"librarySystem/storage/postgres"
	"log"
)

func main() {
	db, err := postgres.NewConnetion()
	if err != nil {
		log.Println(err)
	}
	
}
