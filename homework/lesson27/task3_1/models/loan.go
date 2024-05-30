package models

import "time"

type Loan struct {
	LoanID     string
	BookID     string
	BorrowerID string
	LoanDate   time.Time
	ReturnDate time.Time
}
