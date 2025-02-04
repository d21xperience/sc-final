package services

import (
	"errors"
	"os"
)

// Config untuk menentukan jenis blockchain dan URL
type Config struct {
	NetworkId      uint32
	BlockchainType string // "ethereum" atau "quorum"
	RPCURL         string // URL RPC node
}

// LoadConfig membaca konfigurasi dari environment variables
func LoadConfig() (*Config, error) {
	blockchainType := os.Getenv("BLOCKCHAIN_TYPE") // ethereum atau quorum
	rpcURL := os.Getenv("RPC_URL")                 // URL untuk RPC

	if blockchainType == "" || rpcURL == "" {
		return nil, errors.New("BLOCKCHAIN_TYPE dan RPC_URL harus diset")
	}

	return &Config{
		BlockchainType: blockchainType,
		RPCURL:         rpcURL,
	}, nil
}
