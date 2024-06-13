package handlers

import (
	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/services"
	"github.com/gin-gonic/gin"
)

type LessonHandlerPlan interface {
	CreateLesson(c *gin.Context)
	GetLessonByID(c *gin.Context)
	UpdateLesson(c *gin.Context)
	DeleteLesson(c *gin.Context)
	GetAllLessons(c *gin.Context)
}

type LessonHandler struct {
	LessonService *services.LessonService
}

func NewLessonHandler(service *services.LessonService) *LessonHandler {
	return &LessonHandler{
		LessonService: service,
	}
}

func (lh LessonHandler) CreateLesson(c *gin.Context) {
	// Implementation here
}

func (lh LessonHandler) GetLessonByID(c *gin.Context) {
	// Implementation here
}

func (lh LessonHandler) UpdateLesson(c *gin.Context) {
	// Implementation here
}

func (lh LessonHandler) DeleteLesson(c *gin.Context) {
	// Implementation here
}

func (lh LessonHandler) GetAllLessons(c *gin.Context) {
	// Implementation here
}
