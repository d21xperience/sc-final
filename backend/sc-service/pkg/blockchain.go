package pkg

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Blockchain interface {
	Connect() error
	SendCrypto(address string, ammount *big.Int) error
	DeployContract() (common.Address, *big.Int, error) 
	InterctWithContract(dataID string) (string, string, error)
}