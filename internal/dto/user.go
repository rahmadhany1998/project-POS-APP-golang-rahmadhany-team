package dto

 HEAD
type UpdateAdminAccessRequest struct {
	TargetUserID int    `json:"target_user_id" binding:"required"`
	NewRole      string `json:"new_role" binding:"required"` // "admin", "superadmin"

import "time"

type CreateUserRequest struct {
	Name       string    `json:"name" validate:"required"`
	Email      string    `json:"email" validate:"required,email"`
	Password   string    `json:"password" validate:"required,min=6"`
	Role       string    `json:"role" validate:"required"`
	Photo      string    `json:"photo"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	Salary     float64   `json:"salary"`
	DOB        string  `json:"dob"`
	ShiftStart string    `json:"shift_start"`
	ShiftEnd   string    `json:"shift_end"`
	Detail     string    `json:"detail"`
	IsActive   *bool     `json:"is_active"`
}

type UpdateUserRequest struct {
	Name       string    `json:"name"`
	Email      string    `json:"email" validate:"required,email"`
	Password   string    `json:"password"`
	Role       string    `json:"role"`
	Photo      string    `json:"photo"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	Salary     float64   `json:"salary"`
	DOB        *time.Time `json:"dob"` 
	ShiftStart string    `json:"shift_start"`
	ShiftEnd   string    `json:"shift_end"`
	Detail     string    `json:"detail"`
	IsActive   *bool      `json:"is_active"` 
}

type UserResponse struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
	Photo      string    `json:"photo"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address"`
	Salary     float64   `json:"salary"`
	DOB        *time.Time `json:"dob"` 
	Age        int       `json:"age"`  
	ShiftStart string    `json:"shift_start"`
	ShiftEnd   string    `json:"shift_end"`
	Detail     string    `json:"detail"`
	IsActive   bool      `json:"is_active"`
}

type UserFilter struct {
	Search    string `form:"search"`
	Name      string `form:"name"`
	Email     string `form:"email"`
	SortBy    string `form:"sort_by"`
	SortDesc  bool   `form:"sort_desc"`
	Page      int    `form:"page"`
	PageSize  int    `form:"page_size"`
}

type UserRequest struct {
	Name       string  `json:"name" validate:"required"`
	Email      string  `json:"email" validate:"required,email"`
	Password   string  `json:"password" validate:"required,min=6"`
	Role       string  `json:"role" validate:"required"`
	Photo      string  `json:"photo"`
	Phone      string  `json:"phone"`
	Address    string  `json:"address"`
	Salary     float64 `json:"salary"`
	DOB        string  `json:"dob" validate:"required"` 
	ShiftStart string  `json:"shift_start"`
	ShiftEnd   string  `json:"shift_end"`
	Detail     string  `json:"detail"`
	IsActive   bool    `json:"is_active"`
}

type CreateUserRequestConverted struct {
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Password   string     `json:"password"`
	Role       string     `json:"role"`
	Photo      string     `json:"photo"`
	Phone      string     `json:"phone"`
	Address    string     `json:"address"`
	Salary     float64    `json:"salary"`
	DOB        *time.Time `json:"dob"` 
	ShiftStart string     `json:"shift_start"`
	ShiftEnd   string     `json:"shift_end"`
	Detail     string     `json:"detail"`
	IsActive   *bool      `json:"is_active"`
 da90468 (wip: simpan semua perubahan lokal)
}
