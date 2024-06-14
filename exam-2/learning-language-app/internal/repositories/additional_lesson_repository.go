package repositories

import (
	"fmt"

	"github.com/google/uuid"
)

func (l *LessonRepository) GetLessonByCourse(courseID uuid.UUID) (*CourseLessons, error) {
	query := `
		SELECT lesson_id, title, content 
		FROM lessons
		WHERE course_id = $1 AND deleted_at IS NULL
	`

	rows, err := l.db.Query(query, courseID)
	if err != nil {
		return nil, fmt.Errorf("no rows found" + err.Error())
	}

	var lessons []Lesson
	for rows.Next() {
		var lesson Lesson
		if err := rows.Scan(&lesson.Lesson_id, &lesson.Title, &lesson.Content); err != nil {
			return nil, fmt.Errorf("faield while iterating through rows")
		}
		lessons = append(lessons, lesson)
	}

	return &CourseLessons{
		CourseID: courseID,
		Lessons:  lessons,
	}, nil

}
