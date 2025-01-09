package main

import (
	"sc-service/config"
	"sc-service/controller"
	"sc-service/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://example.com"}, // Tambahkan domain yang diizinkan
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Inisialisasi Database
	config.InitDatabase()
	db := config.DB
	// config.SeedAccountBC()
	// config.SeedBCNetworks()
	// config.SeedStudent()
	// config.InitETHClient() // Inisialisasi client Ethereum
	// controller
	// ac := controller.NewAccountController(db)
	ro := routes.Route{
		// Route untuk account
		Ac:    controller.NewAccountController(db),
		IPFSc: controller.NewIPFSController(db),
		Bc:    controller.NewBlockchainController(db),
		Cc:    controller.NewCertificateController(db),
	}
	routes.RegisterRoutes(r, ro) // Inisialisasi routes API
	r.Run(":8081")               // Jalankan server pada port 8080
}
