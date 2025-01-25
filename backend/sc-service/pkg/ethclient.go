package pkg

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ETHClientInfo struct {

}

type ETHClient struct {
	NetURL        string
	LocalURL        string
	// PvKeyHex      string
	// PubAddressKey common.Address
	// Nonce         uint64
	GasLimit      uint64
	// PvKey         *ecdsa.PrivateKey
	// EthClient     *ethclient.Client
}

// func NewETHClient(netURL, pvKeyHex string) *ETHClient {
// 	return &ETHClient{
// 		Client: ETHClientInfo{
// 			NetURL:        netURL,
// 			PvKeyHex:      pvKeyHex,
// 			PubAddressKey: publicAddress,
// 			ChainID:       chainID,
// 			GasPrice:      gasPrice,
// 			PvKey:         privateKey,
// 			GasLimit:      5000000,
// 			EthClient:     client,
// 		},
// 	}
// }

// Membuat transaksi dengan signed transactor
func (e *ETHClient) AuthOfSC() *bind.TransactOpts {
	auth, err := bind.NewKeyedTransactorWithChainID(e.Client.PvKey, e.Client.ChainID)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}
	auth.GasPrice = e.Client.GasPrice
	auth.GasLimit = e.Client.GasLimit
	auth.From = e.Client.PubAddressKey
	return auth
}
