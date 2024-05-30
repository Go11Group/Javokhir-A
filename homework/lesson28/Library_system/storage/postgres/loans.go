package postgres

import (
	"database/sql"
	"fmt"
	"librarySystem/models"
)

type LoansRepository struct {
	Db *sql.DB
}

func NewLoansRepo(db *sql.DB) *LoansRepository {
	return &LoansRepository{
		Db: db,
	}
}

func (l *LoansRepository) GetLoanById(loanId string) (*models.Loan, error) {

}

func (l *LoansRepository) Create(loan models.Loan) error {

	query := `INSERT INTO loans(bookid, borrowerid, loandate, returndate) 
	VALUES($1, $2, $3, $4)
	`

	_, err := l.Db.Exec(query, loan.BookID, loan.BorrowerID, loan.LoanDate, loan.ReturnDate)
	if err != nil {
		return fmt.Errorf("failed while creating %w", err)
	}

	return nil
}

func (l *LoansRepository) Update(loan models.Loan)
