package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

type PesertaDidikRepository interface {
	Save(pesertaDidik *models.PesertaDidik) error
	FindByID(pesertaDidikID string) (*models.PesertaDidik, error)
	Update(pesertaDidik *models.PesertaDidik) error
	Delete(pesertaDidikID string) error
}
type pesertaDidikRepositoryImpl struct {
	db *gorm.DB
}

func NewPesertaDidikRepository(db *gorm.DB) PesertaDidikRepository {
	return &pesertaDidikRepositoryImpl{db: db}
}

// Create (Menyimpan Data Peserta Didik)
func (r *pesertaDidikRepositoryImpl) Save(pesertaDidik *models.PesertaDidik) error {
	return r.db.Create(pesertaDidik).Error
}

// Read (Mencari Peserta Didik berdasarkan ID)
func (r *pesertaDidikRepositoryImpl) FindByID(pesertaDidikID string) (*models.PesertaDidik, error) {
	var pesertaDidik models.PesertaDidik
	err := r.db.First(&pesertaDidik, "peserta_didik_id = ?", pesertaDidikID).Error
	if err != nil {
		return nil, err
	}
	return &pesertaDidik, nil
}

// Update (Memperbarui Data Peserta Didik)
func (r *pesertaDidikRepositoryImpl) Update(pesertaDidik *models.PesertaDidik) error {
	return r.db.Save(pesertaDidik).Error
}

// Delete (Menghapus Data Peserta Didik berdasarkan ID)
func (r *pesertaDidikRepositoryImpl) Delete(pesertaDidikID string) error {
	return r.db.Delete(&models.PesertaDidik{}, "peserta_didik_id = ?", pesertaDidikID).Error
}
