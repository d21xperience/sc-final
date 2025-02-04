package services

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

// Default implementasi QuorumClient menggunakan go-ethereum dan RPC
type DefaultQuorumClient struct {
	rpcClient *rpc.Client
	ethClient EthClient
}

// Constructor untuk DefaultQuorumClient
func NewDefaultQuorumClient(rawUrl string, ethClient EthClient) (*DefaultQuorumClient, error) {
	client, err := rpc.Dial(rawUrl)
	if err != nil {
		return nil, err
	}
	return &DefaultQuorumClient{rpcClient: client, ethClient: ethClient}, nil
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

// Implement CallContractMethod agar memenuhi EthClient
func (c *DefaultQuorumClient) CallContractMethod(ctx context.Context, contractAddress, abi, method string, params []string) (string, error) {
	return c.ethClient.CallContractMethod(ctx, contractAddress, abi, method, params)
}

// Implement GetTokenBalance
//
//	func (c *DefaultQuorumClient) GetTokenBalance(ctx context.Context, tokenAddress, ownerAddress string) (string, error) {
//		return c.ethClient.GetTokenBalance(ctx, tokenAddress, ownerAddress)
//	}
func (c *DefaultQuorumClient) GetTokenBalance(ctx context.Context, tokenAddress, ownerAddress string) (*big.Int, error) {
	return c.ethClient.GetTokenBalance(ctx, tokenAddress, ownerAddress)
}

func (c *DefaultQuorumClient) DeployContract(ctx context.Context, bytecode string, privateKey string, gasLimit uint64) (string, string, error) {
	return c.ethClient.DeployContract(ctx, bytecode, privateKey, gasLimit)
}
func (c *DefaultQuorumClient) GetContractEvents(ctx context.Context, contractAddress, abi, eventName string, fromBlock, toBlock uint64) ([]string, error) {
	return c.ethClient.GetContractEvents(ctx, contractAddress, abi, eventName, fromBlock, toBlock)
}
func (c *DefaultQuorumClient) SendETH(ctx context.Context, privateKeyHex, toAddress string, amount *big.Int) (string, error) {
	return c.ethClient.SendETH(ctx, privateKeyHex, toAddress, amount)
}

func (c *DefaultQuorumClient) TransferToken(ctx context.Context, tokenAddress, from, to string, amount string, privateKey string, gasLimit uint64) (string, error) {
	return c.ethClient.TransferToken(ctx, tokenAddress, from, to, amount, privateKey, gasLimit)
}

func (c *DefaultQuorumClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return c.ethClient.PendingNonceAt(ctx, account)
}
func (c *DefaultQuorumClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return c.ethClient.SendTransaction(ctx, tx)
}
func (c *DefaultQuorumClient) GetContract(ctx context.Context, contractAddress string) (string, string, error) {
	return c.ethClient.GetContract(ctx, contractAddress)

}
func (c *DefaultQuorumClient) GenerateNewAccount(ctx context.Context, userId int32, password string) (map[string]interface{}, error) {
	return c.ethClient.GenerateNewAccount(ctx, userId, password)
}

// func (c *DefaultQuorumClient) GetAccounts(ctx context.Context, userId int32, schemaname string) ([]*models.Account, error) {
// 	return c.ethClient.GetAccounts(ctx, userId, schemaname)

// }
func (c *DefaultQuorumClient) DeployIjazahContract(ctx context.Context, pvKey string) (contracAddress string, txHash string, err error) {
	return c.ethClient.DeployIjazahContract(ctx, pvKey)

}
