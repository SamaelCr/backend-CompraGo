package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/toor/backend/internal/utils"
)

func Ping(c *gin.Context) {
	utils.WriteJSON(c, http.StatusOK, gin.H{"message": "pong"})
}
