package repositories

import (
	"context"
	"time"

	"github.com/falasefemi2/peopleos/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DepartmentRepository struct {
	pool *pgxpool.Pool
}

func NewDepartmentRepository(pool *pgxpool.Pool) *DepartmentRepository {
	return &DepartmentRepository{
		pool: pool,
	}
}

func (d *DepartmentRepository) CreateDepartment(ctx context.Context, department *models.Department) (*models.Department, error) {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}

	query := `
	INSERT INTO departments (tenant_id, name, status)
	VALUES ($1, $2, $3)
	RETURNING id, tenant_id, name, hod_id, status, created_at, updated_at
	`

	row := d.pool.QueryRow(ctx, query, department.TenantID, department.Name, department.Status)

	var createdDepartment models.Department
	err := row.Scan(
		&createdDepartment.ID,
		&createdDepartment.TenantID,
		&createdDepartment.Name,
		&createdDepartment.HodID,
		&createdDepartment.Status,
		&createdDepartment.CreatedAt,
		&createdDepartment.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &createdDepartment, nil
}

func (d *DepartmentRepository) GetDepartmentByID(ctx context.Context, id int) (*models.Department, error) {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}

	query := `
	SELECT id, tenant_id, name, hod_id, status, created_at, updated_at
	FROM departments
	WHERE id = $1
	`

	row := d.pool.QueryRow(ctx, query, id)

	var department models.Department
	err := row.Scan(
		&department.ID,
		&department.TenantID,
		&department.Name,
		&department.HodID,
		&department.Status,
		&department.CreatedAt,
		&department.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &department, nil
}

func (d *DepartmentRepository) UpdateDepartment(ctx context.Context, departmentID int, department *models.Department) (*models.Department, error) {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}

	query := `
	UPDATE departments
	SET name = $1, hod_id = $2, status = $3, updated_at = CURRENT_TIMESTAMP
	WHERE id = $4
	RETURNING id, tenant_id, name, hod_id, status, created_at, updated_at
	`

	row := d.pool.QueryRow(ctx, query, department.Name, department.HodID, department.Status, departmentID)

	var updatedDepartment models.Department
	err := row.Scan(
		&updatedDepartment.ID,
		&updatedDepartment.TenantID,
		&updatedDepartment.Name,
		&updatedDepartment.HodID,
		&updatedDepartment.Status,
		&updatedDepartment.CreatedAt,
		&updatedDepartment.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &updatedDepartment, nil
}

func (d *DepartmentRepository) DeleteDepartment(ctx context.Context, departmentID int) error {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}

	query := `
	DELETE FROM departments
	WHERE id = $1
	`

	result, err := d.pool.Exec(ctx, query, departmentID)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return nil
	}

	return nil
}
