package models

import (
	"gorm.io/gorm"
)

// Model untuk data akun blockchain
type Account struct {
	gorm.Model

	Address   string `gorm:"uniqueIndex;not null" json:"address"`
	Name      string `json:"name" gorm:"type:varchar(255);null"`
	NetworkID uint   `gorm:"index" json:"network_id"`
	// Blockchain   string  `gorm:"not null" json:"blockchain"` // Ethereum, Quorum, Fabric
	Organization string  `json:"organization,omitempty"` // Untuk Hyperledger Fabric
	Amount       float64 `json:"amount,omitempty"`
	PublicKey    string  `gorm:"type:text" json:"public_key"`
	PrivateKey   string  `gorm:"type:text" json:"private_key"`
	// ammount untuk ethereum
	IsActive bool `json:"isActive,omitempty"`
}
