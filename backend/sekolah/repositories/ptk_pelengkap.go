package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewPTKPelengkapRepository(db *gorm.DB) *GenericRepository[models.PtkPelengkap] {
	return NewGenericRepository[models.PtkPelengkap](db, "tabel_ptk_pelengkap")
}
