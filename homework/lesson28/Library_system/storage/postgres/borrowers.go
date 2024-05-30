package postgres

import (
	"database/sql"
	"fmt"
	"librarySystem/models"
)

type BorrowersRepository struct {
	Db *sql.DB
}

func NewBorrowersRepo(db *sql.DB) *BorrowersRepository {
	return &BorrowersRepository{Db: db}
}

func (b *BorrowersRepository) GetAllBorrowers() ([]models.Borrower, error) {
	query := `
        SELECT borrowerid, name, email, phone FROM borrowers
        ORDER BY name
    `

	rows, err := b.Db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed while querying: %w", err)
	}
	defer rows.Close()

	var borrowers []models.Borrower
	for rows.Next() {
		var borrower models.Borrower
		err := rows.Scan(&borrower.Id, &borrower.Name, &borrower.Email, &borrower.Phone)
		if err != nil {
			return nil, fmt.Errorf("failed while scanning: %w", err)
		}
		borrowers = append(borrowers, borrower)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed while iterating over rows: %w", err)
	}

	return borrowers, nil
}

func (b *BorrowersRepository) Create(borrower models.Borrower) error {
	query := `
	INSERT INTO borrowers (name, email, phone)
	VALUES ($1, $2, $3)
	`

	_, err := b.Db.Exec(query, borrower.Name, borrower.Email, borrower.Phone)
	if err != nil {
		return fmt.Errorf("failed to create borrower: %w", err)
	}

	return nil
}

func (b *BorrowersRepository) Update(updatedBorrower models.Borrower) error {
	query := `
		UPDATE borrowers 
		SET name = $1, 
		    email = $2, 
		    phone = $3 
		WHERE borrowerid = $4
	`

	_, err := b.Db.Exec(query, updatedBorrower.Name, updatedBorrower.Email, updatedBorrower.Phone, updatedBorrower.Id)
	if err != nil {
		return fmt.Errorf("failed to update borrower: %w", err)
	}

	return nil
}

func (b *BorrowersRepository) Delete(borrowerId string) error {
	query := `
		DELETE FROM borrowers WHERE borrowerid = $1
	`

	_, err := b.Db.Exec(query, borrowerId)
	if err != nil {
		return fmt.Errorf("failed to delete borrower: %w", err)
	}

	return nil
}

func (b *BorrowersRepository) GetById(borrowerId string) (*models.Borrower, error) {
	query := `
		SELECT borrowerid, name, email, phone 
		FROM borrowers 
		WHERE borrowerid = $1
	`

	borrower := models.Borrower{}

	row := b.Db.QueryRow(query, borrowerId)

	err := row.Scan(&borrower.Id, &borrower.Name, &borrower.Email, &borrower.Phone)
	if err != nil {
		return nil, fmt.Errorf("failed while scanning borrower: %w", err)
	}

	return &borrower, nil
}
