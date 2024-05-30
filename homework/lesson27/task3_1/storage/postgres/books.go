package postgres

import (
	"database/sql"
	"fmt"
	"librarySystem/models"
)

type BookRepository struct {
	Db *sql.DB
}

func NewBookRepo(db *sql.DB) *BookRepository {
	return &BookRepository{Db: db}
}

func (b *BookRepository) GetAllBooks() ([]models.Book, error) {
	query := `
        SELECT BookId, Title, Author, Genre, PublishedYear FROM books
        ORDER BY Title
    `

	rows, err := b.Db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed while quering %w", err)
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Genre, &book.PublishedYear)
		if err != nil {
			return nil, fmt.Errorf("faild while scanning %w", err)
		}
		books = append(books, book)
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed while iterating over rows %w", err)
	}

	return books, nil
}

func (b *BookRepository) Create(book models.Book) error {

	query := `
	INSERT INTO books (
		Title, 
		Author,
		Genre, 
		PublishedYear) 
	VALUES($1, $2, $3, $4)
	`

	_, err := b.Db.Exec(query, book.Title, book.Author, book.Genre, book.PublishedYear)
	if err != nil {
		return fmt.Errorf("failed to create book %w", err)
	}

	return nil
}
