package postgres

import (
	"database/sql"
	"github.com/Go11Group/at_lesson/lesson28/model"
)

type CourseRepo struct {
	DB *sql.DB
}

func NewCourseRepo(DB *sql.DB) *CourseRepo {
	return &CourseRepo{DB}
}

func (c *CourseRepo) Create(course *model.Course) error {

	return nil
}
