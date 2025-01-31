package services

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// ContractService adalah service untuk interaksi dengan smart contract
type ContractService struct {
	client EthClient
}

type SenderInfo struct {
}

// Constructor untuk ContractService
func NewContractService(client EthClient) *ContractService {
	return &ContractService{client: client}
}

// issueDegree mengeluarkan ijazah dengan dependency injection
func (s *ContractService) IssueDegree(contractAddress string, degreeHash [32]byte, sekolah string, issueDate uint64, privateKey string, gasLimit uint64) {
	//  Load ABI
	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		log.Fatalf("Error parsing ABI: %v", err)
	}

	//  Encode data untuk fungsi `issueDegree`
	data, err := parsedABI.Pack("issueDegree", degreeHash, sekolah, big.NewInt(int64(issueDate)))
	if err != nil {
		log.Fatalf("Error packing data: %v", err)
	}

	//  Kirim transaksi menggunakan SendTransactionToContract
	txHash, err := SendTransactionToContract(context.Background(), s.client, contractAddress, data, privateKey, gasLimit)
	if err != nil {
		log.Fatalf("Transaction failed: %v", err)
	}

	fmt.Printf("Ijazah berhasil dikeluarkan! TxHash: %s\n", txHash)
}

// Fungsi untuk menambahkan transkrip nilai
// func (s *ContractService) AddTranscript(degreeHash [32]byte, mataPelajaran []string, nilai []uint8) {
// 	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
// 	if err != nil {
// 		log.Fatalf("Error parsing ABI: %v", err)
// 	}

// 	data, err := parsedABI.Pack("addTranscript", degreeHash, mataPelajaran, nilai)
// 	if err != nil {
// 		log.Fatalf("Error packing data: %v", err)
// 	}

// 	txHash, err := sendTransaction(s.client, data)
// 	if err != nil {
// 		log.Fatalf("Transaction failed: %v", err)
// 	}

// 	fmt.Printf("Transkrip berhasil ditambahkan! TxHash: %s\n", txHash.Hex())
// }

// func DeployContract(auth *bind.TransactOpts, client EthClient) (common.Address, string, error) {
// 	contractAddress, tx, _, err := verval_ijazah.DeployVervalIjazah(auth, client)
// 	if err != nil {
// 		return common.Address{}, "", err
// 	}
// 	return contractAddress, tx.Hash().Hex(), nil
// }

// func DeploySmartContract(client EthClient, privateKeyHex string, chainID *big.Int) (common.Address, string, error) {
// 	// Buat transactor
// 	auth, err := CreateTransactor(privateKeyHex, chainID)
// 	if err != nil {
// 		return common.Address{}, "", err
// 	}

// 	// Deploy smart contract
// 	contractAddress, txHash, err := DeployContract(auth, client)
// 	if err != nil {
// 		return common.Address{}, "", err
// 	}

// 	return contractAddress, txHash, nil
// }

// Fungsi untuk membuat transaksi dan menandatangani
// func sendTransaction(client EthClient, data []byte) (common.Hash, error) {
// 	// privateKeyHex := client.
// 	privateKey, err := crypto.HexToECDSA(privateKeyHex)
// 	if err != nil {
// 		return common.Hash{}, err
// 	}

// 	publicKey := privateKey.Public().(*ecdsa.PublicKey)
// 	fromAddress := crypto.PubkeyToAddress(*publicKey)

// 	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
// 	if err != nil {
// 		return common.Hash{}, err
// 	}

// 	gasPrice, err := client.SuggestGasPrice(context.Background())
// 	if err != nil {
// 		return common.Hash{}, err
// 	}

// 	tx := types.NewTransaction(nonce, common.HexToAddress(contractAddress), big.NewInt(0), 3000000, gasPrice, data)
// 	chainID, _ := client.NetworkID(context.Background())
// 	signedTx, _ := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)

// 	err = client.SendTransaction(context.Background(), signedTx)
// 	if err != nil {
// 		return common.Hash{}, err
// 	}

// 	return signedTx.Hash(), nil
// }

// func DeploySmartContract(client *ethclient.Client, privateKeyHex string) (common.Address, string, error) {
// 	privateKey, err := crypto.HexToECDSA(privateKeyHex)
// 	if err != nil {
// 		return common.Address{}, "", err
// 	}
// 	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337)) // Ganti Chain ID sesuai kebutuhan
// 	if err != nil {
// 		return common.Address{}, "", err
// 	}
// 	contractAddress, tx, _, err := verval_ijazah.DeployVervalIjazah(auth, client)
// 	if err != nil {
// 		return common.Address{}, "", err
// 	}
// 	return contractAddress, tx.Hash().Hex(), nil
// }

//	func CallSmartContract(client *ethclient.Client, contractAddress, dataID string) (string, error) {
//		contractAddr := common.HexToAddress(contractAddress)
//		// Replace with your contract binding
//		instance, err := verval_ijazah.NewVervalIjazah(contractAddr, client)
//		if err != nil {
//			return "", err
//		}
//		result, err := instance.Get(&bind.CallOpts{
//			From: contractAddr,
//		}, dataID)
//		if err != nil {
//			return "", err
//		}
//		return result, nil
//	}
