package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/falasefemi2/peopleos/dto"
	"github.com/falasefemi2/peopleos/models"
)

func TestLogin(t *testing.T) {
	t.Run("returns 200 and token when credentials are valid", func(t *testing.T) {
		mockAuthService := &MockAuthService{
			LoginResult: "valid-jwt-token",
		}

		reqBody := dto.LoginRequest{
			Email:    "user@test.com",
			Password: "password123",
		}

		body, _ := json.Marshal(reqBody)
		request, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		handler := &AuthHandler{authService: mockAuthService}
		handler.Login(response, request)

		if response.Code != http.StatusOK {
			t.Errorf("got status %d, want %d", response.Code, http.StatusOK)
		}

		var apiResponse models.APIResponse
		json.NewDecoder(response.Body).Decode(&apiResponse)

		if !apiResponse.Success {
			t.Errorf("got success %v, want true", apiResponse.Success)
		}
	})

	t.Run("returns 401 when credentials are invalid", func(t *testing.T) {
		mockAuthService := &MockAuthService{
			LoginError: fmt.Errorf("invalid email or password"),
		}

		reqBody := dto.LoginRequest{
			Email:    "user@test.com",
			Password: "wrongpassword",
		}

		body, _ := json.Marshal(reqBody)
		request, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(body))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		handler := &AuthHandler{authService: mockAuthService}
		handler.Login(response, request)

		if response.Code != http.StatusUnauthorized {
			t.Errorf("got status %d, want %d", response.Code, http.StatusUnauthorized)
		}

		var apiResponse models.APIResponse
		json.NewDecoder(response.Body).Decode(&apiResponse)

		if apiResponse.Success {
			t.Errorf("got success %v, want false", apiResponse.Success)
		}
	})

	t.Run("returns 400 when request body is invalid", func(t *testing.T) {
		mockAuthService := &MockAuthService{}

		request, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader([]byte("invalid json")))
		request.Header.Set("Content-Type", "application/json")

		response := httptest.NewRecorder()

		handler := &AuthHandler{authService: mockAuthService}
		handler.Login(response, request)

		if response.Code != http.StatusBadRequest {
			t.Errorf("got status %d, want %d", response.Code, http.StatusBadRequest)
		}

		var apiResponse models.APIResponse
		json.NewDecoder(response.Body).Decode(&apiResponse)

		if apiResponse.Success {
			t.Errorf("got success %v, want false", apiResponse.Success)
		}
	})
}

type MockAuthService struct {
	LoginResult string
	LoginError  error
}

func (m *MockAuthService) Login(ctx context.Context, req *dto.LoginRequest) (string, error) {
	if m.LoginError != nil {
		return "", m.LoginError
	}
	return m.LoginResult, nil
}
