package configs

import (
	"database/sql"

	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/handlers"
	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/repositories"
	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/routers"
	"github.com/Go11Group/Javokhir-A/exam-2/learning-language-app/internal/services"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db *sql.DB
}

func NewServer(db *sql.DB) *Server {
	return &Server{
		db: db,
	}
}

func (s Server) Start(addr string) {
	router := gin.Default()

	userRepo := repositories.NewUserRepository(s.db)
	userSer := services.NewUserService(userRepo)
	userHand := handlers.NewUserHandler(userSer)

	lessonRepo := repositories.NewLessonRepository(s.db)
	lessonSer := services.NewLessonService(lessonRepo)
	lessonHand := handlers.NewLessonHandler(lessonSer)

	courseRepo := repositories.NewCourseRepository(s.db)
	courseSer := services.NewCourseService(courseRepo)
	courseHand := handlers.NewCourseHandler(courseSer)

	enrollmentRepo := repositories.NewEnrollmentRepository(s.db)
	enrollmentSer := services.NewEnrollmentService(enrollmentRepo)
	enrollmentHand := handlers.NewEnrollmentHandler(enrollmentSer)

	routers.SetupRouter(router, userHand, lessonHand, courseHand, enrollmentHand)

	router.Run(addr)
}
