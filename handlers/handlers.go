package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/falasefemi2/peopleos/models"
)

// HealthCheck is a simple health check endpoint
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := models.APIResponse{
		Success: true,
		Message: "Server is running",
		Data: map[string]string{
			"status": "ok",
			"time":   "2026-01-26",
		},
	}

	json.NewEncoder(w).Encode(response)
}

// AdminHandler is a placeholder for an admin-only endpoint
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := models.APIResponse{
		Success: true,
		Message: "Welcome, Super Admin!",
	}

	json.NewEncoder(w).Encode(response)
}
