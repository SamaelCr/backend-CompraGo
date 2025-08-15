package repository

import (
	"github.com/toor/backend/internal/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *models.Order) (*models.Order, error)
	GetAllOrders() ([]models.Order, error)
	GetOrderById(id uint) (*models.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(order *models.Order) (*models.Order, error) {
	if err := r.db.Create(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (r *orderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.Order("created_at desc").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) GetOrderById(id uint) (*models.Order, error) {
	var order models.Order
	// db.First buscará por clave primaria. Es crucial devolver el error
	// para que podamos manejar el 'not found' en la capa superior.
	if err := r.db.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}