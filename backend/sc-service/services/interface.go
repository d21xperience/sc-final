package services

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

// Antarmuka umum untuk Ethereum dan Quorum
type EthClient interface {
	NetworkID(ctx context.Context) (*big.Int, error)
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	GenerateNewAccount() (string, string, error) 
}

// Antarmuka untuk fitur Quorum tambahan
type QuorumClient interface {
	EthClient
	SendPrivateTransaction(ctx context.Context, tx *types.Transaction) (string, error) // Menggunakan *types.Transaction
	GetPrivateContract(ctx context.Context, address string, payload []byte) ([]byte, error)
	GetConsensusAlgorithm(ctx context.Context) (string, error)
}
