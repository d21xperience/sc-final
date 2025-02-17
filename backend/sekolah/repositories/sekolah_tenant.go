package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewsekolahTenantRepository(db *gorm.DB) *GenericRepository[models.SekolahTabelTenant] {
	return NewGenericRepository[models.SekolahTabelTenant](db, "sekolah_tabel_tenants")
}

// var ErrRecordNotFound = errors.New("record not found")

// type SekolahTenantRepository interface {
// 	Save(sekolahTenant *models.SekolahTabelTenant) error
// 	FindByID(sekolahTenantID int) (*models.SekolahTabelTenant, error)
// 	FindBySchemaName(schemaName string) (*models.SekolahTabelTenant, error)
// 	Update(sekolahTenant *models.SekolahTabelTenant) error
// 	Delete(sekolahTenantID string) error
// }

// type sekolahTenantRepositoryImpl struct {
// 	db *gorm.DB
// }

// func NewsekolahTenantRepository(db *gorm.DB) SekolahTenantRepository {
// 	return &sekolahTenantRepositoryImpl{db: db}
// }

// // Create (Menyimpan Data sekolahTenant)
// func (r *sekolahTenantRepositoryImpl) Save(st *models.SekolahTabelTenant) error {
// 	return r.db.Create(st).Error
// }

// // Read (Mencari Sekolah berdasarkan ID)
// func (r *sekolahTenantRepositoryImpl) FindByID(sekolahTenantID int) (*models.SekolahTabelTenant, error) {
// 	var sekolah models.SekolahTabelTenant
// 	err := r.db.First(&sekolah, "sekolah_id = ?", sekolahTenantID).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &sekolah, nil
// }

// // Update (Memperbarui Data Sekolah)
// func (r *sekolahTenantRepositoryImpl) Update(sekolahTenant *models.SekolahTabelTenant) error {
// 	return r.db.Save(sekolahTenant).Error
// }

// // Delete (Menghapus Data Sekolah berdasarkan ID)
// func (r *sekolahTenantRepositoryImpl) Delete(sekolahTenantID string) error {
// 	return r.db.Delete(&models.Sekolah{}, "sekolah_id = ?", sekolahTenantID).Error
// }
// func (r *sekolahTenantRepositoryImpl) FindBySchemaName(schemaName string) (*models.SekolahTabelTenant, error) {
// 	var sekolah models.SekolahTabelTenant
// 	err := r.db.Where("schema_name = ?", schemaName).First(&sekolah).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return nil, fmt.Errorf("schema %s tidak ditemukan", schemaName)
// 		}
// 		return nil, err
// 	}
// 	return &sekolah, nil
// }
