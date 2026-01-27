package models

import (
	"time"

	"github.com/falasefemi2/peopleos/dto"
)

type Designation struct {
	ID          int       `db:"id" json:"id"`
	TenantID    int       `db:"tenant_id" json:"tenant_id"`
	Name        string    `db:"name" json:"name"`
	Level       int       `db:"level" json:"level"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

func (d *Designation) ToResponse() *dto.DesignationResponse {
	return &dto.DesignationResponse{
		ID:          d.ID,
		TenantID:    d.TenantID,
		Name:        d.Name,
		Level:       d.Level,
		Description: d.Description,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}
