package repository

import (
	"github.com/toor/backend/internal/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrderInTx(tx *gorm.DB, order *models.Order) (*models.Order, error)
	GetAllOrders() ([]models.Order, error)
	GetOrderById(id uint) (*models.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrderInTx(tx *gorm.DB, order *models.Order) (*models.Order, error) {
	if err := tx.Omit("Items", "AccountPoint", "SignedBy").Create(order).Error; err != nil {
		return nil, err
	}
	for i := range order.Items {
		order.Items[i].OrderID = order.ID
	}
	if len(order.Items) > 0 {
		if err := tx.Create(&order.Items).Error; err != nil {
			return nil, err
		}
	}

	var fullOrder models.Order
	if err := tx.Preload("AccountPoint").Preload("Items").Preload("SignedBy").First(&fullOrder, order.ID).Error; err != nil {
		return nil, err
	}
	return &fullOrder, nil
}

func (r *orderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order // <-- LA COMILLA EXTRA HA SIDO ELIMINADA DE AQUÍ
	if err := r.db.Preload("AccountPoint").Preload("Items").Preload("SignedBy").Order("created_at desc").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *orderRepository) GetOrderById(id uint) (*models.Order, error) {
	var order models.Order
	if err := r.db.Preload("AccountPoint").Preload("Items").Preload("SignedBy").First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
