package repository

import (
	"auth_service/models"
	"errors"

	"gorm.io/gorm"
)

type SekolahRepository interface {
	CreateSekolah(*models.Sekolah) error
	// CreateSchemaForSekolah(schemaName string) error
	GetSekolahByNpsn(npsn string) (*models.Sekolah, error)
	// FindByID(userID string) (*model.User, error)
	// Save(user *model.User) error
}

type sekolahRepositoryImpl struct {
	DB *gorm.DB
}

// ErrRecordNotFound is returned when no records are found
var ErrRecordNotFound = errors.New("record not found")

func NewSekolahRepository(db *gorm.DB) SekolahRepository {
	return &sekolahRepositoryImpl{DB: db}
}

func (sri sekolahRepositoryImpl) CreateSekolah(s *models.Sekolah) error {
	result := sri.DB.Create(&s)
	if result.Error != nil {
		// Penangann Error jika terjadi duplikate
		return result.Error
	}
	return nil
}
func (sri sekolahRepositoryImpl) GetSekolahByNpsn(npsn string) (*models.Sekolah, error) {
	var sekolah models.Sekolah
	err := sri.DB.Where("npsn = ?", npsn).First(&sekolah).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	return &sekolah, nil
}

// func (sr sekolahRepositoryImpl) CreateSchemaForSekolah(schemaName string) error {
// 	return sr.DB.Exec("CREATE SCHEMA IF NOT EXISTS ?", schemaName).Error
// }
