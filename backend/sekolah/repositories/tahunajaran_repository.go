package repositories

import (
	"context"
	"fmt"
	"strings"
	"time"

	"sekolah/models"

	"gorm.io/gorm"
)

type TahunAjaranRepository interface {
	Save(ctx context.Context, TahunAjaran *models.TahunAjaran, schemaName string) error
	FindAll(ctx context.Context, schemaName string, limit, offset int) ([]*models.TahunAjaran, error)
	FindByID(ctx context.Context, TahunAjaranID string, schemaName string) (*models.TahunAjaran, error)
	Update(ctx context.Context, TahunAjaran *models.TahunAjaran, schemaName string) error
	Delete(ctx context.Context, TahunAjaranID string, schemaName string) error
}

type TahunAjaranRepositoryImpl struct {
	// schemaRepository SchemaRepository
	db *gorm.DB
}

// NewTahunAjaranRepository membuat instance baru dari TahunAjaranRepository
func NewTahunAjaranRepository(dB *gorm.DB) TahunAjaranRepository {
	return &TahunAjaranRepositoryImpl{
		db: dB,
	}
}

var tabelTahunAjaran = "tahun_ajaran"

func (r *TahunAjaranRepositoryImpl) Save(ctx context.Context, tahunAjaran *models.TahunAjaran, schemaName string) error {
	// Gunakan transaksi agar atomic
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		//  Pastikan schema diubah dalam transaksi
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}
		if err := tx.Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), tabelTahunAjaran)).Create(tahunAjaran).Error; err != nil {
			return fmt.Errorf("failed to save school in schema %s: %w", schemaName, err)
		}

		return nil
	})
}

func (r *TahunAjaranRepositoryImpl) FindAll(ctx context.Context, schemaName string, limit, offset int) ([]*models.TahunAjaran, error) {
	var tahunAjaranList []*models.TahunAjaran

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := r.db.WithContext(ctx).
		Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), tabelTahunAjaran)).
		Limit(limit).
		Offset(offset).
		Find(&tahunAjaranList).Error; err != nil {
		return nil, fmt.Errorf("failed to find records in schema %s: %w", schemaName, err)
	}

	return tahunAjaranList, nil
}
func (r *TahunAjaranRepositoryImpl) FindByID(ctx context.Context, TahunAjaranID string, schemaName string) (*models.TahunAjaran, error) {
	var TahunAjaran models.TahunAjaran

	//  Pastikan schema diubah sebelum query
	if err := r.db.WithContext(ctx).Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
		return nil, fmt.Errorf("failed to set schema: %w", err)
	}

	//  Gunakan `tx.Table(schemaName + ".tabel_TahunAjaran")` agar GORM tahu schema yang benar
	if err := r.db.WithContext(ctx).
		Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), tabelTahunAjaran)).
		First(&TahunAjaran, "tahun_ajaran_id = ?", TahunAjaranID).Error; err != nil {
		return nil, fmt.Errorf("failed to find school in schema %s: %w", schemaName, err)
	}

	return &TahunAjaran, nil
}

// Update (Memperbarui Data TahunAjaran)
func (r *TahunAjaranRepositoryImpl) Update(ctx context.Context, TahunAjaran *models.TahunAjaran, schemaName string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		//  Set schema sebelum query
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		//  Lakukan update dalam transaksi
		if err := tx.Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), tabelTahunAjaran)).
			Where("TahunAjaran_id = ?", TahunAjaran.TahunAjaranID).
			Updates(TahunAjaran).Error; err != nil {
			return fmt.Errorf("failed to update school in schema %s: %w", schemaName, err)
		}

		return nil // Commit transaksi jika tidak ada error
	})
}

// Delete (Menghapus Data TahunAjaran berdasarkan ID)
func (r *TahunAjaranRepositoryImpl) Delete(ctx context.Context, TahunAjaranID string, schemaName string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		//  Set schema sebelum query
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		//  Hapus data dalam transaksi
		if err := tx.Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), tabelTahunAjaran)).
			Where("TahunAjaran_id = ?", TahunAjaranID).
			Delete(nil).Error; err != nil {
			return fmt.Errorf("failed to delete school in schema %s: %w", schemaName, err)
		}

		return nil // Commit transaksi jika tidak ada error
	})
}
