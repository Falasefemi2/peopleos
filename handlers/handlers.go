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

// CreateCompany creates a new company (placeholder)
func CreateCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Success: false,
			Error:   "Method not allowed",
			Code:    http.StatusMethodNotAllowed,
		})
		return
	}

	var req models.CreateCompanyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.ErrorResponse{
			Success: false,
			Error:   "Invalid request body",
			Code:    http.StatusBadRequest,
		})
		return
	}

	// TODO: Implement company creation logic
	w.WriteHeader(http.StatusCreated)
	response := models.APIResponse{
		Success: true,
		Message: "Company creation endpoint ready (implementation coming)",
	}
	json.NewEncoder(w).Encode(response)
}
