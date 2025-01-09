package repositories

import (
	"fmt"
	"sekolah/models"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

type SekolahRepository interface {
	Save(sekolah *models.Sekolah) error
	FindByID(sekolahID string) (*models.Sekolah, error)
	Update(sekolah *models.Sekolah) error
	Delete(sekolahID string) error
	// CreateSekolahDatabase(schemaName string) error
}

type sekolahRepositoryImpl struct {
	db *gorm.DB
}

func NewSekolahRepository(db *gorm.DB) SekolahRepository {
	return &sekolahRepositoryImpl{db: db}
}

// Create (Menyimpan Data Sekolah)
func (r *sekolahRepositoryImpl) Save(sekolah *models.Sekolah) error {
	sekolah.SekolahID = uuid.New() // Generate UUID baru untuk setiap sekolah
	return r.db.Create(sekolah).Error
}

// Read (Mencari Sekolah berdasarkan ID)
func (r *sekolahRepositoryImpl) FindByID(sekolahID string) (*models.Sekolah, error) {
	var sekolah models.Sekolah
	err := r.db.First(&sekolah, "sekolah_id = ?", sekolahID).Error
	if err != nil {
		return nil, err
	}
	return &sekolah, nil
}

// Update (Memperbarui Data Sekolah)
func (r *sekolahRepositoryImpl) Update(sekolah *models.Sekolah) error {
	return r.db.Save(sekolah).Error
}

// Delete (Menghapus Data Sekolah berdasarkan ID)
func (r *sekolahRepositoryImpl) Delete(sekolahID string) error {
	return r.db.Delete(&models.Sekolah{}, "sekolah_id = ?", sekolahID).Error
}

func (r *sekolahRepositoryImpl) CreateSekolahDatabase(namaSekolah string) error {
	// return r.db.Exec("CREATE SCHEMA IF NOT EXISTS ?", schemaName).Error
	// Use GORM to create a new database
	sql := fmt.Sprintf("CREATE DATABASE %s;", namaSekolah)
	if err := r.db.Exec(sql).Error; err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}
	return nil
}
