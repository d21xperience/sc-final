package repositories

import (
	"sc-service/models"

	"gorm.io/gorm"
)

func NewAccountRepository(db *gorm.DB) *GenericRepository[models.Account] {
	return NewGenericRepository[models.Account](db, "accounts")
}
