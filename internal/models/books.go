package models

import (
	"database/sql"
)

// Book struct now lives here
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// BookModel wraps the connection pool
type BookModel struct {
	DB *sql.DB
}

// Get all books
func (m *BookModel) GetAll() ([]Book, error) {
	rows, err := m.DB.Query("SELECT id, title, author FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

// Insert a new book
func (m *BookModel) Insert(title, author string) (int, error) {
	stmt := "INSERT INTO books(title, author) VALUES(?, ?)"
	result, err := m.DB.Exec(stmt, title, author)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
