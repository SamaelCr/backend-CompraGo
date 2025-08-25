package models

import (
	"time"

	"gorm.io/gorm"
)

// Product representa un ítem en el catálogo de productos y servicios.
type Product struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name     string `gorm:"uniqueIndex;not null" json:"name"`
	Unit     string `gorm:"not null" json:"unit"` // Ej: "Unidad", "Servicio", "Kg", "Caja"
	IsActive bool   `gorm:"default:true" json:"isActive"`
}
