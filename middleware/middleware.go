package middleware

import (
<<<<<<< HEAD
	"context"
	"encoding/json"
	"fmt"
=======
>>>>>>> c84adbd5fea15cfee43772c5a62f177c37a8ebec
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"time"
<<<<<<< HEAD

	"github.com/dgrijalva/jwt-go"
=======
	"github.com/falasefemi2/peopleos/utils"
>>>>>>> c84adbd5fea15cfee43772c5a62f177c37a8ebec
)

type contextKey string

const userContextKey = contextKey("user")

// Claims represents the JWT claims
type Claims struct {
	ID   int    `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}

var jwtKey = []byte("your-secret-key") // Replace with a secure key in production

// LoggingMiddleware logs all HTTP requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
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
<<<<<<< HEAD
				fmt.Fprintf(w, `{"success":false,"error":"Internal server error","code":500}`)
=======

				// In production, don't expose stack traces
				// fmt.Fprintf(w, `{"success":false,"error":"Internal server error","code":500}`)
				utils.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
>>>>>>> c84adbd5fea15cfee43772c5a62f177c37a8ebec
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
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// AuthenticationMiddleware verifies the JWT token
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"success":false,"error":"Authorization header required"}`, http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, `{"success":false,"error":"Invalid token"}`, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RoleMiddleware checks for a specific role
func RoleMiddleware(role string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value(userContextKey).(*Claims)
			if !ok {
				http.Error(w, `{"success":false,"error":"Could not retrieve user claims"}`, http.StatusInternalServerError)
				return
			}

			if claims.Role != role {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"success": false,
					"error":   fmt.Sprintf("Forbidden: You must be a %s to access this resource", role),
					"code":    http.StatusForbidden,
				})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// ChainMiddleware chains multiple middlewares
func ChainMiddleware(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}
