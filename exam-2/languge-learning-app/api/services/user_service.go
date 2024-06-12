package services

import (
	"fmt"
	"time"

	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/dtos"
	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/models"
	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (us *UserService) CreateUser(userDTO dtos.UserCreateDTO) (*dtos.UserResponseDTO, error) {
	// Parse the birthday string using the correct layout
	birthday, err := time.Parse("02.01.2006", userDTO.Birthday)
	if err != nil {
		return nil, fmt.Errorf("parsing birthday failed: %v", err)
	}

	// Get the current time
	now := time.Now()

	user := models.User{
		Name:      userDTO.Name,
		Email:     userDTO.Email,
		Birthday:  birthday,
		Password:  userDTO.Password,
		CreatedAt: now, // Set the created time to the current time
		UpdatedAt: now, // Set the updated time to the current time
	}
	var user_id *string

	if user_id, err = us.userRepo.CreateUser(&user); err != nil {
		return nil, err
	}

	// generate the response DTO
	responseDTO := dtos.UserResponseDTO{
		UserID:    *user_id, // UserID is set after creation
		Name:      user.Name,
		Email:     user.Email,
		Birthday:  user.Birthday.String(),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return &responseDTO, nil
}

func (us *UserService) GetUser(userID string) (*dtos.UserResponseDTO, error) {
	// Validate if the user ID is a valid UUID
	if ok := isUUID(userID); !ok {
		return nil, fmt.Errorf("%s is not a valid UUID", userID)
	}

	// Retrieve the user model from the repository
	user, err := us.userRepo.GetUser(userID)
	if err != nil {
		return nil, err
	}

	// Populate the response DTO
	responseDTO := dtos.UserResponseDTO{
		UserID:    user.UserId,
		Name:      user.Name,
		Email:     user.Email,
		Birthday:  user.Birthday.String(),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return &responseDTO, nil
}

func (us *UserService) GetAllUsers(filter repositories.UserFilter) ([]models.User, error) {
	users, err := us.userRepo.GetAllUsers(filter)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (us *UserService) UpdateUser(updateUser dtos.UserUpdateDTO, userID string) error {
	// Check if the user ID is a valid UUID
	if ok := isUUID(userID); !ok {
		return fmt.Errorf("%s is not a valid UUID", userID)
	}
	birthday, err := time.Parse("02.01.2006", *updateUser.Birthday)
	if err != nil {
		return fmt.Errorf("birthday layout is not vaild: " + err.Error())
	}
	// Convert DTO to repository's UpdateUser struct
	repoUpdateUser := repositories.UpdateUser{
		Name:     updateUser.Name,
		Email:    updateUser.Email,
		Birthday: &birthday,
		Password: updateUser.Password,
	}

	// Call the repository method to update the user
	err = us.userRepo.UpdateUser(userID, repoUpdateUser)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserService) DeleteUser(userID string) error {
	// Check if the user ID is a valid UUID
	if ok := isUUID(userID); !ok {
		return fmt.Errorf("%s is not a valid UUID", userID)
	}

	// Call the repository method to delete the user
	err := us.userRepo.DeleteUser(userID)
	if err != nil {
		return err
	}

	return nil
}
