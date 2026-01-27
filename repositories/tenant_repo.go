package repositories

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/falasefemi2/peopleos/models"
)

type TenanatRepository struct {
	pool *pgxpool.Pool
}

func NewTenantRepository(pool *pgxpool.Pool) *TenanatRepository {
	return &TenanatRepository{
		pool: pool,
	}
}

func (t *TenanatRepository) CreateTenant(ctx context.Context, tenant *models.Tenant) (*models.Tenant, error) {
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

	row := t.pool.QueryRow(ctx, query, tenant.CompanyID, tenant.SuperAdminID)

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

func (t *TenanatRepository) UpdateTenantSuperAdmin(ctx context.Context, tenantID int, superAdminID int) error {
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

	_, err := t.pool.Exec(ctx, query, superAdminID, tenantID)
	return err
}
