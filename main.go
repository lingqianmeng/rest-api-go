package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Book represents our data model
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// In-memory database
var books = []Book{
	{ID: "1", Title: "The Go Programming Language", Author: "Alan Donovan"},
}

// Handler: Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Handler: Create a book
func createBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newBook Book
	json.NewDecoder(r.Body).Decode(&newBook)
	books = append(books, newBook)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newBook)
}

func main() {
	// Routing
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getBooks(w, r)
		case http.MethodPost:
			createBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server starting at :8080...")
	http.ListenAndServe(":8080", nil)
}
