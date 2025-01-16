package repositories

import (
	"context"
	"fmt"
	"sekolah/models"
	"strings"

	"gorm.io/gorm"
)

type NilaiAkhirRepository interface {
	Save(ctx context.Context, NilaiAkhir *models.TabelNilaiAkhir, schemaName string) error
	FindByID(ctx context.Context, NilaiAkhirID string, schemaName string) (*models.TabelNilaiAkhir, error)
	Update(ctx context.Context, NilaiAkhir *models.TabelNilaiAkhir, schemaName string) error
	Delete(ctx context.Context, NilaiAkhirID string, schemaName string) error
}

type nilaiAkhirRepositoryImpl struct {
	// schemaRepository SchemaRepository
	db *gorm.DB
}

// NewNilaiAkhirRepository membuat instance baru dari NilaiAkhirRepository
func NewNilaiAkhirRepository(dB *gorm.DB) NilaiAkhirRepository {
	return &nilaiAkhirRepositoryImpl{
		db: dB,
		// schemaRepository: NewSchemaRepository(dB),
	}
}

var tabelNilaiAkhir = "tabel_nilaiakhir"

func (r *nilaiAkhirRepositoryImpl) Save(ctx context.Context, NilaiAkhir *models.TabelNilaiAkhir, schemaName string) error {
	// Gunakan transaksi agar atomic
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 🔥 Pastikan schema diubah dalam transaksi
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		// 🔥 Gunakan `tx.Table(schemaName + ".NilaiAkhirs")` agar GORM tahu schema yang benar
		if err := tx.Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), tabelNilaiAkhir)).Create(NilaiAkhir).Error; err != nil {
			return fmt.Errorf("failed to save school in schema %s: %w", schemaName, err)
		}

		return nil
	})
}

func (r *nilaiAkhirRepositoryImpl) FindByID(ctx context.Context, nilaiAkhirID string, schemaName string) (*models.TabelNilaiAkhir, error) {
	var nilaiAkhirModel models.TabelNilaiAkhir

	// 🔥 Pastikan schema diubah sebelum query
	if err := r.db.WithContext(ctx).Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
		return nil, fmt.Errorf("failed to set schema: %w", err)
	}

	// 🔥 Gunakan `tx.Table(schemaName + ".tabel_NilaiAkhir")` agar GORM tahu schema yang benar
	if err := r.db.WithContext(ctx).
		Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), tabelNilaiAkhir)).
		First(&nilaiAkhirModel, "NilaiAkhir_id = ?", nilaiAkhirID).Error; err != nil {
		return nil, fmt.Errorf("failed to find school in schema %s: %w", schemaName, err)
	}

	return &nilaiAkhirModel, nil
}

// Update (Memperbarui Data NilaiAkhir)
func (r *nilaiAkhirRepositoryImpl) Update(ctx context.Context, nilaiAkhirModel *models.TabelNilaiAkhir, schemaName string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 🔥 Set schema sebelum query
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		// 🔥 Lakukan update dalam transaksi
		if err := tx.Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), tabelNilaiAkhir)).
			Where("NilaiAkhir_id = ?", nilaiAkhirModel.IDNilaiAkhir).
			Updates(nilaiAkhirModel).Error; err != nil {
			return fmt.Errorf("failed to update school in schema %s: %w", schemaName, err)
		}

		return nil // Commit transaksi jika tidak ada error
	})
}

// Delete (Menghapus Data NilaiAkhir berdasarkan ID)
func (r *nilaiAkhirRepositoryImpl) Delete(ctx context.Context, nilaiAkhirID string, schemaName string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 🔥 Set schema sebelum query
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		// 🔥 Hapus data dalam transaksi
		if err := tx.Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), tabelNilaiAkhir)).
			Where("NilaiAkhir_id = ?", nilaiAkhirID).
			Delete(nil).Error; err != nil {
			return fmt.Errorf("failed to delete school in schema %s: %w", schemaName, err)
		}

		return nil // Commit transaksi jika tidak ada error
	})
}
