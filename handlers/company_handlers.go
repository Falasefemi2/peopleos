package handlers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/falasefemi2/peopleos/dto"
	"github.com/falasefemi2/peopleos/models"
	"github.com/falasefemi2/peopleos/services"
)

type CompanyHandler struct {
	companyService services.ICompanyService
}

func NewCompanyHandler(companyService services.ICompanyService) *CompanyHandler {
	return &CompanyHandler{
		companyService: companyService,
	}
}

func (ch *CompanyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req dto.CreateCompanyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "Invalid request body",
		})
		return
	}

	if err := validateCreateCompanyRequest(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	companyResponse, err := ch.companyService.CreateCompany(r.Context(), &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Message: "Company created successfully",
		Data:    companyResponse,
	})
}

func validateCreateCompanyRequest(req *dto.CreateCompanyRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return &ValidationError{Field: "name", Message: "Company name is required"}
	}

	if strings.TrimSpace(req.Country) == "" {
		return &ValidationError{Field: "country", Message: "Country is required"}
	}

	if strings.TrimSpace(req.Timezone) == "" {
		return &ValidationError{Field: "timezone", Message: "Timezone is required"}
	}

	if strings.TrimSpace(req.AdminName) == "" {
		return &ValidationError{Field: "admin_name", Message: "Admin name is required"}
	}

	if strings.TrimSpace(req.AdminEmail) == "" {
		return &ValidationError{Field: "admin_email", Message: "Admin email is required"}
	}

	if !isValidEmail(req.AdminEmail) {
		return &ValidationError{Field: "admin_email", Message: "Admin email is invalid"}
	}

	if len(req.AdminPassword) < 8 {
		return &ValidationError{Field: "admin_password", Message: "Password must be at least 8 characters"}
	}

	return nil
}

func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}

type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func (ch *CompanyHandler) GetCompanyByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "Company name is required",
		})
		return
	}

	company, err := ch.companyService.GetCompanyByName(r.Context(), name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "Company not found",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Message: "Company found",
		Data:    company,
	})
}

func (ch *CompanyHandler) GetCompanyByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars["id"]

	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "Company ID is required",
		})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "Invalid company ID",
		})
		return
	}

	company, err := ch.companyService.GetCompanyByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "Company not found",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Message: "Company found",
		Data:    company,
	})
}

func (ch *CompanyHandler) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars["id"]

	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "Company ID is required",
		})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "Invalid company ID",
		})
		return
	}

	var req dto.UpdateCompanyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "Invalid request body",
		})
		return
	}

	company, err := ch.companyService.UpdateCompany(r.Context(), id, &req)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "Company not found",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Message: "Company updated successfully",
		Data:    company,
	})
}

func (ch *CompanyHandler) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idStr := vars["id"]

	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "Company ID is required",
		})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "Invalid company ID",
		})
		return
	}

	err = ch.companyService.DeleteCompany(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "Company not found",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Message: "Company deleted successfully",
	})
}
