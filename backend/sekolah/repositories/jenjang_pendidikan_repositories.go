package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewJenjangPendidikanRepository(db *gorm.DB) *GenericRepository[models.JenjangPendidikan] {
	return NewGenericRepository[models.JenjangPendidikan](db, "jenjang_pendidikan")
}
