package models

import "gorm.io/gorm"

// Position representa un Cargo dentro de la organización.
type Position struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt int64          `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt int64          `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name     string `gorm:"uniqueIndex;not null" json:"name"`
	IsActive bool   `gorm:"default:true" json:"isActive"`
}
