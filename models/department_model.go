package models

import (
	"time"

	"github.com/falasefemi2/peopleos/dto"
)

type Department struct {
	ID        int       `db:"id" json:"id"`
	TenantID  int       `db:"tenant_id" json:"tenant_id"`
	Name      string    `db:"name" json:"name"`
	HodID     *int      `db:"hod_id" json:"hod_id"`
	Status    string    `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (d *Department) ToResponse() *dto.DepartmentResponse {
	return &dto.DepartmentResponse{
		ID:        d.ID,
		TenantID:  d.TenantID,
		Name:      d.Name,
		HodID:     d.HodID,
		Status:    d.Status,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}
