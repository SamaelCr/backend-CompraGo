package models

import (
	"time"
	"gorm.io/gorm"
)

// Order representa el modelo de datos para una orden de compra o servicio.
// Los campos se han extraído de los requisitos de la minuta y del frontend.
type Order struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // Para soft delete

	// --- Paso 1: Requisición ---
	MemoDate            time.Time `json:"memoDate"`
	MemoNumber          string    `json:"memoNumber"`
	RequestingUnit      string    `json:"requestingUnit"`
	ResponsibleOfficial string    `json:"responsibleOfficial"`
	Concept             string    `gorm:"type:text" json:"concept"`

	// --- Paso 2: Cotización ---
	Provider      string    `json:"provider"`
	DocumentType  string    `json:"documentType"`
	BudgetNumber  string    `json:"budgetNumber"`
	BudgetDate    time.Time `json:"budgetDate"`
	BaseAmount    float64   `json:"baseAmount"`
	IvaAmount     float64   `json:"ivaAmount"` // Lo calcularemos en el backend
	TotalAmount   float64   `json:"totalAmount"` // (Base + IVA)
	DeliveryTime  string    `json:"deliveryTime"`
	OfferQuality  string    `json:"offerQuality"`

	// --- Paso 3: Punto de Cuenta ---
	AccountPointDate     time.Time `gorm:"autoCreateTime" json:"accountPointDate"` // Se genera automáticamente
	PriceInquiryType     string    `json:"priceInquiryType"`
	Subject              string    `json:"subject"`
	Synthesis            string    `gorm:"type:text" json:"synthesis"`
	ProgrammaticCategory string    `json:"programmaticCategory"`
	UEL                  string    `json:"uel"`

	// --- Paso 4: Orden ---
	// Ítems se manejarán en una tabla separada más adelante.
	Status string `gorm:"default:'En Proceso'" json:"status"`
}