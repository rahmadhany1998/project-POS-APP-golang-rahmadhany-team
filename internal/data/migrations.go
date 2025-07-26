package data

import (
	"project-POS-APP-golang-be-team/internal/data/entity"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&entity.User{},
		&entity.PasswordResetToken{},
		&entity.LoginToken{},
		&entity.Category{},
		&entity.Table{},
		&entity.UserAccess{},
		&entity.Product{},
		&entity.Order{},
		&entity.OrderItem{},
	)
}
