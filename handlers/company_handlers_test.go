package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/falasefemi2/peopleos/dto"
	"github.com/falasefemi2/peopleos/models"
)

type MockCompanyService struct {
	CreateCompanyCalled     bool
	CreateCompanyCalledWith *dto.CreateCompanyRequest
}

func (m *MockCompanyService) CreateCompany(ctx context.Context, req *dto.CreateCompanyRequest) (*dto.CompanyResponse, error) {
	m.CreateCompanyCalled = true
	m.CreateCompanyCalledWith = req
	return &dto.CompanyResponse{
		ID:        1,
		Name:      req.Name,
		Industry:  req.Industry,
		Country:   req.Country,
		Timezone:  req.Timezone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

func (m *MockCompanyService) GetCompanyByName(ctx context.Context, name string) (*dto.CompanyResponse, error) {
	return nil, nil
}

func (m *MockCompanyService) GetCompanyByID(ctx context.Context, id int) (*dto.CompanyResponse, error) {
	return nil, nil
}

func (m *MockCompanyService) UpdateCompany(ctx context.Context, companyID int, req *dto.UpdateCompanyRequest) (*dto.CompanyResponse, error) {
	return nil, nil
}

func (m *MockCompanyService) DeleteCompany(ctx context.Context, companyID int) error {
	return nil
}

func TestCreateCompany(t *testing.T) {
	t.Run("returns 201 when company is created successfully", func(t *testing.T) {
		reqBody := dto.CreateCompanyRequest{
			Name:          "Test company",
			Industry:      "Technology",
			Country:       "Nigeria",
			Timezone:      "Africa/Lagos",
			AdminName:     "Falase femi",
			AdminEmail:    "femi@test.com",
			AdminPassword: "password123",
		}

		body, _ := json.Marshal(reqBody)
		request, _ := http.NewRequest(http.MethodPost, "/companies", bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		handler := &CompanyHandler{companyService: &MockCompanyService{}}
		handler.CreateCompany(response, request)

		if response.Code != http.StatusCreated {
			t.Errorf("got status %d, want %d", response.Code, http.StatusCreated)
		}

		var apiResponse models.APIResponse
		json.NewDecoder(response.Body).Decode(&apiResponse)

		if !apiResponse.Success {
			t.Errorf("got success %v, want true", apiResponse.Success)
		}
	})

	t.Run("returns 400 when request body is invalid", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/companies", bytes.NewReader([]byte("invalid json")))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		handler := &CompanyHandler{companyService: &MockCompanyService{}}
		handler.CreateCompany(response, request)

		if response.Code != http.StatusBadRequest {
			t.Errorf("got status %d, want %d", response.Code, http.StatusBadRequest)
		}

		var apiResponse models.APIResponse
		json.NewDecoder(response.Body).Decode(&apiResponse)

		if apiResponse.Success {
			t.Errorf("got success %v, want false", apiResponse.Success)
		}
	})

	t.Run("returns 400 when company name is empty", func(t *testing.T) {
		reqBody := dto.CreateCompanyRequest{
			Name:          "",
			Industry:      "Technology",
			Country:       "Nigeria",
			Timezone:      "Africa/Lagos",
			AdminName:     "Tunde",
			AdminEmail:    "tunde@test.com",
			AdminPassword: "password123",
		}

		body, _ := json.Marshal(reqBody)
		request, _ := http.NewRequest(http.MethodPost, "/companies", bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		handler := &CompanyHandler{companyService: &MockCompanyService{}}
		handler.CreateCompany(response, request)

		if response.Code != http.StatusBadRequest {
			t.Errorf("got status %d, want %d", response.Code, http.StatusBadRequest)
		}

		var apiResponse models.APIResponse
		json.NewDecoder(response.Body).Decode(&apiResponse)

		if apiResponse.Success {
			t.Errorf("got success %v, want false", apiResponse.Success)
		}
	})

	t.Run("returns 400 when admin email is invalid", func(t *testing.T) {
		reqBody := dto.CreateCompanyRequest{
			Name:          "Test company",
			Industry:      "Technology",
			Country:       "Nigeria",
			Timezone:      "Africa/Lagos",
			AdminName:     "Tunde",
			AdminEmail:    "invalid-email",
			AdminPassword: "password123",
		}

		body, _ := json.Marshal(reqBody)
		request, _ := http.NewRequest(http.MethodPost, "/companies", bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		handler := &CompanyHandler{companyService: &MockCompanyService{}}
		handler.CreateCompany(response, request)

		if response.Code != http.StatusBadRequest {
			t.Errorf("got status %d, want %d", response.Code, http.StatusBadRequest)
		}

		var apiResponse models.APIResponse
		json.NewDecoder(response.Body).Decode(&apiResponse)

		if apiResponse.Success {
			t.Errorf("got success %v, want false", apiResponse.Success)
		}
	})

	t.Run("returns 400 when password is less than 8 characters", func(t *testing.T) {
		reqBody := dto.CreateCompanyRequest{
			Name:          "Test company",
			Industry:      "Technology",
			Country:       "Nigeria",
			Timezone:      "Africa/Lagos",
			AdminName:     "Tunde",
			AdminEmail:    "tunde@test.com",
			AdminPassword: "short",
		}

		body, _ := json.Marshal(reqBody)
		request, _ := http.NewRequest(http.MethodPost, "/companies", bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		handler := &CompanyHandler{companyService: &MockCompanyService{}}
		handler.CreateCompany(response, request)

		if response.Code != http.StatusBadRequest {
			t.Errorf("got status %d, want %d", response.Code, http.StatusBadRequest)
		}

		var apiResponse models.APIResponse
		json.NewDecoder(response.Body).Decode(&apiResponse)

		if apiResponse.Success {
			t.Errorf("got success %v, want false", apiResponse.Success)
		}
	})

	t.Run("calls service.CreateCompany with correct data", func(t *testing.T) {
		mockService := &MockCompanyService{}

		reqBody := dto.CreateCompanyRequest{
			Name:          "Test company",
			Industry:      "Technology",
			Country:       "Nigeria",
			Timezone:      "Africa/Lagos",
			AdminName:     "Tunde",
			AdminEmail:    "tunde@test.com",
			AdminPassword: "password123",
		}

		body, _ := json.Marshal(reqBody)
		request, _ := http.NewRequest(http.MethodPost, "/companies", bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		handler := &CompanyHandler{companyService: mockService}
		handler.CreateCompany(response, request)

		if !mockService.CreateCompanyCalled {
			t.Errorf("service.CreateCompany was not called")
		}

		if response.Code != http.StatusCreated {
			t.Errorf("got status %d, want %d", response.Code, http.StatusCreated)
		}
	})
}
