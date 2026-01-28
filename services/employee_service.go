package services

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/falasefemi2/peopleos/dto"
	"github.com/falasefemi2/peopleos/models"
	"github.com/falasefemi2/peopleos/repositories"
)

type IEmployeeService interface {
	CreateEmployee(ctx context.Context, req *dto.CreateEmployeeRequest) (*dto.EmployeeResponse, error)
}

type EmployeeService struct {
	employeeRepo *repositories.EmployeeRepository
	roleRepo     *repositories.RoleRepository
}

func NewEmployeeService(employeeRepo *repositories.EmployeeRepository, roleRepo *repositories.RoleRepository) *EmployeeService {
	return &EmployeeService{
		employeeRepo: employeeRepo,
		roleRepo:     roleRepo,
	}
}

func (es *EmployeeService) CreateEmployee(ctx context.Context, req *dto.CreateEmployeeRequest) (*dto.EmployeeResponse, error) {
	existingEmployee, _ := es.employeeRepo.GetEmployeeByEmail(ctx, req.Email)
	if existingEmployee != nil {
		return nil, fmt.Errorf("employee with this email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	employee := &models.Employee{
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		Email:         req.Email,
		PasswordHash:  string(hashedPassword),
		DepartmentID:  req.DepartmentID,
		DesignationID: req.DesignationID,
		Status:        "active",
	}

	createdEmployee, err := es.employeeRepo.CreateEmployee(ctx, employee)
	if err != nil {
		return nil, fmt.Errorf("error creating employee: %w", err)
	}

	err = es.employeeRepo.AssignRoleToEmployee(ctx, createdEmployee.ID, req.RoleID)
	if err != nil {
		return nil, fmt.Errorf("error assigning role: %w", err)
	}

	return &dto.EmployeeResponse{
		ID:    createdEmployee.ID,
		Email: createdEmployee.Email,
		Name:  createdEmployee.FirstName + " " + createdEmployee.LastName,
		Role:  "Assigned",
	}, nil
}
