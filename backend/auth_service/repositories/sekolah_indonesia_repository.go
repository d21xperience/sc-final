package repositories

import (
	"auth_service/models"

	"gorm.io/gorm"
)

func NewSekolahIndonesiaRepository(db *gorm.DB) *GenericRepository[models.SekolahIndonesia] {
	return NewGenericRepository[models.SekolahIndonesia](db, "sekolah_indonesia")
}
