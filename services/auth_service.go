package services

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/falasefemi2/peopleos/dto"
	"github.com/falasefemi2/peopleos/repositories"
)

type IAuthService interface {
	Login(ctx context.Context, req *dto.LoginRequest) (string, error)
}

type AuthService struct {
	employeeRepo *repositories.EmployeeRepository
}

func NewAuthService(employeeRepo *repositories.EmployeeRepository) *AuthService {
	return &AuthService{
		employeeRepo: employeeRepo,
	}
}

type Claims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

var jwtKey = []byte("your-secret-key-change-in-production")

func (as *AuthService) Login(ctx context.Context, req *dto.LoginRequest) (string, error) {
	employee, roleName, err := as.employeeRepo.GetEmployeeByEmailWithRole(ctx, req.Email)
	if err != nil {
		return "", fmt.Errorf("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(employee.PasswordHash), []byte(req.Password))
	if err != nil {
		return "", fmt.Errorf("invalid email or password")
	}

	if roleName == "" {
		roleName = "Employee"
	}

	claims := &Claims{
		ID:    employee.ID,
		Email: employee.Email,
		Role:  roleName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}

	return tokenString, nil
}
