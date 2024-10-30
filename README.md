# Bookstore API

A simple RESTful API for managing books using Go and SQLite.

## Features

- CRUD operations for books
- SQLite database
- RESTful API design
- Middleware for logging and error handling
- Graceful shutdown
- Input validation
- JSON response formatting

## Prerequisites

- Go 1.16 or higher
- SQLite3

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/bookstore-api
cd bookstore-api
```

2. Install dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run cmd/api/main.go
```

## API Endpoints
```bash
| Method | Path               | Description        |
|--------|---------------     |--------------------|
| POST   | /api/v1/books      | Create a new book  |
| GET    | /api/v1/books      | Get all books      |
| GET    | /api/v1/books/{id} | Get a book by ID   |
| PUT    | /api/v1/books/{id} | Update a book      |
| DELETE | /api/v1/books/{id} | Delete a book      |
```
## API Usage

### Create a Book
```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"id":"1", "title":"The Go Programming Language", "author":"Alan A. A. Donovan", "price":49.99}' \
  http://localhost:8080/api/v1/books
```

### Get All Books
```bash
curl http://localhost:8080/api/v1/books
```

### Get a Book
```bash
curl http://localhost:8080/api/v1/books/1
```

### Update a Book
```bash
curl -X PUT -H "Content-Type: application/json" \
  -d '{"title":"The Go Programming Language", "author":"Alan A. A. Donovan", "price":39.99}' \
  http://localhost:8080/api/v1/books/1
```

### Delete a Book
```bash
curl -X DELETE http://localhost:8080/api/v1/books/1
```

## Project Structure

```bash
bookstore-api/
├── cmd/
│   └── api/
│       └── main.go            # Application entry point
├── internal/
│   ├── config/
│   │   └── database.go        # Database configuration
│   ├── handlers/
│   │   └── book_handler.go    # HTTP handlers
│   ├── middleware/
│   │   ├── logging.go         # Logging middleware
│   │   └── error_handler.go   # Error handling middleware
│   ├── models/
│   │   └── book.go           # Data models
│   └── repository/
│       └── book_repository.go # Database operations
├── pkg/
│   ├── response/
│   │   └── json.go           # Response helpers
│   └── validator/
│       └── validator.go      # Input validation
└── README.md
```

## Error Handling

The API returns errors in the following format:
```json
{
    "success": false,
    "error": "Error message"
}
```

## Success Response

Successful responses are returned in the following format:
```json
{
    "success": true,
    "data": {
        // Response data
    }
}
```