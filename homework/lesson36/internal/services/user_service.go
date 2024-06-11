package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Go11Group/Javokhir-A/homework/lesson36/internal/models"
	"github.com/Go11Group/Javokhir-A/homework/lesson36/internal/repositories"
	"github.com/gin-gonic/gin"
)

type UsersService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UsersService {
	return &UsersService{
		userRepo: userRepo,
	}
}

func (u *UsersService) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "binding request body's json to struct failed: " + err.Error()})
		return
	}

	if err := u.userRepo.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Creating user data in database failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User successfully created"})
}

func (u *UsersService) GetAllUsers(ctx *gin.Context) {
	var filter repositories.UserFilter

	if err := ctx.BindJSON(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, &filter)
		log.Println("Getting filter from request failed: " + err.Error())
		return
	}

	users, err := u.userRepo.GetAllUsers(filter)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Failed to fetch users: " + err.Error()})
	}

	ctx.JSON(http.StatusOK, users)

}

func (u *UsersService) GetUser(ctx *gin.Context) {

	userId := ctx.Param("id")

	if userId == "" {
		ctx.JSON(
			http.StatusBadRequest, gin.H{"error": "check for manual url request"},
		)
		log.Println("id must be string: ")
		return
	}

	user, err := u.userRepo.GetUser(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, user)

}

func (u *UsersService) UpdateUser(ctx *gin.Context) {
	var updatingUser repositories.UpdateUser

	if err := ctx.BindJSON(&updatingUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "updating user by query failed: " + err.Error()})
		return
	}

	id := ctx.Param("id")

	if err := u.userRepo.UpdateUser(id, updatingUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal updating failed: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user data successfully updated"})

}

func (u *UsersService) DeleteUser(c *gin.Context) {

	id := c.Param("id")

	if err := u.userRepo.DeleteUser(id); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "deleting user failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("user id %s deleted", id)})
}
