package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Host string
	Port int
	// User            string
	// Password        string
	DBName          string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	BaseURL         string
	GRPCHost        string
	GRPCPort        int
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port, _ := strconv.Atoi(os.Getenv("DAPODIK_PORT"))
	grpcPort, _ := strconv.Atoi(os.Getenv("GRPC_PORT"))
	return Config{
		GRPCHost: os.Getenv("GRPC_HOST"),
		GRPCPort: grpcPort,
		BaseURL:  os.Getenv("BASE_URL"),
		Host:     os.Getenv("DAPODIK_HOST"),
		// Password:        os.Getenv("DAPODIK_PASSWORD"),
		Port: port,
		// User:            os.Getenv("DAPODIK_USER"),
		DBName:          os.Getenv("DAPODIK_DB"),
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: 30 * time.Minute,
	}
}
