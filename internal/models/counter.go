IGNORE_WHEN_COPYING_START
IGNORE_WHEN_COPYING_END

    
// internal/models/counter.go
package models

import "gorm.io/gorm"

// SystemCounter almacena el estado de los contadores de documentos del sistema.
type SystemCounter struct {
	ID           uint   `gorm:"primarykey"`
	DocumentType string `gorm:"uniqueIndex;not null"` // Ej: "MEMO", "ORDER", "ACCOUNT_POINT"
	CurrentYear  int    `gorm:"not null"`
	LastSequence uint   `gorm:"not null"`
	gorm.Model   `gorm:"-"` // Para no incluir los campos por defecto de gorm si no los necesitas
}