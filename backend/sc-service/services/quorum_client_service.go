package services

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
)

// QuorumClientService adalah implementasi untuk fitur Quorum
type QuorumClientService struct {
	client QuorumClient
}

// Constructor untuk QuorumClientService
func NewQuorumClientService(client QuorumClient) *QuorumClientService {
	return &QuorumClientService{client: client}
}

// Fungsi untuk mengirim transaksi privat
func (s *QuorumClientService) SendPrivateTransaction(ctx context.Context, tx *types.Transaction) (string, error) {
	return s.client.SendPrivateTransaction(ctx, tx)
}

// Fungsi untuk membaca kontrak privat
func (s *QuorumClientService) GetPrivateContract(ctx context.Context, address string, payload []byte) ([]byte, error) {
	return s.client.GetPrivateContract(ctx, address, payload)
}

// Fungsi untuk mendapatkan algoritma konsensus
func (s *QuorumClientService) GetConsensusAlgorithm(ctx context.Context) (string, error) {
	return s.client.GetConsensusAlgorithm(ctx)
}

// Default implementasi QuorumClient menggunakan go-ethereum dan RPC
type DefaultQuorumClient struct {
	rpcClient *rpc.Client
}

// Constructor untuk DefaultQuorumClient
func NewDefaultQuorumClient(rawUrl string) (*DefaultQuorumClient, error) {
	client, err := rpc.Dial(rawUrl)
	if err != nil {
		return nil, err
	}
	return &DefaultQuorumClient{rpcClient: client}, err
}

// Implementasi NetworkID (inherit dari Ethereum)
func (c *DefaultQuorumClient) NetworkID(ctx context.Context) (*big.Int, error) {
	var result string
	err := c.rpcClient.CallContext(ctx, &result, "net_version")
	if err != nil {
		return nil, err
	}
	id := new(big.Int)
	id.SetString(result, 10)
	return id, nil
}

// Implementasi SuggestGasPrice (inherit dari Ethereum)
func (c *DefaultQuorumClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	var result big.Int
	err := c.rpcClient.CallContext(ctx, &result, "eth_gasPrice")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Implementasi SendPrivateTransaction
func (c *DefaultQuorumClient) SendPrivateTransaction(ctx context.Context, tx *types.Transaction) (string, error) {
	var txHash string
	err := c.rpcClient.CallContext(ctx, &txHash, "eth_sendRawTransaction", tx)
	return txHash, err
}

// Implementasi GetPrivateContract
func (c *DefaultQuorumClient) GetPrivateContract(ctx context.Context, address string, payload []byte) ([]byte, error) {
	var result []byte
	err := c.rpcClient.CallContext(ctx, &result, "eth_call", map[string]interface{}{
		"to":   address,
		"data": payload,
	})
	return result, err
}

// Implementasi GetConsensusAlgorithm
func (c *DefaultQuorumClient) GetConsensusAlgorithm(ctx context.Context) (string, error) {
	var result string
	err := c.rpcClient.CallContext(ctx, &result, "istanbul_getSnapshot") // Untuk IBFT
	if err != nil {
		err = c.rpcClient.CallContext(ctx, &result, "raft_cluster") // Untuk Raft
	}
	return result, err
}

func (c *DefaultQuorumClient) GenerateNewAccount() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", fmt.Errorf("failed to generate key: %v", err)
	}
	privateKeyHex := hexutil.Encode(crypto.FromECDSA(privateKey))
	publicAddress := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	return privateKeyHex, publicAddress, nil
}
