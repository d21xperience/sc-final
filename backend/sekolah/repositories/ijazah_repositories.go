package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewIjazahRepository(db *gorm.DB) *GenericRepository[models.Ijazah] {
	return NewGenericRepository[models.Ijazah](db, "tabel_anggotakelas")
}
