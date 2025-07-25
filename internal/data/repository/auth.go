package repository

import (
	"context"
	"project-POS-APP-golang-be-team/internal/data/entity"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthRepository interface {
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	SaveLoginToken(ctx context.Context, token *entity.LoginToken) error
	FindUserByToken(ctx context.Context, token string) (*entity.LoginToken, error)
}

type authRepositoryImpl struct {
	DB  *gorm.DB
	Log *zap.Logger
}

func NewAuthRepository(db *gorm.DB, log *zap.Logger) AuthRepository {
	return &authRepositoryImpl{
		DB:  db,
		Log: log,
	}
}

func (r *authRepositoryImpl) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		r.Log.Error("failed to find user by email : ", zap.String("error", err.Error()))
		return nil, err
	}
	return &user, nil
}

func (r *authRepositoryImpl) SaveLoginToken(ctx context.Context, token *entity.LoginToken) error {
	return r.DB.WithContext(ctx).Create(token).Error
}

func (r *authRepositoryImpl) FindUserByToken(ctx context.Context, token string) (*entity.LoginToken, error) {
	var loginToken entity.LoginToken
	err := r.DB.WithContext(ctx).
		Preload("User").
		Where("token = ?", token).
		First(&loginToken).Error
	if err != nil {
		return nil, err
	}
	return &loginToken, nil
}
