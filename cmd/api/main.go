package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bookstore-api/internal/config"
	"bookstore-api/internal/handlers"
	"bookstore-api/internal/middleware"
	"bookstore-api/internal/repository"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize logger
	logger := log.New(os.Stdout, "BOOKSTORE-API ", log.LstdFlags)

	// Initialize database
	dbConfig := config.NewDBConfig()
	db, err := dbConfig.Connect()
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	// Initialize repository
	bookRepo := repository.NewBookRepository(db)
	if err := bookRepo.Initialize(); err != nil {
		logger.Fatal(err)
	}

	// Initialize handlers
	bookHandler := handlers.NewBookHandler(bookRepo)

	// Initialize router
	router := mux.NewRouter()

	// Set custom error handlers
	router.NotFoundHandler = http.HandlerFunc(middleware.NotFoundHandler)
	router.MethodNotAllowedHandler = http.HandlerFunc(middleware.MethodNotAllowedHandler)

	// Apply global middleware
	router.Use(middleware.Logging(logger))
	router.Use(middleware.ErrorHandler(logger))

	// API routes
	api := router.PathPrefix("/api/v1").Subrouter()

	// Book routes
	api.HandleFunc("/books", bookHandler.Create).Methods(http.MethodPost)
	api.HandleFunc("/books", bookHandler.GetAll).Methods(http.MethodGet)
	api.HandleFunc("/books/{id}", bookHandler.GetByID).Methods(http.MethodGet)
	api.HandleFunc("/books/{id}", bookHandler.Update).Methods(http.MethodPut)
	api.HandleFunc("/books/{id}", bookHandler.Delete).Methods(http.MethodDelete)

	// Configure server
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start server
	go func() {
		logger.Printf("Starting server on port %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal(err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Println("Server shutting down...")
	srv.Close()
}
