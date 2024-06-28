package repositories

import (
	"context"
	"database/sql"
)

type CreateBook struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	YearPublished string `json:"year_published"`
}

type BookRepository interface {
	CreateBook(ctx context.Context, user *CreateBook)
}

type bookRepository struct {
	Db *sql.DB
}

func NewBookRepository(db *sql.DB) *bookRepository {
	return &bookRepository{Db: db}
}

func (l *bookRepository) CreateBook(ctx context.Context, user *CreateBook) {
    _, err := l.Db.ExecContext(ctx, "INSERT INTO books (title, author, year_published) VALUES ($1, $2, $3)", user.Title, user.Author, user.YearPublished)s
}
