package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"

	"github.com/falasefemi2/peopleos/dto"
	"github.com/falasefemi2/peopleos/models"
)

type MockCompanyService struct {
	CreateCompanyCalled     bool
	CreateCompanyCalledWith *dto.CreateCompanyRequest
	GetCompanyByNameResult  *dto.CompanyResponse
	GetCompanyByNameError   error
	GetCompanyByIDResult    *dto.CompanyResponse
	GetCompanyByIDError     error
	UpdateCompanyResult     *dto.CompanyResponse
	UpdateCompanyError      error
	DeleteCompanyError      error
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
	if m.GetCompanyByNameError != nil {
		return nil, m.GetCompanyByNameError
	}
	return m.GetCompanyByNameResult, nil
}

func (m *MockCompanyService) GetCompanyByID(ctx context.Context, id int) (*dto.CompanyResponse, error) {
	if m.GetCompanyByIDError != nil {
		return nil, m.GetCompanyByIDError
	}
	return m.GetCompanyByIDResult, nil
}

func (m *MockCompanyService) UpdateCompany(ctx context.Context, companyID int, req *dto.UpdateCompanyRequest) (*dto.CompanyResponse, error) {
	if m.UpdateCompanyError != nil {
		return nil, m.UpdateCompanyError
	}
	return m.UpdateCompanyResult, nil
}

