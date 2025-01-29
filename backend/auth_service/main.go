package main

import (
	"auth_service/config"
	"auth_service/models"
	"auth_service/repository"
	"auth_service/server"
	"auth_service/services"
)

func main() {
	// Load konfigurasi database
	cfg := config.LoadConfig()

	// Inisialisasi database
	config.InitDatabase(cfg)

	// Migrasi model
	config.DB.AutoMigrate(&models.User{}, &models.Sekolah{}, &models.UserProfile{}, &models.SekolahIndonesia{})
	userRepo := repository.NewUserRepository(config.DB)
	userService := services.NewAuthService(userRepo)
	userProfileRepo := repository.NewUserProfileRepository(config.DB)
	sekolahRepo := repository.NewSekolahRepository(config.DB)
	sekolahService := services.NewSekolahService(sekolahRepo)
	sekolahIndonesiaRepo := repository.NewSekolahIndonesiaRepository(config.DB)
	sekolahIndonesiaService := services.NewSekolahIndonesiaService(sekolahIndonesiaRepo)
	userProfileService := services.NewUserProfileService(userProfileRepo)

	// Start GRPC Server
	server.StartGRPCServer(userService, sekolahService, userProfileService, sekolahIndonesiaService)
}
