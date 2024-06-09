package services

import (
	"fmt"
	"net/http"

	"github.com/Go11Group/Javokhir-A/homework/lesson35/internal/models"
	"github.com/Go11Group/Javokhir-A/homework/lesson35/internal/repositories"
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
	u.userRepo.CreateUser(&models.User{})
	fmt.Println("Create user handler is working")
}
func (u *UsersService) UpdateUser(w http.ResponseWriter, r *http.Request) {

}
func (u *UsersService) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
func (u *UsersService) GetUser(w http.ResponseWriter, r *http.Request) {

}

func (u *UsersService) GetAllUsers(w http.ResponseWriter, r *http.Request) {

}
