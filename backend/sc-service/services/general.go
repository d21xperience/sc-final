package services

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"math/big"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Struktur untuk menyimpan respons ABI dari Etherscan
type EtherscanABIResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

func CreateTransactor(privateKeyHex string, chainID *big.Int) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	return auth, nil
}
func WeiToEther(wei *big.Int) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(math.Pow10(18)))
}
func IsValidAddress(address string) bool {
	return common.IsHexAddress(address)
}

// convertParams mengonversi string params ke tipe data sesuai untuk ABI
func convertParams(params []string) []interface{} {
	converted := make([]interface{}, len(params))
	for i, param := range params {
		converted[i] = param // Bisa dikembangkan untuk tipe data lain
	}
	return converted
}

// SendTransactionToContract mengirim transaksi ke smart contract
func SendTransactionToContract(ctx context.Context, client *ethclient.Client, contractAddress string, data []byte, privateKeyHex string, gasLimit uint64) (string, error) {
	//  Konversi private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", errors.New("gagal mengonversi private key: " + err.Error())
	}

	//  Ambil alamat pengirim dari private key
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)

	//  Ambil nonce akun pengirim
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return "", errors.New("gagal mendapatkan nonce: " + err.Error())
	}

	//  Dapatkan harga gas
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return "", errors.New("gagal mendapatkan harga gas: " + err.Error())
	}

	//  Buat transaksi
	toAddress := common.HexToAddress(contractAddress)
	tx := types.NewTransaction(nonce, toAddress, big.NewInt(0), gasLimit, gasPrice, data)

	//  Tanda tangani transaksi
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(big.NewInt(1)), privateKey)
	if err != nil {
		return "", errors.New("gagal menandatangani transaksi: " + err.Error())
	}

	//  Kirim transaksi ke jaringan
	err = client.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", errors.New("gagal mengirim transaksi: " + err.Error())
	}

	//  Kembalikan hash transaksi
	return signedTx.Hash().Hex(), nil
}

func GetAddressFromPublicKey(publicKey ecdsa.PublicKey) common.Address {
	return crypto.PubkeyToAddress(publicKey)
}

// SignTransaction: Menandatangani transaksi menggunakan private key
func SignTransaction(privateKeyHex string, tx *types.Transaction) (*types.Transaction, error) {
	// Dekode private key dari hex string
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, errors.New("gagal mengonversi private key ke ECDSA")
	}

	// Buat signer sesuai dengan chain ID
	chainID := big.NewInt(1) // Gantilah dengan chain ID yang sesuai
	signer := types.LatestSignerForChainID(chainID)

	// Tandatangani transaksi
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		return nil, errors.New("gagal menandatangani transaksi")
	}

	return signedTx, nil
}

func GetECDSAPrivateKey(privateKeyHex string) (*ecdsa.PrivateKey, error) {
	// Hapus "0x" jika ada di awal
	privateKeyHex = strings.TrimPrefix(privateKeyHex, "0x")

	// Validasi panjang private key
	if len(privateKeyHex) != 64 {
		return nil, fmt.Errorf("panjang private key tidak valid: %d karakter (harus 64 karakter)", len(privateKeyHex))
	}

	// Konversi ke private key
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, errors.New("gagal mengonversi private key: " + err.Error())
	}

	return privateKey, nil
}

// Fungsi untuk mendapatkan ABI dari Etherscan
func GetABIFromEtherscan(contractAddress string, EtherscanURL, EtherscanAPIKey string) (string, error) {
	// Format URL API Etherscan
	url := fmt.Sprintf("%s?module=contract&action=getabi&address=%s&apikey=%s",
		EtherscanURL, contractAddress, EtherscanAPIKey)

	// Kirim HTTP request ke Etherscan
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("gagal mengambil ABI: %w", err)
	}
	defer resp.Body.Close()

	// Baca response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("gagal membaca response: %w", err)
	}

	// Parsing JSON response
	var etherscanResp EtherscanABIResponse
	if err := json.Unmarshal(body, &etherscanResp); err != nil {
		return "", fmt.Errorf("gagal parsing JSON: %w", err)
	}

	// Cek status response
	if etherscanResp.Status != "1" {
		return "", errors.New("Etherscan API error: " + etherscanResp.Message)
	}

	return etherscanResp.Result, nil
}

func GetABIFromFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("gagal membaca file ABI: %w", err)
	}

	var abiData map[string]interface{}
	if err := json.Unmarshal(data, &abiData); err != nil {
		return "", fmt.Errorf("gagal parsing ABI JSON: %w", err)
	}

	abiJSON, err := json.Marshal(abiData["abi"])
	if err != nil {
		return "", fmt.Errorf("gagal mengubah ke string: %w", err)
	}

	return string(abiJSON), nil
}
func GetAddressFromPrivateKey(privateKey string) (common.Address, error) {
	// Dekode private key dari string hex
	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return common.Address{}, errors.New("gagal mengonversi private key ke ECDSA")
	}

	// Ambil public key dari private key
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("gagal mendapatkan public key dari private key")
	}

	// Konversi public key menjadi Ethereum address
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address, nil
}

func TransactOptsAuth(key *keystore.Key, chainID, gasPrice *big.Int, nonce, gasLimit uint64) *bind.TransactOpts {
	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal("Gagal membuat Transactor karena ", err.Error())
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice
	return auth
}

// ImportPrivateKey mengimpor private key dan mengonversinya menjadi alamat Ethereum
// func ImportPrivateKey(hexKey string) (common.Address, error) {
// 	// Hapus "0x" jika ada di awal
// 	hexKey = strings.TrimPrefix(hexKey, "0x")

// 	// Validasi panjang private key
// 	if len(hexKey) != 64 {
// 		return common.Address{}, fmt.Errorf("panjang private key tidak valid: %d karakter (harus 64 karakter)", len(hexKey))
// 	}
// 	// Dekode private key dari hex ke byte array
// 	privateKeyBytes, err := hex.DecodeString(hexKey)
// 	if err != nil {
// 		return common.Address{}, fmt.Errorf("gagal mendekode private key: %v", err)
// 	}

// 	// Parse private key menjadi objek ECDSA
// 	privateKey, err := crypto.ToECDSA(privateKeyBytes)
// 	if err != nil {
// 		return common.Address{}, fmt.Errorf("private key tidak valid: %v", err)
// 	}

// 	// Ambil public key dari private key
// 	publicKey := privateKey.Public().(*ecdsa.PublicKey)

// 	// Hitung alamat Ethereum dari public key
// 	address := crypto.PubkeyToAddress(*publicKey)

//		return address, nil
//	}
func ImportPrivateKey(hexKey string) (common.Address, error) {
	//  Hapus spasi yang tidak diinginkan
	hexKey = strings.TrimSpace(hexKey)

	//  Hapus "0x" jika ada di awal
	hexKey = strings.TrimPrefix(hexKey, "0x")

	//  Validasi panjang private key
	if len(hexKey) != 64 {
		return common.Address{}, errors.New("panjang private key tidak valid: harus 64 karakter")
	}

	//  Validasi format hanya berisi karakter hex
	matched, _ := regexp.MatchString("^[0-9a-fA-F]+$", hexKey)
	if !matched {
		return common.Address{}, errors.New("private key mengandung karakter tidak valid")
	}

	//  Dekode private key dari hex ke byte array
	privateKeyBytes, err := hex.DecodeString(hexKey)
	if err != nil {
		return common.Address{}, fmt.Errorf("gagal mendekode private key: %v", err)
	}

	//  Parse private key menjadi objek ECDSA
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return common.Address{}, fmt.Errorf("private key tidak valid: %v", err)
	}

	//  Ambil public key dari private key
	publicKey := privateKey.Public().(*ecdsa.PublicKey)

	//  Hitung alamat Ethereum dari public key
	address := crypto.PubkeyToAddress(*publicKey)

	log.Printf("Private key berhasil diimpor, menghasilkan address: %s", address.Hex())

	return address, nil
}
func ConvertStringPrivateKey(privateKeyHex string) (*ecdsa.PrivateKey, common.Address, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to import private key: %v", err)
	}
	publicAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	return privateKey, publicAddress, nil
}
