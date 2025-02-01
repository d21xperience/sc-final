package repositories

import (
	"sc-service/models"

	"gorm.io/gorm"
)

func NewWalletTableRepository(db *gorm.DB) *GenericRepository[models.WalletTable] {
	return NewGenericRepository[models.WalletTable](db, "wallet_tables")
}
