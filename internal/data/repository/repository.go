package repository

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	AuthRepo AuthRepository
}

func NewRepository(db *gorm.DB, log *zap.Logger) Repository {
	return Repository{
		AuthRepo: NewAuthRepository(db, log),
	}
}
