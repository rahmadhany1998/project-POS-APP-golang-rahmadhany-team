package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	AuthRepo    AuthRepository
	RevenueRepo RevenueRepository
	MenuRepo    MenuRepository
	UserRepo    UserRepository
	RegisRepo   ResgiterAdminRepository
}

func NewRepository(db *gorm.DB, log *zap.Logger) Repository {
	return Repository{
		AuthRepo:    NewAuthRepository(db, log),
		RevenueRepo: NewRevenueRepository(db, log),
		MenuRepo:    NewMenuRepository(db),
		UserRepo:    NewUserRepository(db),
		RegisRepo:   NewRegisterAdminRepository(db),
	}
}
