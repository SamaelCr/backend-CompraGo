package service

import (
	"github.com/toor/backend/internal/models"
	"github.com/toor/backend/internal/repository"
)

type OrderService interface {
	CreateOrder(order *models.Order) (*models.Order, error)
	GetAllOrders() ([]models.Order, error)
	GetOrderById(id uint) (*models.Order, error)
}

type orderService struct {
	repo           repository.OrderRepository
	counterService CounterService // <-- AÑADIDO: Inyección del servicio de contadores
}

// <-- MODIFICADO: El constructor ahora requiere el counterService
func NewOrderService(repo repository.OrderRepository, counterService CounterService) OrderService {
	return &orderService{
		repo:           repo,
		counterService: counterService,
	}
}

func (s *orderService) CreateOrder(order *models.Order) (*models.Order, error) {
	// --- LÓGICA DE NEGOCIO PARA GENERAR CORRELATIVO ---
	// Se genera un nuevo número de memorando automáticamente.
	// El valor que venga del frontend en `order.MemoNumber` será ignorado.
	newMemoNumber, err := s.counterService.GenerateNextID("MEMO")
	if err != nil {
		return nil, fmt.Errorf("could not generate memo number: %w", err)
	}
	order.MemoNumber = newMemoNumber
	// --------------------------------------------------

	// --- LÓGICA DE NEGOCIO EXISTENTE ---
	// Calcular IVA y Total automáticamente
	order.IvaAmount = order.BaseAmount * 0.16 // Asumimos 16% de IVA
	order.TotalAmount = order.BaseAmount + order.IvaAmount

	// Arrastrar información
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

func (s *orderService) GetOrderById(id uint) (*models.Order, error) {
	return s.repo.GetOrderById(id)
}