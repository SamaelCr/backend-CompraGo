package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DSN string
}

func Load() *Config {
	// Intentamos cargar el .env, pero no tratamos la ausencia como un error fatal,
	// ya que las variables pueden ser provistas por el sistema (ej. Docker Compose).
	// Se elimina el log de aquí para que el punto de entrada (main) controle todo el logging.
	_ = godotenv.Load()

	return &Config{
		DSN: os.Getenv("DSN"),
	}
}
