package repositories

import (
	"context"
	"fmt"
	"sekolah/models"
	"strings"

	"gorm.io/gorm"
)

type RombonganBelajarRepository interface {
	Save(ctx context.Context, rombonganBelajar *models.RombonganBelajar, schemaName string) error
	FindByID(ctx context.Context, rombonganBelajarID string, schemaName string) (*models.RombonganBelajar, error)
	Update(ctx context.Context, rombonganBelajar *models.RombonganBelajar, schemaName string) error
	Delete(ctx context.Context, rombonganBelajarID string, schemaName string) error
}

type rombonganBelajarRepositoryImpl struct {
	// schemaRepository SchemaRepository
	db *gorm.DB
}

// NewrombonganBelajarRepository membuat instance baru dari rombonganBelajarRepository
func NewrombonganBelajarRepository(dB *gorm.DB) RombonganBelajarRepository {
	return &rombonganBelajarRepositoryImpl{
		db: dB,
		// schemaRepository: NewSchemaRepository(dB),
	}
}

var tabelrombonganBelajar = "tabel_kelas"

func (r *rombonganBelajarRepositoryImpl) Save(ctx context.Context, rombonganBelajar *models.RombonganBelajar, schemaName string) error {
	// Gunakan transaksi agar atomic
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 🔥 Pastikan schema diubah dalam transaksi
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		// 🔥 Gunakan `tx.Table(schemaName + ".rombonganBelajars")` agar GORM tahu schema yang benar
		if err := tx.Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), tabelrombonganBelajar)).Create(rombonganBelajar).Error; err != nil {
			return fmt.Errorf("failed to save school in schema %s: %w", schemaName, err)
		}

		return nil
	})
}

func (r *rombonganBelajarRepositoryImpl) FindByID(ctx context.Context, rombonganBelajarID string, schemaName string) (*models.RombonganBelajar, error) {
	var rombonganBelajar models.RombonganBelajar

	// 🔥 Pastikan schema diubah sebelum query
	if err := r.db.WithContext(ctx).Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
		return nil, fmt.Errorf("failed to set schema: %w", err)
	}

	// 🔥 Gunakan `tx.Table(schemaName + ".tabel_rombonganBelajar")` agar GORM tahu schema yang benar
	if err := r.db.WithContext(ctx).
		Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), tabelrombonganBelajar)).
		First(&rombonganBelajar, "rombonganBelajar_id = ?", rombonganBelajarID).Error; err != nil {
		return nil, fmt.Errorf("failed to find school in schema %s: %w", schemaName, err)
	}

	return &rombonganBelajar, nil
}

// Update (Memperbarui Data rombonganBelajar)
func (r *rombonganBelajarRepositoryImpl) Update(ctx context.Context, rombel *models.RombonganBelajar, schemaName string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 🔥 Set schema sebelum query
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		// 🔥 Lakukan update dalam transaksi
		if err := tx.Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), tabelrombonganBelajar)).
			Where("rombongan_belajar_id = ?", rombel.RombonganBelajarID).
			Updates(rombel).Error; err != nil {
			return fmt.Errorf("failed to update school in schema %s: %w", schemaName, err)
		}

		return nil // Commit transaksi jika tidak ada error
	})
}

// Delete (Menghapus Data rombonganBelajar berdasarkan ID)
func (r *rombonganBelajarRepositoryImpl) Delete(ctx context.Context, rombelID string, schemaName string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 🔥 Set schema sebelum query
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		// 🔥 Hapus data dalam transaksi
		if err := tx.Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), tabelrombonganBelajar)).
			Where("rombongan_belajar_id = ?", rombelID).
			Delete(nil).Error; err != nil {
			return fmt.Errorf("failed to delete school in schema %s: %w", schemaName, err)
		}

		return nil // Commit transaksi jika tidak ada error
	})
}
