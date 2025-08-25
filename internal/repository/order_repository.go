package repository

import (
	"github.com/toor/backend/internal/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrderInTx(tx *gorm.DB, order *models.Order) (*models.Order, error)
	FindAll(params OrderSearchParams) ([]models.Order, int64, error)
	GetOrderById(id uint) (*models.Order, error)
	GetOrdersByAccountPointID(apID uint) ([]models.Order, error) // NUEVO
}

// OrderSearchParams define los parámetros para la búsqueda y paginación.
type OrderSearchParams struct {
	Keyword  string
	Provider string
	DateFrom string
	DateTo   string
	Limit    int
	Offset   int
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

func (r *orderRepository) FindAll(params OrderSearchParams) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64

	query := r.db.Model(&models.Order{})

	if params.Provider != "" {
		query = query.Where("provider = ?", params.Provider)
	}
	if params.DateFrom != "" {
		query = query.Where("memo_date >= ?", params.DateFrom)
	}
	if params.DateTo != "" {
		query = query.Where("memo_date <= ?", params.DateTo)
	}

	if params.Keyword != "" {
		searchTerm := "%" + params.Keyword + "%"
		query = query.Where(
			"memo_number ILIKE ? OR concept ILIKE ? OR provider ILIKE ? OR price_inquiry_type ILIKE ?",
			searchTerm, searchTerm, searchTerm, searchTerm,
		)
	}

	countQuery := query

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload("AccountPoint").Preload("SignedBy").
		Order("created_at desc").
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&orders).Error

	return orders, total, err
}

func (r *orderRepository) GetOrderById(id uint) (*models.Order, error) {
	var order models.Order
	if err := r.db.Preload("AccountPoint").Preload("Items").Preload("SignedBy").First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

// NUEVA FUNCIÓN
func (r *orderRepository) GetOrdersByAccountPointID(apID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("account_point_id = ?", apID).Order("created_at desc").Find(&orders).Error
	return orders, err
}
