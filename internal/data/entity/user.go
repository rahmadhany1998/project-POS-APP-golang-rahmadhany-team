package entity

import (
	"project-POS-APP-golang-be-team/pkg/utils"
	"time"
)

type User struct {
	Model
	Name       string    `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	Email      string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email" validate:"required,email"`
	Password   string    `gorm:"type:varchar(100);not null" json:"password" validate:"required,min=6"`
	Role       string    `gorm:"type:varchar(50);not null" json:"role" validate:"required"`
	Photo      string    `gorm:"type:text" json:"photo"`
	Phone      string    `gorm:"type:varchar(20)" json:"phone"`
	Address    string    `gorm:"type:text" json:"address"`
	Salary     float64   `json:"salary"`
	DOB        time.Time `json:"dob"`
	ShiftStart string    `json:"shift_start"`
	ShiftEnd   string    `json:"shift_end"`
	Detail     string    `gorm:"type:text" json:"detail"`
	IsActive   bool      `gorm:"default:true"`
}

func SeedUsers() []User {
	users := []User{
		{
			Name:     "Budi Santoso",
			Email:    "budi@example.com",
			Password: utils.HashPassword("password123"),
			Role:     "superadmin",
		},
	}

	return users
}
