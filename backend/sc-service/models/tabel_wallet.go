package models

import "gorm.io/gorm"

type WalletTable struct {
	gorm.Model
	UserId         string `gorm:"unique;not null"`
	Password       string
	Address        string 
	Name           string //Nama wallet
	WalletFilename string
	
	// PrivateKey string
	// PublicKey string
}
