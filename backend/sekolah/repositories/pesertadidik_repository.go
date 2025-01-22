package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewpesertaDidikRepository(db *gorm.DB) *GenericRepository[models.PesertaDidik] {
	return NewGenericRepository[models.PesertaDidik](db, "tabel_siswa")
}
