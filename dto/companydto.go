package dto

import "time"

type CreateCompanyRequest struct {
	Name          string `json:"name" validate:"required"`
	Industry      string `json:"industry"`
	Country       string `json:"country" validate:"required"`
	Timezone      string `json:"timezone" validate:"required"`
	AdminEmail    string `json:"admin_email" validate:"required,email"`
	AdminName     string `json:"admin_name" validate:"required"`
	AdminPassword string `json:"admin_password" validate:"required,min=8"`
}

type UpdateCompanyRequest struct {
	Name     string `json:"name" validate:"required"`
	Industry string `json:"industry"`
	Country  string `json:"country" validate:"required"`
	Timezone string `json:"timezone" validate:"required"`
}

type CompanyResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Industry  string    `json:"industry"`
	Country   string    `json:"country"`
	Timezone  string    `json:"timezone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
