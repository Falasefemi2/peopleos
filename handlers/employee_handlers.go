package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/falasefemi2/peopleos/dto"
	"github.com/falasefemi2/peopleos/models"
	"github.com/falasefemi2/peopleos/services"
)

type EmployeeHandler struct {
	employeeService services.IEmployeeService
}

func NewEmployeeHandler(employeeService services.IEmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		employeeService: employeeService,
	}
}

func (eh *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req dto.CreateEmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "Invalid request body",
		})
		return
	}

	if strings.TrimSpace(req.Email) == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "Email is required",
		})
		return
	}

	if strings.TrimSpace(req.FirstName) == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "First name is required",
		})
		return
	}

	if len(req.Password) < 8 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "Password must be at least 8 characters",
		})
		return
	}

	employee, err := eh.employeeService.CreateEmployee(r.Context(), &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Message: "Employee created successfully",
		Data:    employee,
	})
}
