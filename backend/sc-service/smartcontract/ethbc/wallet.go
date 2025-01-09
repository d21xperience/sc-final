package ethbc

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

type WalletInfo struct {
	PrivateKey     []byte
	PublicKey      []byte
	PubblicAddress string
	Pwd            string
	PathDir        string
}
type Wallet struct {
	WlInfo WalletInfo
}

func (wallet WalletInfo) NewWallet() Wallet {
	return Wallet{wallet}
}

func (wal Wallet) CreateWallet() {
	dir := wal.WlInfo.PathDir
	key := keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)
	// pwd dienkripsi dan bandingkan dengan database
	a, err := key.NewAccount(wal.WlInfo.Pwd)
	if err != nil {
		log.Fatal("Wallet gagal dibuat ", err)
	}
	fmt.Println("Wallet berhasil dibuat, dengan Public Address", a.Address) //Public Address

	// Simpan User baru di database
}

// Read wallet
func (wal Wallet) ReadWallet() {
	// Baca alamat direktori di database
	b, err := os.ReadFile("./wallet/")
	if err != nil {
		log.Fatal("Gagal memuat wallet, ", err)
	}
	// Dekrip key
	key, err := keystore.DecryptKey(b, wal.WlInfo.Pwd)
	if err != nil {
		log.Fatal("Password salah, ", err)
	}
	wal.WlInfo.PrivateKey = crypto.FromECDSA(key.PrivateKey)
	wal.WlInfo.PublicKey = crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	wal.WlInfo.PubblicAddress = crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()
	// Simpan ke database
	// Private Key, Public Key dan Public Address
}
