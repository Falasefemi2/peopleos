package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
	"time"
	"github.com/falasefemi2/peopleos/utils"
)

// LoggingMiddleware logs all HTTP requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call next handler
		next.ServeHTTP(w, r)

		// Log after request completes
		duration := time.Since(start)
		log.Printf("[%s] %s %s - %v", r.Method, r.RequestURI, r.RemoteAddr, duration)
	})
}

// RecoveryMiddleware catches panics and returns a proper error response
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("PANIC: %v\n%s", err, debug.Stack())

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)

				// In production, don't expose stack traces
				// fmt.Fprintf(w, `{"success":false,"error":"Internal server error","code":500}`)
				utils.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// CORSMiddleware enables CORS
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// ChainMiddleware chains multiple middlewares
func ChainMiddleware(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, mw := range middlewares {
		h = mw(h)
	}
	return h
}
