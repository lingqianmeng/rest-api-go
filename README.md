[![Go Bookstore CI](https://github.com/lingqianmeng/rest-api-go/actions/workflows/go.yml/badge.svg)](https://github.com/lingqianmeng/rest-api-go/actions/workflows/go.yml)
# Go Bookstore REST API (Pure Go SQLite)

A robust, beginner-friendly REST API built with Go. This project demonstrates how to structure a professional Go application, handle database persistence without external C-compilers, and separate concerns using the Standard Go Project Layout.

## 📖 Project Context & Architecture

### 1. The Architecture
Moved away from a single `main.go` file to a scalable structure:
* **`cmd/api/`**: Contains the entry point (`main.go`) and HTTP handlers.
* **`internal/models/`**: Contains the database logic. This separates "how we store data" from "how we serve data."

### 2. The Database Driver
**Critical Decision:** We use `github.com/glebarez/go-sqlite` instead of the standard `mattn/go-sqlite3`.
* **Why?** The standard driver requires a C-compiler (GCC), which causes frequent errors on Windows. The `glebarez` driver is "Pure Go," meaning it compiles instantly on Windows, Mac, and Linux with zero setup.

## 📂 Directory Structure
```text
go-rest-api/
├── cmd/
│   └── api/
│       └── main.go       # Application entry point & HTTP handlers
├── internal/
│   └── models/
│       └── books.go      # DB logic (Insert, GetAll)
├── .gitignore            # Ignores books.db and binaries
├── go.mod                # Module definition
├── go.sum                # Dependency checksums
├── books.db              # (Created automatically on run)
└── README.md             # This file
```

## 🚀 Getting Started
### 1. Prerequisites
Go installed (version 1.20+ recommended).

No other tools required (SQLite is embedded).

### 2. Installation
Clone the repo and install dependencies.

Bash
```
# Initialize the module (if starting from scratch)
go mod init bookstore-api

# Tidy up dependencies (downloads the pure-go sqlite driver)
go mod tidy
```
### 3. Running the Server
Because the main file is inside cmd/api, use this command:

Bash
```
go run ./cmd/api
```
You should see:
```
Server starting at :8080...
```

## 🔌 API Documentation
### 1. Get All Books (GET)
Retrieves the list of books from the SQLite database.

Endpoint: `/books`

Method: `GET`

PowerShell Command:
```
Invoke-RestMethod -Uri "http://localhost:8080/books"
```
cURL Command:
```
curl http://localhost:8080/books
```
### 2. Add a New Book (POST)
Inserts a book into the database. The ID is auto-generated.

Endpoint: `/books`

Method: `POST`

Headers: `Content-Type: application/json`

Body: `{"title": "String", "author": "String"}`

PowerShell Command:
```
Invoke-RestMethod -Method Post -Uri "http://localhost:8080/books" `
  -Body '{"title":"The Alchemist", "author":"Paulo Coelho"}' `
  -ContentType "application/json"
```
  
## 🛠️ Development Notes
### Database Persistence
The application automatically creates a file named `books.db` in the root directory if it does not exist.

**Do not commit books.db to Git. It is your local development data.**

### Git Ignore Rules
Ensure your `.gitignore` file contains the following to keep the repo clean:
```
# Ignore the local database
*.db

# Ignore build artifacts
*.exe
*.test
```
## Common Issues
### 1. Error: `exec: "gcc": executable file not found`

Fix: You are likely using `mattn/go-sqlite3`. Switch to `glebarez/go-sqlite` in your imports and run `go mod tidy`.

### 2. Error: `package cmd/api is not in GOROOT`

Fix: Ensure you are running `go run ./cmd/api` from the root folder, not inside the cmd folder itself.

## 📄 License
Open Source.
