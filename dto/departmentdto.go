package dto

import "time"

type DepartmentResponse struct {
	ID        int       `json:"id"`
	TenantID  int       `json:"tenant_id"`
	Name      string    `json:"name"`
	HodID     *int      `json:"hod_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
