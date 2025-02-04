package repositories

import (
	"sc-service/models"

	"gorm.io/gorm"
)

func NewNetworkRepository(db *gorm.DB) *GenericRepository[models.Network] {
	return NewGenericRepository[models.Network](db, "networks")
}
