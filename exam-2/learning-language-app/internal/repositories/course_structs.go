package repositories

import (
	"database/sql"
	"time"

	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/models"
	"github.com/google/uuid"
)

// plan for course repo
type CourseRepositoryPlan interface {
	CreateUser(course models.Course) error
	GetUserByID(courseId string) (models.Course, error)
	UpdateUser(course models.Course) error
	DeleteUser(courseId string) error
	GetAllUsers() ([]models.Course, error)
}

type CourseFilter struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Limit       *int       `json:"limit"`
	Offset      *int       `json:"offset"`
}

// UserCourses is used to find courses that a user take
type Course struct {
	CourseID    uuid.UUID `json:"course_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}
type UserCourses struct {
	UserID  uuid.UUID `json:"user_id"`
	Courses []Course  `json:"courses"`
}

// These structs are used to find the most popular course in a time of interval
type TimePeriod struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
type PopularCourse struct {
	CourseID         uuid.UUID `json:"course_id"`
	CourseTitle      string    `json:"course_title"`
	EnrollmentsCount int       `json:"enrollements_count"`
}
type ResponseCourse struct {
	TimePeriod     TimePeriod      `json:"time_period"`
	PopularCourses []PopularCourse `json:"popular_courses"`
}

type CourseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{db: db}
}
