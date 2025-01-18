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
	tahunAjaranRepo := repositories.NewTahunAjaranRepository(config.DB)
	tahunAjaranService := services.NewTahunAjaranService(tahunAjaranRepo)
	semesterRepo := repositories.NewSemesterRepository(config.DB)
	semesterService := services.NewSemesterService(semesterRepo)
	// =============================================================================
	pesertaDidikRepo := repositories.NewpesertaDidikRepository(config.DB)
	pesertaDidikService := services.NewPesertaDidikService(pesertaDidikRepo)

	server.StartGRPCServer(schemaService, sekolahService, tahunAjaranService, semesterService, pesertaDidikService)
}
