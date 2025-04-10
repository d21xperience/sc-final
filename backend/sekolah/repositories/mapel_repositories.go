package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewMapelRepository(db *gorm.DB) *GenericRepository[models.MataPelajaran] {
	return NewGenericRepository[models.MataPelajaran](db, "mata_pelajaran")
}
