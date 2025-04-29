package main

import (
	"sc-service/config"
	"sc-service/server"
)

func main() {
	// Load konfigurasi database
	cfg := config.LoadConfig()
	// Inisialisasi database
	config.InitDatabase(cfg)
	// config.DB.AutoMigrate(&models.SchemaLog{}, &models.SekolahTenant{})

	// err := models.Migrate(config.DB)
	// if err != nil {
	// 	log.Printf("gagal: %v", err)
	// }
	server.StartServer()
}
