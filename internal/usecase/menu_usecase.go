package usecase

import (
	"errors"
	"project-POS-APP-golang-be-team/internal/data/entity"
	"project-POS-APP-golang-be-team/internal/data/repository"
	"project-POS-APP-golang-be-team/pkg/utils"

	"go.uber.org/zap"
)

type MenuUsecase interface {
	CreateMenu(product *entity.Product) error
	GetAllMenus() ([]entity.Product, error)
	GetMenuByID(id uint) (*entity.Product, error)
	UpdateMenu(product *entity.Product) error
	DeleteMenu(id uint) error
	// Filter(id uint)(*entity.Product)
	GetAllCategory()
}

type menuUsecase struct {
	repo   repository.MenuRepository
	logger *zap.Logger
	config utils.Configuration
}

func NewMenuUsecase(r repository.MenuRepository, logger *zap.Logger, config utils.Configuration) MenuUsecase {
	return &menuUsecase{repo: r, logger: logger, config: config}
}

var ErrInvalidMenuData = errors.New("data menu tidak valid")

func (uc *menuUsecase) CreateMenu(product *entity.Product) error {
	if product.Name == "" || product.Price <= 0 {
		return ErrInvalidMenuData
	}
	return uc.repo.Create(product)
}

func (uc *menuUsecase) GetAllMenus() ([]entity.Product, error) {
	return uc.repo.FindAll()
}

func (uc *menuUsecase) GetMenuByID(id uint) (*entity.Product, error) {
	return uc.repo.FindByID(id)
}

func (uc *menuUsecase) UpdateMenu(product *entity.Product) error {
	return uc.repo.Update(product)
}

func (uc *menuUsecase) DeleteMenu(id uint) error {
	return uc.repo.Delete(id)
}
