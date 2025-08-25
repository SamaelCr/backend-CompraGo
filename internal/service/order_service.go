package service

import (
	"bytes"
	"fmt"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/toor/backend/internal/models"
	"github.com/toor/backend/internal/repository"
	"gorm.io/gorm"
)

type OrderService interface {
	CreateOrder(order *models.Order) (*models.Order, error)
	GetAllOrders() ([]models.Order, error)
	GetOrderById(id uint) (*models.Order, error)
	GenerateOrderPDF(id uint) (*bytes.Buffer, error)
}

type orderService struct {
	repo           repository.OrderRepository
	counterService CounterService
	db             *gorm.DB
}

func NewOrderService(repo repository.OrderRepository, counterService CounterService, db *gorm.DB) OrderService {
	return &orderService{
		repo:           repo,
		counterService: counterService,
		db:             db,
	}
}

func (s *orderService) CreateOrder(order *models.Order) (*models.Order, error) {
	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	newMemoNumber, err := s.counterService.GenerateNextID("MEMO")
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("could not generate memo number: %w", err)
	}
	order.MemoNumber = newMemoNumber

	var baseAmount float64
	for i := range order.Items {
		item := &order.Items[i]
		item.Total = item.Quantity * item.UnitPrice
		baseAmount += item.Total
	}
	order.BaseAmount = baseAmount
	order.IvaAmount = baseAmount * 0.16
	order.TotalAmount = order.BaseAmount + order.IvaAmount

	createdOrder, err := s.repo.CreateOrderInTx(tx, order)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return createdOrder, nil
}

func (s *orderService) GetAllOrders() ([]models.Order, error) {
	return s.repo.GetAllOrders()
}

func (s *orderService) GetOrderById(id uint) (*models.Order, error) {
	return s.repo.GetOrderById(id)
}

func (s *orderService) GenerateOrderPDF(id uint) (*bytes.Buffer, error) {
	order, err := s.repo.GetOrderById(id)
	if err != nil {
		return nil, err
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	pdf.Cell(0, 10, "Orden de Compra / Servicio")
	pdf.Ln(12)

	writeRow := func(label, value string) {
		pdf.SetFont("Arial", "B", 10)
		pdf.CellFormat(40, 7, label, "", 0, "L", false, 0, "")
		pdf.SetFont("Arial", "", 10)
		pdf.MultiCell(0, 7, value, "", "L", false)
		pdf.Ln(2)
	}

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 10, "1. Requisicion")
	pdf.Ln(8)

	writeRow("Numero Memo:", order.MemoNumber)
	writeRow("Fecha Memo:", time.Time(order.MemoDate).Format("02-01-2006"))
	writeRow("Unidad Solicitante:", order.RequestingUnit)
	writeRow("Funcionario:", order.ResponsibleOfficial)
	writeRow("Concepto:", order.Concept)
	pdf.Ln(5)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 10, "2. Cotizacion")
	pdf.Ln(8)
	writeRow("Proveedor:", order.Provider)
	writeRow("Nro Presupuesto:", order.BudgetNumber)
	writeRow("Fecha Presupuesto:", time.Time(order.BudgetDate).Format("02-01-2006"))
	pdf.Ln(5)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 10, "3. Punto de Cuenta")
	pdf.Ln(8)
	writeRow("Numero PC:", order.AccountPoint.AccountNumber)
	writeRow("Asunto:", order.AccountPoint.Subject)
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 10, "Resumen Monetario")
	pdf.Ln(8)
	writeRow("Monto Base:", fmt.Sprintf("%.2f", order.BaseAmount))
	writeRow("Monto IVA (16%):", fmt.Sprintf("%.2f", order.IvaAmount))
	writeRow("Monto Total:", fmt.Sprintf("%.2f", order.TotalAmount))

	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}
