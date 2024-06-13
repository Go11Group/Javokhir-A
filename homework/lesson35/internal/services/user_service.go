package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Go11Group/Javokhir-A/homework/lesson35/internal/models"
	"github.com/Go11Group/Javokhir-A/homework/lesson35/internal/repositories"
	"github.com/gorilla/mux"
)

type UsersService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UsersService {
	return &UsersService{
		userRepo: userRepo,
	}
}

func (u *UsersService) CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := models.User{}
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := u.userRepo.CreateUser(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (u *UsersService) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var filter repositories.UserFilter

	w.Header().Set("content-type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&filter); err != nil {
		log.Println("Failed while validating user filter:" + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users, err := u.userRepo.GetAllUsers(filter)
	if err != nil {
		log.Fatal("getting all users by filter failed: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Println("Failed while transfering date into respons: " + err.Error())
		http.Error(w, err.Error(), http.StatusNoContent)
		return
	}
}

func (u *UsersService) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	w.Header().Set("content-type", "application/json")

	user, err := u.userRepo.GetUser(userId)
	if err != nil {
		log.Println("Getting user form database failed:" + err.Error())
		http.Error(w, "User not found"+err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(&user); err != nil {
		log.Println("Writing to response failed:" + err.Error())
		http.Error(w, "Writing to response body failed:"+err.Error(), http.StatusBadRequest)
	}

}

func (u *UsersService) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updateFilter repositories.UpdateUser

	if err := json.NewDecoder(r.Body).Decode(&updateFilter); err != nil {
		http.Error(w, "decoding to update filter failed: "+err.Error(), http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	if err := u.userRepo.UpdateUser(id, updateFilter); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Failed udatading user: " + err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User successfully updated"))
}

func (u *UsersService) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := u.userRepo.DeleteUser(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotImplemented)
		return
	}

	w.WriteHeader(http.StatusOK)
}
