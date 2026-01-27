package repositories

import (
	"context"
	"time"

	"github.com/falasefemi2/peopleos/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CompanyRepository struct {
	pool *pgxpool.Pool
}

func NewCompanyRepository(pool *pgxpool.Pool) *CompanyRepository {
	return &CompanyRepository{
		pool: pool,
	}
}

func (c *CompanyRepository) CreateCompany(ctx context.Context, company *models.Company) (*models.Company, error) {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}

	query := `
	INSERT INTO companies (name, industry, country, timezone)
	VALUES ($1, $2, $3, $4)
	RETURNING id, name, industry, country, timezone, created_at, updated_at
	`

	row := c.pool.QueryRow(ctx, query, company.Name, company.Industry, company.Country, company.Timezone)

	var createdCompany models.Company
	err := row.Scan(
		&createdCompany.ID,
		&createdCompany.Name,
		&createdCompany.Industry,
		&createdCompany.Country,
		&createdCompany.Timezone,
		&createdCompany.CreatedAt,
		&createdCompany.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &createdCompany, nil
}

func (c *CompanyRepository) GetCompanyByName(ctx context.Context, name string) (*models.Company, error) {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}

	query := `
	SELECT id, name, industry, country, timezone, created_at, updated_at
	FROM companies
	WHERE name = $1
	`

	row := c.pool.QueryRow(ctx, query, name)

	var company models.Company
	err := row.Scan(
		&company.ID,
		&company.Name,
		&company.Industry,
		&company.Country,
		&company.Timezone,
		&company.CreatedAt,
		&company.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &company, nil
}

func (c *CompanyRepository) GetCompanyByID(ctx context.Context, id int) (*models.Company, error) {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}

	query := `
	SELECT id, name, industry, country, timezone, created_at, updated_at
	FROM companies
	WHERE id = $1
	`

	row := c.pool.QueryRow(ctx, query, id)

	var company models.Company
	err := row.Scan(
		&company.ID,
		&company.Name,
		&company.Industry,
		&company.Country,
		&company.Timezone,
		&company.CreatedAt,
		&company.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &company, nil
}

func (c *CompanyRepository) CreateTenant(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error) {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}

	query := `
	INSERT INTO tenants (company_id, super_admin_id)
	VALUES ($1, $2)
	RETURNING id, company_id, super_admin_id, created_at, updated_at
	`

	row := c.pool.QueryRow(ctx, query, tenant.CompanyID, tenant.SuperAdminID)

	var createdTenant models.Tenant
	err := row.Scan(
		&createdTenant.ID,
		&createdTenant.CompanyID,
		&createdTenant.SuperAdminID,
		&createdTenant.CreatedAt,
		&createdTenant.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &createdTenant, nil
}

func (c *CompanyRepository) CreateEmployee(ctx context.Context, employee *models.Employee) (*models.Employee, error) {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}

	query := `
	INSERT INTO employees (tenant_id, first_name, last_name, email, phone, department_id, designation_id, manager_id, status, hire_date)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING id, tenant_id, first_name, last_name, email, phone, department_id, designation_id, manager_id, status, hire_date, created_at, updated_at
	`

	row := c.pool.QueryRow(ctx, query, employee.TenantID, employee.FirstName, employee.LastName, employee.Email, employee.Phone, employee.DepartmentID, employee.DesignationID, employee.ManagerID, employee.Status, employee.HireDate)

	var createdEmployee models.Employee
	err := row.Scan(
		&createdEmployee.ID,
		&createdEmployee.TenantID,
		&createdEmployee.FirstName,
		&createdEmployee.LastName,
		&createdEmployee.Email,
		&createdEmployee.Phone,
		&createdEmployee.DepartmentID,
		&createdEmployee.DesignationID,
		&createdEmployee.ManagerID,
		&createdEmployee.Status,
		&createdEmployee.HireDate,
		&createdEmployee.CreatedAt,
		&createdEmployee.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &createdEmployee, nil
}

func (c *CompanyRepository) CreateRole(ctx context.Context, role *models.Role) (*models.Role, error) {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}

	query := `
	INSERT INTO roles (tenant_id, name, description)
	VALUES ($1, $2, $3)
	RETURNING id, tenant_id, name, description, created_at, updated_at
	`

	row := c.pool.QueryRow(ctx, query, role.TenantID, role.Name, role.Description)

	var createdRole models.Role
	err := row.Scan(
		&createdRole.ID,
		&createdRole.TenantID,
		&createdRole.Name,
		&createdRole.Description,
		&createdRole.CreatedAt,
		&createdRole.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &createdRole, nil
}

func (c *CompanyRepository) UpdateTenantSuperAdmin(ctx context.Context, tenantID int, superAdminID int) error {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}

	query := `
	UPDATE tenants
	SET super_admin_id = $1, updated_at = CURRENT_TIMESTAMP
	WHERE id = $2
	`

	_, err := c.pool.Exec(ctx, query, superAdminID, tenantID)
	return err
}
