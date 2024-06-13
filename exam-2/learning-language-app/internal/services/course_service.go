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

type CourseServicePlan interface {
	CreateCourse(course models.Course) error
	GetCourseByID(courseID string) (models.Course, error)
	UpdateCourse(course models.Course) error
	DeleteCourse(courseID string) error
	GetAllCourses() ([]models.Course, error)
}

type CourseService struct {
	CourseRepository *repositories.CourseRepository
}

type CourseCreate struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UpdateCourse struct {
	Title       string `json:"title" binding:"omitempty"`
	Description string `json:"description" binding:"omitempty"`
}

type CourseResponse struct {
	CourseID    uuid.UUID `json:"course_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewCourseService(repo *repositories.CourseRepository) *CourseService {
	return &CourseService{
		CourseRepository: repo,
	}
}

func (cs CourseService) CreateCourse(course *CourseCreate) (*CourseResponse, error) {
	newId := uuid.New()
	newCourse := models.Course{
		CourseID:    newId.String(),
		Title:       course.Title,
		Description: course.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := cs.CourseRepository.CreateCourse(&newCourse); err != nil {
		return nil, fmt.Errorf("failed to create course: %v", err)
	}

	courseRes := &CourseResponse{
		CourseID:    newId,
		Title:       course.Title,
		Description: course.Description,
		CreatedAt:   newCourse.CreatedAt,
		UpdatedAt:   newCourse.UpdatedAt,
	}
	return courseRes, nil
}

func (cs CourseService) GetCourseByID(courseID string) (*CourseResponse, error) {
	course, err := cs.CourseRepository.GetCourseByID(courseID)
	if err != nil {
		return nil, errors.New("getting course failed: " + err.Error())
	}

	res := CourseResponse{
		CourseID:    uuid.MustParse(course.CourseID),
		Title:       course.Title,
		Description: course.Description,
		CreatedAt:   course.CreatedAt,
		UpdatedAt:   course.UpdatedAt,
	}

	return &res, nil
}

func (cs CourseService) UpdateCourse(course UpdateCourse, courseId uuid.UUID) (*CourseResponse, error) {
	c, err := cs.CourseRepository.GetCourseByID(courseId.String())
	if err != nil {
		return nil, errors.New("getting course to update failed: " + err.Error())
	}

	cRes := CourseResponse{
		CourseID:    courseId,
		Title:       c.Title,
		Description: c.Description,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}

	if course.Title != "" {
		c.Title = course.Title
		cRes.Title = course.Title
	}

	if course.Description != "" {
		c.Description = course.Description
		cRes.Description = course.Description
	}

	if err := cs.CourseRepository.UpdateCourse(c); err != nil {
		return nil, err
	}

	return &cRes, nil
}

func (cs CourseService) DeleteCourse(courseID string) error {
	if err := cs.CourseRepository.DeleteCourse(courseID); err != nil {
		if sql.ErrNoRows == err {
			return fmt.Errorf("not found")
		}
		return fmt.Errorf("deleting course failed: " + err.Error())
	}

	return nil
}

func (cs CourseService) GetAllCourses(filter *repositories.CourseFilter) ([]models.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	resultCh := make(chan []models.Course)
	errCh := make(chan error)

	go func() {
		courses, err := cs.CourseRepository.GetAllCourses(filter, ctx)
		if err != nil {
			errCh <- fmt.Errorf("getting courses failed: %v", err)
			return
		}
		resultCh <- courses
	}()

	select {
	case <-ctx.Done():
		return nil, errors.New("timeout")
	case courses := <-resultCh:
		return courses, nil
	case err := <-errCh:
		return nil, err
	}
}
