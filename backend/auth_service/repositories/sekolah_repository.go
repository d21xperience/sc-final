package repositories

import (
	"auth_service/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type SekolahRepository interface {
	CreateSekolah(*models.SekolahTenant) error
	// GetSekolah(query SekolahQuery) (*models.SekolahTenant, error)
	GetSekolahByNPSN(npsn string) (*models.SekolahTenant, error)
	GetSekolahByTenantId(tenantId uint32) (*models.SekolahTenant, error)
}

type sekolahRepositoryImpl struct {
	DB *gorm.DB
}

type SekolahQuery struct {
	Npsn      string
	SekolahID int
}

// ErrRecordNotFound is returned when no records are found
var ErrRecordNotFound = errors.New("record not found")

func NewSekolahRepository(db *gorm.DB) SekolahRepository {
	return &sekolahRepositoryImpl{DB: db}
}

func (sri *sekolahRepositoryImpl) CreateSekolah(s *models.SekolahTenant) error {
	result := sri.DB.Create(&s)
	if result.Error != nil {
		// Penangann Error jika terjadi duplikate
		return result.Error
	}
	return nil
}
func (sri *sekolahRepositoryImpl) GetSekolah(query SekolahQuery) (*models.SekolahTenant, error) {
	// Validasi: Pastikan minimal satu parameter ada
	if query.Npsn == "" && query.SekolahID == 0 {
		return nil, fmt.Errorf("minimal salah satu parameter (npsn atau sekolah_id) harus disediakan")
	}

	var sekolah models.SekolahTenant

	// Gunakan kedua parameter jika keduanya ada
	dbQuery := sri.DB
	if query.Npsn != "" && query.SekolahID != 0 {
		dbQuery = dbQuery.Where("npsn = ? AND id = ?", query.Npsn, query.SekolahID)
	} else {
		// Gunakan salah satu parameter jika hanya salah satu yang ada
		if query.Npsn != "" {
			dbQuery = dbQuery.Where("npsn = ?", query.Npsn)
		}
		if query.SekolahID != 0 {
			dbQuery = dbQuery.Where("id = ?", query.SekolahID)
		}
	}

	// Eksekusi query
	err := dbQuery.First(&sekolah).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}

	return &sekolah, nil
}
func (sri *sekolahRepositoryImpl) GetSekolahByNPSN(npsn string) (*models.SekolahTenant, error) {
	// Validasi: Pastikan minimal satu parameter ada
	var sekolah models.SekolahTenant
	dbQuery := sri.DB
	dbQuery = dbQuery.Where("npsn = ?", npsn)

	// Eksekusi query
	err := dbQuery.First(&sekolah).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}

	return &sekolah, nil
}
func (sri *sekolahRepositoryImpl) GetSekolahByTenantId(tenantId uint32) (*models.SekolahTenant, error) {
	// Validasi: Pastikan minimal satu parameter ada
	var sekolah models.SekolahTenant
	dbQuery := sri.DB
	dbQuery = dbQuery.Where("id = ?", tenantId)

	// Eksekusi query
	err := dbQuery.First(&sekolah).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}

	return &sekolah, nil
}
