package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/models"
)

type CourseRepository struct {
	db *sql.DB
}

func NewCourseRepo(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

func (cr *CourseRepository) CreateCourse(course *models.Course) error {
	_, err := cr.db.Exec(`
		INSERT INTO courses (course_id, title, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`, course.CourseID, course.Title, course.Description, course.CreatedAt, course.UpdatedAt)
	if err != nil {
		return fmt.Errorf("creating course failed: %v", err)
	}

	return nil
}

func (cr *CourseRepository) GetCourse(courseID string) (*models.Course, error) {
	query := `
		SELECT course_id, title, description, created_at, updated_at
		FROM courses
		WHERE course_id = $1
	`

	row := cr.db.QueryRow(query, courseID)
	course := models.Course{}

	err := row.Scan(&course.CourseID, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("getting course failed: %v", err)
	}

	return &course, nil
}

func (cr *CourseRepository) UpdateCourse(courseID string, course *models.Course) error {
	query := `
		UPDATE courses 
		SET title = $2, description = $3, updated_at = $4
		WHERE course_id = $1
	`

	_, err := cr.db.Exec(query, courseID, course.Title, course.Description, time.Now())
	if err != nil {
		return fmt.Errorf("updating course failed: %v", err)
	}

	return nil
}

func (cr *CourseRepository) DeleteCourse(courseID string) error {
	query := `
		DELETE FROM courses
		WHERE course_id = $1
	`

	_, err := cr.db.Exec(query, courseID)
	if err != nil {
		return fmt.Errorf("deleting course failed: %v", err)
	}

	return nil
}
