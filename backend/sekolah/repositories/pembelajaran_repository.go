package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewPembelajaranRepository(db *gorm.DB) *GenericRepository[models.Pembelajaran] {
	return NewGenericRepository[models.Pembelajaran](db, "tabel_pembelajaran")
}
