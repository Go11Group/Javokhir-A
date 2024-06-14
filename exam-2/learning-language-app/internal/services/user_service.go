package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/models"
	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/repositories"
	"github.com/google/uuid"
)

type UserServicePlan interface {
	CreateUser(user models.User) error
	GetUserByID(userID string) (models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(userID string) error
	GetAllUsers() ([]models.User, error)
}

type UserService struct {
	UserRepository *repositories.UserRepository
}

type UserCreate struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Birthday string `json:"birthday" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type UpdateUser struct {
	Name     string `json:"name" binding:"omitempty"`
	Email    string `json:"email" binding:"omitempty,email"`
	Birthday string `json:"birthday" binding:"omitempty"`
	Password string `json:"password" binding:"omitempty"`
}

type UserResponse struct {
	UserID    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Birthday  string    `json:"birthday"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{
		UserRepository: repo,
	}
}

func (us UserService) CreateUser(user *UserCreate) (*UserResponse, error) {
	newId := uuid.New()
	birthday, err := time.Parse("02.01.2006", user.Birthday)
	if err != nil {
		return nil, fmt.Errorf("Layout:02.01.2006")
	}
	newUser := models.User{
		UserID:   newId,
		Name:     user.Name,
		Email:    user.Email,
		Birthday: birthday,
		Password: user.Password,
	}

	if err := us.UserRepository.CreateUser(&newUser); err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	userRes := &UserResponse{
		UserID:    newId,
		Name:      user.Name,
		Email:     user.Email,
		Birthday:  user.Birthday,
		Password:  user.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return userRes, nil
}

func (us UserService) GetUserByID(userID string) (*UserResponse, error) {

	user, err := us.UserRepository.GetUserByID(userID)
	if err != nil {
		return nil, errors.New("creating user failed: " + err.Error())
	}

	res := UserResponse{
		UserID:    user.UserID,
		Name:      user.Name,
		Email:     user.Email,
		Birthday:  user.Birthday.Format("02.01.2006"),
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return &res, nil
}

func (us UserService) UpdateUser(user UpdateUser, userId uuid.UUID) (*UserResponse, error) {
	u, err := us.UserRepository.GetUserByID(userId.String())
	if err != nil {
		return nil, errors.New("getting user to update failed: " + err.Error())
	}

	uRes := UserResponse{
		UserID:    userId,
		Name:      u.Name,
		Email:     u.Email,
		Birthday:  u.Birthday.Format("02.01.2006"),
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}

	if user.Name != "" {
		u.Name = user.Name
		uRes.Name = user.Name
	}

	if user.Email != "" {
		u.Email = user.Email
		uRes.Email = user.Email
	}

	if user.Birthday != "" {
		birthday, err := time.Parse("02.01.2006", user.Birthday)
		if err != nil {
			return nil, errors.New("time layout is not valid: " + err.Error())
		}
		u.Birthday = birthday
		uRes.Birthday = user.Birthday
	}

	if user.Password != "" {
		u.Password = user.Password
		uRes.Password = user.Password
	}

	if err := us.UserRepository.UpdateUser(u); err != nil {
		return nil, err
	}

	return &uRes, nil
}

func (us UserService) DeleteUser(userID string) error {

	if err := us.UserRepository.DeleteUser(userID); err != nil {
		if sql.ErrNoRows == err {
			return fmt.Errorf("not found")
		}
		return fmt.Errorf("deleteting user failed: " + err.Error())
	}

	return nil
}
func (us UserService) GetAllUsers(filter *repositories.UserFilter) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// Using a channel to handle the result and error
	resultCh := make(chan []models.User)
	errCh := make(chan error)

	go func() {
		users, err := us.UserRepository.GetAllUsers(filter, ctx)
		if err != nil {
			errCh <- fmt.Errorf("getting users failed: %v", err)
			return
		}
		resultCh <- users
	}()

	select {
	case <-ctx.Done():
		return nil, errors.New("timeout")
	case users := <-resultCh:
		return users, nil
	case err := <-errCh:
		return nil, err
	}
}

func (us *UserService) GetUserProgressByUserID(userID uuid.UUID) (*repositories.UserProgressResponse, error) {
	return us.UserRepository.GetUserProgressByUserID(userID)
}
