package service

import (
	"github.com/toor/backend/internal/models"
	"github.com/toor/backend/internal/repository"
)

type ProviderService interface {
	CreateProvider(provider *models.Provider) (*models.Provider, error)
	GetAllProviders() ([]models.Provider, error)
	GetProviderByID(id uint) (*models.Provider, error)
	UpdateProvider(provider *models.Provider) (*models.Provider, error)
	DeleteProvider(id uint) error
}

type providerService struct {
	repo repository.ProviderRepository
}

func NewProviderService(repo repository.ProviderRepository) ProviderService {
	return &providerService{repo: repo}
}

func (s *providerService) CreateProvider(provider *models.Provider) (*models.Provider, error) {
	if err := s.repo.Create(provider); err != nil {
		return nil, err
	}
	return provider, nil
}

func (s *providerService) GetAllProviders() ([]models.Provider, error) {
	return s.repo.GetAll()
}

func (s *providerService) GetProviderByID(id uint) (*models.Provider, error) {
	return s.repo.GetByID(id)
}

func (s *providerService) UpdateProvider(provider *models.Provider) (*models.Provider, error) {
	if err := s.repo.Update(provider); err != nil {
		return nil, err
	}
	return provider, nil
}

func (s *providerService) DeleteProvider(id uint) error {
	return s.repo.Delete(id)
}