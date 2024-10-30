package repository

import (
    "fmt"
    "time"
    "github.com/jmoiron/sqlx"
    "bookstore-api/internal/models"
)

type BookRepository struct {
    db *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) *BookRepository {
    return &BookRepository{db: db}
}

func (r *BookRepository) Initialize() error {
    schema := `
        CREATE TABLE IF NOT EXISTS books (
            id TEXT PRIMARY KEY,
            title TEXT NOT NULL,
            author TEXT NOT NULL,
            price REAL NOT NULL,
            created_at TEXT NOT NULL
        );`

    _, err := r.db.Exec(schema)
    return err
}

func (r *BookRepository) Create(book *models.Book) error {
    query := `
        INSERT INTO books (id, title, author, price, created_at)
        VALUES (?, ?, ?, ?, ?)`

    book.CreatedAt = time.Now().Format(time.RFC3339)
    
    result, err := r.db.Exec(query, book.ID, book.Title, book.Author, book.Price, book.CreatedAt)
    if err != nil {
        return fmt.Errorf("error creating book: %v", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error checking affected rows: %v", err)
    }

    if rows == 0 {
        return fmt.Errorf("no rows affected")
    }

    return nil
}

func (r *BookRepository) GetByID(id string) (*models.Book, error) {
    var book models.Book
    
    query := `SELECT * FROM books WHERE id = ?`

    err := r.db.Get(&book, query, id)
    if err != nil {
        return nil, fmt.Errorf("error getting book: %v", err)
    }

    return &book, nil
}

func (r *BookRepository) GetAll() ([]models.Book, error) {
    var books []models.Book
    
    query := `SELECT * FROM books ORDER BY created_at DESC`

    err := r.db.Select(&books, query)
    if err != nil {
        return nil, fmt.Errorf("error getting books: %v", err)
    }

    return books, nil
}

func (r *BookRepository) Update(book *models.Book) error {
    query := `
        UPDATE books
        SET title = ?, author = ?, price = ?
        WHERE id = ?`

    result, err := r.db.Exec(query, book.Title, book.Author, book.Price, book.ID)
    if err != nil {
        return fmt.Errorf("error updating book: %v", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error checking affected rows: %v", err)
    }

    if rows == 0 {
        return fmt.Errorf("book not found")
    }

    return nil
}

func (r *BookRepository) Delete(id string) error {
    query := `DELETE FROM books WHERE id = ?`

    result, err := r.db.Exec(query, id)
    if err != nil {
        return fmt.Errorf("error deleting book: %v", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("error checking affected rows: %v", err)
    }

    if rows == 0 {
        return fmt.Errorf("book not found")
    }

    return nil
}