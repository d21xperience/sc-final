package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewKenaikanRepository(db *gorm.DB) *GenericRepository[models.Kenaikan] {
	return NewGenericRepository[models.Kenaikan](db, "tabel_kenaikan")
}
