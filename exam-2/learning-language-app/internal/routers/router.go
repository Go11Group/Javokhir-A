package routers

import (
	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, userHandler *handlers.UserHandler, lessonHandler *handlers.LessonHandler, courseHandler *handlers.CourseHandler, enrollmentHandler *handlers.EnrollmentHandler) {
	api := router.Group("/api")
	{
		users := api.Group("/user")
		{
			users.POST("/", userHandler.CreateUser)
			users.GET("/", userHandler.GetAllUsers)
			users.GET("/:id", userHandler.GetUserByID)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}

		lessons := api.Group("/lesson")
		{
			lessons.POST("/", lessonHandler.CreateLesson)
			lessons.GET("/", lessonHandler.GetAllLessons)
			lessons.GET("/:id", lessonHandler.GetLessonByID)
			lessons.PUT("/:id", lessonHandler.UpdateLesson)
			lessons.DELETE("/:id", lessonHandler.DeleteLesson)
			lessons.GET("/:id/lessons", lessonHandler.GetLessonsByCourse)

		}

		courses := api.Group("/course")
		{
			courses.POST("/", courseHandler.CreateCourse)
			courses.GET("/", courseHandler.GetAllCourses)
			courses.GET("/:id", courseHandler.GetCourseByID)
			courses.PUT("/:id", courseHandler.UpdateCourse)
			courses.DELETE("/:id", courseHandler.DeleteCourse)
			courses.GET("/:id/courses", courseHandler.GetCourseByUser)
			courses.GET("/:id/enrollments", courseHandler.GetEnroleldUsersByCourse)
			courses.GET("/popular", courseHandler.GetMostPopularCoursesHandler)
		}

		enrollments := api.Group("/enrollment")
		{
			enrollments.POST("/", enrollmentHandler.EnrollUser)
			enrollments.GET("/", enrollmentHandler.GetAllEnrollments)
			enrollments.GET("/:id", enrollmentHandler.GetEnrollmentByID)
			enrollments.DELETE("/:id", enrollmentHandler.DeleteEnrollment)
		}
	}
}
