package handlers

import (
	"net/http"
	"strings"

	"github.com/falasefemi2/peopleos/dto"
	"github.com/falasefemi2/peopleos/services"
	"github.com/falasefemi2/peopleos/utils"
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
	var req dto.CreateCompanyRequest
	if err := utils.DecodeJSONBody(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := validateCreateCompanyRequest(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	companyResponse, err := ch.companyService.CreateCompany(r.Context(), &req)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, utils.APIResponse{
		Success: true,
		Message: "Company created successfully",
		Data:    companyResponse,
	})
}

func validateCreateCompanyRequest(req *dto.CreateCompanyRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return &utils.ValidationError{Field: "name", Message: "Company name is required"}
	}

	if strings.TrimSpace(req.Country) == "" {
		return &utils.ValidationError{Field: "country", Message: "Country is required"}
	}

	if strings.TrimSpace(req.Timezone) == "" {
		return &utils.ValidationError{Field: "timezone", Message: "Timezone is required"}
	}

	if strings.TrimSpace(req.AdminName) == "" {
		return &utils.ValidationError{Field: "admin_name", Message: "Admin name is required"}
	}

	if strings.TrimSpace(req.AdminEmail) == "" {
		return &utils.ValidationError{Field: "admin_email", Message: "Admin email is required"}
	}

	if !utils.IsValidEmail(req.AdminEmail) {
		return &utils.ValidationError{Field: "admin_email", Message: "Admin email is invalid"}
	}

	if len(req.AdminPassword) < 8 {
		return &utils.ValidationError{Field: "admin_password", Message: "Password must be at least 8 characters"}
	}

	return nil
}

func (ch *CompanyHandler) GetCompanyByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Company name is required")
		return
	}

	company, err := ch.companyService.GetCompanyByName(r.Context(), name)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Company not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, utils.APIResponse{
		Success: true,
		Message: "Company found",
		Data:    company,
	})
}

func (ch *CompanyHandler) GetCompanyByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIntParam(r, "id")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid company ID")
		return
	}

	company, err := ch.companyService.GetCompanyByID(r.Context(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Company not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, utils.APIResponse{
		Success: true,
		Message: "Company found",
		Data:    company,
	})
}

func (ch *CompanyHandler) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIntParam(r, "id")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid company ID")
		return
	}

	var req dto.UpdateCompanyRequest
	if err := utils.DecodeJSONBody(r, &req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	company, err := ch.companyService.UpdateCompany(r.Context(), id, &req)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Company not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, utils.APIResponse{
		Success: true,
		Message: "Company updated successfully",
		Data:    company,
	})
}

func (ch *CompanyHandler) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseIntParam(r, "id")
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid company ID")
		return
	}

	err = ch.companyService.DeleteCompany(r.Context(), id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Company not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, utils.APIResponse{
		Success: true,
		Message: "Company deleted successfully",
	})
}
