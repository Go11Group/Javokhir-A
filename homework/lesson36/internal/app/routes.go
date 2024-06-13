package app

import (
	"database/sql"

	"github.com/Go11Group/Javokhir-A/homework/lesson36/internal/repositories"
	"github.com/Go11Group/Javokhir-A/homework/lesson36/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	userRepo := repositories.NewUserRepo(db)
	problemRepo := repositories.NewProblemRepo(db)

	// problemService := services.NewProblemService(problemRepo)
	// userService := services.NewUserService(userRepo)

	SetupHandlers(router, userService, problemService)
}
