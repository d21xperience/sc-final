package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewKurikulumRepository(db *gorm.DB) *GenericRepository[models.Kurikulum] {
	return NewGenericRepository[models.Kurikulum](db, "kurikulum")
}
