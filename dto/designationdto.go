package dto

import "time"

type DesignationResponse struct {
	ID          int       `json:"id"`
	TenantID    int       `json:"tenant_id"`
	Name        string    `json:"name"`
	Level       int       `json:"level"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
