package main

import (
	"auth_service/config"
	"auth_service/server"
	"auth_service/models"
)

func main() {
	// Load konfigurasi database
	cfg := config.LoadConfig()

	// Inisialisasi database
	config.InitDatabase(cfg)
	config.DB.AutoMigrate(&models.Sekolah{}, &models.User{}, &models.UserProfile{})
	// Start GRPC Server
	server.StartGRPCServer()
}
