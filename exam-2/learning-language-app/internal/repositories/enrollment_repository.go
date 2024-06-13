package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

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

func NewEnrollmentRepository(db *sql.DB) *EnrollmentRepository {
	return &EnrollmentRepository{db: db}
}

func (r *EnrollmentRepository) EnrollUser(enrollment models.Enrollment) error {
	query := `INSERT INTO enrollments (enrollment_id, user_id, course_id, enrollment_date)
              VALUES ($1, $2, $3, $4)`

	enDate, err := time.Parse("02.01.2006", enrollment.EnrollmentDate)
	if err != nil {
		return fmt.Errorf("enrollment date is not valid")
	}
	_, err = r.db.Exec(query, enrollment.EnrollmentID, enrollment.UserID, enrollment.CourseID, enDate)
	return err
}

func (r *EnrollmentRepository) GetEnrollmentByID(enrollmentID string) (models.Enrollment, error) {
	query := `SELECT enrollment_id, user_id, course_id, enrollment_date, created_at, updated_at, deleted_at 
              FROM enrollments WHERE enrollment_id = $1`
	row := r.db.QueryRow(query, enrollmentID)

	var enrollment models.Enrollment
	err := row.Scan(&enrollment.EnrollmentID, &enrollment.UserID, &enrollment.CourseID, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdatedAt, &enrollment.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return enrollment, errors.New("enrollment not found")
		}
		return enrollment, err
	}
	return enrollment, nil
}

func (r *EnrollmentRepository) DeleteEnrollment(enrollmentID string) error {
	query := `DELETE FROM enrollments WHERE enrollment_id = $1`
	_, err := r.db.Exec(query, enrollmentID)
	return err
}

func (r *EnrollmentRepository) GetAllEnrollments(ctx *context.Context, f EnrollmentFilter) ([]models.Enrollment, error) {
	query := `SELECT enrollment_id, user_id, course_id, enrollment_date, created_at, updated_at 
	FROM enrollments WHERE deleted_at is null`
	conditions := []string{}
	args := []interface{}{}

	if f.CourseID != nil {
		conditions = append(conditions, fmt.Sprintf("course_id = $%d", len(args)+1))
		args = append(args, *f.CourseID)
	}
	if f.UserID != nil {
		conditions = append(conditions, fmt.Sprintf("user_id = $%d", len(args)+1))
		args = append(args, *f.UserID)
	}

	if f.EnrollmentDate != nil {
		enDate, err := time.Parse("02.01.2006", *f.EnrollmentDate)
		if err != nil {
			return nil, fmt.Errorf("parsing the date failed")
		}
		conditions = append(conditions, fmt.Sprintf(" DATE(enrollment_date) = DATE($%d) ", len(args)+1))
		args = append(args, enDate)
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	if f.Limit != nil {
		query += fmt.Sprintf("LIMIT %d", *f.Limit)
	}

	if f.Offset != nil {
		query += fmt.Sprintf("OFFSET %d", *f.Offset)
	}
	fmt.Println(query)
	rows, err := r.db.QueryContext(*ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []models.Enrollment
	for rows.Next() {
		var enrollment models.Enrollment
		err := rows.Scan(&enrollment.EnrollmentID, &enrollment.UserID, &enrollment.CourseID, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdatedAt)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}
	return enrollments, nil
}
