package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewRombelAnggotaRepository(db *gorm.DB) *GenericRepository[models.RombelAnggota] {
	return NewGenericRepository[models.RombelAnggota](db, "tabel_anggotakelas")
}
