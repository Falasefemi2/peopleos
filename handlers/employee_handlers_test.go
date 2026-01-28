package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/falasefemi2/peopleos/dto"
	"github.com/falasefemi2/peopleos/models"
)

func TestCreateEmployee(t *testing.T) {
	t.Run("returns 201 when employee is created successfully", func(t *testing.T) {
		mockEmployeeService := &MockEmployeeService{
			CreateEmployeeResult: &dto.EmployeeResponse{
				ID:    2,
				Email: "hr@company.com",
				Name:  "HR Manager",
				Role:  "HR",
			},
		}

		reqBody := dto.CreateEmployeeRequest{
			Email:         "hr@company.com",
			FirstName:     "HR",
			LastName:      "Manager",
			Password:      "password123",
			DepartmentID:  1,
			DesignationID: 1,
			RoleID:        2,
		}

		body, _ := json.Marshal(reqBody)
		request, _ := http.NewRequest(http.MethodPost, "/employees", bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		handler := &EmployeeHandler{employeeService: mockEmployeeService}
		handler.CreateEmployee(response, request)

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
		mockEmployeeService := &MockEmployeeService{}

		request, _ := http.NewRequest(http.MethodPost, "/employees", bytes.NewReader([]byte("invalid json")))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		handler := &EmployeeHandler{employeeService: mockEmployeeService}
		handler.CreateEmployee(response, request)

		if response.Code != http.StatusBadRequest {
			t.Errorf("got status %d, want %d", response.Code, http.StatusBadRequest)
		}

		var apiResponse models.APIResponse
		json.NewDecoder(response.Body).Decode(&apiResponse)

		if apiResponse.Success {
			t.Errorf("got success %v, want false", apiResponse.Success)
		}
	})

	t.Run("returns 400 when email is empty", func(t *testing.T) {
		mockEmployeeService := &MockEmployeeService{}

		reqBody := dto.CreateEmployeeRequest{
			Email:         "",
			FirstName:     "HR",
			LastName:      "Manager",
			Password:      "password123",
			DepartmentID:  1,
			DesignationID: 1,
			RoleID:        2,
		}

		body, _ := json.Marshal(reqBody)
		request, _ := http.NewRequest(http.MethodPost, "/employees", bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		handler := &EmployeeHandler{employeeService: mockEmployeeService}
		handler.CreateEmployee(response, request)

		if response.Code != http.StatusBadRequest {
			t.Errorf("got status %d, want %d", response.Code, http.StatusBadRequest)
		}
	})
}

type MockEmployeeService struct {
	CreateEmployeeResult *dto.EmployeeResponse
	CreateEmployeeError  error
}

func (m *MockEmployeeService) CreateEmployee(ctx context.Context, req *dto.CreateEmployeeRequest) (*dto.EmployeeResponse, error) {
	if m.CreateEmployeeError != nil {
		return nil, m.CreateEmployeeError
	}
	return m.CreateEmployeeResult, nil
}
