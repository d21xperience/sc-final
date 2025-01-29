package pkg

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"

	"sc-service/smartcontract/ethbc/gen/verval_ijazah"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// ===========================================
// 1. Fungsi untuk Mengelola Akun
// ===========================================

// Membuat Akun Baru (Key Pair)
func GenerateNewAccount() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", fmt.Errorf("failed to generate key: %v", err)
	}
	privateKeyHex := hexutil.Encode(crypto.FromECDSA(privateKey))
	publicAddress := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	return privateKeyHex, publicAddress, nil
}

// Import Private Key
func ImportPrivateKey(privateKeyHex string) (*ecdsa.PrivateKey, common.Address, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to import private key: %v", err)
	}
	publicAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	return privateKey, publicAddress, nil
}

// Hitung Address dari Public Key
func GetAddressFromPublicKey(publicKey ecdsa.PublicKey) common.Address {
	return crypto.PubkeyToAddress(publicKey)
}

// ===========================================

// ===========================================
// 2. Fungsi untuk Berinteraksi dengan Blockchain
// ===========================================

// Mendapatkan Saldo
func GetBalance(client *ethclient.Client, address string) (*big.Float, error) {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}
	// Konversi saldo dari wei ke ETH
	ethBalance := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(math.Pow10(18)))
	return ethBalance, nil
}

// Mendapatkan Nonce
func GetNonce(client *ethclient.Client, address string) (uint64, error) {
	account := common.HexToAddress(address)
	nonce, err := client.PendingNonceAt(context.Background(), account)
	if err != nil {
		return 0, err
	}
	return nonce, nil
}

// Mendapatkan Informasi Blok
func GetLatestBlock(client *ethclient.Client) (*big.Int, error) {
	block, err := client.BlockByNumber(context.Background(), nil) // `nil` untuk blok terbaru
	if err != nil {
		return nil, err
	}
	return block.Number(), nil
}

// ===========================================

// ===========================================
// 3. Fungsi untuk Membuat dan Mengirim Transaksi

// ===========================================
// Mengirim ETH
func SendETH(client *ethclient.Client, privateKeyHex, toAddress string, amount *big.Int) (string, error) {
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

// Memanggil Fungsi di Smart Contract
// func CallSmartContract(client *ethclient.Client, contractAddress, dataID string) (string, error) {
// 	contractAddr := common.HexToAddress(contractAddress)
// 	// Replace with your contract binding
// 	instance, err := verval_ijazah.NewVervalIjazah(contractAddr, client)
// 	if err != nil {
// 		return "", err
// 	}
// 	result, err := instance.SomeFunction(&bind.CallOpts{
// 		From: contractAddr,
// 	}, dataID)
// 	if err != nil {
// 		return "", err
// 	}
// 	return result, nil
// }

// Deploy Smart Contract
func DeploySmartContract(client *ethclient.Client, privateKeyHex string) (common.Address, string, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return common.Address{}, "", err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337)) // Ganti Chain ID sesuai kebutuhan
	if err != nil {
		return common.Address{}, "", err
	}
	contractAddress, tx, _, err := verval_ijazah.DeployVervalIjazah(auth, client)
	if err != nil {
		return common.Address{}, "", err
	}
	return contractAddress, tx.Hash().Hex(), nil
}

// ===========================================

// ===========================================
// 4. Fungsi untuk Monitoring

// ===========================================

// Subscribing ke Log Event
func SubscribeToEvents(client *ethclient.Client, contractAddress string) {
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

// ===========================================

// ===========================================
// 5. Utilitas Lain

// ===========================================

// Konversi dari Wei ke ETH
func WeiToEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(math.Pow10(18)))
}

// Validasi Address
func IsValidAddress(address string) bool {
	return common.IsHexAddress(address)
}

// ===========================================
