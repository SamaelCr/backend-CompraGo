package models

import "gorm.io/gorm"

// Unit representa una Unidad Solicitante (ej. Servicios Generales).
type Unit struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt int64          `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt int64          `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name     string `gorm:"uniqueIndex:idx_units_name_active,where:deleted_at IS NULL;not null" json:"name"`
	IsActive bool   `gorm:"default:true" json:"isActive"`
}
