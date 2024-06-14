package handlers

import (
	"fmt"
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
	var input map[string]interface{}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to bind request body"})
		return
	}

	courseIDStr, ok := input["course_id"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "course_id required and must be a string"})
		return
	}

	courseID, err := uuid.Parse(courseIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course_id format"})
		return
	}

	title, ok := input["title"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title required and must be a string"})
		return
	}

	content, ok := input["content"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "content required and must be a string"})
		return
	}

	lesson, err := lh.LessonService.CreateLesson(courseID, title, content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, lesson)
}

func (lh LessonHandler) GetLessonByID(c *gin.Context) {
	lessonID := c.Param("id")
	lessonIDuuid, err := uuid.Parse(lessonID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course_id format"})
		return
	}

	lesson, err := lh.LessonService.GetLessonByID(lessonIDuuid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lesson)
}

func (lh LessonHandler) UpdateLesson(c *gin.Context) {
	lessonID := c.Param("id")

	var input repositories.UpdateLesson

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lessonId, _ := uuid.Parse(lessonID)
	lesson, err := lh.LessonService.UpdateLesson(lessonId, *input.Title, *input.Content)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, lesson)
}

func (lh LessonHandler) DeleteLesson(c *gin.Context) {
	lessonID := c.Param("id")
	lessonIDuuid, err := uuid.Parse(lessonID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid lesson_id format"})
		return
	}

	if err := lh.LessonService.DeleteLesson(lessonIDuuid); err != nil {
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

func (lh LessonHandler) GetLessonsByCourse(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	courseID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be uuid"})
		return
	}
	courseLessons, err := lh.LessonService.GetLessonsByCourse(courseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusFound, courseLessons)
}
