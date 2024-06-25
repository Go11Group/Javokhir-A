package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

type User struct {
	UserId    string    `json:"user_id"` // id is required
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"` // email is required
	Age       int       `json:"age"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"` // default value for this is "user"
	Password  string    `json:"password"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type Response struct {
	Message *string `json:"message"`
	Error   *string `json:"error"`
	User    *User   `json:"user"`
}

func (h *Handler) UserHandler(c *gin.Context) {
	method := c.Request.Method
	url := c.Request.URL.Path
	body := c.Request.Body
	defer body.Close()

	// userID := ""

	// if method == "GET" || method == "PUT" {
	// 	userID = c.Param("id")
	// } else {
	// 	userID = ""
	// }

	client := http.Client{}
	request, err := http.NewRequest(method, "http://localhost:8085"+url, body)
	fmt.Println(url)
	if err != nil {
		log.Println("Error creating new request:", err)
		c.JSON(http.StatusBadGateway, gin.H{"error": "Invalid client request: " + err.Error()})
		return
	}

	request.Header.Set("Content-Type", "application/json")
	res, err := client.Do(request)
	if err != nil {
		log.Println("Error sending request:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldn't send request: " + err.Error()})
		return
	}
	defer res.Body.Close()

	r := Response{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Println("Error decoding response body:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode response: " + err.Error()})
		return
	}

	if r.Error != nil {
		log.Println("Response error:", *r.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": *r.Error})
		return
	}

	if r.User != nil {
		c.JSON(http.StatusOK, gin.H{"user": *r.User})
		return
	}

	if r.Message != nil {
		c.JSON(http.StatusOK, gin.H{"message": *r.Message})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Request processed successfully"})
	}
}
