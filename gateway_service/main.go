package main

// import (
// 	"gateway_service/routes"
// 	"time"

// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	// Initialize Gin
// 	r := gin.Default()
// 	r.Use(cors.New(cors.Config{
// 		AllowOrigins:     []string{"http://localhost:5173", "http://example.com"}, // Tambahkan domain yang diizinkan
// 		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
// 		AllowHeaders:     []string{"Content-Type", "Authorization"},
// 		ExposeHeaders:    []string{"Content-Length"},
// 		AllowCredentials: true,
// 		MaxAge:           12 * time.Hour,
// 	}))
// 	// Set up routes
// 	routes.SetupGatewayRoutes(r)

// 	// Start Gateway Service
// 	r.Run(":8081") // Gateway listens on port 8080
// }
