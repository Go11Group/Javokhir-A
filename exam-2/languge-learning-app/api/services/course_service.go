package services

import (
	"time"

	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/dtos"
	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/models"
	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/repositories"
	"github.com/google/uuid"
)

type CourseService struct {
	courseRepo *repositories.CourseRepository
}

func NewCourseService(courseRepo *repositories.CourseRepository) *CourseService {
	return &CourseService{
		courseRepo: courseRepo,
	}
}

func (cs *CourseService) CreateCourse(courseDTO dtos.CourseCreateDTO) (*dtos.CourseResponseDTO, error) {
	// Generate UUID for the course
	courseID := uuid.NewString()

	// Get the current time
	now := time.Now()

	course := models.Course{
		CourseID:    courseID,
		Title:       courseDTO.Title,
		Description: courseDTO.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := cs.courseRepo.CreateCourse(&course); err != nil {
		return nil, err
	}

	// Generate the response DTO
	responseDTO := dtos.CourseResponseDTO{
		CourseID:    course.CourseID,
		Title:       course.Title,
		Description: course.Description,
		CreatedAt:   course.CreatedAt,
		UpdatedAt:   course.UpdatedAt,
	}

	return &responseDTO, nil
}

func (cs *CourseService) GetCourse(courseID string) (*dtos.CourseResponseDTO, error) {
	course, err := cs.courseRepo.GetCourse(courseID)
	if err != nil {
		return nil, err
	}

	// Populate the response DTO
	responseDTO := dtos.CourseResponseDTO{
		CourseID:    course.CourseID,
		Title:       course.Title,
		Description: course.Description,
		CreatedAt:   course.CreatedAt,
		UpdatedAt:   course.UpdatedAt,
	}

	return &responseDTO, nil
}

func (cs *CourseService) UpdateCourse(courseID string, courseDTO dtos.CourseCreateDTO) error {
	// Get the current time
	now := time.Now()

	course := models.Course{
		Title:       courseDTO.Title,
		Description: courseDTO.Description,
		UpdatedAt:   now,
	}

	if err := cs.courseRepo.UpdateCourse(courseID, &course); err != nil {
		return err
	}

	return nil
}

func (cs *CourseService) DeleteCourse(courseID string) error {
	return cs.courseRepo.DeleteCourse(courseID)
}
