package repositories

import (
	"errors"
	"log"
	"sc-service/models"

	"gorm.io/gorm"
)

var ErrRecordNotFound = errors.New("record not found")

// func NewSekolahTenantRepository(db *gorm.DB) *GenericRepository[models.SekolahTenant] {
// 	return NewGenericRepository[models.SekolahTenant](db, "sekolah_tenants")
// }

type SekolahTenantRepository interface {
	Save(sekolahTenant *models.SekolahTenant) error
	FindByID(sekolahTenantID int) (*models.SekolahTenant, error)
	Update(sekolahTenant *models.SekolahTenant) error
	FindByName(schemaname string) (*models.SekolahTenant, error)
	// Delete(sekolahTenantID string) error
}

type sekolahTenantRepositoryImpl struct {
	db *gorm.DB
}

func NewsekolahTenantRepository(db *gorm.DB) SekolahTenantRepository {
	return &sekolahTenantRepositoryImpl{db: db}
}

// Create (Menyimpan Data sekolahTenant)
func (r *sekolahTenantRepositoryImpl) Save(st *models.SekolahTenant) error {
	return r.db.Create(st).Error
}

// Read (Mencari Sekolah berdasarkan ID)
func (r *sekolahTenantRepositoryImpl) FindByID(sekolahTenantID int) (*models.SekolahTenant, error) {
	var sekolah models.SekolahTenant
	err := r.db.First(&sekolah, "sekolah_id = ?", sekolahTenantID).Error
	if err != nil {
		return nil, err
	}
	return &sekolah, nil
}

// Read (Mencari Sekolah berdasarkan nama)
func (r *sekolahTenantRepositoryImpl) FindByName(schemaname string) (*models.SekolahTenant, error) {
	var sekolah models.SekolahTenant

	//  Gunakan Where().First() untuk menangani query dengan lebih baik
	err := r.db.Where("schema_name = ?", schemaname).First(&sekolah).Error

	//  Tangani kasus jika data tidak ditemukan
	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("FindByName: Data tidak ditemukan untuk schemaname = %s", schemaname)
		return nil, nil // Return nil tanpa error untuk menunjukkan data tidak ada
	}

	//  Tangani error lain dari database
	if err != nil {
		log.Printf("FindByName: Gagal mengambil data dari database: %v", err)
		return nil, err
	}

	return &sekolah, nil
}

// Update (Memperbarui Data Sekolah)
func (r *sekolahTenantRepositoryImpl) Update(sekolahTenant *models.SekolahTenant) error {
	return r.db.Save(sekolahTenant).Error
}

// Delete (Menghapus Data Sekolah berdasarkan ID)
// func (r *sekolahTenantRepositoryImpl) Delete(sekolahTenantID string) error {
// 	return r.db.Delete(&models.Sekolah{}, "sekolah_id = ?", sekolahTenantID).Error
// }
