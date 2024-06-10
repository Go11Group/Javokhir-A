package app

import (
	"database/sql"

	"github.com/Go11Group/Javokhir-A/homework/lesson35/internal/repositories"
	"github.com/Go11Group/Javokhir-A/homework/lesson35/internal/services"
	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router, db *sql.DB) {
	userRepo := repositories.NewUserRepo(db)
	problemRepo := repositories.NewProblemRepo(db)

	problemService := services.NewProblemService(problemRepo)
	userService := services.NewUserService(userRepo)

	SetupHandlers(router, userService, problemService)
}
