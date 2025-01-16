package main

import (
	"sekolah/config"
	"sekolah/models"
	"sekolah/repositories"
	"sekolah/server"
	"sekolah/services"
)

func main() {
	// Load konfigurasi database
	cfg := config.LoadConfig()
	// Inisialisasi database
	config.InitDatabase(cfg)
	// Migrasi model
	config.DB.AutoMigrate(&models.SekolahTabelTenant{})
	// =========================Inisialisasi Awal =======
	sekolahRepo := repositories.NewSekolahRepository(config.DB)
	sekolahService := services.NewSekolahService(sekolahRepo)
	schemaRepo := repositories.NewSchemaRepository(config.DB)
	sekolahTenantRepo := repositories.NewsekolahTenantRepository(config.DB)
	schemaService := services.NewSchemaService(schemaRepo, sekolahTenantRepo)
	// =============================================================================
	pesertaDidikRepo := repositories.NewpesertaDidikRepository(config.DB)
	pesertaDidikService := services.NewPesertaDidikService(pesertaDidikRepo)
	// // nilaiAkhirRepo := repositories.NewNilaiAkhirRepository(config.DB)
	// // nilaiAkhirService := services.NewNilaiAkhirService(nilaiAkhirRepo)
	// // Start GRPC Server
	server.StartGRPCServer(schemaService, sekolahService, pesertaDidikService)
}
