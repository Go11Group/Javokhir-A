package models

import "time"

type SolvedProblem struct {
	Id        string    `json:"id"`
	ProblemId string    `json:"problem_id"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_ate"`
	DeletedAt time.Time `json:"deleted_at"`
}
