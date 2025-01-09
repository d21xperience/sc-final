package ethbc

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"math/big"
	verval_ijazah "sc-service/smartcontract/ethbc/gen"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	// client "github.com/influxdata/influxdb1-client"
)

type ETHClientInfo struct {
	NetURL           string
	PvKeyHex         string
	PubKeyHex        string
	PubAddressKeyHex string
	PubAddressKey    common.Address
	Nonce            uint64
	GasLimit         uint64
	ChainID          *big.Int
	GasPrice         *big.Int
	PvKey            *ecdsa.PrivateKey
	// PubKey              string
	EthClient *ethclient.Client
}

type ETH_Client struct {
	Client ETHClientInfo
}

func NewETHClient(cbc ETHClientInfo) ETH_Client {
	// var err error
	// a := common.HexToAddress(cbc.PvKeyHex)
	aByte, err := hex.DecodeString(cbc.PvKeyHex)
	if err != nil {
		log.Fatal("Gagal konversi ke []byte ", err)
	}
	cbc.PvKey, err = crypto.ToECDSA(aByte)
	if err != nil {
		log.Fatal("Gagal konversi ", err)
	}
	cbc.PubAddressKey = crypto.PubkeyToAddress(cbc.PvKey.PublicKey)
	cbc.PubKeyHex = crypto.PubkeyToAddress(cbc.PvKey.PublicKey).Hex()

	cbc.EthClient, err = ethclient.Dial(cbc.NetURL)
	if err != nil {
		log.Fatal("Gagal menghubungkan ke Server ", err)
	}
	defer cbc.EthClient.Close()

	// Nonce
	cbc.Nonce, err = cbc.EthClient.PendingNonceAt(context.Background(), cbc.PubAddressKey)
	if err != nil {
		log.Fatal("Gagal mendapatkan nonce ", err.Error())
	}
	// ChainID

	if cbc.ChainID == nil {
		cbc.ChainID, err = cbc.EthClient.NetworkID(context.Background())
		if err != nil {
			log.Fatal("Gagal mendapatkan Chain ID, ", err)
		}

	}
	// Gas Price
	cbc.GasPrice, err = cbc.EthClient.SuggestGasPrice(context.Background())
	// Gas Limit
	cbc.GasLimit = uint64(5000000)
	if err != nil {
		log.Fatal("Gagal mendapatkan harga Gas, ", err)
	}
	return ETH_Client{cbc}
}

func (cbc ETH_Client) SendCrypto(to string, amount *big.Int) {
	a2 := common.HexToAddress(to)
	cbc.Client.Nonce, _ = cbc.Client.EthClient.PendingNonceAt(context.Background(), cbc.Client.PubAddressKey)
	trx := types.NewTransaction(cbc.Client.Nonce, a2, amount, cbc.Client.GasLimit, cbc.Client.GasPrice, nil)

	tx, err := types.SignTx(trx, types.NewEIP155Signer(cbc.Client.ChainID), cbc.Client.PvKey)
	if err != nil {
		log.Fatal("Gagal mendapatkan Chain ID, ", err)
	}
	// Mengirim crypto
	re := cbc.Client.EthClient.SendTransaction(context.Background(), tx)
	if re != nil {
		log.Fatal("Gagal mengirim transaksi ", re.Error())
	}
	fmt.Println("Transaksi berhasi dengan hash ", tx.Hash().Hex())
}
func (cbc ETH_Client) CheckBalance(pubAddress string) (*big.Float, error) {
	a := common.HexToAddress(pubAddress)
	balance, err := cbc.Client.EthClient.BalanceAt(context.Background(), a, nil)
	if err != nil {
		return nil, err
	}
	// Konvesi ke ETH
	// 1 ETH = 10^18 wei
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	value := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	return value, nil
}

func (cbc ETH_Client) AuthOfSC() *bind.TransactOpts {
	auth, err := bind.NewKeyedTransactorWithChainID(cbc.Client.PvKey, cbc.Client.ChainID)
	if err != nil {
		log.Fatal("Gagal membuat Transactor karena ", err.Error())
	}
	auth.Nonce = big.NewInt(int64(cbc.Client.Nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = cbc.Client.GasLimit
	auth.GasPrice = cbc.Client.GasPrice
	return auth
}

// Check Block
func (cbc ETH_Client) CekBlock() *big.Int {
	block, err := cbc.Client.EthClient.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	return block.Number()
}

func (cbc ETH_Client) DeployContract() (common.Address, *big.Int, error) {
	auth := cbc.AuthOfSC()
	contractAddress, tx, _, err := verval_ijazah.DeployVervalIjazah(auth, cbc.Client.EthClient)
	if err != nil {
		return common.Address{}, nil, err
	}
	fmt.Printf("Contract deployed at: %s, tx hash: %s\n", contractAddress.Hex(), tx.Hash().Hex())
	return contractAddress, tx.Hash().Big(), nil
}

// Fungsi untuk Interaksi Kontrak
func (cbc ETH_Client) InteractWithContract(dataID string) (string, string, error) {
	cAdd := common.HexToAddress("0x1e2f4432bfef9e9ad39da6d272f4aff33629c770")
	ver, err := verval_ijazah.NewVervalIjazah(cAdd, cbc.Client.EthClient)
	if err != nil {
		return "", "", err
	}

	_, nama, _, noIjazah, _, _, err := ver.Get(&bind.CallOpts{From: cbc.Client.PubAddressKey}, dataID)
	if err != nil {
		return "", "", err
	}
	return nama, noIjazah, nil
}
