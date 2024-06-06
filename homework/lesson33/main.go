package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Go11Group/Javokhir-A/homework/lesson33/repositories"
	"github.com/Go11Group/Javokhir-A/homework/lesson33/storage/postgres"
)

var (
	db = postgres.DB
)
var userRepo repositories.UserRepository

func main() {
	dsn := "host=localhost port=5432 user=postgres dbname=testing sslmode=disable password=1702"

	err := postgres.ConnectDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer postgres.CloseDB()

	db = postgres.DB

	userRepo = repositories.NewUserRepository(db)

	user, err := userRepo.GetUserByID(10)
	if err != nil {
		log.Printf("Failed to get user: %v", err)
		return
	}

	fmt.Println(user)

	// APIs
	// ``
	mux := http.NewServeMux()
	mux.HandleFunc("GET /user/{id}", UserFuncHandler) //first API to fetch user data by its id
	mux.HandleFunc("DELETE /user/{id}", DeleteUser)
	err = http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

func UserFuncHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")
	// w.Header().
	userId, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Fatal(err)
	}

	user, err := userRepo.GetUserByID(userId)
	if err != nil {
		fmt.Fprintf(w, "User by id %d not found", userId)
	}

	userInfo, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(userInfo)
	if err != nil {
		log.Fatal(err)
	}

}
