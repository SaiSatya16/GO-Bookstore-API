package config

import (
    "fmt"
    "github.com/jmoiron/sqlx"
    _ "github.com/mattn/go-sqlite3"
)

type DBConfig struct {
    DBPath string
}

func NewDBConfig() *DBConfig {
    return &DBConfig{
        DBPath: "bookstore.db",
    }
}

func (c *DBConfig) Connect() (*sqlx.DB, error) {
    db, err := sqlx.Connect("sqlite3", c.DBPath)
    if err != nil {
        return nil, fmt.Errorf("error connecting to the database: %v", err)
    }

    // Enable foreign keys
    _, err = db.Exec("PRAGMA foreign_keys = ON")
    if err != nil {
        return nil, fmt.Errorf("error enabling foreign keys: %v", err)
    }

    // Configure connection pool
    db.SetMaxOpenConns(1) // SQLite only supports one writer at a time
    db.SetMaxIdleConns(1)

    return db, nil
}