package main

import (
	"github.com/Go11Group/Javokhir-A/homework/lesson35/internal/app"
	"github.com/Go11Group/Javokhir-A/homework/lesson35/internal/db"
)

func main() {
	db.NewDBConnection()

	server := app.NewServer(db.DB)

	server.Start(":8080")
}
