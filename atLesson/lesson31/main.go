package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bxcodec/faker"
	_ "github.com/lib/pq"
)

type User struct {
	Id                string
	UserName          string
	Email             string
	Password          string
	FirstName         string
	LastName          string
	PhoneNumber       string
	Address           string
	ProfilePictureUrl string
	Roles             string
	LastLogin         string
	CreatedAt         string
	UpdatedAt         string
	Active            bool
}

func main() {
	fmt.Println()
	dns := "host=localhost port=5432 password=1702 user=postgres databse=testing ssmolde=disable"
	db, err := sql.Open("postgres", dns)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 100; i++ {
		query := `INSERT INTO users(
			username,
			email,
			password,
			first_name,
			last_name,
			phone_number,
			address,
			profile_picture_url,
			roles,
			last_login,
			created_at,
			updated_at,
			active
		) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		`
		db.Exec(query, "")
		if _, err := db.Exec(query, faker.UserName); err != nil {
			fmt.Println(err)
		}
	}

}
