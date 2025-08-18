package service

import (
	"github.com/toor/backend/internal/models"
	"github.com/toor/backend/internal/repository"
)

// Interfaces separadas para claridad, implementadas por un solo servicio.
type UnitService interface {
	CreateUnit(unit *models.Unit) (*models.Unit, error)
	GetAllUnits() ([]models.Unit, error)
	UpdateUnit(id uint, req *models.Unit) (*models.Unit, error)
}

// ... (se podrían crear interfaces para Position y Official si crecen en complejidad)

type MasterDataService interface {
	UnitService // Embeber la interfaz
	// Positions
	CreatePosition(pos *models.Position) (*models.Position, error)
	GetAllPositions() ([]models.Position, error)
	UpdatePosition(id uint, req *models.Position) (*models.Position, error)
	// Officials
	CreateOfficial(off *models.Official) (*models.Official, error)
	GetAllOfficials() ([]models.Official, error)
	UpdateOfficial(id uint, req *models.Official) (*models.Official, error)
}

type masterDataService struct {
	repo repository.MasterDataRepository
}

func NewMasterDataService(repo repository.MasterDataRepository) MasterDataService {
	return &masterDataService{repo: repo}
}

// Implementaciones...
func (s *masterDataService) CreateUnit(unit *models.Unit) (*models.Unit, error) {
	err := s.repo.CreateUnit(unit)
	return unit, err
}
func (s *masterDataService) GetAllUnits() ([]models.Unit, error) { return s.repo.GetAllUnits() }
func (s *masterDataService) UpdateUnit(id uint, req *models.Unit) (*models.Unit, error) {
	req.ID = id
	err := s.repo.UpdateUnit(req)
	return req, err
}

func (s *masterDataService) CreatePosition(pos *models.Position) (*models.Position, error) {
	err := s.repo.CreatePosition(pos)
	return pos, err
}
func (s *masterDataService) GetAllPositions() ([]models.Position, error) {
	return s.repo.GetAllPositions()
}
func (s *masterDataService) UpdatePosition(id uint, req *models.Position) (*models.Position, error) {
	req.ID = id
	err := s.repo.UpdatePosition(req)
	return req, err
}
func (s *masterDataService) CreateOfficial(off *models.Official) (*models.Official, error) {
	err := s.repo.CreateOfficial(off)
	return off, err
}
func (s *masterDataService) GetAllOfficials() ([]models.Official, error) {
	return s.repo.GetAllOfficials()
}
func (s *masterDataService) UpdateOfficial(id uint, req *models.Official) (*models.Official, error) {
	req.ID = id
	err := s.repo.UpdateOfficial(req)
	// Recargar para obtener los datos de Unit y Position
	updated, _ := s.repo.GetAllOfficials() // Simplificación, idealmente sería GetByID
	for _, o := range updated {
		if o.ID == id {
			return &o, err
		}
	}
	return req, err
}
