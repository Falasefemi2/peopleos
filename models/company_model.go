package models

import "time"

type Company struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Industry  string    `db:"industry" json:"industry"`
	Country   string    `db:"country" json:"country"`
	Timezone  string    `db:"timezone" json:"timezone"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type CreateCompanyRequest struct {
	Name          string `json:"name" validate:"required"`
	Industry      string `json:"industry"`
	Country       string `json:"country" validate:"required"`
	Timezone      string `json:"timezone" validate:"required"`
	AdminEmail    string `json:"admin_email" validate:"required,email"`
	AdminName     string `json:"admin_name" validate:"required"`
	AdminPassword string `json:"admin_password" validate:"required,min=8"`
}
