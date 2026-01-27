package services

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/falasefemi2/peopleos/dto"
	"github.com/falasefemi2/peopleos/models"
	"github.com/falasefemi2/peopleos/repositories"
)

type CompanyService struct {
	companyRepo  *repositories.CompanyRepository
	tenantRepo   *repositories.TenanatRepository
	roleRepo     *repositories.RoleRepository
	employeeRepo *repositories.EmployeeRepository
}

func NewCompanyService(companyRepo *repositories.CompanyRepository, tenantRepo *repositories.TenanatRepository, roleRepo *repositories.RoleRepository, employeeRepo *repositories.EmployeeRepository) *CompanyService {
	return &CompanyService{
		companyRepo:  companyRepo,
		tenantRepo:   tenantRepo,
		roleRepo:     roleRepo,
		employeeRepo: employeeRepo,
	}
}

func (cs *CompanyService) CreateCompany(ctx context.Context, req *dto.CreateCompanyRequest) (*dto.CompanyResponse, error) {
	existingCompany, err := cs.companyRepo.GetCompanyByName(ctx, req.Name)
	if err == nil && existingCompany != nil {
		return nil, fmt.Errorf("company name already exists")
	}

	company := &models.Company{
		Name:     req.Name,
		Industry: req.Industry,
		Country:  req.Country,
		Timezone: req.Timezone,
	}

	createdCompany, err := cs.companyRepo.CreateCompany(ctx, company)
	if err != nil {
		return nil, fmt.Errorf("error creating company: %w", err)
	}

	tenant := &models.Tenant{
		CompanyID: createdCompany.ID,
	}

	createdTenant, err := cs.tenantRepo.CreateTenant(ctx, tenant)
	if err != nil {
		return nil, fmt.Errorf("error creating tenant: %w", err)
	}

	superAdminRole := &models.Role{
		TenantID:    createdTenant.ID,
		Name:        "Super Admin",
		Description: "Company owner with full access",
	}

	_, err = cs.roleRepo.CreateRole(ctx, superAdminRole)
	if err != nil {
		return nil, fmt.Errorf("error creating super admin role: %w", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.AdminPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}

	superAdmin := &models.Employee{
		TenantID:      createdTenant.ID,
		FirstName:     req.AdminName,
		LastName:      "",
		Email:         req.AdminEmail,
		PasswordHash:  string(hashedPassword),
		DepartmentID:  1,
		DesignationID: 1,
		Status:        "active",
	}

	createdAdmin, err := cs.employeeRepo.CreateEmployee(ctx, superAdmin)
	if err != nil {
		return nil, fmt.Errorf("error creating super admin employee: %w", err)
	}

	err = cs.tenantRepo.UpdateTenantSuperAdmin(ctx, createdTenant.ID, createdAdmin.ID)
	if err != nil {
		return nil, fmt.Errorf("error updating tenant super admin: %w", err)
	}

	return createdCompany.ToResponse(), nil
}
