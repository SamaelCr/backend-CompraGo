package service

import (
	"errors"
	"strconv"

	"github.com/toor/backend/internal/models"
	"github.com/toor/backend/internal/repository"
	"gorm.io/gorm"
)

const IVAPercentageKey = "iva_percentage"

type SettingService interface {
	GetIVAPercentage() (float64, error)
	UpdateIVAPercentage(value float64) error
}

type settingService struct {
	repo repository.SettingRepository
}

func NewSettingService(repo repository.SettingRepository) SettingService {
	return &settingService{repo: repo}
}

func (s *settingService) GetIVAPercentage() (float64, error) {
	setting, err := s.repo.GetValueByKey(IVAPercentageKey)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Si no existe, creamos el valor por defecto de 16%
			if err := s.UpdateIVAPercentage(16.0); err != nil {
				return 0, err
			}
			return 16.0, nil
		}
		return 0, err
	}

	iva, err := strconv.ParseFloat(setting.Value, 64)
	if err != nil {
		return 0, errors.New("invalid IVA percentage stored in database")
	}
	return iva, nil
}

func (s *settingService) UpdateIVAPercentage(value float64) error {
	setting := &models.Setting{
		Key:   IVAPercentageKey,
		Value: strconv.FormatFloat(value, 'f', -1, 64),
	}
	return s.repo.Upsert(setting)
}
