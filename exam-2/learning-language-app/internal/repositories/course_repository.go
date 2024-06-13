package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
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

type CourseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

func (r *CourseRepository) CreateCourse(course *models.Course) error {
	query := `INSERT INTO courses (course_id, title, description)
              VALUES ($1, $2, $3)`
	newId := uuid.New()
	_, err := r.db.Exec(query, newId, course.Title, course.Description)
	if err != nil {
		return fmt.Errorf("failed execute the query")
	}
	return nil
}

func (r *CourseRepository) GetCourseByID(courseId string) (models.Course, error) {
	query := `SELECT course_id, title, description, created_at, updated_at 
              FROM courses WHERE course_id = $1 and deleted_at IS NULL`
	row := r.db.QueryRow(query, courseId)

	var course models.Course
	err := row.Scan(&course.CourseID, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return course, errors.New("course not found")
		}
		return course, err
	}

	return course, nil
}

func (r *CourseRepository) UpdateCourse(course models.Course) error {
	query := `UPDATE courses SET title = $1, description = $2, updated_at = $3 WHERE course_id = $4 and deleted_at IS NULL`
	_, err := r.db.Exec(query, course.Title, course.Description, time.Now(), course.CourseID)
	return err
}

func (r *CourseRepository) DeleteCourse(courseID string) error {
	query := `UPDATE FROM courses SET deleted_at = CURRENT_TIMESTAMP WHERE course_id = $1 and deleted_at IS NULL`
	_, err := r.db.Exec(query, courseID)
	return err
}
func (r *CourseRepository) GetAllCourses(f *CourseFilter, ctx context.Context) ([]models.Course, error) {
	query := `SELECT course_id, title, description, created_at, updated_at FROM courses WHERE deleted_at IS NULL`

	var conditions []string
	var args []interface{}

	if f.Title != nil {
		conditions = append(conditions, fmt.Sprintf("title = $%d", len(args)+1))
		args = append(args, *f.Title)
	}
	if f.Description != nil {
		conditions = append(conditions, fmt.Sprintf("description = $%d", len(args)+1))
		args = append(args, *f.Description)
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	if f.Limit != nil {
		query += fmt.Sprintf(" LIMIT %d", *f.Limit)
	}

	if f.Offset != nil {
		query += fmt.Sprintf(" OFFSET %d", *f.Offset)
	}

	fmt.Println(query, args)
	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []models.Course
	for rows.Next() {
		var course models.Course
		if err := rows.Scan(&course.CourseID, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}
