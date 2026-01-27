package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/falasefemi2/peopleos/config"
	"github.com/falasefemi2/peopleos/database"
	"github.com/falasefemi2/peopleos/handlers"
	"github.com/falasefemi2/peopleos/middleware"
)

func main() {
	fmt.Println("Connecting to database...")
	pool, err := config.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer pool.Close()

	fmt.Println("Running migrations...")
	if err := database.RunMigrations(pool); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	router := mux.NewRouter()

	handler := middleware.ChainMiddleware(
		router,
		middleware.RecoveryMiddleware,
		middleware.LoggingMiddleware,
		middleware.CORSMiddleware,
	)

	fmt.Println("Registering routes...")
	router.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/companies", handlers.CreateCompany).Methods("POST")

	port := ":8080"
	fmt.Printf("\n✓ Server starting on http://localhost%s\n", port)
	fmt.Println("Press Ctrl+C to stop the server\n")

	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
