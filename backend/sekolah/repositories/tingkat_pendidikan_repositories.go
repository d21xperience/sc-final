package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewTingkatPendidikanRepository(db *gorm.DB) *GenericRepository[models.TingkatPendidikan] {
	return NewGenericRepository[models.TingkatPendidikan](db, "tingkat_pendidikan")
}
