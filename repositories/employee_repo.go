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
	INSERT INTO employees (tenant_id, first_name, last_name, email, phone, department_id, designation_id, manager_id, status, hire_date)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING id, tenant_id, first_name, last_name, email, phone, department_id, designation_id, manager_id, status, hire_date, created_at, updated_at
	`

	row := e.pool.QueryRow(ctx, query, employee.TenantID, employee.FirstName, employee.LastName, employee.Email, employee.Phone, employee.DepartmentID, employee.DesignationID, employee.ManagerID, employee.Status, employee.HireDate)

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
