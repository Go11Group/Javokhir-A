package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Database connection parameters
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1702"
	dbname   = "lesson23"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error verifying connection: %q", err)
	}
	fmt.Println("Successfully connected to the database!")

	createTableSQL := `
    CREATE TABLE IF NOT EXISTS cars (
        id SERIAL PRIMARY KEY,
        brand VARCHAR(50),
        model VARCHAR(50),
        created_year INT,
        price INT
    );`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Error creating table: %q", err)
	}
	fmt.Println("Table created successfully!")

	// Insert data
	insertSQL := `
    INSERT INTO cars (brand, model, created_year, price) VALUES
    ('Toyota', 'Camry', 2010, 15000),
    ('Honda', 'Accord', 2012, 16000),
    ('Ford', 'Mustang', 2015, 25000),
    ('Chevrolet', 'Malibu', 2013, 14000),
    ('Nissan', 'Altima', 2014, 14500),
    ('BMW', '3 Series', 2018, 30000),
    ('Audi', 'A4', 2017, 28000),
    ('Mercedes-Benz', 'C-Class', 2019, 35000),
    ('Hyundai', 'Elantra', 2016, 13000),
    ('Kia', 'Optima', 2020, 20000);
    `
	_, err = db.Exec(insertSQL)
	if err != nil {
		log.Fatalf("Error inserting data: %q", err)
	}
	fmt.Println("Data inserted successfully!")

	// Query data
	querySQL := `SELECT id, brand, model, created_year, price FROM cars;`
	rows, err := db.Query(querySQL)
	if err != nil {
		log.Fatalf("Error querying data: %q", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var brand, model string
		var createdYear, price int

		err := rows.Scan(&id, &brand, &model, &createdYear, &price)
		if err != nil {
			log.Fatalf("Error scanning row: %q", err)
		}
		fmt.Printf("ID: %d, Brand: %s, Model: %s, Created Year: %d, Price: %d\n", id, brand, model, createdYear, price)
	}

	// Check for errors during iteration
	err = rows.Err()
	if err != nil {
		log.Fatalf("Error during iteration: %q", err)
	}
}
