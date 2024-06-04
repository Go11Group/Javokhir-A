package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/jaswdr/faker"
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
	fmt.Println("Starting data generation...")

	dsn := "host=localhost port=5432 user=postgres password=1702 dbname=testing sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rand.Seed(time.Now().UnixNano()) // Set the random seed

	faker := faker.New()
	for i := 0; i <= 1_000_000; i++ {
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
		) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

		person := faker.Person()
		address := faker.Address()

		lastLogin := faker.Time().Time(time.Now().Add(-time.Duration(faker.IntBetween(1, 30000)) * time.Hour))
		createdAt := time.Now()
		updatedAt := time.Now()

		if _, err := db.Exec(query,
			faker.Internet().User(),
			faker.Internet().Email(),
			faker.Internet().Password(),
			person.FirstName(),
			person.LastName(),
			person.Contact().Phone,
			address.Address(),
			faker.Internet().URL(),
			person.Title(),
			lastLogin,
			createdAt,
			updatedAt,
			true,
		); err != nil {
			log.Printf("Error inserting data: %v", err)
		} else if i%100000 == 0 {
			fmt.Println(i)
		}

	}

	fmt.Println("Data generation completed.")
}
