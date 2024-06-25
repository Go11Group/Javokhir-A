package main

import (
	"log"

	"github.com/Go11Group/Javokhir-A/homework/lesson43/atto-app/api-gateway/server"
)

func main() {
	ser := server.NewServer(":8083")
	err := ser.Start()
	if err != nil {
		log.Println(err)
	}
}
