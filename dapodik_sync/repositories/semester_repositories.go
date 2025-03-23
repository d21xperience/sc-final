package repositories

import (
	"dapodik_sync/models"

	"gorm.io/gorm"
)

func NewSemesterRepository(db *gorm.DB) *GenericRepository[models.Semester] {
	return NewGenericRepository[models.Semester](db, "semesters")
}
