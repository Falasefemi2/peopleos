package models

import "time"

type Tenant struct {
	ID           int       `db:"id" json:"id"`
	CompanyID    int       `db:"company_id" json:"company_id"`
	SuperAdminID *int      `db:"super_admin_id" json:"super_admin_id"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}
