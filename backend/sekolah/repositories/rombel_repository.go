package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewrombonganBelajarRepository(db *gorm.DB) *GenericRepository[models.RombonganBelajar] {
	return NewGenericRepository[models.RombonganBelajar](db, "tabel_kelas")
}
