package services

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// Implementasi generik Ethereum client
type EthClientService struct {
	client EthClient
}

// Constructor untuk EthClientService
func NewEthClientService(client EthClient) *EthClientService {
	return &EthClientService{client: client}
}

// NetworkID menggunakan metode standar
func (s *EthClientService) NetworkID(ctx context.Context) (*big.Int, error) {
	return s.client.NetworkID(ctx)
}

// SuggestGasPrice menggunakan metode standar
func (s *EthClientService) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return s.client.SuggestGasPrice(ctx)
}
func (s *EthClientService) GenerateNewAccount() (string, string, error) {
	return s.client.GenerateNewAccount()
}

// Default implementasi EthClient menggunakan go-ethereum dan RPC
type DefaultEthClient struct {
	rpcClient *rpc.Client
}

// Constructor untuk DefaultEthClient
// Constructor untuk DefaultEthClient
func NewDefaultEthClient(rawUrl string) (*DefaultEthClient, error) {
	client, err := rpc.Dial(rawUrl)
	if err != nil {
		return nil, err
	}
	return &DefaultEthClient{rpcClient: client}, nil
}

// Implementasi NetworkID
func (c *DefaultEthClient) NetworkID(ctx context.Context) (*big.Int, error) {
	var result string
	err := c.rpcClient.CallContext(ctx, &result, "net_version")
	if err != nil {
		return nil, err
	}
	id := new(big.Int)
	id.SetString(result, 10)
	return id, nil
}

// Implementasi SuggestGasPrice
func (c *DefaultEthClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	var result big.Int
	err := c.rpcClient.CallContext(ctx, &result, "eth_gasPrice")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *DefaultEthClient) GenerateNewAccount() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", fmt.Errorf("failed to generate key: %v", err)
	}
	privateKeyHex := hexutil.Encode(crypto.FromECDSA(privateKey))
	publicAddress := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	return privateKeyHex, publicAddress, nil
}

func (c *DefaultEthClient) ImportPrivateKey(privateKeyHex string) (*ecdsa.PrivateKey, common.Address, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to import private key: %v", err)
	}
	publicAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	return privateKey, publicAddress, nil
}
func (c *DefaultEthClient) GetAddressFromPublicKey(publicKey ecdsa.PublicKey) common.Address {
	return crypto.PubkeyToAddress(publicKey)
}
func (c *DefaultEthClient) GetBalance(client *ethclient.Client, address string) (*big.Float, error) {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}
	// Konversi saldo dari wei ke ETH
	ethBalance := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(math.Pow10(18)))
	return ethBalance, nil
}
func (c *DefaultEthClient) GetLatestBlock(client *ethclient.Client) (*big.Int, error) {
	block, err := client.BlockByNumber(context.Background(), nil) // `nil` untuk blok terbaru
	if err != nil {
		return nil, err
	}
	return block.Number(), nil
}
func (c *DefaultEthClient) SendETH(client *ethclient.Client, privateKeyHex, toAddress string, amount *big.Int) (string, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", err
	}
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	tx := types.NewTransaction(nonce, common.HexToAddress(toAddress), amount, uint64(21000), gasPrice, nil)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}

//	func (c *DefaultEthClient) CallSmartContract(client *ethclient.Client, contractAddress, dataID string) (string, error) {
//		contractAddr := common.HexToAddress(contractAddress)
//		// Replace with your contract binding
//		instance, err := verval_ijazah.NewVervalIjazah(contractAddr, client)
//		if err != nil {
//			return "", err
//		}
//		result, err := instance.SomeFunction(&bind.CallOpts{
//			From: contractAddr,
//		}, dataID)
//		if err != nil {
//			return "", err
//		}
//		return result, nil
//	}
func (c *DefaultEthClient) DeploySmartContract(client *ethclient.Client, privateKeyHex string) (common.Address, string, error) {
	res, add, err := DeploySmartContract(client, privateKeyHex)
	return res, add, err
}
func (c *DefaultEthClient) SubscribeToEvents(client *ethclient.Client, contractAddress string) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe: %v", err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case err := <-sub.Err():
			log.Printf("Subscription error: %v", err)
		case vLog := <-logs:
			log.Printf("New log: %+v", vLog)
		}
	}
}
func (c *DefaultEthClient) WeiToEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(math.Pow10(18)))
}
func (c *DefaultEthClient) IsValidAddress(address string) bool {
	return common.IsHexAddress(address)
}
