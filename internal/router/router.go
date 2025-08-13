package router

import (
	"github.com/gin-gonic/gin"
	"github.com/toor/backend/internal/handlers"
)

func New(db interface{}) *gin.Engine {
	r := gin.Default()
	r.GET("/api/ping", handlers.Ping)
	return r
}