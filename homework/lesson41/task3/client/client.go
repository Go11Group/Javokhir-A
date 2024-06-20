package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type UserReq struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	url := "http://127.0.0.1:8080/user"
	user := UserReq{
		Id:    1,
		Name:  "John Doe",
		Email: "johndor@gmail.com",
		Age:   46,
	}

	createUser(url, user)

	getUser(url, 1)
}

func createUser(url string, user UserReq) {

	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("failed to marshal: %v", err)
	}

	res, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer res.Body.Close()

	fmt.Println(res.Body, res.Status)

}

func getUser(url string, id int) {
	url += "?id=" + strconv.Itoa(id)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to get user: received status code %d", resp.StatusCode)
	}

	var user UserReq
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		log.Fatalf("Failed to decode user: %v", err)
	}

	fmt.Println("Get user response:", user, "Status:", resp.Status)
}

func UpdateUser(url string, id int) {
	url += "?id=" + strconv.Itoa(id)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to get user: received status code %d", resp.StatusCode)
	}

}
