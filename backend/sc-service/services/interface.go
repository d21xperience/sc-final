package services

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Antarmuka umum untuk Ethereum dan Quorum
type EthClient interface {
	NetworkID(ctx context.Context) (*big.Int, error)
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	GenerateNewAccount(username, password string) ( string, error) 
	DeployContract(ctx context.Context, bytecode string, privateKey string, gasLimit uint64) (string, string, error)
	// SendTransactionToContract(ctx context.Context, contractAddress, abi, method string, params []string, privateKey string, gasLimit uint64) (string, error)
	SendTransaction(ctx context.Context, tx *types.Transaction) error
	CallContractMethod(ctx context.Context, contractAddress, abi, method string, params []string) (string, error)
	// GetTokenBalance(ctx context.Context, tokenAddress, ownerAddress string) (string, error)
	GetTokenBalance(ctx context.Context, tokenAddress, ownerAddress string) (*big.Int, error)
	TransferToken(ctx context.Context, tokenAddress, from, to string, amount string, privateKey string, gasLimit uint64) (string, error)
	GetContractEvents(ctx context.Context, contractAddress, abi, eventName string, fromBlock, toBlock uint64) ([]string, error)
	SendETH(ctx context.Context, privateKeyHex, toAddress string, amount *big.Int) (string, error)
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	GetContract(ctx context.Context, contractAddress string) (string, string, error)
}

// Antarmuka untuk fitur Quorum tambahan
type QuorumClient interface {
	SendPrivateTransaction(ctx context.Context, tx *types.Transaction) (string, error) // Menggunakan *types.Transaction
	GetPrivateContract(ctx context.Context, address string, payload []byte) ([]byte, error)
	GetConsensusAlgorithm(ctx context.Context) (string, error)
}
