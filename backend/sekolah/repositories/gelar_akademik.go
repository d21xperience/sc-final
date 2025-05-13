package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewGelarAkademikRepository(db *gorm.DB) *GenericRepository[models.GelarAkademik] {
	return NewGenericRepository[models.GelarAkademik](db, "gelar_akademik")
}
