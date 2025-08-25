package repository

import (
	"github.com/toor/backend/internal/models"
	"gorm.io/gorm"
)

type AccountPointRepository interface {
	Create(accountPoint *models.AccountPoint) error
	GetAll() ([]models.AccountPoint, error)
	GetByID(id uint) (*models.AccountPoint, error)
	Update(accountPoint *models.AccountPoint) error
	Delete(id uint) error
}

type accountPointRepository struct {
	db *gorm.DB
}

func NewAccountPointRepository(db *gorm.DB) AccountPointRepository {
	return &accountPointRepository{db: db}
}

func (r *accountPointRepository) Create(accountPoint *models.AccountPoint) error {
	return r.db.Create(accountPoint).Error
}

func (r *accountPointRepository) GetAll() ([]models.AccountPoint, error) {
	var accountPoints []models.AccountPoint
	err := r.db.Order("date desc").Find(&accountPoints).Error
	return accountPoints, err
}

func (r *accountPointRepository) GetByID(id uint) (*models.AccountPoint, error) {
	var accountPoint models.AccountPoint
	err := r.db.First(&accountPoint, id).Error
	return &accountPoint, err
}

func (r *accountPointRepository) Update(accountPoint *models.AccountPoint) error {
	return r.db.Save(accountPoint).Error
}

func (r *accountPointRepository) Delete(id uint) error {
	return r.db.Delete(&models.AccountPoint{}, id).Error
}
