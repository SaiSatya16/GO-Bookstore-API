package middleware

import (
	"bookstore-api/pkg/response"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// ErrorHandler wraps an http.Handler and handles any panics
func ErrorHandler(logger *log.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.Printf("panic: %v", err)
					response.Error(w, http.StatusInternalServerError, "Internal server error")
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

// NotFoundHandler handles 404 errors
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	response.Error(w, http.StatusNotFound, "Resource not found")
}

// MethodNotAllowedHandler handles 405 errors
func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	response.Error(w, http.StatusMethodNotAllowed, "Method not allowed")
}
