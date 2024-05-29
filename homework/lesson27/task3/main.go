package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	user   = "postgres"
	dbname = "univercity"
	pass   = "1702"
	host   = "localhost"
	port   = 5432
)

type Student struct {
	Id   string
	Name string
	Age  int
}

func main() {
	connectionString := fmt.Sprintf("host=%s port=%d database=%s user=%s password=%s sslmode=disable",
		host, port, dbname, user, pass)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	query := `
		SELECT *FROM STUDENTS 
	`
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}

	students := []Student{}

	for rows.Next() {
		stud := Student{}

		err := rows.Scan(&stud.Id, &stud.Name, &stud.Age)
		if err != nil {
			panic(err)
		}

		students = append(students, stud)
	}

	fmt.Println(students)
}
