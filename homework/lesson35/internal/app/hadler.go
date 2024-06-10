package app

import (
	"github.com/Go11Group/Javokhir-A/homework/lesson35/internal/services"
	"github.com/gorilla/mux"
)

func SetupHandlers(router *mux.Router, userSer *services.UsersService, problemSer *services.ProblemService) {

	router.HandleFunc("/user/{id}", userSer.GetUser).Methods("GET")
	router.HandleFunc("/user/", userSer.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", userSer.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", userSer.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/", userSer.GetAllUsers).Methods("GET")

	router.HandleFunc("/problem/{id}", problemSer.GetProblem).Methods("GET")
	router.HandleFunc("/problem/", problemSer.CreateProblem).Methods("POST")
	router.HandleFunc("/problem/{id}", problemSer.DeleteProblem).Methods("DELETE")
	router.HandleFunc("/problems/", problemSer.GetAllProblems).Methods("GET")

}
