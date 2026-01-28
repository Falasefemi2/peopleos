package repositories

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/falasefemi2/peopleos/models"
)

type EmployeeRepository struct {
	pool *pgxpool.Pool
}

func NewEmployeeRepository(pool *pgxpool.Pool) *EmployeeRepository {
	return &EmployeeRepository{
		pool: pool,
	}
}

func (e *EmployeeRepository) CreateEmployee(ctx context.Context, employee *models.Employee) (*models.Employee, error) {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}

	query := `
	INSERT INTO employees (tenant_id, first_name, last_name, email, phone, department_id, designation_id, manager_id, status, hire_date, password_hash)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	RETURNING id, tenant_id, first_name, last_name, email, phone, department_id, designation_id, manager_id, status, hire_date, password_hash, created_at, updated_at
	`

	row := e.pool.QueryRow(ctx, query, employee.TenantID, employee.FirstName, employee.LastName, employee.Email, employee.Phone, employee.DepartmentID, employee.DesignationID, employee.ManagerID, employee.Status, employee.HireDate, employee.PasswordHash)

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
		&createdEmployee.PasswordHash,
		&createdEmployee.CreatedAt,
		&createdEmployee.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &createdEmployee, nil
}

func (e *EmployeeRepository) GetEmployeeByEmail(ctx context.Context, email string) (*models.Employee, error) {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}

	query := `
	SELECT id, tenant_id, first_name, last_name, email, phone, department_id, designation_id, manager_id, status, hire_date, password_hash, created_at, updated_at
	FROM employees
	WHERE email = $1
	`

	row := e.pool.QueryRow(ctx, query, email)

	var employee models.Employee
	err := row.Scan(
		&employee.ID,
		&employee.TenantID,
		&employee.FirstName,
		&employee.LastName,
		&employee.Email,
		&employee.Phone,
		&employee.DepartmentID,
		&employee.DesignationID,
		&employee.ManagerID,
		&employee.Status,
		&employee.HireDate,
		&employee.PasswordHash,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (e *EmployeeRepository) GetEmployeeByEmailWithRole(ctx context.Context, email string) (*models.Employee, string, error) {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}

	query := `
	SELECT e.id, e.tenant_id, e.first_name, e.last_name, e.email, e.phone, e.department_id, e.designation_id, e.manager_id, e.status, e.hire_date, e.password_hash, e.created_at, e.updated_at, r.name
	FROM employees e
	LEFT JOIN roles r ON e.role_id = r.id
	WHERE e.email = $1
	`

	row := e.pool.QueryRow(ctx, query, email)

	var employee models.Employee
	var roleName string

	err := row.Scan(
		&employee.ID,
		&employee.TenantID,
		&employee.FirstName,
		&employee.LastName,
		&employee.Email,
		&employee.Phone,
		&employee.DepartmentID,
		&employee.DesignationID,
		&employee.ManagerID,
		&employee.Status,
		&employee.HireDate,
		&employee.PasswordHash,
		&employee.CreatedAt,
		&employee.UpdatedAt,
		&roleName,
	)

	if err != nil {
		return nil, "", err
	}

	return &employee, roleName, nil
}

func (e *EmployeeRepository) AssignRoleToEmployee(ctx context.Context, employeeID int, roleID int) error {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}

	query := `
	UPDATE employees
	SET role_id = $1, updated_at = CURRENT_TIMESTAMP
	WHERE id = $2
	`

	_, err := e.pool.Exec(ctx, query, roleID, employeeID)
	return err
}
