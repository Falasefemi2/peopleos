package services

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/falasefemi2/peopleos/dto"
	"github.com/falasefemi2/peopleos/models"
	"github.com/falasefemi2/peopleos/repositories"
)

type ICompanyService interface {
	CreateCompany(ctx context.Context, req *dto.CreateCompanyRequest) (*dto.CompanyResponse, error)
	GetCompanyByName(ctx context.Context, name string) (*dto.CompanyResponse, error)
	GetCompanyByID(ctx context.Context, id int) (*dto.CompanyResponse, error)
	UpdateCompany(ctx context.Context, companyID int, req *dto.UpdateCompanyRequest) (*dto.CompanyResponse, error)
	DeleteCompany(ctx context.Context, companyID int) error
}

type CompanyService struct {
	companyRepo     *repositories.CompanyRepository
	tenantRepo      *repositories.TenanatRepository
	roleRepo        *repositories.RoleRepository
	employeeRepo    *repositories.EmployeeRepository
	departmentRepo  *repositories.DepartmentRepository
	designationRepo *repositories.DesignationRepository
}

func NewCompanyService(
	companyRepo *repositories.CompanyRepository,
	tenantRepo *repositories.TenanatRepository,
	roleRepo *repositories.RoleRepository,
	employeeRepo *repositories.EmployeeRepository,
	departmentRepo *repositories.DepartmentRepository,
	designationRepo *repositories.DesignationRepository,
) *CompanyService {
	return &CompanyService{
		companyRepo:     companyRepo,
		tenantRepo:      tenantRepo,
		roleRepo:        roleRepo,
		employeeRepo:    employeeRepo,
		departmentRepo:  departmentRepo,
		designationRepo: designationRepo,
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

	defaultDepartment := &models.Department{
		TenantID: createdTenant.ID,
		Name:     "General",
		Status:   "active",
	}

	createdDept, err := cs.departmentRepo.CreateDepartment(ctx, defaultDepartment)
	if err != nil {
		return nil, fmt.Errorf("error creating default department: %w", err)
	}

	defaultDesignation := &models.Designation{
		TenantID:    createdTenant.ID,
		Name:        "Owner",
		Level:       1,
		Description: "Company owner",
	}

	createdDesig, err := cs.designationRepo.CreateDesignation(ctx, defaultDesignation)
	if err != nil {
		return nil, fmt.Errorf("error creating default designation: %w", err)
	}

	superAdmin := &models.Employee{
		TenantID:      createdTenant.ID,
		FirstName:     req.AdminName,
		LastName:      "",
		Email:         req.AdminEmail,
		PasswordHash:  string(hashedPassword),
		DepartmentID:  createdDept.ID,
		DesignationID: createdDesig.ID,
		Status:        "active",
	}

	createdAdmin, err := cs.employeeRepo.CreateEmployee(ctx, superAdmin)
	if err != nil {
		return nil, fmt.Errorf("error creating super admin employee: %w", err)
	}

	// Find the Super Admin role we just created
	superAdminRoles, err := cs.roleRepo.GetRoleByName(ctx, createdTenant.ID, "Super Admin")
	if err == nil && superAdminRoles != nil {
		// Assign the role to the super admin employee
		err = cs.employeeRepo.AssignRoleToEmployee(ctx, createdAdmin.ID, superAdminRoles.ID)
		if err != nil {
			return nil, fmt.Errorf("error assigning super admin role: %w", err)
		}
	}

	err = cs.tenantRepo.UpdateTenantSuperAdmin(ctx, createdTenant.ID, createdAdmin.ID)
	if err != nil {
		return nil, fmt.Errorf("error updating tenant super admin: %w", err)
	}

	return createdCompany.ToResponse(), nil
}

func (cs *CompanyService) GetCompanyByName(ctx context.Context, name string) (*dto.CompanyResponse, error) {
	companyName, err := cs.companyRepo.GetCompanyByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("no company name found: %w", err)
	}
	return companyName.ToResponse(), nil
}

func (cs *CompanyService) GetCompanyByID(ctx context.Context, id int) (*dto.CompanyResponse, error) {
	companyID, err := cs.companyRepo.GetCompanyByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("no company with ID found: %w", err)
	}
	return companyID.ToResponse(), nil
}

func (cs *CompanyService) UpdateCompany(ctx context.Context, companyID int, req *dto.UpdateCompanyRequest) (*dto.CompanyResponse, error) {
	updatedCompany, err := cs.companyRepo.UpdateCompany(ctx, companyID, req)
	if err != nil {
		return nil, fmt.Errorf("error updating company: %w", err)
	}
	return updatedCompany.ToResponse(), nil
}

func (cs *CompanyService) DeleteCompany(ctx context.Context, companyID int) error {
	err := cs.companyRepo.DeleteCompany(ctx, companyID)
	if err != nil {
		return fmt.Errorf("error deleting company: %w", err)
	}
	return nil
}
