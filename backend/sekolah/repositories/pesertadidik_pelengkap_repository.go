package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewSiswaPelengkapRepository(db *gorm.DB) *GenericRepository[models.PesertaDidikPelengkap] {
	return NewGenericRepository[models.PesertaDidikPelengkap](db, "tabel_siswa_pelengkap")
}
