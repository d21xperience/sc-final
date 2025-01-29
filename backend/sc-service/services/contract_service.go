package services

import (
	"math/big"
	"sc-service/smartcontract/ethbc/gen/verval_ijazah"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

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

// func CallSmartContract(client *ethclient.Client, contractAddress, dataID string) (string, error) {
// 	contractAddr := common.HexToAddress(contractAddress)
// 	// Replace with your contract binding
// 	instance, err := verval_ijazah.NewVervalIjazah(contractAddr, client)
// 	if err != nil {
// 		return "", err
// 	}
// 	result, err := instance.Get(&bind.CallOpts{
// 		From: contractAddr,
// 	}, dataID)
// 	if err != nil {
// 		return "", err
// 	}
// 	return result, nil
// }
