package repository

import (
	"github.com/toor/backend/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SettingRepository interface {
	GetValueByKey(key string) (*models.Setting, error)
	Upsert(setting *models.Setting) error
}

type settingRepository struct {
	db *gorm.DB
}

func NewSettingRepository(db *gorm.DB) SettingRepository {
	return &settingRepository{db: db}
}

func (r *settingRepository) GetValueByKey(key string) (*models.Setting, error) {
	var setting models.Setting
	if err := r.db.Where("key = ?", key).First(&setting).Error; err != nil {
		return nil, err
	}
	return &setting, nil
}

func (r *settingRepository) Upsert(setting *models.Setting) error {
	// Inserta o actualiza el valor basado en la clave única.
	return r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},
		DoUpdates: clause.AssignmentColumns([]string{"value"}),
	}).Create(setting).Error
}
