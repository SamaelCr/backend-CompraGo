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
	DeleteUnit(id uint) error
	// Positions
	CreatePosition(pos *models.Position) error
	GetAllPositions() ([]models.Position, error)
	UpdatePosition(pos *models.Position) error
	DeletePosition(id uint) error
	// Officials
	CreateOfficial(off *models.Official) error
	GetAllOfficials() ([]models.Official, error)
	GetOfficialByID(id uint) (*models.Official, error) // <-- MODIFICACIÓN: Añadida nueva función
	UpdateOfficial(off *models.Official) error
	DeleteOfficial(id uint) error

	IsUnitInUse(unitID uint) (bool, error)
	IsPositionInUse(positionID uint) (bool, error)
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
func (r *masterDataRepository) DeleteUnit(id uint) error {
	return r.db.Delete(&models.Unit{}, id).Error
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
func (r *masterDataRepository) DeletePosition(id uint) error {
	return r.db.Delete(&models.Position{}, id).Error
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

// GetOfficialByID recupera un único funcionario y precarga sus relaciones.
// MODIFICACIÓN: Añadida nueva función
func (r *masterDataRepository) GetOfficialByID(id uint) (*models.Official, error) {
	var official models.Official
	// Usamos Preload para traer los datos anidados de Unit y Position
	err := r.db.Preload("Unit").Preload("Position").First(&official, id).Error
	return &official, err
}

func (r *masterDataRepository) UpdateOfficial(off *models.Official) error {
	return r.db.Save(off).Error
}

func (r *masterDataRepository) DeleteOfficial(id uint) error {
	return r.db.Delete(&models.Official{}, id).Error
}
func (r *masterDataRepository) IsUnitInUse(unitID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.Official{}).Where("unit_id = ?", unitID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *masterDataRepository) IsPositionInUse(positionID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.Official{}).Where("position_id = ?", positionID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
