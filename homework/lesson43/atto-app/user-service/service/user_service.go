package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"github.com/Go11Group/Javokhir-A/homework/lesson43/atto-app/user-service/repository"
	"github.com/Go11Group/Javokhir-A/homework/lesson43/atto-app/user-service/utils"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{UserRepository: repository.NewUserRepo(db)}
}

// Create creates a new user and returns the user ID or an error
func (us UserService) Create(ctx context.Context, user repository.CreateUser) (string, error) {
	if valError := utils.ValidateUser(user); valError != nil {
		errors := ""
		for _, err := range *valError {
			errors += err.Error()
		}
		return "", fmt.Errorf(errors)
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", err
	}

	user.Password = hashedPassword
	userID := uuid.New().String()

	err = us.UserRepository.CreateUser(user, userID)
	if err != nil {
		return "", fmt.Errorf("error creating user: %w", err)
	}

	return userID, nil
}

func (us UserService) GetUser(userID string) (*repository.UserData, error) {
	if _, err := uuid.Parse(userID); err != nil {
		return nil, fmt.Errorf("invalid user_id: " + err.Error())
	}

	userData, err := us.UserRepository.GetUser(userID)
	if err != nil {
		return nil, fmt.Errorf("getting user failed: " + err.Error())

	}
	return userData, nil
}
