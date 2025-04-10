package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewBentukPendidikanRepository(db *gorm.DB) *GenericRepository[models.BentukPendidikan] {
	return NewGenericRepository[models.BentukPendidikan](db, "bentuk_pendidikan")
}
