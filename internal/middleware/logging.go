package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Logging(logger *log.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Call the next handler
			next.ServeHTTP(w, r)

			// Log the request
			logger.Printf(
				"%s %s %s %s",
				r.RemoteAddr,
				r.Method,
				r.RequestURI,
				time.Since(start),
			)
		})
	}
}

func Recovery(logger *log.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.Printf("panic: %v", err)
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
