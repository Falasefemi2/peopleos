package repositories

import (
	"context"
	"time"

	"github.com/falasefemi2/peopleos/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DesignationRepository struct {
	pool *pgxpool.Pool
}

func NewDesignationRepository(pool *pgxpool.Pool) *DesignationRepository {
	return &DesignationRepository{
		pool: pool,
	}
}

func (d *DesignationRepository) CreateDesignation(ctx context.Context, designation *models.Designation) (*models.Designation, error) {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}

	query := `
	INSERT INTO designations (tenant_id, name, level, description)
	VALUES ($1, $2, $3, $4)
	RETURNING id, tenant_id, name, level, description, created_at, updated_at
	`

	row := d.pool.QueryRow(ctx, query, designation.TenantID, designation.Name, designation.Level, designation.Description)

	var createdDesignation models.Designation
	err := row.Scan(
		&createdDesignation.ID,
		&createdDesignation.TenantID,
		&createdDesignation.Name,
		&createdDesignation.Level,
		&createdDesignation.Description,
		&createdDesignation.CreatedAt,
		&createdDesignation.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &createdDesignation, nil
}

func (d *DesignationRepository) GetDesignationByID(ctx context.Context, id int) (*models.Designation, error) {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}

	query := `
	SELECT id, tenant_id, name, level, description, created_at, updated_at
	FROM designations
	WHERE id = $1
	`

	row := d.pool.QueryRow(ctx, query, id)

	var designation models.Designation
	err := row.Scan(
		&designation.ID,
		&designation.TenantID,
		&designation.Name,
		&designation.Level,
		&designation.Description,
		&designation.CreatedAt,
		&designation.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &designation, nil
}

func (d *DesignationRepository) UpdateDesignation(ctx context.Context, designationID int, designation *models.Designation) (*models.Designation, error) {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}

	query := `
	UPDATE designations
	SET name = $1, level = $2, description = $3, updated_at = CURRENT_TIMESTAMP
	WHERE id = $4
	RETURNING id, tenant_id, name, level, description, created_at, updated_at
	`

	row := d.pool.QueryRow(ctx, query, designation.Name, designation.Level, designation.Description, designationID)

	var updatedDesignation models.Designation
	err := row.Scan(
		&updatedDesignation.ID,
		&updatedDesignation.TenantID,
		&updatedDesignation.Name,
		&updatedDesignation.Level,
		&updatedDesignation.Description,
		&updatedDesignation.CreatedAt,
		&updatedDesignation.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &updatedDesignation, nil
}

func (d *DesignationRepository) DeleteDesignation(ctx context.Context, designationID int) error {
	if _, ok := ctx.Deadline(); !ok {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}

	query := `
	DELETE FROM designations
	WHERE id = $1
	`

	result, err := d.pool.Exec(ctx, query, designationID)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return nil
	}

	return nil
}
