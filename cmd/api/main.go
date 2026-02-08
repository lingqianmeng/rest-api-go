package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/glebarez/go-sqlite"

	// Import own internal package
	"bookstore-api/internal/models"
)

// Application struct holds your dependencies (like the database model)
type application struct {
	books *models.BookModel
}

// In-memory database
/*
var books = []Book{
	{ID: "1", Title: "The Go Programming Language", Author: "Alan Donovan"},
}
*/

// This is now a method of 'application', so it can access 'app.books'
func (app *application) booksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.getBooks(w, r)
	case http.MethodPost:
		app.createBook(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (app *application) getBooks(w http.ResponseWriter, r *http.Request) {
	// Call the model! No SQL here.
	books, err := app.books.GetAll()
	if err != nil {
		http.Error(w, "Server Error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (app *application) createBook(w http.ResponseWriter, r *http.Request) {
	var input models.Book
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Bad Request", 400)
		return
	}

	// Call the model!
	id, err := app.books.Insert(input.Title, input.Author)
	if err != nil {
		http.Error(w, "Server Error", 500)
		return
	}

	input.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(input)
}

func main() {
	// 1. Init Database
	db, err := sql.Open("sqlite", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 2. Setup the model
	// We check if the table exists here for simplicity
	_, _ = db.Exec("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, title TEXT, author TEXT);")

	// 3. Initialize the app struct
	app := &application{
		books: &models.BookModel{DB: db},
	}

	// 4. Start Server
	fmt.Println("Server starting at :8080...")
	http.HandleFunc("/books", app.booksHandler) // Use the method on 'app'
	log.Fatal(http.ListenAndServe(":8080", nil))
}
