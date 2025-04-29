package services

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"path/filepath"
	"sc-service/utils"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// QuorumClient mengimplementasikan BlockchainClient untuk Quorum
type QuorumClient struct {
	rpcClient *rpc.Client
	client    *ethclient.Client
	// client    *EthereumClient
}

// NewQuorumClient membuat koneksi ke jaringan Quorum
func NewQuorumClient(cfg *Config) (BlockchainClient, error) {
	if cfg.RPCURL == "" {
		return nil, fmt.Errorf("quorum RPC URL tidak boleh kosong")
	}

	rpcClient, err := rpc.Dial(cfg.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("gagal menghubungkan ke Quorum RPC: %v", err)
	}

	client := ethclient.NewClient(rpcClient)
	return &QuorumClient{rpcClient: rpcClient, client: client}, nil
}

// Connect menghubungkan ke jaringan Quorum
func (q *QuorumClient) Connect() error {
	_, err := q.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("gagal terhubung ke jaringan Quorum: %w", err)
	}
	log.Println("Terhubung ke Quorum")
	return nil
}

// Implementasi NetworkID (inherit dari Ethereum)
func (q *QuorumClient) NetworkID(ctx context.Context) (*big.Int, error) {
	var result string
	err := q.rpcClient.CallContext(ctx, &result, "net_version")
	if err != nil {
		return nil, err
	}
	id := new(big.Int)
	id.SetString(result, 10)
	return id, nil
}

// Implementasi SuggestGasPrice (inherit dari Ethereum)
func (q *QuorumClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	var result big.Int
	err := q.rpcClient.CallContext(ctx, &result, "eth_gasPrice")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Implementasi SendPrivateTransaction
func (q *QuorumClient) SendPrivateTransaction(ctx context.Context, tx *types.Transaction) (string, error) {
	var txHash string
	err := q.rpcClient.CallContext(ctx, &txHash, "eth_sendRawTransaction", tx)
	return txHash, err
}

// Implementasi GetPrivateContract
func (q *QuorumClient) GetPrivateContract(ctx context.Context, address string, payload []byte) ([]byte, error) {
	var result []byte
	err := q.rpcClient.CallContext(ctx, &result, "eth_call", map[string]interface{}{
		"to":   address,
		"data": payload,
	})
	return result, err
}

// Implementasi GetConsensusAlgorithm
func (q *QuorumClient) GetConsensusAlgorithm(ctx context.Context) (string, error) {
	var result string
	err := q.rpcClient.CallContext(ctx, &result, "istanbul_getSnapshot") // Untuk IBFT
	if err != nil {
		err = q.rpcClient.CallContext(ctx, &result, "raft_cluster") // Untuk Raft
	}
	return result, err
}

// IssueDegree untuk Quorum
func (q *QuorumClient) IssueDegree(ctx context.Context, contractAddress string, degreeHash [32]byte, sekolah string, issueDate uint64, privateKey string, gasLimit uint64) (string, error) {
	//  Load ABI
	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return "", fmt.Errorf("error parsing ABI: %v", err)
	}

	//  Encode data untuk fungsi `issueDegree`
	data, err := parsedABI.Pack("issueDegree", degreeHash, sekolah, big.NewInt(int64(issueDate)))
	if err != nil {
		return "", fmt.Errorf("error packing data: %v", err)
	}

	//  Kirim transaksi menggunakan SendTransactionToContract
	txHash, err := SendTransactionToContract(ctx, q.client, contractAddress, data, privateKey, gasLimit)
	if err != nil {
		return "", fmt.Errorf("transaction failed: %v", err)
	}

	return txHash, nil
}

// Implement CallContractMethod agar memenuhi EthClient
//
//	func (q *QuorumClient) CallContractMethod(ctx context.Context, contractAddress, abi, method string, params []string) (string, error) {
//		return q.client.CallContract(ctx, contractAddress, abi, method, params)
//	}
func (q *QuorumClient) GenerateNewAccount(ctx context.Context, userId int32, password string) (map[string]interface{}, error) {
	key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)

	a, err := key.NewAccount(password)
	if err != nil {
		return nil, err
	}
	// simpan ke database
	pass, err := utils.EncryptPassword(password)
	if err != nil {
		return nil, err
	}
	var results = map[string]interface{}{
		"Password":          pass,
		"KeystrokeFilename": filepath.Base(a.URL.Path),
		"Address":           a.Address.Hex(),
	}
	return results, nil
}

// Implement GetTokenBalance
//
//	func (q *QuorumClient) GetTokenBalance(ctx context.Context, tokenAddress, ownerAddress string) (string, error) {
//		return q..GetTokenBalance(ctx, tokenAddress, ownerAddress)
//	}
// func (q *QuorumClient) GetTokenBalance(ctx context.Context, tokenAddress, ownerAddress string) (*big.Int, error) {
// 	return q..GetTokenBalance(ctx, tokenAddress, ownerAddress)
// }

// func (q *QuorumClient) DeployContract(ctx context.Context, bytecode string, privateKey string, gasLimit uint64) (string, string, error) {
// 	return q..DeployContract(ctx, bytecode, privateKey, gasLimit)
// }
// func (q *QuorumClient) GetContractEvents(ctx context.Context, contractAddress, abi, eventName string, fromBlock, toBlock uint64) ([]string, error) {
// 	return q..GetContractEvents(ctx, contractAddress, abi, eventName, fromBlock, toBlock)
// }
// func (q *QuorumClient) SendETH(ctx context.Context, privateKeyHex, toAddress string, amount *big.Int) (string, error) {
// 	return q..SendETH(ctx, privateKeyHex, toAddress, amount)
// }

// func (q *QuorumClient) TransferToken(ctx context.Context, tokenAddress, from, to string, amount string, privateKey string, gasLimit uint64) (string, error) {
// 	return q..TransferToken(ctx, tokenAddress, from, to, amount, privateKey, gasLimit)
// }

// func (q *QuorumClient) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
// 	return q..PendingNonceAt(ctx, account)
// }
// func (q *QuorumClient) SendTransaction(ctx context.Context, tx *types.Transaction) error {
// 	return q..SendTransaction(ctx, tx)
// }
// func (q *QuorumClient) GetContract(ctx context.Context, contractAddress string) (string, string, error) {
// 	return q..GetContract(ctx, contractAddress)

// }
// func (q *QuorumClient) GenerateNewAccount(ctx context.Context, userId int32, password string) (map[string]interface{}, error) {
// 	return q..GenerateNewAccount(ctx, userId, password)
// }

// func (q *QuorumClient) GetAccounts(ctx context.Context, userId int32, schemaname string) ([]*models.Account, error) {
// 	return q..GetAccounts(ctx, userId, schemaname)

// }
// func (q *QuorumClient) DeployIjazahContract(ctx context.Context, pvKey string) (contracAddress string, txHash string, err error) {
// 	return q..DeployIjazahContract(ctx, pvKey)

// }
