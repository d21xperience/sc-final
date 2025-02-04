package models

import (
	"time"

	"gorm.io/gorm"
)

type NetworkType string

const (
	Mainnet NetworkType = "mainnet"
	Testnet NetworkType = "testnet"
	Private NetworkType = "private"
)

type AccountType string

const (
	ImportAccount NetworkType = "imported"
	Keystore      NetworkType = "keystore"
)

// WalletTransaction menyimpan riwayat transaksi pengguna
type WalletTransaction struct {
	ID          uint      `gorm:"primaryKey"`
	AccountID   uint      `gorm:"not null;index"`       // Relasi ke akun
	Hash        string    `gorm:"uniqueIndex;not null"` // Hash transaksi Ethereum
	Amount      float64   // Jumlah ETH atau token lain
	TokenSymbol string    // Simbol token jika bukan ETH
	Timestamp   time.Time // Waktu transaksi
}

// type WalletTable struct {
// 	gorm.Model
// 	UserId         string `gorm:"unique;not null"`
// 	Password       string
// 	Address        string
// 	Name           string //Nama wallet
// 	WalletFilename string

// 	// PrivateKey string
// 	// PublicKey string
// }

// Network menyimpan informasi jaringan blockchain
type Network struct {
	ID          uint32      `gorm:"primaryKey"`
	Name        string      `gorm:"size:100;not null;unique"`            // Nama jaringan (Ethereum, Polygon, BSC)
	ChainID     int64       `gorm:"not null;unique"`                     // Chain ID jaringan
	RPCURL      string      `gorm:"size:255;not null"`                   // URL RPC jaringan
	ExplorerURL string      `gorm:"size:255"`                            // URL block explorer
	Symbol      string      `gorm:"size:10;not null"`                    // Simbol token utama (ETH, MATIC, BNB)
	Type        NetworkType `gorm:"type:network_type;default:'mainnet'"` // ENUM di PostgreSQL
	CreatedAt   time.Time
	UpdatedAt   time.Time
	// kolom Activate digunakan untuk menampilkan jaringan pada saat pemilihan jaringan
	Activate bool `gorm:"default:false"`
	// kolom NetworkAvailable digunakan jika logic bisnis sudah dibuat, saat ini baru tersedia ethereum, quorum dan hyperledger fabric
	Available bool `gorm:"default:false"`
	// Digunakan untuk menampilkan jaringan secara default
}

// Account menyimpan alamat Ethereum pengguna
type Account struct {
	ID                uint        `gorm:"primaryKey"`
	Address           string      `gorm:"uniqueIndex;not null"` // Alamat Ethereum unik
	Username          string      `gorm:"uniqueIndex;not null"` // Nama pengguna opsional
	Type              AccountType `gorm:"type:account_type;default:'IMPORTED'"`
	UserID            int32       // data dari admin sekolah
	Password          string      `gorm:"size:100"` // digunakan untuk keystore file
	KeystrokeFilename string      // digunakan untuk keystore file
	NetworkID         uint32      `gorm:"not null;index"` // Relasi ke jaringan blockchain
	Network           Network     `gorm:"foreignKey:NetworkID"`
	Organization      string      `json:"organization,omitempty"` // Untuk Hyperledger Fabric
	IsActive          bool        `json:"isActive,omitempty"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

// Contract menyimpan informasi smart contract
type Contract struct {
	ID              uint    `gorm:"primaryKey"`
	Name            string  `gorm:"size:255;not null"`
	ContractAddress string  `gorm:"uniqueIndex;not null"`
	ABI             string  `gorm:"type:text;not null"`
	Bytecode        string  `gorm:"type:text"`
	NetworkID       uint    `gorm:"not null;index"` // Relasi ke jaringan
	Network         Network `gorm:"foreignKey:NetworkID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Transaction menyimpan data transaksi blockchain
type Transaction struct {
	ID         uint      `gorm:"primaryKey"`
	AccountID  uint      `gorm:"not null;index"` // Relasi ke akun pengguna
	Account    Account   `gorm:"foreignKey:AccountID"`
	TxHash     string    `gorm:"uniqueIndex;not null"`
	ContractID *uint     `gorm:"index"` // Opsional, relasi ke kontrak
	Contract   *Contract `gorm:"foreignKey:ContractID"`
	NetworkID  uint      `gorm:"not null;index"` // Relasi ke jaringan blockchain
	Network    Network   `gorm:"foreignKey:NetworkID"`
	Method     string    `gorm:"size:100"`
	InputData  string    `gorm:"type:text"`
	GasUsed    uint64
	Status     string    `gorm:"size:20"` // pending, success, failed
	Timestamp  time.Time // Waktu transaksi
}

// Buat ENUM secara manual sebelum AutoMigrate()
func Migrate(db *gorm.DB) error {
	// Buat ENUM jika belum ada
	query1 := `DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'network_type') THEN CREATE TYPE network_type AS ENUM ('MAINNET', 'TESTNET','PRIVATE'); END IF; END $$;`
	query2 := `DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'account_type') THEN CREATE TYPE account_type AS ENUM ('IMPORTED', 'KEYSTORE'); END IF; END $$;`
	err := db.Exec(query1 + query2).Error
	if err != nil {
		return err
	}
	return db.AutoMigrate(&SekolahTenant{}, &SchemaLog{})
}
