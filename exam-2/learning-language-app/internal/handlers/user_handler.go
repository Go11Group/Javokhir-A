package handlers

import (
	"net/http"

	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/repositories"
	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandlerPlan interface {
	CreateUser(c *gin.Context)
	GetUserByID(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetAllUsers(c *gin.Context)
}

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{
		UserService: service,
	}
}

func (uh UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	ures, err := uh.UserService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "getting response from user service failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, ures)
}

func (uh UserHandler) CreateUser(c *gin.Context) {
	var uc services.UserCreate
	var errCode = http.StatusBadRequest

	if err := c.BindJSON(&uc); err != nil {
		c.JSON(errCode, gin.H{"error": "failed binding json to user struct: " + err.Error()})
		return
	}

	res, err := uh.UserService.CreateUser(&uc)

	if err != nil {
		c.JSON(errCode, gin.H{"error": "creating user failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &res)
}

func (uh UserHandler) UpdateUser(c *gin.Context) {
	updateUser := services.UpdateUser{}
	id := c.Param("id")

	if err := c.BindJSON(&updateUser); err != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": "binding to json failed: " + err.Error()})
		return
	}

	uId, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is not valid" + err.Error()})
		return
	}

	res, err := uh.UserService.UpdateUser(updateUser, uId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "updated user failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (uh UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	errCode := http.StatusBadRequest

	if err := uh.UserService.DeleteUser(id); err != nil {
		if err.Error() == "not found" {
			errCode = http.StatusNotFound
		}

		c.JSON(errCode, gin.H{"error": "deleting user failed: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user successfully deleted"})
}

func (uh UserHandler) GetAllUsers(c *gin.Context) {
	filter := repositories.UserFilter{}

	if err := c.BindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "binding to json failed: " + err.Error()})
		return
	}

	users, err := uh.UserService.GetAllUsers(&filter)
	if err != nil {
		if err.Error() == "timeout" {
			c.JSON(http.StatusRequestTimeout, gin.H{"message": "getting users time exeeded, use limit to limit the data"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "getting all users faield: " + err.Error()})
		return
	}

	c.JSON(http.StatusFound, &users)
}

func (uh *UserHandler) GetUserProgressByUserID(c *gin.Context) {
	userID := c.Param("id")

	// Parse the userID string into a UUID
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id format"})
		return
	}

	response, err := uh.UserService.GetUserProgressByUserID(parsedUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
