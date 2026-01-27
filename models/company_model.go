package models

import (
	"time"

	"github.com/falasefemi2/peopleos/dto"
)

type Company struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Industry  string    `db:"industry" json:"industry"`
	Country   string    `db:"country" json:"country"`
	Timezone  string    `db:"timezone" json:"timezone"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (c *Company) ToResponse() *dto.CompanyResponse {
	return &dto.CompanyResponse{
		ID:        c.ID,
		Name:      c.Name,
		Industry:  c.Industry,
		Country:   c.Country,
		Timezone:  c.Timezone,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
