package models

import "errors"

type Book struct {
    ID        string  `json:"id" db:"id"`
    Title     string  `json:"title" db:"title"`
    Author    string  `json:"author" db:"author"`
    Price     float64 `json:"price" db:"price"`
    CreatedAt string  `json:"created_at,omitempty" db:"created_at"`
}

func (b *Book) Validate() error {
    if b.ID == "" {
        return errors.New("id is required")
    }
    if b.Title == "" {
        return errors.New("title is required")
    }
    if b.Author == "" {
        return errors.New("author is required")
    }
    if b.Price <= 0 {
        return errors.New("price must be greater than 0")
    }
    return nil
}