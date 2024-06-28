package models

type Book struct {
	BookId        string `json:"book_id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	YearPublished string `json:"year_published"`
}
