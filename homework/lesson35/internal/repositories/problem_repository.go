package repositories

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Go11Group/Javokhir-A/homework/lesson35/internal/models"
	"github.com/google/uuid"
)

type ProblemRepository struct {
	db *sql.DB
}

func NewProblemRepo(db *sql.DB) *ProblemRepository {
	return &ProblemRepository{
		db: db,
	}
}

type UpdateProblem struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Complexity  *string `json:"compexity"`
}

type ProblemFilter struct {
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Complexity  *string    `json:"compexity"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Limit       *int       `json:"limit"`
	Offset      *int       `json:"offset"`
}

func (p ProblemRepository) GetAllProblems(filter ProblemFilter) ([]models.Problem, error) {
	var problems []models.Problem
	var conditions []string
	var args []interface{}

	query := `
        SELECT id, title, description, complexity, created_at, updated_at
        FROM problems
    `

	if filter.Title != nil {
		conditions = append(conditions, fmt.Sprintf("title = $%d", len(args)+1))
		args = append(args, *filter.Title)
	}

	if filter.Description != nil {
		conditions = append(conditions, fmt.Sprintf("description = $%d", len(args)+1))
		args = append(args, *filter.Description)
	}

	if filter.Complexity != nil {
		conditions = append(conditions, fmt.Sprintf("complexity = $%d", len(args)+1))
		args = append(args, *filter.Complexity)
	}

	if filter.CreatedAt != nil {
		conditions = append(conditions, fmt.Sprintf("created_at = $%d", len(args)+1))
		args = append(args, *filter.CreatedAt)
	}

	if filter.UpdatedAt != nil {
		conditions = append(conditions, fmt.Sprintf("updated_at = $%d", len(args)+1))
		args = append(args, *filter.UpdatedAt)
	}

	conditions = append(conditions, "deleted_at IS NULL")

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	if filter.Limit != nil {
		query += fmt.Sprintf(" LIMIT %d", *filter.Limit)
	}

	if filter.Offset != nil {
		query += fmt.Sprintf(" OFFSET %d", *filter.Offset)
	}

	rows, err := p.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed while querying: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var problem models.Problem
		if err := rows.Scan(&problem.Id, &problem.Title, &problem.Description, &problem.Complexity, &problem.CreatedAt, &problem.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed while scanning data to slice: %v", err)
		}
		problems = append(problems, problem)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return problems, nil
}

func (p ProblemRepository) CreateProblem(problem *models.Problem) error {
	id := uuid.NewString()
	query := `
        INSERT INTO problems (id, title, description, complexity)
        VALUES ($1, $2, $3, $4)
    `
	_, err := p.db.Exec(query, id, problem.Title, problem.Description, problem.Complexity)
	if err != nil {
		return fmt.Errorf("creating problem failed: %v", err)
	}

	return nil
}

func (p ProblemRepository) GetProblem(problemID string) (*models.Problem, error) {
	query := `
        SELECT id, title, description, complexity, created_at, updated_at
        FROM problems
        WHERE deleted_at IS NULL AND id = $1
    `
	row := p.db.QueryRow(query, problemID)
	problem := models.Problem{}

	err := row.Scan(&problem.Id, &problem.Title, &problem.Description, &problem.Complexity, &problem.CreatedAt, &problem.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("getting problem failed: %v", err)
	}

	return &problem, nil
}

func (p ProblemRepository) DeleteProblem(problemId string) error {
	query := `
        UPDATE problems 
        SET deleted_at = CURRENT_TIMESTAMP 
        WHERE id = $1 AND deleted_at IS NULL
    `

	if _, err := p.db.Exec(query, problemId); err != nil {
		return fmt.Errorf("failed to delete problem by this id: %v", err)
	}

	return nil
}

func (p ProblemRepository) UpdateProblem(problemID string, updateFilter UpdateProblem) error {
	var conditions []string
	var args []interface{}

	query := `
        SELECT id
        FROM problems
        WHERE deleted_at IS NULL AND id = $1
    `

	if err := p.db.QueryRow(query, problemID).Err(); err != nil {
		return fmt.Errorf("problem by this id not found: %v", err)
	}

	query = `
        UPDATE problems SET 
    `

	if updateFilter.Title != nil {
		conditions = append(conditions, fmt.Sprintf("title = $%d", len(args)+1))
		args = append(args, *updateFilter.Title)
	}

	if updateFilter.Description != nil {
		conditions = append(conditions, fmt.Sprintf("description = $%d", len(args)+1))
		args = append(args, *updateFilter.Description)
	}

	if updateFilter.Complexity != nil {
		conditions = append(conditions, fmt.Sprintf("complexity = $%d", len(args)+1))
		args = append(args, *updateFilter.Complexity)
	}

	if len(conditions) == 0 {
		return fmt.Errorf("no fields to update")
	}

	args = append(args, problemID)
	query += strings.Join(conditions, ", ") + fmt.Sprintf(" WHERE id = $%d AND deleted_at IS NULL", len(args))

	if _, err := p.db.Exec(query, args...); err != nil {
		return fmt.Errorf("failed executing query: %v", err)
	}
	return nil
}
