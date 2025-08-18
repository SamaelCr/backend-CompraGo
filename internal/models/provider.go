package models

import "gorm.io/gorm"

// Provider representa el modelo de datos para un proveedor.
type Provider struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt int64          `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt int64          `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name    string `gorm:"not null" json:"name"`
	RIF     string `gorm:"uniqueIndex" json:"rif"`
	Address string `json:"address"`
}