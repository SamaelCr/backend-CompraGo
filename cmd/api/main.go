package main

import (
	"log"
	"github.com/toor/backend/internal/config"
	"github.com/toor/backend/internal/storage"
	"github.com/toor/backend/internal/router"
)

func main() {
	cfg := config.Load()
	db := storage.MustInit(cfg.DSN)
	r := router.New(db)
	log.Fatal(r.Run(":8080"))
}