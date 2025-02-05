package services

import (
	"errors"
	"fmt"
)

// BlockchainClient interface umum untuk semua blockchain
type BlockchainClient interface {
	Connect() error
}

// BlockchainClientFactory mendefinisikan factory function
type BlockchainClientFactory func(cfg *Config) (BlockchainClient, error)

// Peta blockchain factories
var blockchainFactories = map[string]BlockchainClientFactory{
	"ethereum": NewEthereumClient,
	// "quorum":   NewQuorumClient,
	// "hyperledger": NewHyperledgerFabricClient,
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

// // NewQuorumClient membuat client Quorum
// func NewQuorumClient(cfg *Config) (BlockchainClient, error) {
// 	if cfg.RPCURL == "" {
// 		return nil, errors.New("Quorum node URL is required")
// 	}
// 	return &QuorumClient{cfg.RPCURL}, nil
// }


// // QuorumClient adalah implementasi BlockchainClient untuk Quorum
// type QuorumClient struct {
// 	NodeURL string
// }

// // Connect implementasi koneksi ke Quorum
// func (q *QuorumClient) Connect() error {
// 	fmt.Println("Connecting to Quorum node at", q.NodeURL)
// 	return nil
// }
