package services

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"math"
	"math/big"

	"sc-service/pkg"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthClient interface {
}

type ethClientImpl struct {
	client *pkg.ETHClient
}

// NewethClientImpl membuat instance ethClientImpl
func NewEthClientImpl(client *pkg.ETHClient) EthClient {
	return &ethClientImpl{
		client: client,
	}
}

func (s *ethClientImpl) CreateNewETHClient(netURL, pvKeyHex string, chainID *big.Int, contractAddress string) {
	// Buat client
	// Inisialisasi private key
	// privateKey, err := crypto.HexToECDSA(pvKeyHex)
	// if err != nil {
	// 	log.Fatalf("Error parsing private key: %v", err)
	// }

	// // Public key dan address
	// publicAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
}

func (s *ethClientImpl) LogintToETHNetwork(netURL string) (*ethclient.Client, error) {
	client, err := ethclient.DialContext(context.Background(), netURL)
	if err != nil {
		log.Fatalf("Error connecting to Ethereum client: %v", err)
		return nil, err
	}
	defer client.Close()
	return client, nil
}

func (s *ethClientImpl) GetChainId(client *ethclient.Client) (*big.Int, error) {
	// Ambil ChainID
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return nil, err
		log.Fatalf("Error getting Chain ID: %v", err)
	}
	return chainID, err
}
func (s *ethClientImpl) GetGasPrice(client *ethclient.Client) (*big.Int, error) {
	// Ambil gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Error getting Gas Price: %v", err)
		return nil, err
	}
	return gasPrice, nil
}

// Check balance
func (s *ethClientImpl) GetBalance(client *ethclient.Client, pubAddress string) (*big.Float, error) {
	address := common.HexToAddress(pubAddress)
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		return nil, err
	}
	// Konversi dari wei ke ETH
	fBalance := new(big.Float).SetInt(balance)
	ethValue := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	return ethValue, nil
}

// Check Block
func (s *ethClientImpl) GetBlock(client *ethclient.Client) (*big.Int, error) {
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return block.Number(), nil
}

func generatePrivateKey() string {
	// Generate private key
	pvKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	pData := crypto.FromECDSA(pvKey)
	return hexutil.Encode(pData)
}
func generatePublicKey(pvKey *ecdsa.PrivateKey) string {
	// Generate public key
	pubKey := crypto.FromECDSAPub(&pvKey.PublicKey)
	return hexutil.Encode(pubKey)
}
func generatePublicAddress(pvKey *ecdsa.PrivateKey) string {
	// Generate public address key
	pubAddress := crypto.PubkeyToAddress(pvKey.PublicKey)
	return pubAddress.Hex()
}
