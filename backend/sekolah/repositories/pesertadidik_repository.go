package repositories

import (
	"context"
	"fmt"
	"sekolah/models"
	"strings"

	"gorm.io/gorm"
)

type PesertaDidikRepository interface {
	Save(ctx context.Context, pesertaDidik *models.PesertaDidik, schemaName string) error
	FindByID(ctx context.Context, pesertaDidikID string, schemaName string) (*models.PesertaDidik, error)
	Update(ctx context.Context, pesertaDidik *models.PesertaDidik, pdPelengkap *models.PesertaDidikPelengkap, schemaName string) error
	Delete(ctx context.Context, pesertaDidikID string, schemaName string) error
}

type pesertaDidikRepositoryImpl struct {
	db *gorm.DB
}

// NewpesertaDidikRepository membuat instance baru dari pesertaDidikRepository
func NewpesertaDidikRepository(dB *gorm.DB) PesertaDidikRepository {
	return &pesertaDidikRepositoryImpl{
		db: dB,
		// schemaRepository: NewSchemaRepository(dB),
	}
}

var tabelPesertaDidik = "tabel_pesertaDidik"

func (r *pesertaDidikRepositoryImpl) Save(ctx context.Context, pesertaDidik *models.PesertaDidik, schemaName string) error {
	// Gunakan transaksi agar atomic
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 🔥 Pastikan schema diubah dalam transaksi
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		// 🔥 Gunakan `tx.Table(schemaName + ".pesertaDidiks")` agar GORM tahu schema yang benar
		if err := tx.Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), tabelPesertaDidik)).Create(pesertaDidik).Error; err != nil {
			return fmt.Errorf("failed to save school in schema %s: %w", schemaName, err)
		}

		return nil
	})
}

func (r *pesertaDidikRepositoryImpl) FindByID(ctx context.Context, pesertaDidikID string, schemaName string) (*models.PesertaDidik, error) {
	var pesertaDidik models.PesertaDidik

	// 🔥 Pastikan schema diubah sebelum query
	if err := r.db.WithContext(ctx).Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
		return nil, fmt.Errorf("failed to set schema: %w", err)
	}

	// 🔥 Gunakan `tx.Table(schemaName + ".tabel_pesertaDidik")` agar GORM tahu schema yang benar
	if err := r.db.WithContext(ctx).
		Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), tabelPesertaDidik)).
		First(&pesertaDidik, "pesertaDidik_id = ?", pesertaDidikID).Error; err != nil {
		return nil, fmt.Errorf("failed to find school in schema %s: %w", schemaName, err)
	}

	return &pesertaDidik, nil
}

// Update (Memperbarui Data pesertaDidik)
func (r *pesertaDidikRepositoryImpl) Update(ctx context.Context, pesertaDidik *models.PesertaDidik, pdPelengkap *models.PesertaDidikPelengkap, schemaName string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 🔥 Set schema sebelum query
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		// 🔥 Lakukan update dalam transaksi
		if err := tx.Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), tabelPesertaDidik)).
			Where("pesertaDidik_id = ?", pesertaDidik.PesertaDidikID).
			Updates(pesertaDidik).Error; err != nil {
			return fmt.Errorf("failed to update school in schema %s: %w", schemaName, err)
		}

		return nil // Commit transaksi jika tidak ada error
	})
}

// Delete (Menghapus Data pesertaDidik berdasarkan ID)
func (r *pesertaDidikRepositoryImpl) Delete(ctx context.Context, pesertaDidikID string, schemaName string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 🔥 Set schema sebelum query
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		// 🔥 Hapus data dalam transaksi
		if err := tx.Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), tabelPesertaDidik)).
			Where("pesertaDidik_id = ?", pesertaDidikID).
			Delete(nil).Error; err != nil {
			return fmt.Errorf("failed to delete school in schema %s: %w", schemaName, err)
		}

		return nil // Commit transaksi jika tidak ada error
	})
}
