package services

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

// Default implementasi QuorumClient menggunakan go-ethereum dan RPC
// type QuorumClient struct {
// 	rpcClient *rpc.Client
// 	ethClient EthClient
// }

// // Constructor untuk QuorumClient
//
//	func NewQuorumClient(rawUrl string, ethClient EthClient) (*QuorumClient, error) {
//		client, err := rpc.Dial(rawUrl)
//		if err != nil {
//			return nil, err
//		}
//		return &QuorumClient{rpcClient: client, ethClient: ethClient}, nil
//	}
//
// QuorumClient struct untuk menyimpan koneksi ke Quorum
type QuorumClient struct {
	rpcClient *rpc.Client
	ethClient EthClient
}

// Connect ke jaringan Quorum dan validasi koneksi
func (q *QuorumClient) Connect() error {
	var blockNumber string
	err := q.rpcClient.Call(&blockNumber, "eth_blockNumber")
	if err != nil {
		return fmt.Errorf("gagal mengambil block number: %w", err)
	}
	log.Printf("Terhubung ke Quorum - Block terbaru: %s", blockNumber)
	return nil
}

// NewQuorumClient membuat klien Quorum berdasarkan konfigurasi runtime
func NewQuorumClient(cfg *Config) (BlockchainClient, error) {
	client, err := rpc.Dial(cfg.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat klien Quorum: %w", err)
	}

	return &QuorumClient{
		rpcClient: client,
	}, nil
}

// Implementasi NetworkID (inherit dari Ethereum)
func (c *QuorumClient) NetworkID(ctx context.Context) (*big.Int, error) {
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
func (c *QuorumClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	var result big.Int
	err := c.rpcClient.CallContext(ctx, &result, "eth_gasPrice")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Implementasi SendPrivateTransaction
func (c *QuorumClient) SendPrivateTransaction(ctx context.Context, tx *types.Transaction) (string, error) {
	var txHash string
	err := c.rpcClient.CallContext(ctx, &txHash, "eth_sendRawTransaction", tx)
	return txHash, err
}

// Implementasi GetPrivateContract
func (c *QuorumClient) GetPrivateContract(ctx context.Context, address string, payload []byte) ([]byte, error) {
	var result []byte
	err := c.rpcClient.CallContext(ctx, &result, "eth_call", map[string]interface{}{
		"to":   address,
		"data": payload,
	})
	return result, err
}

// Implementasi GetConsensusAlgorithm
func (c *QuorumClient) GetConsensusAlgorithm(ctx context.Context) (string, error) {
	var result string
	err := c.rpcClient.CallContext(ctx, &result, "istanbul_getSnapshot") // Untuk IBFT
	if err != nil {
		err = c.rpcClient.CallContext(ctx, &result, "raft_cluster") // Untuk Raft
	}
	return result, err
}

// Implement CallContractMethod agar memenuhi EthClient
func (c *QuorumClient) CallContractMethod(ctx context.Context, contractAddress, abi, method string, params []string) (string, error) {
	return c.ethClient.CallContractMethod(ctx, contractAddress, abi, method, params)
}

// Implement GetTokenBalance
//
//	func (c *QuorumClient) GetTokenBalance(ctx context.Context, tokenAddress, ownerAddress string) (string, error) {
//		return c.ethClient.GetTokenBalance(ctx, tokenAddress, ownerAddress)
//	}
func (c *QuorumClient) GetTokenBalance(ctx context.Context, tokenAddress, ownerAddress string) (*big.Int, error) {
	return c.ethClient.GetTokenBalance(ctx, tokenAddress, ownerAddress)
}

func (c *QuorumClient) DeployContract(ctx context.Context, bytecode string, privateKey string, gasLimit uint64) (string, string, error) {
	return c.ethClient.DeployContract(ctx, bytecode, privateKey, gasLimit)
}
func (c *QuorumClient) GetContractEvents(ctx context.Context, contractAddress, abi, eventName string, fromBlock, toBlock uint64) ([]string, error) {
	return c.ethClient.GetContractEvents(ctx, contractAddress, abi, eventName, fromBlock, toBlock)
}
func (c *QuorumClient) SendETH(ctx context.Context, privateKeyHex, toAddress string, amount *big.Int) (string, error) {
	return c.ethClient.SendETH(ctx, privateKeyHex, toAddress, amount)
}

func (c *QuorumClient) TransferToken(ctx context.Context, tokenAddress, from, to string, amount string, privateKey string, gasLimit uint64) (string, error) {
	return c.ethClient.TransferToken(ctx, tokenAddress, from, to, amount, privateKey, gasLimit)
}

func (c *QuorumClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return c.ethClient.PendingNonceAt(ctx, account)
}
func (c *QuorumClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return c.ethClient.SendTransaction(ctx, tx)
}
func (c *QuorumClient) GetContract(ctx context.Context, contractAddress string) (string, string, error) {
	return c.ethClient.GetContract(ctx, contractAddress)

}
func (c *QuorumClient) GenerateNewAccount(ctx context.Context, userId int32, password string) (map[string]interface{}, error) {
	return c.ethClient.GenerateNewAccount(ctx, userId, password)
}

// func (c *QuorumClient) GetAccounts(ctx context.Context, userId int32, schemaname string) ([]*models.Account, error) {
// 	return c.ethClient.GetAccounts(ctx, userId, schemaname)

// }
func (c *QuorumClient) DeployIjazahContract(ctx context.Context, pvKey string) (contracAddress string, txHash string, err error) {
	return c.ethClient.DeployIjazahContract(ctx, pvKey)

}
