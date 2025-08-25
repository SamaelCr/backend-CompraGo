package models

import (
	"time"

	"gorm.io/gorm"
)

// Order representa el modelo de datos para una orden de compra o servicio.
type Order struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// --- Requisición ---
	MemoDate            DateOnly `json:"memoDate"`
	MemoNumber          string   `json:"memoNumber"`
	RequestingUnit      string   `json:"requestingUnit"`
	ResponsibleOfficial string   `json:"responsibleOfficial"`
	Concept             string   `gorm:"type:text" json:"concept"`

	// --- Cotización ---
	Provider         string   `json:"provider"`
	DocumentType     string   `json:"documentType"`
	BudgetNumber     string   `json:"budgetNumber"`
	BudgetDate       DateOnly `json:"budgetDate"`
	DeliveryTime     string   `json:"deliveryTime"`
	OfferQuality     string   `json:"offerQuality"`
	PriceInquiryType string   `json:"priceInquiryType"` // Campo añadido

	// --- Campos Adicionales ---
	Observations    string `gorm:"type:text" json:"observations"`
	HasIvaRetention bool   `json:"hasIvaRetention"`
	HasIslr         bool   `json:"hasIslr"`
	HasItf          bool   `json:"hasItf"`

	// --- Firmante ---
	SignedByID uint     `json:"signedById"`
	SignedBy   Official `gorm:"foreignKey:SignedByID" json:"signedBy"`

	// --- Asociación ---
	AccountPointID uint         `json:"accountPointId"`
	AccountPoint   AccountPoint `json:"accountPoint"`

	// --- Ítems de la Orden ---
	Items []OrderItem `gorm:"foreignKey:OrderID" json:"items"`

	// --- Montos Calculados ---
	BaseAmount  float64 `json:"baseAmount"`
	IvaAmount   float64 `json:"ivaAmount"`
	TotalAmount float64 `json:"totalAmount"`

	Status string `gorm:"default:'En Proceso'" json:"status"`
}
