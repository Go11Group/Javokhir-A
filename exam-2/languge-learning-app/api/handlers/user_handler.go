package handlers

import (
	"net/http"

	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/dtos"
	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(UserService *services.UserService) *UserHandler {
	return &UserHandler{
		UserService: UserService,
	}
}

func (uh *UserHandler) CreateUser(c *gin.Context) {
	var userData dtos.UserCreateDTO

	if err := c.BindJSON(&userData); err != nil {
		c.JSON(403, gin.H{"error": "binding json data failed: " + err.Error()})
		return
	}

	user, err := uh.UserService.CreateUser(userData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (uh *UserHandler) GetUser(c *gin.Context) {
	userId := c.Param("id")

	codeErr := http.StatusBadRequest

	user, err := uh.UserService.GetUser(userId)
	if err != nil {
		if err.Error() == "not found" {
			codeErr = http.StatusNotFound
		}
		c.JSON(codeErr, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &user)

}

func (uh *UserHandler) UpdateUser(c *gin.Context) {
	userId := c.Param("id")

	var updateUserDTO dtos.UserUpdateDTO
	if err := c.BindJSON(&updateUserDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "binding JSON data failed: " + err.Error()})
		return
	}

	err := uh.UserService.UpdateUser(updateUserDTO, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (uh *UserHandler) DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	codeErr := http.StatusBadRequest
	errMsg := ""
	err := uh.UserService.DeleteUser(userId)
	if err != nil {
		if err.Error() == "not found" {
			codeErr = http.StatusNotFound
			errMsg = "no such user by this id"
		}
		c.JSON(codeErr, gin.H{"error": errMsg + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
