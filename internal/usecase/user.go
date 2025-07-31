package usecase

import (
	"context"
	"errors"
	"project-POS-APP-golang-be-team/internal/data/entity"
	"project-POS-APP-golang-be-team/internal/data/repository"
	"project-POS-APP-golang-be-team/internal/dto"
	"project-POS-APP-golang-be-team/pkg/utils"

	"go.uber.org/zap"
)

type UserService interface {
	GetProfile(userID int) (dto.ProfileResponse, error)
	UpdateProfile(userID int, req dto.UpdateProfileRequest) error
	GetAdminList() ([]entity.User, error)
	UpdateAdminAccess(superAdminID int, req dto.UpdateAdminAccessRequest) error
}

type userService struct {
	userRepo repository.Repository
	logger   *zap.Logger
	config   utils.Configuration
}

func NewUserService(repo repository.Repository, logger *zap.Logger, config utils.Configuration) UserService {
	return &userService{
		userRepo: repo,
		logger:   logger,
		config:   config,
	}
}

func (s *userService) GetProfile(userID int) (dto.ProfileResponse, error) {
	ctx := context.Background()
	user, err := s.userRepo.UserRepo.FindByID(ctx, uint(userID))
	if err != nil {
		return dto.ProfileResponse{}, err
	}

	return dto.ProfileResponse{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Role:    user.Role,
		Phone:   user.Phone,
		Address: user.Address,
	}, nil
}

func (s *userService) UpdateProfile(userID int, req dto.UpdateProfileRequest) error {
	ctx := context.Background()
	user, err := s.userRepo.UserRepo.FindByID(ctx, uint(userID))
	if err != nil {
		return err
	}

	user.Name = req.Name
	user.Email = req.Email
	user.Phone = req.Phone
	user.Address = req.Address
	user.DOB = req.DOB

	return s.userRepo.UserRepo.Update(ctx, user)
}

func (s *userService) GetAdminList() ([]entity.User, error) {
	ctx := context.Background()
	admins, err := s.userRepo.UserRepo.GetAdmins(ctx)
	if err != nil {
		return nil, err
	}
	return admins, nil
}

func (s *userService) UpdateAdminAccess(superAdminID int, req dto.UpdateAdminAccessRequest) error {
	ctx := context.Background()

	user, err := s.userRepo.UserRepo.FindByID(ctx, uint(superAdminID))
	if err != nil {
		return errors.New("user tidak ditemukan")
	}
	if user.Role != "superadmin" {
		return errors.New("akses ditolak: hanya superadmin yang dapat mengubah akses admin")
	}

	return s.userRepo.UserRepo.UpdateUserRole(ctx, uint(req.TargetUserID), req.NewRole)
}
