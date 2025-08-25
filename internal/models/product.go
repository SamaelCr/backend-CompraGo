package models

import (
	"time"
)

// Product representa un ítem en el catálogo de productos y servicios.
type Product struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	// CAMBIO: Se eliminó gorm.DeletedAt para deshabilitar el borrado lógico

	Name       string `gorm:"uniqueIndex;not null" json:"name"` // Volvemos a un uniqueIndex simple
	Unit       string `gorm:"not null" json:"unit"`
	IsActive   bool   `gorm:"default:true" json:"isActive"`
	AppliesIVA bool   `gorm:"default:true;not null" json:"appliesIva"`
}
