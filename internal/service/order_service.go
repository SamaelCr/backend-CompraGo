package service

import (
	"github.com/toor/backend/internal/models"
	"github.com/toor/backend/internal/repository"
)

type OrderService interface {
	CreateOrder(order *models.Order) (*models.Order, error)
	GetAllOrders() ([]models.Order, error)
}

type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{repo: repo}
}

// Aquí irá la lógica de negocio. Por ahora, solo llamamos al repositorio.
func (s *orderService) CreateOrder(order *models.Order) (*models.Order, error) {
    // --- LÓGICA DE NEGOCIO AÑADIDA ---
    // Propuesta de mejora: Calcular IVA y Total automáticamente
    order.IvaAmount = order.BaseAmount * 0.16 // Asumimos 16% de IVA
    order.TotalAmount = order.BaseAmount + order.IvaAmount

    // Propuesta de mejora: Arrastrar información
    if order.Subject == "" {
        order.Subject = order.Concept
    }
    if order.Synthesis == "" {
        order.Synthesis = "Basado en la necesidad de: " + order.Concept
    }
    // ------------------------------------

	return s.repo.CreateOrder(order)
}

func (s *orderService) GetAllOrders() ([]models.Order, error) {
	return s.repo.GetAllOrders()
}