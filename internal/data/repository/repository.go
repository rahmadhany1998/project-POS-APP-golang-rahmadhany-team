package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	AuthRepo    AuthRepository
	RevenueRepo RevenueRepository
	UserRepo    UserRepository
}

func NewRepository(db *gorm.DB, log *zap.Logger) Repository {
	return Repository{
		AuthRepo:    NewAuthRepository(db, log),
		RevenueRepo: NewRevenueRepository(db, log),
		UserRepo:    NewUserRepository(db),
	}
}
