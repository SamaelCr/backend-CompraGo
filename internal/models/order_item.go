package models

import "gorm.io/gorm"

// OrderItem representa una línea de ítem dentro de una orden.
type OrderItem struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	OrderID     uint           `gorm:"not null" json:"orderId"`
	Description string         `gorm:"not null" json:"description"`
	Unit        string         `gorm:"not null" json:"unit"`
	Quantity    float64        `gorm:"not null" json:"quantity"`
	UnitPrice   float64        `gorm:"not null" json:"unitPrice"`
	Total       float64        `gorm:"not null" json:"total"` // Calculado: Quantity * UnitPrice
}
