package service

import (
	"fmt"

	"github.com/toor/backend/internal/models"
	"github.com/toor/backend/internal/repository"
)

type AccountPointService interface {
	CreateAccountPoint(accountPoint *models.AccountPoint) (*models.AccountPoint, error)
	GetAllAccountPoints() ([]models.AccountPoint, error)
	GetAccountPointByID(id uint) (*models.AccountPoint, error)
	UpdateAccountPoint(accountPoint *models.AccountPoint) (*models.AccountPoint, error)
	DeleteAccountPoint(id uint) error
}

type accountPointService struct {
	repo           repository.AccountPointRepository
	counterService CounterService
}

func NewAccountPointService(repo repository.AccountPointRepository, counterService CounterService) AccountPointService {
	return &accountPointService{
		repo:           repo,
		counterService: counterService,
	}
}

func (s *accountPointService) CreateAccountPoint(ap *models.AccountPoint) (*models.AccountPoint, error) {
	newAccountNumber, err := s.counterService.GenerateNextID("PC")
	if err != nil {
		return nil, fmt.Errorf("could not generate account point number: %w", err)
	}
	ap.AccountNumber = newAccountNumber

	if err := s.repo.Create(ap); err != nil {
		return nil, err
	}
	return ap, nil
}

func (s *accountPointService) GetAllAccountPoints() ([]models.AccountPoint, error) {
	return s.repo.GetAll()
}

func (s *accountPointService) GetAccountPointByID(id uint) (*models.AccountPoint, error) {
	return s.repo.GetByID(id)
}

func (s *accountPointService) UpdateAccountPoint(ap *models.AccountPoint) (*models.AccountPoint, error) {
	if err := s.repo.Update(ap); err != nil {
		return nil, err
	}
	return ap, nil
}

func (s *accountPointService) DeleteAccountPoint(id uint) error {
	return s.repo.Delete(id)
}
