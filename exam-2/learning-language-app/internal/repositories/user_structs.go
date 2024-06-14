package repositories

import (
	"database/sql"
	"time"

	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/models"
	"github.com/google/uuid"
)

// plan for user repo
type UserRepositoryPlan interface {
	CreateUser(user models.User) error
	GetUserByID(userID string) (models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(userID string) error
	GetAllUsers() ([]models.User, error)
}

type UserFilter struct {
	Name      *string    `json:"name"`
	Email     *string    `json:"email"`
	Birthday  *time.Time `json:"birthday"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Limit     *int       `json:"limit"`
	Offset    *int       `json:"offset"`
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// UserProgress, it is the progress of a user in a course
type UserProgress struct {
	CourseID         uuid.UUID `json:"course_id"`
	CourseTitle      string    `json:"course_title"`
	CompletedLessons int       `json:"completed_lessons"`
	TotalLessons     int       `json:"total_lessons"`
}

// UserProgressResponse that is the response format for user progress
type UserProgressResponse struct {
	UserID   uuid.UUID      `json:"user_id"`
	Progress []UserProgress `json:"progress"`
}
