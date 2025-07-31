package repository

import (
<<<<<<< HEAD
	"context"
	"project-POS-APP-golang-be-team/internal/data/entity"

=======
	"project-POS-APP-golang-be-team/internal/data/entity"
	"project-POS-APP-golang-be-team/internal/dto"

	"go.uber.org/zap"
>>>>>>> da90468 (wip: simpan semua perubahan lokal)
	"gorm.io/gorm"
)

type UserRepository interface {
<<<<<<< HEAD
	FindByID(ctx context.Context, id uint) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	Save(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id uint) error
	GetAdmins(ctx context.Context) ([]entity.User, error)
	UpdateUserRole(ctx context.Context, userID uint, newRole string) error
=======
	Create(user *entity.User) error
	GetByID(id uint) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id uint) error
	ListStaffs(limit, offset int, sortBy, sort string) ([]entity.User, int64, error)
	GetAllUsers(filter *dto.UserFilter) ([]*dto.UserResponse, dto.Pagination, error)
	
>>>>>>> da90468 (wip: simpan semua perubahan lokal)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByID(ctx context.Context, id uint) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	return &user, err
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) Save(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) Update(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&entity.User{}, id).Error
}

func (r *userRepository) GetAdmins(ctx context.Context) ([]entity.User, error) {
	var admins []entity.User
	if err := r.db.Where("role IN ?", []string{"admin", "superadmin"}).Find(&admins).Error; err != nil {
		return nil, err
	}
	return admins, nil
}

func (r *userRepository) UpdateUserRole(ctx context.Context, userID uint, newRole string) error {
	if err := r.db.WithContext(ctx).Model(&entity.User{}).
		Where("id = ? AND role IN ?", userID, []string{"admin", "superadmin"}).
		Update("role", newRole).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepositoryImpl) Create(user *entity.User) error {
	if err := r.DB.Create(user).Error; err != nil {
		r.Log.Error("Failed to create user", zap.Error(err))
		return err
	}
	return nil
}

func (r *userRepositoryImpl) GetByID(id uint) (*entity.User, error) {
	var user entity.User
	if err := r.DB.First(&user, id).Error; err != nil {
		r.Log.Error("Failed to get user by ID", zap.Uint("id", id), zap.Error(err))
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) GetByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		r.Log.Error("Failed to get user by email", zap.String("email", email), zap.Error(err))
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) Update(user *entity.User) error {
	if err := r.DB.Save(user).Error; err != nil {
		r.Log.Error("Failed to update user", zap.Error(err))
		return err
	}
	return nil
}

func (r *userRepositoryImpl) Delete(id uint) error {
	if err := r.DB.Delete(&entity.User{}, id).Error; err != nil {
		r.Log.Error("Failed to delete user", zap.Uint("id", id), zap.Error(err))
		return err
	}
	return nil
}

func (r *userRepositoryImpl) ListStaffs(limit, offset int, sortBy, sort string) ([]entity.User, int64, error) {
	var users []entity.User
	var total int64

	query := r.DB.Model(&entity.User{}).Where("role = ?", "staff")

	if err := query.Count(&total).Error; err != nil {
		r.Log.Error("Failed to count staff users", zap.Error(err))
		return nil, 0, err
	}

	if sortBy == "" {
		sortBy = "name"
	}
	if sort == "" {
		sort = "asc"
	}

	if err := query.Order(sortBy + " " + sort).
		Limit(limit).Offset(offset).
		Find(&users).Error; err != nil {
		r.Log.Error("Failed to list staff users", zap.Error(err))
		return nil, 0, err
	}

	return users, total, nil
}


func (r *userRepositoryImpl) GetAllUsers(filter *dto.UserFilter) ([]*dto.UserResponse, dto.Pagination, error) {
	var users []entity.User
	var total int64

	query := r.DB.Model(&entity.User{})

	// Filtering by name
	if filter.Name != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Name+"%")
	}

	// Filtering by email
	if filter.Email != "" {
		query = query.Where("email ILIKE ?", "%"+filter.Email+"%")
	}

	// Count total rows
	if err := query.Count(&total).Error; err != nil {
		return nil, dto.Pagination{}, err
	}

	// Sorting
	sortField := filter.SortBy
	if sortField == "" {
		sortField = "created_at"
	}
	sortOrder := "asc"
	if filter.SortDesc {
		sortOrder = "desc"
	}

	query = query.Order(sortField + " " + sortOrder)

	// Pagination
	limit := filter.PageSize
	offset := (filter.Page - 1) * filter.PageSize

	if err := query.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, dto.Pagination{}, err
	}

	// Mapping to response
	var result []*dto.UserResponse
	for _, u := range users {
		result = append(result, &dto.UserResponse{
			ID:       uint(u.ID), // Konversi int ke uint
			Name:     u.Name,
			Email:    u.Email,
			Role:     u.Role,
			IsActive: u.IsActive,
		})
	}

	// Hitung total halaman
	totalPages := int((total + int64(limit) - 1) / int64(limit))

	pagination := dto.Pagination{
		CurrentPage:  filter.Page,
		Limit:        limit,
		TotalPages:   totalPages,
		TotalRecords: int(total),
	}

	return result, pagination, nil
}
