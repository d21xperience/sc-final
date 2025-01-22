package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewNilaiAkhirRepository(db *gorm.DB) *GenericRepository[models.NilaiAkhir] {
	return NewGenericRepository[models.NilaiAkhir](db, "tabel_nilaiakhir")
}
