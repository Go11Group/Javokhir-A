package repositories

import (
	"fmt"

	"github.com/google/uuid"
)

// GetUserProgressByUserID is used to find the progress of a user in all courses.
func (c *UserRepository) GetUserProgressByUserID(userID uuid.UUID) (*UserProgressResponse, error) {
	query := `
		SELECT c.course_id, c.title, COUNT(l.lesson_id) as completed_lessons, COUNT(*) as total_lessons
		FROM courses c 
		LEFT JOIN lessons l ON c.course_id = l.course_id
		JOIN enrollments e ON c.course_id = e.course_id
		WHERE e.user_id = $1
		GROUP BY c.course_id, c.title
	`
	rows, err := c.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var progress []UserProgress
	for rows.Next() {
		var p UserProgress
		if err := rows.Scan(&p.CourseID, &p.CourseTitle, &p.CompletedLessons, &p.TotalLessons); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		progress = append(progress, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error encountered during row iteration: %v", err)
	}

	return &UserProgressResponse{
		UserID:   userID,
		Progress: progress,
	}, nil
}
