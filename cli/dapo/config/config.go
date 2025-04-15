package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	BaseURL     string
	Token       string
	OutputDir   string
	SemesterID  string
	NPSN        string
	GRPCHost    string
	GRPCPort    string
	GRPCTimeout time.Duration
}

func LoadConfig() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	timeoutSeconds, err := strconv.Atoi(os.Getenv("GRPC_TIMEOUT_SECONDS"))
	if err != nil {
		log.Fatalf("Invalid GRPC_TIMEOUT_SECONDS: %v", err)
	}

	return &AppConfig{
		BaseURL:     os.Getenv("BASE_URL"),
		Token:       os.Getenv("TOKEN"),
		OutputDir:   os.Getenv("OUTPUT_DIR"),
		SemesterID:  os.Getenv("SEMESTER_ID"),
		NPSN:        os.Getenv("NPSN"),
		GRPCHost:    os.Getenv("GRPC_HOST"),
		GRPCPort:    os.Getenv("GRPC_PORT"),
		GRPCTimeout: time.Duration(timeoutSeconds) * time.Second,
	}
}
