package repository

import (
	"errors"
	"fmt"

	"github.com/toor/backend/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CounterRepository interface {
	GetNextSequence(docType string, year int) (uint, error)
	ResetAllCounters(newYear int) error
}

type counterRepository struct {
	db *gorm.DB
}

func NewCounterRepository(db *gorm.DB) CounterRepository {
	return &counterRepository{db: db}
}

// GetNextSequence obtiene el siguiente número de forma segura (transaccional).
func (r *counterRepository) GetNextSequence(docType string, year int) (uint, error) {
	var counter models.SystemCounter
	var nextSequence uint

	err := r.db.Transaction(func(tx *gorm.DB) error {
		// Bloquea la fila para evitar que dos peticiones incrementen el contador al mismo tiempo
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("document_type = ? AND current_year = ?", docType, year).
			First(&counter).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Si no existe un contador para este tipo y año, lo creamos
			counter = models.SystemCounter{
				DocumentType: docType,
				CurrentYear:  year,
				LastSequence: 0,
			}
			if createErr := tx.Create(&counter).Error; createErr != nil {
				return createErr
			}
		} else if err != nil {
			return err
		}

		// Incrementa la secuencia y la guarda
		counter.LastSequence++
		if saveErr := tx.Save(&counter).Error; saveErr != nil {
			return saveErr
		}

		nextSequence = counter.LastSequence
		return nil
	})

	if err != nil {
		return 0, err
	}

	return nextSequence, nil
}

// ResetAllCounters actualiza el año y reinicia la secuencia para todos los contadores.
func (r *counterRepository) ResetAllCounters(newYear int) error {
	// Esto es una simplificación. Una lógica más robusta crearía los contadores del nuevo año si no existen.
	// Por ahora, asumimos que se crean al primer uso.
	// Esta función puede expandirse para "pre-crear" contadores para el nuevo año.
	// Por simplicidad, el reinicio es implícito: cuando `GetNextSequence` sea llamado con un nuevo año, creará un nuevo registro.
	// Dejamos esta función como un placeholder para lógica de cierre más compleja si es necesaria.
	fmt.Printf("Lógica de reinicio para el año %d activada. Los nuevos contadores se crearán bajo demanda.\n", newYear)
	return nil
}