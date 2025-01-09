package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

type NilaiAkhirRepository interface {
	Save(nilaiAkhir *models.TabelNilaiAkhir) error
	FindByID(idNilaiAkhir string) (*models.TabelNilaiAkhir, error)
	Update(nilaiAkhir *models.TabelNilaiAkhir) error
	Delete(idNilaiAkhir string) error
}
type nilaiAkhirRepositoryImpl struct {
	db *gorm.DB
}

func NewNilaiAkhirRepository(db *gorm.DB) NilaiAkhirRepository {
	return &nilaiAkhirRepositoryImpl{db: db}
}

// Create (Menyimpan Data Nilai Akhir)
func (r *nilaiAkhirRepositoryImpl) Save(nilaiAkhir *models.TabelNilaiAkhir) error {
	return r.db.Create(nilaiAkhir).Error
}

// Read (Mencari Nilai Akhir berdasarkan ID)
func (r *nilaiAkhirRepositoryImpl) FindByID(idNilaiAkhir string) (*models.TabelNilaiAkhir, error) {
	var nilaiAkhir models.TabelNilaiAkhir
	err := r.db.First(&nilaiAkhir, "id_nilai_akhir = ?", idNilaiAkhir).Error
	if err != nil {
		return nil, err
	}
	return &nilaiAkhir, nil
}

// Update (Memperbarui Data Nilai Akhir)
func (r *nilaiAkhirRepositoryImpl) Update(nilaiAkhir *models.TabelNilaiAkhir) error {
	return r.db.Save(nilaiAkhir).Error
}

// Delete (Menghapus Data Nilai Akhir berdasarkan ID)
func (r *nilaiAkhirRepositoryImpl) Delete(idNilaiAkhir string) error {
	return r.db.Delete(&models.TabelNilaiAkhir{}, "id_nilai_akhir = ?", idNilaiAkhir).Error
}
