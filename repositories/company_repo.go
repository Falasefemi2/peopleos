package repositories

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/falasefemi2/peopleos/models"
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
