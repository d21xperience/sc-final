package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewJurusanRepository(db *gorm.DB) *GenericRepository[models.Jurusan] {
	return NewGenericRepository[models.Jurusan](db, "jurusan")
}
