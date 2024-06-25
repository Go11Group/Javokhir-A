package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Go11Group/Javokhir-A/homework/lesson45/grpc/translator"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := translator.NewTranslatorServiceClient(conn)
	for {

		var word string

		fmt.Print("Enter an uzbek word: ")
		fmt.Scan(&word)

		req := &translator.TranslatingRequest{
			UzWord: word,
		}

		if err != nil {
			log.Fatalf("Failed to call Translator: %v", err)
		}

		resp, err := client.Translator(context.Background(), req)
		if err != nil {
			log.Fatalf("Failed to call Translator: %v", err)
		}

		log.Printf("Received: %v", resp)
	}
}
