package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewSemesterRepository(db *gorm.DB) *GenericRepository[models.Semester] {
	return NewGenericRepository[models.Semester](db, "semester")
}
