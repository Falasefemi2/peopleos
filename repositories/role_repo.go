package repositories

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/falasefemi2/peopleos/models"
)

type RoleRepository struct {
	pool *pgxpool.Pool
}

func NewRoleRepository(pool *pgxpool.Pool) *RoleRepository {
	return &RoleRepository{
		pool: pool,
	}
}

func (r *RoleRepository) CreateRole(ctx context.Context, role *models.Role) (*models.Role, error) {
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

	row := r.pool.QueryRow(ctx, query, role.TenantID, role.Name, role.Description)

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
