package handlers

import (
	"net/http"

	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/repositories"
	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	var input services.LessonCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lesson, err := lh.LessonService.CreateLesson(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, lesson)
}

func (lh LessonHandler) GetLessonByID(c *gin.Context) {
	lessonID := c.Param("id")

	lesson, err := lh.LessonService.GetLessonByID(lessonID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lesson)
}

func (lh LessonHandler) UpdateLesson(c *gin.Context) {
	lessonID := c.Param("id")

	var input services.UpdateLesson
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lessonId, _ := uuid.Parse(lessonID)
	lesson, err := lh.LessonService.UpdateLesson(input, lessonId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lesson)
}

func (lh LessonHandler) DeleteLesson(c *gin.Context) {
	lessonID := c.Param("id")

	if err := lh.LessonService.DeleteLesson(lessonID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (lh LessonHandler) GetAllLessons(c *gin.Context) {
	filter := repositories.LessonFilter{}

	if err := c.BindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "binding to json failed: " + err.Error()})
		return
	}
	lessons, err := lh.LessonService.GetAllLessons(&filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lessons)
}
