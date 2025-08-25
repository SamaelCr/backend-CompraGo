package models

import (
	"time"

	"gorm.io/gorm"
)

// AccountPoint representa el modelo de datos para un Punto de Cuenta.
type AccountPoint struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	AccountNumber        string   `gorm:"uniqueIndex;not null" json:"accountNumber"`
	Date                 DateOnly `json:"date"` // <-- CAMBIO AQUÍ
	Subject              string   `json:"subject"`
	Synthesis            string   `gorm:"type:text" json:"synthesis"`
	ProgrammaticCategory string   `json:"programmaticCategory"`
	UEL                  string   `json:"uel"`
	Status               string   `gorm:"default:'Disponible'" json:"status"` // Ej: Disponible, Utilizado
}
