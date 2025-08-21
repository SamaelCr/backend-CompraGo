package models

import (
	"time" // <-- AÑADIR IMPORT

	"gorm.io/gorm"
)

// Provider representa el modelo de datos para un proveedor.
type Provider struct {
	ID uint `gorm:"primarykey" json:"id"`
	// MODIFICAR ESTAS DOS LÍNEAS
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name    string `gorm:"not null" json:"name"`
	RIF     string `gorm:"uniqueIndex:idx_providers_rif_active,where:deleted_at IS NULL" json:"rif"`
	Address string `json:"address"`
}
