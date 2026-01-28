package dto

type CreateEmployeeRequest struct {
	Email         string `json:"email" validate:"required,email"`
	FirstName     string `json:"first_name" validate:"required"`
	LastName      string `json:"last_name"`
	Password      string `json:"password" validate:"required,min=8"`
	DepartmentID  int    `json:"department_id" validate:"required"`
	DesignationID int    `json:"designation_id" validate:"required"`
	RoleID        int    `json:"role_id" validate:"required"`
}

type EmployeeResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}
