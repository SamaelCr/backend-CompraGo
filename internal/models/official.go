package models

import (
	"time" // <-- AÑADIR IMPORT

	"gorm.io/gorm"
)

// Official representa a un Funcionario responsable.
type Official struct {
	ID uint `gorm:"primarykey" json:"id"`
	// MODIFICAR ESTAS DOS LÍNEAS
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	FullName string `gorm:"not null" json:"fullName"`
	IsActive bool   `gorm:"default:true" json:"isActive"`

	UnitID uint `json:"unitId"`
	Unit   Unit `json:"unit"` // Relación para precargar datos

	PositionID uint     `json:"positionId"`
	Position   Position `json:"position"` // Relación para precargar datos
}
