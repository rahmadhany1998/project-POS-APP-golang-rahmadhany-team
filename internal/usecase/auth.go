package usecase

import (
	"context"
	"errors"
	"project-POS-APP-golang-be-team/internal/data/entity"
	"project-POS-APP-golang-be-team/internal/data/repository"
	"project-POS-APP-golang-be-team/internal/dto"
	"project-POS-APP-golang-be-team/pkg/utils"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type AuthService interface {
	Login(ctx context.Context, req dto.LoginRequest) (*dto.ResponseUser, error)
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
