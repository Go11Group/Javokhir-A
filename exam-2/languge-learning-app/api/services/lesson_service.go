package services

import (
	"fmt"
	"time"

	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/dtos"
	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/models"
	"github.com/Go11Group/Javokhir-A/exam-2/languge-learning-app/api/repositories"
	"github.com/google/uuid"
)

type LessonService struct {
	lessonRepo *repositories.LessonRepository
}

func NewLessonService(lessonRepo *repositories.LessonRepository) *LessonService {
	return &LessonService{
		lessonRepo: lessonRepo,
	}
}

func (ls *LessonService) CreateLesson(lessonDTO dtos.LessonCreateDTO) (*dtos.LessonResponseDTO, error) {
	// Parse the course ID string to UUID
	courseID, err := uuid.Parse(lessonDTO.CourseID)
	if err != nil {
		return nil, fmt.Errorf("parsing course ID failed: %v", err)
	}

	// Get the current time
	now := time.Now()

	lesson := models.Lesson{
		CourseID:  courseID.String(),
		Title:     lessonDTO.Title,
		Content:   lessonDTO.Content,
		CreatedAt: now,
		UpdatedAt: now,
		DeletedAt: time.Time{}, // Not deleted yet
	}

	lessonID, err := ls.lessonRepo.CreateLesson(&lesson)
	if err != nil {
		return nil, err
	}

	// Generate the response DTO
	responseDTO := dtos.LessonResponseDTO{
		LessonID:  *lessonID,
		CourseID:  lesson.CourseID,
		Title:     lesson.Title,
		Content:   lesson.Content,
		CreatedAt: lesson.CreatedAt,
		UpdatedAt: lesson.UpdatedAt,
	}

	return &responseDTO, nil
}

func (ls *LessonService) GetLesson(lessonID string) (*dtos.LessonResponseDTO, error) {
	lesson, err := ls.lessonRepo.GetLesson(lessonID)
	if err != nil {
		return nil, err
	}

	// Populate the response DTO
	responseDTO := dtos.LessonResponseDTO{
		LessonID:  lesson.LessonID,
		CourseID:  lesson.CourseID,
		Title:     lesson.Title,
		Content:   lesson.Content,
		CreatedAt: lesson.CreatedAt,
		UpdatedAt: lesson.UpdatedAt,
		DeletedAt: lesson.DeletedAt,
	}

	return &responseDTO, nil
}

func (ls *LessonService) UpdateLesson(lessonID string, lessonUpdateDTO dtos.LessonUpdateDTO) (*dtos.LessonResponseDTO, error) {
	// Check if the lesson ID is a valid UUID
	if ok := isUUID(lessonID); !ok {
		return nil, fmt.Errorf("%s is not a valid UUID", lessonID)
	}

	// Parse the lesson update time string using the correct layout
	updatedAt, err := time.Parse("2006-01-02T15:04:05Z07:00", lessonUpdateDTO.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("parsing updated_at failed: %v", err)
	}

	// Convert DTO to repository's UpdateLesson struct
	repoUpdateLesson := repositories.UpdateLesson{
		Title:     lessonUpdateDTO.Title,
		Content:   lessonUpdateDTO.Content,
		UpdatedAt: &updatedAt,
	}

	// Call the repository method to update the lesson
	err = ls.lessonRepo.UpdateLesson(lessonID, repoUpdateLesson)
	if err != nil {
		return nil, err
	}

	// Retrieve the updated lesson from the repository
	lesson, err := ls.lessonRepo.GetLesson(lessonID)
	if err != nil {
		return nil, err
	}

	// Populate the response DTO
	responseDTO := dtos.LessonResponseDTO{
		LessonID:  lesson.LessonID,
		CourseID:  lesson.CourseID,
		Title:     lesson.Title,
		Content:   lesson.Content,
		CreatedAt: lesson.CreatedAt,
		UpdatedAt: lesson.UpdatedAt,
	}

	return &responseDTO, nil
}

func (ls *LessonService) DeleteLesson(lessonID string) error {
	return ls.lessonRepo.DeleteLesson(lessonID)
}

// Helper function to validate UUID
func isUUID(s string) bool {
	_, err := uuid.Parse(s)
	return err == nil
}
