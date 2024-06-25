package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Go11Group/Javokhir-A/homework/lesson43/atto-app/user-service/repository"
	"github.com/Go11Group/Javokhir-A/homework/lesson43/atto-app/user-service/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserService *service.UserService
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{UserService: service.NewUserService(db)}
}

func (h *Handler) CreateUser(ctx *gin.Context) {
	user := repository.CreateUser{}

	if err := ctx.BindJSON(&user); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "Invalid response: " + err.Error()})
		return
	}

	id, err := h.UserService.Create(ctx, user)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"error": "Creating user failed: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created with id: " + id})
}

func (h *Handler) GetUserByID(c *gin.Context) {

	userID := c.Param("id")

	userReqBody, err := h.UserService.GetUser(userID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "getting user failed: " + err.Error()})
		return
	}

	if userReqBody == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": userReqBody})
}
