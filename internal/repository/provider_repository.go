package repository

import (
	"github.com/toor/backend/internal/models"
	"gorm.io/gorm"
)

type ProviderRepository interface {
	Create(provider *models.Provider) error
	GetAll() ([]models.Provider, error)
	GetByID(id uint) (*models.Provider, error)
	Update(provider *models.Provider) error
	Delete(id uint) error
}

type providerRepository struct {
	db *gorm.DB
}

func NewProviderRepository(db *gorm.DB) ProviderRepository {
	return &providerRepository{db: db}
}

func (r *providerRepository) Create(provider *models.Provider) error {
	return r.db.Create(provider).Error
}

func (r *providerRepository) GetAll() ([]models.Provider, error) {
	var providers []models.Provider
	err := r.db.Order("name asc").Find(&providers).Error
	return providers, err
}

func (r *providerRepository) GetByID(id uint) (*models.Provider, error) {
	var provider models.Provider
	err := r.db.First(&provider, id).Error
	return &provider, err
}

func (r *providerRepository) Update(provider *models.Provider) error {
	return r.db.Save(provider).Error
}

func (r *providerRepository) Delete(id uint) error {
	return r.db.Delete(&models.Provider{}, id).Error
}