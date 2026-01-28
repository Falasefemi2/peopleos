package handlers

import (
	"net/http"

<<<<<<< HEAD
	"github.com/falasefemi2/peopleos/models"
=======
	"github.com/falasefemi2/peopleos/dto"
	"github.com/falasefemi2/peopleos/utils"
>>>>>>> c84adbd5fea15cfee43772c5a62f177c37a8ebec
)

// HealthCheck is a simple health check endpoint
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, utils.APIResponse{
		Success: true,
		Message: "Server is running",
		Data: map[string]string{
			"status": "ok",
			"time":   "2026-01-26",
		},
	})
}

<<<<<<< HEAD
// AdminHandler is a placeholder for an admin-only endpoint
func AdminHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := models.APIResponse{
		Success: true,
		Message: "Welcome, Super Admin!",
	}

	json.NewEncoder(w).Encode(response)
=======
// CreateCompany creates a new company (placeholder)
func CreateCompany(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var req dto.CreateCompanyRequest
	if err := utils.DecodeJSONBody(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// TODO: Implement company creation logic
	utils.RespondWithJSON(w, http.StatusCreated, utils.APIResponse{
		Success: true,
		Message: "Company creation endpoint ready (implementation coming)",
	})
>>>>>>> c84adbd5fea15cfee43772c5a62f177c37a8ebec
}
