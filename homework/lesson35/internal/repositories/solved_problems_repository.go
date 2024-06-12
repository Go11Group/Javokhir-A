package repositories

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Go11Group/Javokhir-A/homework/lesson35/internal/models"
	"github.com/google/uuid"
)

type FilterSP struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	ProblemId string    `json:"problem_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SolvedProblemRepository struct {
	db *sql.DB
}

func NewProbSolvRepo(db *sql.DB) *SolvedProblemRepository {
	return &SolvedProblemRepository{
		db: db,
	}
}

func (s SolvedProblemRepository) CreateSP(sp *models.SolvedProblem) error {
	query := `
		INSERT INTO solved_problems(id, user_id, problem_id)
		VALUES($1, $2, $3)
	`
	newId := uuid.NewString()

	if _, err := s.db.Exec(query, newId, sp.UserId, sp.ProblemId); err != nil {
		return fmt.Errorf("Executing query failed: " + err.Error())
	}
	return nil
}

func (s SolvedProblemRepository) UpdateSP(filter FilterSP) error {

}
