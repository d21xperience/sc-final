package main

import (
	"sekolah/config"
	"sekolah/models"
)

func main() {
	// Load konfigurasi database
	cfg := config.LoadConfig()

	// Inisialisasi database
	config.InitDatabase(cfg)

	// Migrasi model
	config.DB.AutoMigrate(
		&models.Sekolah{},
		&models.PesertaDidik{},
		&models.RegistrasiPesertaDidik{},
		&models.RombonganBelajar{},
		// &models.RombonganBelajar{},
		&models.TabelNilaiAkhir{},
	)
	// sekolahRepo := repositories.NewSekolahRepository(config.DB)
	// userService := services.
	// userProfileRepo := repository.NewUserProfileRepository(config.DB)
	// sekolahRepo := repository.NewSekolahRepository(config.DB)
	// sekolahService := services.NewSekolahService(sekolahRepo)
	// userProfileService := services.NewUserProfileService(userProfileRepo)
	// Start GRPC Server
	// server.StartGRPCServer(userService, sekolahService, userProfileService)
}
