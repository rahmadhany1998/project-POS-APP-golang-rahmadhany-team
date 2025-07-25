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
	CreatePasswordResetToken(ctx context.Context, token *entity.PasswordResetToken) error
	FindValidOtpToken(ctx context.Context, email string, otp string) (*entity.PasswordResetToken, error)
	UpdateUserPassword(ctx context.Context, email string, newHashed string) error
	MarkOtpAsUsed(ctx context.Context, id int) error
	DeleteLoginToken(ctx context.Context, token string) error
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
		r.Log.Error("FindByEmail failed: " + err.Error())
		return nil, err
	}
	return &user, nil
}

func (r *authRepositoryImpl) SaveLoginToken(ctx context.Context, token *entity.LoginToken) error {
	if err := r.DB.WithContext(ctx).Create(token).Error; err != nil {
		r.Log.Error("SaveLoginToken failed: " + err.Error())
		return err
	}
	return nil
}

func (r *authRepositoryImpl) FindUserByToken(ctx context.Context, token string) (*entity.LoginToken, error) {
	var loginToken entity.LoginToken
	err := r.DB.WithContext(ctx).
		Preload("User").
		Where("token = ?", token).
		First(&loginToken).Error
	if err != nil {
		r.Log.Error("FindUserByToken failed: " + err.Error())
		return nil, err
	}
	return &loginToken, nil
}

func (r *authRepositoryImpl) CreatePasswordResetToken(ctx context.Context, token *entity.PasswordResetToken) error {
	if err := r.DB.WithContext(ctx).Create(token).Error; err != nil {
		r.Log.Error("CreatePasswordResetToken failed: " + err.Error())
		return err
	}
	return nil
}

func (r *authRepositoryImpl) FindValidOtpToken(ctx context.Context, email string, otp string) (*entity.PasswordResetToken, error) {
	var token entity.PasswordResetToken
	err := r.DB.WithContext(ctx).
		Joins("JOIN users ON users.id = password_reset_tokens.user_id").
		Where("users.email = ? AND otp_code = ? AND is_used = false AND expires_at > NOW()", email, otp).
		Preload("User").
		First(&token).Error

	if err != nil {
		r.Log.Error("FindValidOtpToken failed: " + err.Error())
		return nil, err
	}
	return &token, nil
}

func (r *authRepositoryImpl) UpdateUserPassword(ctx context.Context, email string, newHashed string) error {
	if err := r.DB.WithContext(ctx).
		Model(&entity.User{}).
		Where("email = ?", email).
		Update("password", newHashed).Error; err != nil {
		r.Log.Error("UpdateUserPassword failed: " + err.Error())
		return err
	}
	return nil
}

func (r *authRepositoryImpl) MarkOtpAsUsed(ctx context.Context, id int) error {
	if err := r.DB.WithContext(ctx).
		Model(&entity.PasswordResetToken{}).
		Where("id = ?", id).
		Update("is_used", true).Error; err != nil {
		r.Log.Error("MarkOtpAsUsed failed: " + err.Error())
		return err
	}
	return nil
}

func (r *authRepositoryImpl) DeleteLoginToken(ctx context.Context, token string) error {
	if err := r.DB.WithContext(ctx).Where("token = ?", token).Delete(&entity.LoginToken{}).Error; err != nil {
		r.Log.Error("DeleteLoginToken failed: " + err.Error())
		return err
	}
	return nil
}
