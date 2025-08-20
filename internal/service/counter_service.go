package service

import (
	"fmt"
	"time"

	"github.com/toor/backend/internal/repository"
)

type CounterService interface {
	GenerateNextID(docType string) (string, error)
	PerformAnnualReset(newYear int) error
}

type counterService struct {
	repo repository.CounterRepository
}

func NewCounterService(repo repository.CounterRepository) CounterService {
	return &counterService{repo: repo}
}

func (s *counterService) GenerateNextID(docType string) (string, error) {
	currentYear := time.Now().Year()
	sequence, err := s.repo.GetNextSequence(docType, currentYear)
	if err != nil {
		return "", err
	}
	// Formato: TIPO-AÑO-SECUENCIA (ej. MEMO-2024-00001)
	return fmt.Sprintf("%s-%d-%05d", docType, currentYear, sequence), nil
}

func (s *counterService) PerformAnnualReset(newYear int) error {
	// La lógica real está en el repositorio, que crea nuevos contadores por año.
	// El servicio podría añadir lógica adicional, como validar que el newYear sea futuro.
	return s.repo.ResetAllCounters(newYear)
}