func (m *MockCompanyService) DeleteCompany(ctx context.Context, companyID int) error {
	return m.DeleteCompanyError
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

func TestGetCompanyName(t *testing.T) {
	t.Run("returns 200 and company data when company exists", func(t *testing.T) {
		mockService := &MockCompanyService{
			GetCompanyByNameResult: &dto.CompanyResponse{
				ID:       1,
				Name:     "Acme Corp",
				Industry: "Technology",
				Country:  "Nigeria",
				Timezone: "Africa/Lagos",
			},
		}

		request, _ := http.NewRequest(http.MethodGet, "/companies/search?name=Acme%20Corp", nil)
		response := httptest.NewRecorder()

		handler := &CompanyHandler{companyService: mockService}
		handler.GetCompanyByName(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("got status %d, want %d", response.Code, http.StatusOK)
		}

		var apiResponse models.APIResponse
		json.NewDecoder(response.Body).Decode(&apiResponse)

		if !apiResponse.Success {
			t.Errorf("got success %v, want true", apiResponse.Success)
		}
	})

	t.Run("returns 400 when name query parameter is missing", func(t *testing.T) {
		mockService := &MockCompanyService{}

		request, _ := http.NewRequest(http.MethodGet, "/companies/search", nil)
		response := httptest.NewRecorder()

		handler := &CompanyHandler{companyService: mockService}
		handler.GetCompanyByName(response, request)

		if response.Code != http.StatusBadRequest {
			t.Errorf("got status %d, want %d", response.Code, http.StatusBadRequest)
		}

		var apiResponse models.APIResponse
		json.NewDecoder(response.Body).Decode(&apiResponse)

		if apiResponse.Success {
			t.Errorf("got success %v, want false", apiResponse.Success)
		}
	})

	t.Run("returns 404 when company not found", func(t *testing.T) {
		mockService := &MockCompanyService{
			GetCompanyByNameError: fmt.Errorf("not found"),
		}

		request, _ := http.NewRequest(http.MethodGet, "/companies/search?name=NonExistent", nil)
		response := httptest.NewRecorder()

		handler := &CompanyHandler{companyService: mockService}
		handler.GetCompanyByName(response, request)

		if response.Code != http.StatusNotFound {
			t.Errorf("got status %d, want %d", response.Code, http.StatusNotFound)
		}

		var apiResponse models.APIResponse
		json.NewDecoder(response.Body).Decode(&apiResponse)

		if apiResponse.Success {
			t.Errorf("got success %v, want false", apiResponse.Success)
		}
	})
}
func TestCompanyByID(t *testing.T) {
	t.Run("get company with ID not found", func(t *testing.T) {
		mockService := &MockCompanyService{
			GetCompanyByIDError: fmt.Errorf("not found"),
		}

		request, _ := http.NewRequest(http.MethodGet, "/companies/1", nil)
		request = mux.SetURLVars(request, map[string]string{
			"id": "1",
		})

		response := httptest.NewRecorder()
		handler := &CompanyHandler{companyService: mockService}
		handler.GetCompanyByID(response, request)

		if response.Code != http.StatusNotFound {
			t.Errorf("got status %d, want %d", response.Code, http.StatusNotFound)
		}

		var apiResponse models.APIResponse
		json.NewDecoder(response.Body).Decode(&apiResponse)

		if apiResponse.Success {
			t.Errorf("got success %v, want false", apiResponse.Success)
		}
	})

	t.Run("returns 200 and company data when company exists", func(t *testing.T) {
		mockService := &MockCompanyService{
			GetCompanyByIDResult: &dto.CompanyResponse{
				ID:       1,
				Name:     "Acme Corp",
				Industry: "Technology",
				Country:  "Nigeria",
				Timezone: "Africa/Lagos",
			},
		}

		request, _ := http.NewRequest(http.MethodGet, "/companies/1", nil)

		// Set the route variables manually for testing
		request = mux.SetURLVars(request, map[string]string{
			"id": "1",
		})

		response := httptest.NewRecorder()
		handler := &CompanyHandler{companyService: mockService}
		handler.GetCompanyByID(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("got status %d, want %d", response.Code, http.StatusOK)
		}

		var apiResponse models.APIResponse
		json.NewDecoder(response.Body).Decode(&apiResponse)

		if !apiResponse.Success {
			t.Errorf("got success %v, want true", apiResponse.Success)
		}
	})
}

func TestUpdateCompany(t *testing.T) {
	t.Run("update company name", func(t *testing.T) {
		mockService := &MockCompanyService{
			UpdateCompanyResult: &dto.CompanyResponse{
				ID:       1,
				Name:     "Updated Company",
				Industry: "Technology",
				Country:  "Nigeria",
				Timezone: "Africa/Lagos",
			},
		}

		reqBody := dto.UpdateCompanyRequest{
			Name:     "Updated Company",
			Industry: "Technology",
			Country:  "Nigeria",
			Timezone: "Africa/Lagos",
		}

		body, _ := json.Marshal(reqBody)
		request, _ := http.NewRequest(http.MethodPut, "/companies/1", bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")
		request = mux.SetURLVars(request, map[string]string{
			"id": "1",
		})

		response := httptest.NewRecorder()
		handler := &CompanyHandler{companyService: mockService}
		handler.UpdateCompany(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("got status %d, want %d", response.Code, http.StatusOK)
		}

		var apiResponse models.APIResponse
		json.NewDecoder(response.Body).Decode(&apiResponse)

		if !apiResponse.Success {
			t.Errorf("got success %v, want true", apiResponse.Success)
		}

		if apiResponse.Data == nil {
			t.Errorf("got data nil, want company response")
		}
	})

	t.Run("returns 400 when request body is invalid", func(t *testing.T) {
		mockService := &MockCompanyService{}

		request, _ := http.NewRequest(http.MethodPut, "/companies/1", bytes.NewReader([]byte("invalid json")))
		request.Header.Set("Content-Type", "application/json")
		request = mux.SetURLVars(request, map[string]string{
			"id": "1",
		})

		response := httptest.NewRecorder()
		handler := &CompanyHandler{companyService: mockService}
		handler.UpdateCompany(response, request)

		if response.Code != http.StatusBadRequest {
			t.Errorf("got status %d, want %d", response.Code, http.StatusBadRequest)
		}

		var apiResponse models.APIResponse
		json.NewDecoder(response.Body).Decode(&apiResponse)

		if apiResponse.Success {
			t.Errorf("got success %v, want false", apiResponse.Success)
		}
	})

	t.Run("returns 404 when company not found", func(t *testing.T) {
		mockService := &MockCompanyService{
			UpdateCompanyError: fmt.Errorf("not found"),
		}

		reqBody := dto.UpdateCompanyRequest{
			Name:     "Updated Company",
			Industry: "Technology",
			Country:  "Nigeria",
			Timezone: "Africa/Lagos",
		}

		body, _ := json.Marshal(reqBody)
		request, _ := http.NewRequest(http.MethodPut, "/companies/999", bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")
		request = mux.SetURLVars(request, map[string]string{
			"id": "999",
		})

		response := httptest.NewRecorder()
		handler := &CompanyHandler{companyService: mockService}
		handler.UpdateCompany(response, request)

		if response.Code != http.StatusNotFound {
			t.Errorf("got status %d, want %d", response.Code, http.StatusNotFound)
		}

		var apiResponse models.APIResponse
		json.NewDecoder(response.Body).Decode(&apiResponse)

		if apiResponse.Success {
			t.Errorf("got success %v, want false", apiResponse.Success)
		}
	})
}

func TestDeleteCompany(t *testing.T) {
	t.Run("returns 200 when company is deleted successfully", func(t *testing.T) {
		mockService := &MockCompanyService{}

		request, _ := http.NewRequest(http.MethodDelete, "/companies/1", nil)
		request = mux.SetURLVars(request, map[string]string{
			"id": "1",
		})

		response := httptest.NewRecorder()
		handler := &CompanyHandler{companyService: mockService}
		handler.DeleteCompany(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("got status %d, want %d", response.Code, http.StatusOK)
		}

		var apiResponse models.APIResponse
		json.NewDecoder(response.Body).Decode(&apiResponse)

		if !apiResponse.Success {
			t.Errorf("got success %v, want true", apiResponse.Success)
		}
	})

	t.Run("returns 404 when company not found", func(t *testing.T) {
		mockService := &MockCompanyService{
			DeleteCompanyError: fmt.Errorf("not found"),
		}

		request, _ := http.NewRequest(http.MethodDelete, "/companies/999", nil)
		request = mux.SetURLVars(request, map[string]string{
			"id": "999",
		})

		response := httptest.NewRecorder()
		handler := &CompanyHandler{companyService: mockService}
		handler.DeleteCompany(response, request)

		if response.Code != http.StatusNotFound {
			t.Errorf("got status %d, want %d", response.Code, http.StatusNotFound)
		}

		var apiResponse models.APIResponse
		json.NewDecoder(response.Body).Decode(&apiResponse)

		if apiResponse.Success {
			t.Errorf("got success %v, want false", apiResponse.Success)
		}
	})
}
