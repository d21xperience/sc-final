package services

import (
	"math/big"
	"sc-service/smartcontract/ethbc/gen/verval_ijazah"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

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

func DeployContract(auth *bind.TransactOpts, client *ethclient.Client) (common.Address, string, error) {
	contractAddress, tx, _, err := verval_ijazah.DeployVervalIjazah(auth, client)
	if err != nil {
		return common.Address{}, "", err
	}
	return contractAddress, tx.Hash().Hex(), nil
}

func DeploySmartContract(client *ethclient.Client, privateKeyHex string, chainID *big.Int) (common.Address, string, error) {
	// Buat transactor
	auth, err := CreateTransactor(privateKeyHex, chainID)
	if err != nil {
		return common.Address{}, "", err
	}

	// Deploy smart contract
	contractAddress, txHash, err := DeployContract(auth, client)
	if err != nil {
		return common.Address{}, "", err
	}

	return contractAddress, txHash, nil
}

// // Fungsi untuk mengeluarkan ijazah (issueDegree)
// func issueDegree(client *ethclient.Client, degreeHash [32]byte, sekolah string, issueDate uint64) {
// 	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
// 	if err != nil {
// 		log.Fatalf("Error parsing ABI: %v", err)
// 	}

// 	data, err := parsedABI.Pack("issueDegree", degreeHash, sekolah, big.NewInt(int64(issueDate)))
// 	if err != nil {
// 		log.Fatalf("Error packing data: %v", err)
// 	}

// 	txHash, err := sendTransaction(client, data)
// 	if err != nil {
// 		log.Fatalf("Transaction failed: %v", err)
// 	}

// 	fmt.Printf("Ijazah berhasil dikeluarkan! TxHash: %s\n", txHash.Hex())
// }

// // Fungsi untuk menambahkan transkrip nilai
// func addTranscript(client *ethclient.Client, degreeHash [32]byte, mataPelajaran []string, nilai []uint8) {
// 	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
// 	if err != nil {
// 		log.Fatalf("Error parsing ABI: %v", err)
// 	}

// 	data, err := parsedABI.Pack("addTranscript", degreeHash, mataPelajaran, nilai)
// 	if err != nil {
// 		log.Fatalf("Error packing data: %v", err)
// 	}

// 	txHash, err := sendTransaction(client, data)
// 	if err != nil {
// 		log.Fatalf("Transaction failed: %v", err)
// 	}

// 	fmt.Printf("Transkrip berhasil ditambahkan! TxHash: %s\n", txHash.Hex())
// }
