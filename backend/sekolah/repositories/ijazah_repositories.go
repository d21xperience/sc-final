package repositories

import (
	"context"
	"fmt"
	"sekolah/models"
	"strings"

	"gorm.io/gorm"
)

type IjazahRepository interface {
	Save(ctx context.Context, Ijazah *models.Ijazah, schemaName string) error
	FindByID(ctx context.Context, IjazahID string, schemaName string) (*models.Ijazah, error)
	Update(ctx context.Context, Ijazah *models.Ijazah, schemaName string) error
	Delete(ctx context.Context, IjazahID string, schemaName string) error
}

type IjazahRepositoryImpl struct {
	// schemaRepository SchemaRepository
	db *gorm.DB
}

// NewIjazahRepository membuat instance baru dari IjazahRepository
func NewIjazahRepository(dB *gorm.DB) IjazahRepository {
	return &IjazahRepositoryImpl{
		db: dB,
		// schemaRepository: NewSchemaRepository(dB),
	}
}

var tabelIjazah = "tabel_Ijazah"

func (r *IjazahRepositoryImpl) Save(ctx context.Context, Ijazah *models.Ijazah, schemaName string) error {
	// Gunakan transaksi agar atomic
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 🔥 Pastikan schema diubah dalam transaksi
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		// 🔥 Gunakan `tx.Table(schemaName + ".Ijazahs")` agar GORM tahu schema yang benar
		if err := tx.Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), tabelIjazah)).Create(Ijazah).Error; err != nil {
			return fmt.Errorf("failed to save school in schema %s: %w", schemaName, err)
		}

		return nil
	})
}

func (r *IjazahRepositoryImpl) FindByID(ctx context.Context, IjazahID string, schemaName string) (*models.Ijazah, error) {
	var IjazahModel models.Ijazah

	// 🔥 Pastikan schema diubah sebelum query
	if err := r.db.WithContext(ctx).Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
		return nil, fmt.Errorf("failed to set schema: %w", err)
	}

	// 🔥 Gunakan `tx.Table(schemaName + ".tabel_Ijazah")` agar GORM tahu schema yang benar
	if err := r.db.WithContext(ctx).
		Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), tabelIjazah)).
		First(&IjazahModel, "Ijazah_id = ?", IjazahID).Error; err != nil {
		return nil, fmt.Errorf("failed to find school in schema %s: %w", schemaName, err)
	}

	return &IjazahModel, nil
}

// Update (Memperbarui Data Ijazah)
func (r *IjazahRepositoryImpl) Update(ctx context.Context, IjazahModel *models.Ijazah, schemaName string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 🔥 Set schema sebelum query
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		// 🔥 Lakukan update dalam transaksi
		if err := tx.Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), tabelIjazah)).
			Where("Ijazah_id = ?", IjazahModel.ID).
			Updates(IjazahModel).Error; err != nil {
			return fmt.Errorf("failed to update school in schema %s: %w", schemaName, err)
		}

		return nil // Commit transaksi jika tidak ada error
	})
}

// Delete (Menghapus Data Ijazah berdasarkan ID)
func (r *IjazahRepositoryImpl) Delete(ctx context.Context, IjazahID string, schemaName string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 🔥 Set schema sebelum query
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		// 🔥 Hapus data dalam transaksi
		if err := tx.Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), tabelIjazah)).
			Where("Ijazah_id = ?", IjazahID).
			Delete(nil).Error; err != nil {
			return fmt.Errorf("failed to delete school in schema %s: %w", schemaName, err)
		}

		return nil // Commit transaksi jika tidak ada error
	})
}
