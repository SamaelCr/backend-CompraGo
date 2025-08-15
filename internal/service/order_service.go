package service

import (
	"fmt"
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
	counterService CounterService
}

func NewOrderService(repo repository.OrderRepository, counterService CounterService) OrderService {
	return &orderService{
		repo:           repo,
		counterService: counterService,
	}
}

func (s *orderService) CreateOrder(order *models.Order) (*models.Order, error) {
	// --- LÓGICA DE NEGOCIO PARA GENERAR CORRELATIVO ---
	newMemoNumber, err := s.counterService.GenerateNextID("MEMO")
	if err != nil {
		// Aquí es donde se usa `fmt.Errorf`, causando el error
		return nil, fmt.Errorf("could not generate memo number: %w", err)
	}
	order.MemoNumber = newMemoNumber
	// --------------------------------------------------

	// --- LÓGICA DE NEGOCIO EXISTENTE ---
	order.IvaAmount = order.BaseAmount * 0.16
	order.TotalAmount = order.BaseAmount + order.IvaAmount

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