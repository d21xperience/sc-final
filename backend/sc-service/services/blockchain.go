package services

import (
	"context"
	"errors"
	"fmt"
	"math/big"
)

// BlockchainClient interface umum untuk semua blockchain
type BlockchainClient interface {
	Connect() error
	NetworkID(ctx context.Context) (*big.Int, error)
	GenerateNewAccount(ctx context.Context, userId int32, password string) (map[string]interface{}, error)
}

// BlockchainClientFactory mendefinisikan factory function
type BlockchainClientFactory func(cfg *Config) (BlockchainClient, error)

// Peta blockchain factories
var blockchainFactories = map[string]BlockchainClientFactory{
	"ethereum": NewEthereumClient,
	// "quorum":   NewQuorumClient,
	"hyperledger": NewHyperledgerFabricClient,
}

// CreateClientFactory memilih blockchain berdasarkan runtime config
func CreateClientFactory(cfg *Config) (BlockchainClient, error) {
	if cfg == nil {
		return nil, errors.New("config tidak boleh nil")
	}

	factory, exists := blockchainFactories[cfg.BlockchainType]
	if !exists {
		return nil, fmt.Errorf("BlockchainType tidak dikenali: %s", cfg.BlockchainType)
	}
	return factory(cfg)
}

// SmartContractClient interface untuk semua blockchain yang mendukung smart contract
type SmartContractClient interface {
	IssueDegree(ctx context.Context, contractAddress string, degreeHash [32]byte, sekolah string, issueDate uint64, privateKey string, gasLimit uint64) (string, error)
}
