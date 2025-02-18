package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewPTKTerdaftarRepository(db *gorm.DB) *GenericRepository[models.PTKTerdaftar] {
	return NewGenericRepository[models.PTKTerdaftar](db, "tabel_ptk_terdaftar")
}
