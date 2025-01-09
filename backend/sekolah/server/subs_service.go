package server

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// 	"sekolah/models"

// 	"github.com/go-redis/redis"
// )

// // Fungsi untuk membaca environment variable dengan nilai default
// func getEnv(key, defaultValue string) string {
// 	if value, exists := os.LookupEnv(key); exists {
// 		return value
// 	}
// 	return defaultValue
// }

// // Langganan ke channel Redis
// func subscribeToChannel(rdb *redis.Client, channel string) *redis.PubSub {
// 	sub := rdb.Subscribe(channel)
// 	return sub
// }

// // Handler untuk memproses pesan dari Redis
// func handleMessage(msg *redis.Message) {
// 	var registration SchoolRegistration
// 	if err := json.Unmarshal([]byte(msg.Payload), &registration); err != nil {
// 		log.Printf("Failed to parse message: %v", err)
// 		return
// 	}

// 	fmt.Printf("Received registration: %+v\n", registration)

// 	// Buat database untuk sekolah yang baru
// 	sekolahRepo := repositories.NewSekolahRepository(models.Sekolah)
// 	if err := repositories.SekolahRepository.CreateSekolahDatabase(registration.SchoolName); err != nil {
// 		log.Printf("Failed to create database for school: %v", err)
// 	} else {
// 		fmt.Printf("Database created successfully for school: %s\n", registration.SchoolName)
// 	}
// }
