package services

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// Default implementasi EthClient menggunakan go-ethereum dan RPC
type DefaultEthClient struct {
	rpcClient *rpc.Client
	client    *ethclient.Client
}

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
func (c *DefaultEthClient) GetBalance(address string) (*big.Float, error) {
	account := common.HexToAddress(address)
	balance, err := c.client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, err
	}
	// Konversi saldo dari wei ke ETH
	ethBalance := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(math.Pow10(18)))
	return ethBalance, nil
}
func (c *DefaultEthClient) GetLatestBlock() (*big.Int, error) {
	block, err := c.client.BlockByNumber(context.Background(), nil) // `nil` untuk blok terbaru
	if err != nil {
		return nil, err
	}
	return block.Number(), nil
}
func (c *DefaultEthClient) SendETH(ctx context.Context, privateKeyHex, toAddress string, amount *big.Int) (string, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", err
	}
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := c.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", err
	}
	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	tx := types.NewTransaction(nonce, common.HexToAddress(toAddress), amount, uint64(21000), gasPrice, nil)
	chainID, err := c.client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}
	err = c.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}

//	func (c *DefaultEthClient) CallSmartContract(client *ethclient, contractAddress, dataID string) (string, error) {
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
//
//	func (c *DefaultEthClient) DeploySmartContract(client *ethclient, privateKeyHex string) (common.Address, string, error) {
//		res, add, err := DeploySmartContract(client, privateKeyHex)
//		return res, add, err
//	}
func (c *DefaultEthClient) SubscribeToEvents(contractAddress string) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
	}
	logs := make(chan types.Log)
	sub, err := c.client.SubscribeFilterLogs(context.Background(), query, logs)
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
func (c *DefaultEthClient) DeployContract(ctx context.Context, bytecode string, privateKey string, gasLimit uint64) (string, string, error) {
	if c.client == nil {
		return "", "", errors.New("ethereum client tidak dikonfigurasi")
	}

	// Konversi bytecode dari string ke format bytes
	contractBytecode := common.FromHex(bytecode)

	// Dapatkan nonce untuk transaksi
	fromAddress, _ := GetAddressFromPrivateKey(privateKey) // Buat fungsi ini jika belum ada
	nonce, err := c.client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return "", "", err
	}

	// Dapatkan harga gas saat ini
	gasPrice, err := c.client.SuggestGasPrice(ctx)
	if err != nil {
		return "", "", err
	}

	// Buat transaksi untuk deploy contract
	tx := types.NewContractCreation(nonce, big.NewInt(0), gasLimit, gasPrice, contractBytecode)

	// Tanda tangani transaksi dengan private key
	signedTx, err := SignTransaction(privateKey, tx) // Buat fungsi SignTransaction jika belum ada
	if err != nil {
		return "", "", err
	}

	// Kirim transaksi
	err = c.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", "", err
	}

	// Kembalikan alamat contract & tx hash
	contractAddress := crypto.CreateAddress(fromAddress, nonce) // Buat alamat contract dari nonce
	return contractAddress.Hex(), signedTx.Hash().Hex(), nil
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

// SendTransactionToContract mengirim transaksi ke smart contract
func (c *DefaultEthClient) SendTransactionToContract(ctx context.Context, contractAddress, abiJSON, method string, params []string, privateKeyHex string, gasLimit uint64) (string, error) {
	//  Konversi private key dari hex ke ECDSA
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", errors.New("gagal mengonversi private key: " + err.Error())
	}

	//  Dapatkan alamat pengirim dari private key
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)

	//  Dapatkan nonce akun pengirim
	nonce, err := c.client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return "", errors.New("gagal mendapatkan nonce: " + err.Error())
	}

	//  Load ABI contract
	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return "", errors.New("gagal mem-parsing ABI: " + err.Error())
	}

	//  Konversi parameter menjadi data transaksi
	data, err := parsedABI.Pack(method, convertParams(params)...)
	if err != nil {
		return "", errors.New("gagal mengkodekan data transaksi: " + err.Error())
	}

	//  Dapatkan harga gas
	gasPrice, err := c.client.SuggestGasPrice(ctx)
	if err != nil {
		return "", errors.New("gagal mendapatkan harga gas: " + err.Error())
	}

	//  Buat transaksi
	toAddress := common.HexToAddress(contractAddress)
	tx := types.NewTransaction(nonce, toAddress, big.NewInt(0), gasLimit, gasPrice, data)

	//  Tanda tangani transaksi
	chainID, err := c.client.NetworkID(ctx)
	if err != nil {
		return "", errors.New("gagal mendapatkan chain ID: " + err.Error())
	}

	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	if err != nil {
		return "", errors.New("gagal menandatangani transaksi: " + err.Error())
	}

	//  Kirim transaksi ke jaringan
	err = c.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", errors.New("gagal mengirim transaksi: " + err.Error())
	}

	//  Kembalikan hash transaksi
	return signedTx.Hash().Hex(), nil
}

// convertParams mengonversi string params ke tipe data sesuai untuk ABI
func convertParams(params []string) []interface{} {
	converted := make([]interface{}, len(params))
	for i, param := range params {
		converted[i] = param // Bisa dikembangkan untuk tipe data lain
	}
	return converted
}

