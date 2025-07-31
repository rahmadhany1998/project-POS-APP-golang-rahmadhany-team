package repository

import (
	"context"
	"project-POS-APP-golang-be-team/internal/data/entity"

	"gorm.io/gorm"
)

type ResgiterAdminRepository interface {
	CreateAdmin(ctx context.Context, user *entity.User) error
}

type registerAdminRepository struct {
	db *gorm.DB
}

func NewRegisterAdminRepository(db *gorm.DB) ResgiterAdminRepository {
	return &registerAdminRepository{db: db}
}

func (r *registerAdminRepository) CreateAdmin(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}
