package models

import (
	"time" // <-- AÑADIR IMPORT

	"gorm.io/gorm"
)

// Position representa un Cargo dentro de la organización.
type Position struct {
	ID uint `gorm:"primarykey" json:"id"`
	// MODIFICAR ESTAS DOS LÍNEAS
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name     string `gorm:"uniqueIndex:idx_positions_name_active,where:deleted_at IS NULL;not null" json:"name"`
	IsActive bool   `gorm:"default:true" json:"isActive"`
}