// TransferToken mengirim token ERC-20 ke alamat lain
func (c *DefaultEthClient) TransferToken(ctx context.Context, tokenAddress, from, to, amountStr, privateKeyHex string, gasLimit uint64) (string, error) {
	//  Konversi private key dari hex ke ECDSA
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return "", errors.New("gagal mengonversi private key: " + err.Error())
	}

	//  Dapatkan alamat pengirim dari private key
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)

	//  Pastikan alamat pengirim sesuai dengan private key
	if !strings.EqualFold(fromAddress.Hex(), from) {
		return "", errors.New("private key tidak cocok dengan alamat pengirim")
	}
	

	//  Dapatkan nonce akun pengirim
	nonce, err := c.client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return "", errors.New("gagal mendapatkan nonce: " + err.Error())
	}

	//  Konversi jumlah token dari string ke *big.Int
	amount := new(big.Int)
	amount, ok := amount.SetString(amountStr, 10)
	if !ok {
		return "", errors.New("gagal mengonversi amount ke *big.Int")
	}

	//  ABI fungsi transfer ERC-20: transfer(address,uint256)
	erc20ABI := `[{"constant":false,"inputs":[{"name":"recipient","type":"address"},{"name":"amount","type":"uint256"}],"name":"transfer","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"}]`

	//  Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		return "", errors.New("gagal mem-parsing ABI: " + err.Error())
	}

	//  Encode data untuk fungsi transfer ERC-20
	data, err := parsedABI.Pack("transfer", common.HexToAddress(to), amount)
	if err != nil {
		return "", errors.New("gagal mengkodekan data transaksi: " + err.Error())
	}

	//  Dapatkan harga gas
	gasPrice, err := c.client.SuggestGasPrice(ctx)
	if err != nil {
		return "", errors.New("gagal mendapatkan harga gas: " + err.Error())
	}

	//  Buat transaksi
	tokenContract := common.HexToAddress(tokenAddress)
	tx := types.NewTransaction(nonce, tokenContract, big.NewInt(0), gasLimit, gasPrice, data)

	//  Dapatkan Chain ID
	chainID, err := c.client.NetworkID(ctx)
	if err != nil {
		return "", errors.New("gagal mendapatkan chain ID: " + err.Error())
	}

	//  Tandatangani transaksi
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	if err != nil {
		return "", errors.New("gagal menandatangani transaksi: " + err.Error())
	}

	//  Kirim transaksi ke jaringan
	err = c.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", errors.New("gagal mengirim transaksi: " + err.Error())
	}

	//  Kembalikan hash transaksi
	return signedTx.Hash().Hex(), nil
}

func (c *DefaultEthClient) CallContractMethod(ctx context.Context, contractAddress, abiStr, method string, params []string) (string, error) {
	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		return "", fmt.Errorf("failed to parse ABI: %w", err)
	}

	// Convert parameters
	args := make([]interface{}, len(params))
	for i, param := range params {
		args[i] = param
	}

	// Pack the method call
	data, err := parsedABI.Pack(method, args...)
	if err != nil {
		return "", fmt.Errorf("failed to pack method call: %w", err)
	}

	// Prepare contract call message
	contractAddr := common.HexToAddress(contractAddress) // Buat variabel terlebih dahulu
	msg := ethereum.CallMsg{
		To:   &contractAddr,
		Data: data,
	}

	// Call contract method
	result, err := c.client.CallContract(ctx, msg, nil)
	if err != nil {
		return "", fmt.Errorf("contract call failed: %w", err)
	}

	// Decode return value
	return hexutil.Encode(result), nil
}

func (c *DefaultEthClient) GetTokenBalance(ctx context.Context, tokenAddress, ownerAddress string) (string, error) {
	// ERC-20 ABI (minimal)
	const erc20ABI = `[{"constant":true,"inputs":[{"name":"owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`

	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(erc20ABI))
	if err != nil {
		return "", fmt.Errorf("failed to parse ERC-20 ABI: %w", err)
	}

	// Pack balanceOf(owner)
	data, err := parsedABI.Pack("balanceOf", common.HexToAddress(ownerAddress))
	if err != nil {
		return "", fmt.Errorf("failed to pack balanceOf: %w", err)
	}

	// Prepare contract call message
	tokenAdd := common.HexToAddress(tokenAddress) // Buat variabel terlebih dahulu
	msg := ethereum.CallMsg{
		To:   &tokenAdd,
		Data: data,
	}

	// Call contract method
	result, err := c.client.CallContract(ctx, msg, nil)
	if err != nil {
		return "", fmt.Errorf("failed to get token balance: %w", err)
	}

	// Decode the balance from result
	balance := new(big.Int).SetBytes(result)
	return balance.String(), nil
}
func (c *DefaultEthClient) GetContractEvents(ctx context.Context, contractAddress, abiStr, eventName string, fromBlock, toBlock uint64) ([]string, error) {
	// Parse ABI
	parsedABI, err := abi.JSON(strings.NewReader(abiStr))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: %w", err)
	}

	// Ambil event berdasarkan nama
	event, exists := parsedABI.Events[eventName]
	if !exists {
		return nil, fmt.Errorf("event %s not found in ABI", eventName)
	}

	// Buat filter query
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
		Addresses: []common.Address{common.HexToAddress(contractAddress)},
		Topics:    [][]common.Hash{{event.ID}}, // Event signature hash
	}

	// Ambil log dari blockchain
	logs, err := c.client.FilterLogs(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get logs: %w", err)
	}

	// Decode event logs
	var events []string
	for _, vLog := range logs {
		data, err := parsedABI.Unpack(event.Name, vLog.Data)
		if err != nil {
			log.Printf("failed to decode log data: %v", err)
			continue
		}
		events = append(events, fmt.Sprintf("%v", data))
	}

	return events, nil
}
