package repositories

import (
	"database/sql"

	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/models"
)

type EnrollmentRepositoryPlan interface {
	EnrollUser(enrollment models.Enrollment) error
	GetEnrollmentByID(enrollmentID string) (models.Enrollment, error)
	DeleteEnrollment(enrollmentID string) error
	GetAllEnrollments() ([]models.Enrollment, error)
}

type EnrollmentRepository struct {
	db *sql.DB
}

type EnrollmentFilter struct {
	UserID         *string `json:"user_id"`
	CourseID       *string `json:"course_id"`
	EnrollmentDate *string `json:"enrollment_date"`
	Limit          *int    `json:"limit"`
	Offset         *int    `json:"offset"`
}
