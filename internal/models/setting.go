package models

import "gorm.io/gorm"

// Setting almacena configuraciones globales de la aplicación como pares clave-valor.
type Setting struct {
	gorm.Model
	Key   string `gorm:"uniqueIndex;not null"`
	Value string `gorm:"not null"`
}
