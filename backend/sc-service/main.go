package main

import (
	"log"
	"sc-service/config"
	"sc-service/models"
	"sc-service/server"
)

func main() {
	// Load konfigurasi database
	cfg := config.LoadConfig()
	// Inisialisasi database
	config.InitDatabase(cfg)
	err := models.Migrate(config.DB)
	if err != nil {
		log.Printf("gagal: %v", err)
	}
	server.StartServer()
}
