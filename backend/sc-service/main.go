package main

import (
	"sc-service/config"
	"sc-service/models"
	"sc-service/server"
)

func main() {
	// Load konfigurasi database
	cfg := config.LoadConfig()
	// Inisialisasi database
	config.InitDatabase(cfg)
	config.DB.AutoMigrate(&models.WalletTable{})
	
	server.StartServer()
}
