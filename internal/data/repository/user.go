package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository interface {
}

type userRepositoryImpl struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewUserRepository(db *gorm.DB, log *zap.Logger) UserRepository {
	return &userRepositoryImpl{
		DB:  db,
		Log: log,
	}
}
