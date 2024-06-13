package main

import (
	"github.com/Go11Group/Javokhir-A/homework/lesson36/internal/app"
	"github.com/Go11Group/Javokhir-A/homework/lesson36/internal/db"
)

func main() {
	db.NewDBConnection()

	server := app.NewServer(db.DB)

	server.Start(":8080")

	db.CloseDb()
}
