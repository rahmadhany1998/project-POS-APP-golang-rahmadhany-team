package repository

import (
	"project-POS-APP-golang-be-team/internal/data/entity"

	"gorm.io/gorm"
)

type MenuRepository interface {
	Create(product *entity.Product) error
	FindAll() ([]entity.Product, error)
	FindByID(id uint) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id uint) error
}

type menuRepository struct {
	db *gorm.DB
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &menuRepository{db}
}

func (r *menuRepository) Create(product *entity.Product) error {
	return r.db.Create(product).Error
}

func (r *menuRepository) FindAll() ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.Preload("Category").Find(&products).Error
	return products, err
}

func (r *menuRepository) FindByID(id uint) (*entity.Product, error) {
	var product entity.Product
	err := r.db.Preload("Category").First(&product, id).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *menuRepository) Update(product *entity.Product) error {
	return r.db.Save(product).Error
}

func (r *menuRepository) Delete(id uint) error {
	return r.db.Delete(&entity.Product{}, id).Error
}
