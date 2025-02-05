package services

import (
	"errors"
	"os"
	"strconv"
)

// Config untuk koneksi blockchain
type Config struct {
	NetworkId      uint32 // Untuk Ethereum & Quorum
	BlockchainType string // "ethereum", "quorum", atau "hyperledger"
	RPCURL         string // URL RPC untuk Ethereum/Quorum

	// Untuk Hyperledger Fabric
	FabricConfigPath string
	FabricWallet     string
	FabricIdentity   string
}

// LoadConfig membaca environment variables
func LoadConfig() (*Config, error) {
	blockchainType := os.Getenv("BLOCKCHAIN_TYPE")
	rpcURL := os.Getenv("RPC_URL")
	networkIDStr := os.Getenv("NETWORK_ID")
	fabricConfigPath := os.Getenv("FABRIC_CONFIG_PATH")
	fabricWallet := os.Getenv("FABRIC_WALLET")
	fabricIdentity := os.Getenv("FABRIC_IDENTITY")
	// network.
	// Konversi Network ID ke uint32 (hanya untuk EVM-based)
	var networkID uint32
	if networkIDStr != "" {
		id, err := strconv.Atoi(networkIDStr)
		if err != nil {
			return nil, errors.New("NETWORK_ID harus berupa angka")
		}
		networkID = uint32(id)
	}

	// Validasi parameter berdasarkan jenis blockchain
	switch blockchainType {
	case "ethereum", "quorum":
		if rpcURL == "" {
			return nil, errors.New("RPC_URL harus diisi untuk Ethereum/Quorum")
		}
	case "hyperledger":
		if fabricConfigPath == "" || fabricWallet == "" || fabricIdentity == "" {
			return nil, errors.New("FABRIC_CONFIG_PATH, FABRIC_WALLET, dan FABRIC_IDENTITY harus diisi untuk Hyperledger Fabric")
		}
	default:
		return nil, errors.New("BLOCKCHAIN_TYPE tidak valid: gunakan 'ethereum', 'quorum', atau 'hyperledger'")
	}

	return &Config{
		NetworkId:        networkID,
		BlockchainType:   blockchainType,
		RPCURL:           rpcURL,
		FabricConfigPath: fabricConfigPath,
		FabricWallet:     fabricWallet,
		FabricIdentity:   fabricIdentity,
	}, nil
}
