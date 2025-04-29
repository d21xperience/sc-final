package repositories

import (
	"sc-service/models"

	"gorm.io/gorm"
)

func NewDegreeDataRepository(db *gorm.DB) *GenericRepository[models.DegreeData] {
	return NewGenericRepository[models.DegreeData](db, "degree_data")
}
