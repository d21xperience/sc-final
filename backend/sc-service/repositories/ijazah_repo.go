package repositories

import (
	"sc-service/models"

	"gorm.io/gorm"
)

func NewIjazahBcRepository(db *gorm.DB) *GenericRepository[models.IjazahBc] {
	return NewGenericRepository[models.IjazahBc](db, "ijazah_bc")
}
