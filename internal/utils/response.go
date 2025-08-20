package utils

import (
	"github.com/gin-gonic/gin"
)

// WriteJSON envía una respuesta JSON estándar para casos de éxito.
// Simplifica el código en los handlers y mantiene la consistencia.
func WriteJSON(c *gin.Context, status int, data any) {
	c.JSON(status, data)
}

// APIError define la estructura estándar para los errores en la API.
// Esto permite que el frontend siempre espere el mismo formato.
type APIError struct {
	Message string `json:"message"`
}

// WriteError envía una respuesta de error JSON estandarizada.
// Usa AbortWithStatusJSON para detener la ejecución de otros manejadores en la cadena.
func WriteError(c *gin.Context, status int, message string) {
	errResponse := APIError{Message: message}
	c.AbortWithStatusJSON(status, gin.H{"error": errResponse})
}
