package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewStatusKepemilikanRepository(db *gorm.DB) *GenericRepository[models.StatusKepemilikan] {
	return NewGenericRepository[models.StatusKepemilikan](db, "status_kepemilikan")
}
