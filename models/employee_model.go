package models

import "time"

type Employee struct {
	ID            int        `db:"id" json:"id"`
	TenantID      int        `db:"tenant_id" json:"tenant_id"`
	FirstName     string     `db:"first_name" json:"first_name"`
	LastName      string     `db:"last_name" json:"last_name"`
	Email         string     `db:"email" json:"email"`
	Phone         string     `db:"phone" json:"phone"`
	DepartmentID  int        `db:"department_id" json:"department_id"`
	DesignationID int        `db:"designation_id" json:"designation_id"`
	ManagerID     *int       `db:"manager_id" json:"manager_id"`
	Status        string     `db:"status" json:"status"`
	HireDate      *time.Time `db:"hire_date" json:"hire_date"`
	CreatedAt     time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at" json:"updated_at"`
}
