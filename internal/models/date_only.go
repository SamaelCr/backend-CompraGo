package models

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

// DateOnly es un tipo personalizado para manejar fechas en formato YYYY-MM-DD.
type DateOnly time.Time

const dateOnlyLayout = "2006-01-02"

// UnmarshalJSON implementa la interfaz json.Unmarshaler.
// Esto permite que Gin decodifique "2006-01-02" a nuestro tipo.
func (d *DateOnly) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		return err // No es una cadena JSON válida
	}
	t, err := time.Parse(dateOnlyLayout, s)
	if err != nil {
		return err
	}
	*d = DateOnly(t)
	return nil
}

// MarshalJSON implementa la interfaz json.Marshaler.
// Esto asegura que al enviar JSON, se formatee como "2006-01-02".
func (d DateOnly) MarshalJSON() ([]byte, error) {
	s := time.Time(d).Format(dateOnlyLayout)
	return []byte(`"` + s + `"`), nil
}

// Value implementa la interfaz driver.Valuer para GORM.
// Le dice a GORM cómo guardar este tipo en la base de datos.
func (d DateOnly) Value() (driver.Value, error) {
	return time.Time(d), nil
}

// Scan implementa la interfaz sql.Scanner para GORM.
// Le dice a GORM cómo leer este tipo desde la base de datos.
func (d *DateOnly) Scan(value interface{}) error {
	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("failed to scan DateOnly: value is not time.Time, but %T", value)
	}
	*d = DateOnly(t)
	return nil
}
