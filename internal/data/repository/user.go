package repository

import (
	"context"
	"project-POS-APP-golang-be-team/internal/data/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(ctx context.Context, id uint) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	Save(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id uint) error
	GetAdmins(ctx context.Context) ([]entity.User, error)
	UpdateUserRole(ctx context.Context, userID uint, newRole string) error
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
