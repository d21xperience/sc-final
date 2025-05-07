package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewsekolahTenantRepository(db *gorm.DB) *GenericRepository[models.SekolahTenant] {
	return NewGenericRepository[models.SekolahTenant](db, "sekolah_tenant")
}

// var ErrRecordNotFound = errors.New("record not found")

// type SekolahTenantRepository interface {
// 	Save(sekolahTenant *models.SekolahTenant) error
// 	FindByID(sekolahTenantID int) (*models.SekolahTenant, error)
// 	FindBySchemaName(schemaName string) (*models.SekolahTenant, error)
// 	Update(sekolahTenant *models.SekolahTenant) error
// 	Delete(sekolahTenantID string) error
// }

// type sekolahTenantRepositoryImpl struct {
// 	db *gorm.DB
// }

// func NewsekolahTenantRepository(db *gorm.DB) SekolahTenantRepository {
// 	return &sekolahTenantRepositoryImpl{db: db}
// }

// // Create (Menyimpan Data sekolahTenant)
// func (r *sekolahTenantRepositoryImpl) Save(st *models.SekolahTenant) error {
// 	return r.db.Create(st).Error
// }

// // Read (Mencari Sekolah berdasarkan ID)
// func (r *sekolahTenantRepositoryImpl) FindByID(sekolahTenantID int) (*models.SekolahTenant, error) {
// 	var sekolah models.SekolahTenant
// 	err := r.db.First(&sekolah, "sekolah_id = ?", sekolahTenantID).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &sekolah, nil
// }

// // Update (Memperbarui Data Sekolah)
// func (r *sekolahTenantRepositoryImpl) Update(sekolahTenant *models.SekolahTenant) error {
// 	return r.db.Save(sekolahTenant).Error
// }

// // Delete (Menghapus Data Sekolah berdasarkan ID)
// func (r *sekolahTenantRepositoryImpl) Delete(sekolahTenantID string) error {
// 	return r.db.Delete(&models.Sekolah{}, "sekolah_id = ?", sekolahTenantID).Error
// }
// func (r *sekolahTenantRepositoryImpl) FindBySchemaName(schemaName string) (*models.SekolahTenant, error) {
// 	var sekolah models.SekolahTenant
// 	err := r.db.Where("schema_name = ?", schemaName).First(&sekolah).Error
// 	if err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			return nil, fmt.Errorf("schema %s tidak ditemukan", schemaName)
// 		}
// 		return nil, err
// 	}
// 	return &sekolah, nil
// }
