package model

import "gorm.io/gorm"

type WalletTable struct {
	gorm.Model
	UserId   string
	Password string
	Address   string
	Name      string //Nama wallet
	// UrlWallet  string
	// PrivateKey string
	// PublicKey string
}
