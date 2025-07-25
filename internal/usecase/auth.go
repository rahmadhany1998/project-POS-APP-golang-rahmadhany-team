package usecase

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"project-POS-APP-golang-be-team/internal/data/entity"
	"project-POS-APP-golang-be-team/internal/data/repository"
	"project-POS-APP-golang-be-team/internal/dto"
	"project-POS-APP-golang-be-team/pkg/utils"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type AuthService interface {
	Login(ctx context.Context, req dto.LoginRequest) (*dto.ResponseUser, error)
	ForgotPassword(ctx context.Context, req dto.ForgotPasswordRequest) error
	VerifyOtp(ctx context.Context, req dto.VerifyOtpRequest) error
	ResetPassword(ctx context.Context, req dto.ResetPasswordRequest) error
	Logout(ctx context.Context, token string) error
}

type authService struct {
	Repo   repository.Repository
	Logger *zap.Logger
	Config utils.Configuration
}

func NewAuthService(repo repository.Repository, logger *zap.Logger, config utils.Configuration) AuthService {
	return &authService{
		Repo:   repo,
		Logger: logger,
		Config: config,
	}
}

func (s *authService) Login(ctx context.Context, req dto.LoginRequest) (*dto.ResponseUser, error) {
	user, err := s.Repo.AuthRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		s.Logger.Error("login failed: user not found", zap.String("error", err.Error()))
		return nil, errors.New("invalid username or password")
	}

	isValid := utils.CheckPassword(req.Password, user.Password)
	if !isValid {
		s.Logger.Error("login failed: wrong password", zap.String("email", req.Email))
		return nil, errors.New("invalid username or password")
	}

	token := uuid.New().String()

	err = s.Repo.AuthRepo.SaveLoginToken(ctx, &entity.LoginToken{
		UserID: user.ID,
		Token:  token,
	})
	if err != nil {
		s.Logger.Error("failed to save login token", zap.String("email", req.Email), zap.String("error", err.Error()))
		return nil, errors.New("failed to save login token")
	}

	resp := &dto.ResponseUser{
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}
	return resp, nil
}

func (s *authService) ForgotPassword(ctx context.Context, req dto.ForgotPasswordRequest) error {
	user, err := s.Repo.AuthRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return errors.New("email not registered")
	}

	otp := fmt.Sprintf("%04d", rand.Intn(10000))
	token := &entity.PasswordResetToken{
		UserID:    uint(user.ID),
		OtpCode:   otp,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
	if err := s.Repo.AuthRepo.CreatePasswordResetToken(ctx, token); err != nil {
		return errors.New("failed to save otp")
	}

	s.Logger.Info("OTP for reset password", zap.String("email", req.Email), zap.String("otp", otp))
	return nil
}

func (s *authService) VerifyOtp(ctx context.Context, req dto.VerifyOtpRequest) error {
	_, err := s.Repo.AuthRepo.FindValidOtpToken(ctx, req.Email, req.Otp)
	if err != nil {
		return errors.New("otp bot valid or expired")
	}
	return nil
}

func (s *authService) ResetPassword(ctx context.Context, req dto.ResetPasswordRequest) error {
	token, err := s.Repo.AuthRepo.FindValidOtpToken(ctx, req.Email, req.Otp)
	if err != nil {
		return errors.New("otp not valid")
	}

	hashed := utils.HashPassword(req.NewPassword)
	if err := s.Repo.AuthRepo.UpdateUserPassword(ctx, req.Email, string(hashed)); err != nil {
		return errors.New("failed to change password")
	}
	if err := s.Repo.AuthRepo.MarkOtpAsUsed(ctx, token.ID); err != nil {
		s.Logger.Warn("failed to mark otp as used", zap.Error(err))
	}
	return nil
}

func (s *authService) Logout(ctx context.Context, token string) error {
	if err := s.Repo.AuthRepo.DeleteLoginToken(ctx, token); err != nil {
		return errors.New("logout failed")
	}
	return nil
}
