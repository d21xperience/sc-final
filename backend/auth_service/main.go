package main

import (
	"auth_service/config"
	"auth_service/server"
	"auth_service/worker"
)

func main() {
	// Load konfigurasi database
	cfg := config.LoadConfig()

	// Inisialisasi database
	config.InitDatabase(cfg)
	// config.DB.AutoMigrate(&models.Sekolah{}, &models.User{}, &models.UserProfile{})

	// Start redis server
	config.InitRedisClient(cfg)

	// Jalankan worker di background
	go worker.StartSekolahInitWorker(config.RDB)

	// Start GRPC Server
	server.StartGRPCServer()
}
