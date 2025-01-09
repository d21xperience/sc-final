package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Host            string
	Port            int
	User            string
	Password        string
	DBName          string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))

	return Config{
		Host:            os.Getenv("POSTGRES_HOST"),
		Password:        os.Getenv("POSTGRES_PASSWORD"),
		Port:            port,
		User:            os.Getenv("POSTGRES_USER"),
		DBName:          os.Getenv("POSTGRES_DB"),
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifetime: 30 * time.Minute,
	}
}
