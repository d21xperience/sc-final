package repositories

import (
	"context"
	"fmt"
	"strings"

	"sekolah/models"

	"gorm.io/gorm"
)

type SekolahRepository interface {
	Save(ctx context.Context, sekolah *models.Sekolah, schemaName string) error
	Find(ctx context.Context, schemaName string) (*models.Sekolah, error)
	FindByID(ctx context.Context, sekolahID string, schemaName string) (*models.Sekolah, error)
	Update(ctx context.Context, sekolah *models.Sekolah, schemaName string) error
	Delete(ctx context.Context, sekolahID string, schemaName string) error
}

type sekolahRepositoryImpl struct {
	// schemaRepository SchemaRepository
	db *gorm.DB
}

// NewSekolahRepository membuat instance baru dari SekolahRepository
func NewSekolahRepository(dB *gorm.DB) SekolahRepository {
	return &sekolahRepositoryImpl{
		db: dB,
		// schemaRepository: NewSchemaRepository(dB),
	}
}

var namaTabel = "tabel_sekolah"

func (r *sekolahRepositoryImpl) Save(ctx context.Context, sekolah *models.Sekolah, schemaName string) error {
	// Gunakan transaksi agar atomic
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 🔥 Pastikan schema diubah dalam transaksi
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}
		// 🔥 Gunakan `tx.Table(schemaName + ".sekolahs")` agar GORM tahu schema yang benar
		if err := tx.Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), namaTabel)).Create(sekolah).Error; err != nil {
			return fmt.Errorf("failed to save school in schema %s: %w", schemaName, err)
		}

		return nil
	})
}

func (r *sekolahRepositoryImpl) Find(ctx context.Context, schemaName string) (*models.Sekolah, error) {
	var sekolah models.Sekolah

	// 🔥 Pastikan schema diubah sebelum query
	if err := r.db.WithContext(ctx).Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
		return nil, fmt.Errorf("failed to set schema: %w", err)
	}

	// 🔥 Gunakan `tx.Table(schemaName + ".tabel_sekolah")` agar GORM tahu schema yang benar
	if err := r.db.WithContext(ctx).
		Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), namaTabel)).
		First(&sekolah).Error; err != nil {
		return nil, fmt.Errorf("failed to find school in schema %s: %w", schemaName, err)
	}

	return &sekolah, nil
}
func (r *sekolahRepositoryImpl) FindByID(ctx context.Context, sekolahID string, schemaName string) (*models.Sekolah, error) {
	var sekolah models.Sekolah

	// 🔥 Pastikan schema diubah sebelum query
	if err := r.db.WithContext(ctx).Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
		return nil, fmt.Errorf("failed to set schema: %w", err)
	}

	// 🔥 Gunakan `tx.Table(schemaName + ".tabel_sekolah")` agar GORM tahu schema yang benar
	if err := r.db.WithContext(ctx).
		Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), namaTabel)).
		First(&sekolah, "sekolah_id = ?", sekolahID).Error; err != nil {
		return nil, fmt.Errorf("failed to find school in schema %s: %w", schemaName, err)
	}

	return &sekolah, nil
}

// Update (Memperbarui Data Sekolah)
func (r *sekolahRepositoryImpl) Update(ctx context.Context, sekolah *models.Sekolah, schemaName string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 🔥 Set schema sebelum query
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		// 🔥 Lakukan update dalam transaksi
		if err := tx.Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), namaTabel)).
			Where("sekolah_id = ?", sekolah.SekolahID).
			Updates(sekolah).Error; err != nil {
			return fmt.Errorf("failed to update school in schema %s: %w", schemaName, err)
		}

		return nil // Commit transaksi jika tidak ada error
	})
}

// Delete (Menghapus Data Sekolah berdasarkan ID)
func (r *sekolahRepositoryImpl) Delete(ctx context.Context, sekolahID string, schemaName string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 🔥 Set schema sebelum query
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		// 🔥 Hapus data dalam transaksi
		if err := tx.Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), namaTabel)).
			Where("sekolah_id = ?", sekolahID).
			Delete(nil).Error; err != nil {
			return fmt.Errorf("failed to delete school in schema %s: %w", schemaName, err)
		}

		return nil // Commit transaksi jika tidak ada error
	})
}
