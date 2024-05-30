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
		bookid,
		Title, 
		Author,
		Genre, 
		PublishedYear) 
	VALUES($1, $2, $3, $4, $5)
	`

	_, err := b.Db.Exec(query, book.Id, book.Title, book.Author, book.Genre, book.PublishedYear)
	if err != nil {
		return fmt.Errorf("failed to create book %w", err)
	}

	return nil
}

func (b *BookRepository) Update(updatedBook models.Book) error {

	query := `
		UPDATE books 
		SET Title = $1, 
		    Author = $2, 
		    Genre = $3, 
		    PublishedYear = $4 
		WHERE id = $5
	`

	_, err := b.Db.Exec(query, updatedBook.Title, updatedBook.Author, updatedBook.Genre, updatedBook.PublishedYear, updatedBook.Id)
	if err != nil {
		return fmt.Errorf("failed to update book: %w", err)
	}

	return nil
}

func (b *BookRepository) Delete(deletingBook models.Book) error {

	query := `
		DELETE FROM books WHERE bookid = $1
	`

	_, err := b.Db.Exec(query, deletingBook.Id)

	if err != nil {
		return fmt.Errorf("failed while deleting book %w", err)
	}

	return nil
}

func (b *BookRepository) GetById(bookId string) (*models.Book, error) {
	// Prepare the SQL query using a parameterized query
	query := `
		SELECT bookid, title, author, genre, publishedyear 
		FROM books 
		WHERE bookid = $1
	`

	// Create an empty Book object to hold the result
	book := models.Book{}

	// Execute the query with the bookId as the parameter
	row := b.Db.QueryRow(query, bookId)

	// Scan the result into the Book object
	err := row.Scan(&book.Id, &book.Title, &book.Author, &book.Genre, &book.PublishedYear)
	if err != nil {
		return nil, fmt.Errorf("failed while scanning book: %w", err)
	}

	return &book, nil
}
