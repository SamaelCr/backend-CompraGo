package repository

import (
	"github.com/toor/backend/internal/models"
	"gorm.io/gorm"
)

type MasterDataRepository interface {
	// Units
	CreateUnit(unit *models.Unit) error
	GetAllUnits() ([]models.Unit, error)
	UpdateUnit(unit *models.Unit) error
	// Positions
	CreatePosition(pos *models.Position) error
	GetAllPositions() ([]models.Position, error)
	UpdatePosition(pos *models.Position) error
	// Officials
	CreateOfficial(off *models.Official) error
	GetAllOfficials() ([]models.Official, error)
	UpdateOfficial(off *models.Official) error
}

type masterDataRepository struct {
	db *gorm.DB
}

func NewMasterDataRepository(db *gorm.DB) MasterDataRepository {
	return &masterDataRepository{db: db}
}

// Units
func (r *masterDataRepository) CreateUnit(unit *models.Unit) error {
	return r.db.Create(unit).Error
}
func (r *masterDataRepository) GetAllUnits() ([]models.Unit, error) {
	var units []models.Unit
	err := r.db.Order("name asc").Find(&units).Error
	return units, err
}
func (r *masterDataRepository) UpdateUnit(unit *models.Unit) error {
	return r.db.Save(unit).Error
}

// Positions
func (r *masterDataRepository) CreatePosition(pos *models.Position) error {
	return r.db.Create(pos).Error
}
func (r *masterDataRepository) GetAllPositions() ([]models.Position, error) {
	var positions []models.Position
	err := r.db.Order("name asc").Find(&positions).Error
	return positions, err
}
func (r *masterDataRepository) UpdatePosition(pos *models.Position) error {
	return r.db.Save(pos).Error
}

// Officials
func (r *masterDataRepository) CreateOfficial(off *models.Official) error {
	return r.db.Create(off).Error
}
func (r *masterDataRepository) GetAllOfficials() ([]models.Official, error) {
	var officials []models.Official
	// Usamos Preload para traer los datos de Unit y Position
	err := r.db.Preload("Unit").Preload("Position").Order("full_name asc").Find(&officials).Error
	return officials, err
}
func (r *masterDataRepository) UpdateOfficial(off *models.Official) error {
	return r.db.Save(off).Error
}
