package usecase

import (
	"project-POS-APP-golang-be-team/internal/data/repository"
	"project-POS-APP-golang-be-team/pkg/utils"

	"go.uber.org/zap"
)

type UserService interface {
}

type userService struct {
	Repo   repository.Repository
	Logger *zap.Logger
	Config utils.Configuration
}

func NewUserService(repo repository.Repository, logger *zap.Logger, config utils.Configuration) UserService {
	return &userService{
		Repo:   repo,
		Logger: logger,
		Config: config,
	}
}
