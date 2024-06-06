package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.example.com")
	if err != nil {
		log.Fatalf("Failed to make request: %s\n", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %s\n", err)
	}
	fmt.Println(string(body))
}
