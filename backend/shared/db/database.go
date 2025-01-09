package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config struct to store database configuration
type Config struct {
	Host            string
	Port            int
	User            string
	Password        string
	DBName          string
	MaxIdleConns    int           // Jumlah koneksi idle maksimum
	MaxOpenConns    int           // Jumlah koneksi maksimum
	ConnMaxLifetime time.Duration // Durasi maksimal koneksi hidup
}

// Global variable to store the database connection
var DB *gorm.DB

// InitDatabase initializes the database connection with connection pooling
func InitDatabase(config Config, models ...interface{}) {
	sqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	// Get generic database object sql.DB untuk konfigurasi pooling
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Error getting database instance: ", err)
	}

	// Konfigurasi connection pooling
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)       // Koneksi idle maksimum
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)       // Total koneksi maksimum
	sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime) // Durasi maksimum koneksi

	// Auto-migrate provided models
	if len(models) > 0 {
		err = DB.AutoMigrate(models...)
		if err != nil {
			log.Fatal("Error during migration: ", err)
		}
	}

	log.Println("Database connected successfully with connection pooling!")
}
