package main

import (
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

var Users []UserReq

func main() {
	http.HandleFunc("/user", UserHandler)

	log.Println("Serving on 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to server: " + err.Error())
	}

}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		CreateUser(w, r)
	case http.MethodGet:
		GetUser(w, r)
	case http.MethodPut:
		UpdateUser(w, r)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var uReq UserReq
	err := json.NewDecoder(r.Body).Decode(&uReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to decode request body: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User Created!"))

	Users = append(Users, uReq)

	fmt.Println("User -> ", uReq)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user ID"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	for _, user := range Users {
		if user.Id == id {
			if err := json.NewEncoder(w).Encode(user); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Failed to encode user data: " + err.Error()))
				return
			}
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Cannot find user with this ID"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id is not valid: " + err.Error()))
		return
	}
	user := UserReq{}

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i, user := range Users {
		if id == user.Id {
			Users[i].Age = user.Age
			Users[i].Email = user.Email
			Users[i].Name = user.Name
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("user not found"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("id is not valid: " + err.Error()))
		return
	}

	for i, u := range Users {
		if u.Id == id {
			left := Users[:i]
			right := Users[i:]
			Users = nil
			Users = append(Users, left...)
			Users = append(Users, right...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("user by this id not found"))
	fmt.Println(Users)
}
