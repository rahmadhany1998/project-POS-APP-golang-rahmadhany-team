package usecase

import (
	"context"
	"errors"
	"project-POS-APP-golang-be-team/internal/data/entity"
	"project-POS-APP-golang-be-team/internal/data/repository"
	"project-POS-APP-golang-be-team/internal/dto"
	"project-POS-APP-golang-be-team/pkg/utils"
	"time"

	"go.uber.org/zap"
)

type UserService interface {
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

// Hitung umur dari tanggal lahir
func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}


func toUserResponse(user *entity.User) *dto.UserResponse {
	return &dto.UserResponse{
		ID:         uint(user.ID),
		Name:       user.Name,
		Email:      user.Email,
		Role:       user.Role,
		DOB:        &user.DOB,
		Age:        calculateAge(user.DOB),
		Phone:      user.Phone,
		Address:    user.Address,
		Salary:     user.Salary,
		Photo:      user.Photo,
		Detail:     user.Detail,
		IsActive:   user.IsActive,
		ShiftStart: user.ShiftStart,
		ShiftEnd:   user.ShiftEnd,
	}
}


func (s *userService) CreateUser(req *dto.CreateUserRequest) (*dto.UserResponse, error) {
	parsedDOB, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		s.Logger.Error("Invalid date format", zap.Error(err))
		return nil, err
	}

	isActive := false
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	user := &entity.User{
		Name:       req.Name,
		Email:      req.Email,
		Password:   req.Password,
		Role:       req.Role,
		Phone:      req.Phone,
		Address:    req.Address,
		Salary:     req.Salary,
		Photo:      req.Photo,
		Detail:     req.Detail,
		DOB:        parsedDOB,
		IsActive:   isActive,
		ShiftStart: req.ShiftStart,
		ShiftEnd:   req.ShiftEnd,
	}

	if err := s.Repo.UserRepo.Create(user); err != nil {
		s.Logger.Error("Failed to create user", zap.Error(err))
		return nil, err
	}

	return toUserResponse(user), nil
}

func (s *userService) CreateUserConverted(req *dto.CreateUserRequestConverted) (*dto.UserResponse, error) {
	isActive := false
	if req.IsActive != nil {
		isActive = *req.IsActive
	}

	user := &entity.User{
		Name:       req.Name,
		Email:      req.Email,
		Password:   req.Password,
		Role:       req.Role,
		Phone:      req.Phone,
		Address:    req.Address,
		Salary:     req.Salary,
		Photo:      req.Photo,
		Detail:     req.Detail,
		ShiftStart: req.ShiftStart,
		ShiftEnd:   req.ShiftEnd,
		IsActive:   isActive,
	}

	if req.DOB != nil {
		user.DOB = *req.DOB
	}

	if err := s.Repo.UserRepo.Create(user); err != nil {
		s.Logger.Error("Failed to create user (converted)", zap.Error(err))
		return nil, err
	}

	return toUserResponse(user), nil
}

func (s *userService) GetUserByID(id int) (*dto.UserResponse, error) {
	user, err := s.Repo.UserRepo.GetByID(uint(id))
	if err != nil {
		s.Logger.Error("Failed to get user by ID", zap.Int("id", id), zap.Error(err))
		return nil, err
	}
	return toUserResponse(user), nil
}

func (s *userService) UpdateUser(id int, req *dto.UpdateUserRequest) (*dto.UserResponse, error) {
	user, err := s.Repo.UserRepo.GetByID(uint(id))
	if err != nil {
		s.Logger.Error("User not found", zap.Int("id", id), zap.Error(err))
		return nil, err
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		user.Password = req.Password
	}
	if req.Role != "" {
		user.Role = req.Role
	}
	if req.DOB != nil {
		user.DOB = *req.DOB
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	if req.Address != "" {
		user.Address = req.Address
	}
	if req.Photo != "" {
		user.Photo = req.Photo
	}
	if req.Detail != "" {
		user.Detail = req.Detail
	}
	if req.Salary != 0 {
		user.Salary = req.Salary
	}
	if req.ShiftStart != "" {
		user.ShiftStart = req.ShiftStart
	}
	if req.ShiftEnd != "" {
		user.ShiftEnd = req.ShiftEnd
	}
	if req.IsActive != nil {
		user.IsActive = *req.IsActive
	}

	if err := s.Repo.UserRepo.Update(user); err != nil {
		s.Logger.Error("Failed to update user", zap.Error(err))
		return nil, err
	}

	return toUserResponse(user), nil
}

func (s *userService) DeleteUser(id int) error {
	err := s.Repo.UserRepo.Delete(uint(id))
	if err != nil {
		s.Logger.Error("Failed to delete user", zap.Int("id", id), zap.Error(err))
	}
	return err
}

func (s *userService) ListStaffs(limit, offset int, sortBy, sort string) ([]dto.UserResponse, int64, error) {
	users, total, err := s.Repo.UserRepo.ListStaffs(limit, offset, sortBy, sort)
	if err != nil {
		s.Logger.Error("Failed to list staff users", zap.Error(err))
		return nil, 0, err
	}

	var res []dto.UserResponse
	for _, user := range users {
		res = append(res, *toUserResponse(&user))
	}

	return res, total, nil
}

func (s *userService) GetAllUsers(filter *dto.UserFilter) ([]*dto.UserResponse, dto.Pagination, error) {
	return s.Repo.UserRepo.GetAllUsers(filter)
}
