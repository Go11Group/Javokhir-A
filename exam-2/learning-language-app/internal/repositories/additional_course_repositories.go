package repositories

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (c *CourseRepository) GetCourseByUser(userID uuid.UUID) (*UserCourses, error) {
	query := `
		SELECT  c.course_id, c.title, c.description
		FROM users u JOIN enrollments e
		ON e.user_id = u.user_id
		JOIN courses c ON
		c.course_id = e.course_id
		WHERE u.user_id = $1
		GROUP BY c.course_id, c.title, c.description;
	`
	var courses []Course

	rows, err := c.db.Query(query, userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.CourseID, &course.Title, &course.Description); err != nil {
			return nil, fmt.Errorf("failed while iterating over rows: " + err.Error())
		}
		courses = append(courses, course)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &UserCourses{
		UserID:  userID,
		Courses: courses,
	}, nil
}

type EnrolledUsers struct {
	CourseID      uuid.UUID `json:"course_id"`
	EnrolledUsers []User    `json:"enrolled_users"`
}
type User struct {
	UserID uuid.UUID `json:"user_id"`
	Name   string    `json:"name"`
	Email  string    `json:"email"`
}

func (c *CourseRepository) GetEnrolledUsersByCourse(courseID uuid.UUID) (*EnrolledUsers, error) {
	query := `
		SELECT u.user_id, u.name, u.email
		FROM users u 
		JOIN enrollments e ON
		u.user_id = e.user_id
		WHERE e.course_id = $1
	`
	rows, err := c.db.Query(query, courseID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.UserID, &user.Name, &user.Email); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		// fmt.Printf("Fetched user: %+v\n", user)
		users = append(users, user)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error encountered during row iteration: %v", err)
	}

	// // Debug log to print all users
	// fmt.Printf("All fetched users: %+v\n", users)

	return &EnrolledUsers{
		CourseID:      courseID,
		EnrolledUsers: users,
	}, nil
}

func (c *CourseRepository) GetMostPopularCourses(startDate, endDate time.Time) (*ResponseCourse, error) {
	query := `
		SELECT c.course_id, c.title, COUNT(e.enrollment_id) AS enrollment_count
		FROM courses c 
		JOIN enrollments e ON c.course_id = e.course_id
		WHERE e.enrollment_date BETWEEN $1 AND $2
		GROUP BY c.course_id, c.title
		ORDER BY enrollment_count DESC
		LIMIT 3;
	`

	rows, err := c.db.Query(query, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var popularCourses []PopularCourse
	for rows.Next() {
		var course PopularCourse
		if err := rows.Scan(&course.CourseID, &course.CourseTitle, &course.EnrollmentsCount); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		popularCourses = append(popularCourses, course)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error encountered during row iteration: %v", err)
	}

	return &ResponseCourse{
		TimePeriod: TimePeriod{
			StartDate: startDate.Format("2006-01-02"),
			EndDate:   endDate.Format("2006-01-02"),
		},
		PopularCourses: popularCourses,
	}, nil
}
