package app

import (
	"github.com/Go11Group/Javokhir-A/homework/lesson35/internal/services"
	"github.com/gorilla/mux"
)

func SetupHandlers(router *mux.Router, userSer *services.UsersService) {

	router.HandleFunc("/user/{id}", userSer.GetUser).Methods("GET")
	// userGet.HandleFunc("create")
	router.HandleFunc("/user/", userSer.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", userSer.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", userSer.DeleteUser).Methods("DELETE")
	// userCreate := router.HandleFunc("/user/create", userSer.CreateUser).Methods("POST")

}
