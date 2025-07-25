package entity

import "project-POS-APP-golang-be-team/pkg/utils"

type User struct {
	Model
	Name     string `gorm:"type:varchar(100);not null" json:"name" validate:"required"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email" validate:"required,email"`
	Password string `gorm:"type:varchar(100);not null" json:"password" validate:"required,min=6"`
	Role     string `gorm:"type:varchar(50);not null" json:"role" validate:"required"`
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
