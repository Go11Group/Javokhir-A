package routes

import (
	"database/sql"

	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/handlers"
	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/repositories"
	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	userRepo := repositories.NewUserRepo(db)
	UserService := services.NewUserService(userRepo)
	UserHandler := handlers.NewUserHandler(UserService)

	user := router.Group("/user")

	user.POST("/create", UserHandler.CreateUser)
	user.GET("/get/:id", UserHandler.GetUser)
	user.PUT("/update/:id", UserHandler.UpdateUser)
	user.DELETE("/delete/:id", UserHandler.DeleteUser)

}
