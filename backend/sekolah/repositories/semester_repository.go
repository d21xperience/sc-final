package repositories

import (
	"context"
	"fmt"
	"strings"
	"time"

	"sekolah/models"

	"gorm.io/gorm"
)

type SemesterRepository interface {
	Save(ctx context.Context, Semester *models.Semester, schemaName string) error
	FindByID(ctx context.Context, SemesterID string, schemaName string) (*models.Semester, error)
	FindAll(ctx context.Context, schemaName string, limit, offset int) ([]*models.Semester, error)
	Update(ctx context.Context, Semester *models.Semester, schemaName string) error
	Delete(ctx context.Context, SemesterID string, schemaName string) error
}

type semesterRepositoryImpl struct {
	// schemaRepository SchemaRepository
	db *gorm.DB
}

// NewSemesterRepository membuat instance baru dari SemesterRepository
func NewSemesterRepository(dB *gorm.DB) SemesterRepository {
	return &semesterRepositoryImpl{
		db: dB,
	}
}

var tabelSemester = "semester"

func (r *semesterRepositoryImpl) Save(ctx context.Context, semesterModel *models.Semester, schemaName string) error {
	// Gunakan transaksi agar atomic
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		//  Pastikan schema diubah dalam transaksi
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}
		if err := tx.Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), tabelSemester)).Create(semesterModel).Error; err != nil {
			return fmt.Errorf("failed to save school in schema %s: %w", schemaName, err)
		}

		return nil
	})
}
func (r *semesterRepositoryImpl) FindByID(ctx context.Context, SemesterID string, schemaName string) (*models.Semester, error) {
	var Semester models.Semester

	//  Pastikan schema diubah sebelum query
	if err := r.db.WithContext(ctx).Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
		return nil, fmt.Errorf("failed to set schema: %w", err)
	}

	//  Gunakan `tx.Table(schemaName + ".tabel_Semester")` agar GORM tahu schema yang benar
	if err := r.db.WithContext(ctx).
		Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), tabelSemester)).
		First(&Semester, "Semester_id = ?", SemesterID).Error; err != nil {
		return nil, fmt.Errorf("failed to find school in schema %s: %w", schemaName, err)
	}

	return &Semester, nil
}

// Update (Memperbarui Data Semester)
func (r *semesterRepositoryImpl) Update(ctx context.Context, Semester *models.Semester, schemaName string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		//  Set schema sebelum query
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		//  Lakukan update dalam transaksi
		if err := tx.Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), tabelSemester)).
			Where("Semester_id = ?", Semester.SemesterID).
			Updates(Semester).Error; err != nil {
			return fmt.Errorf("failed to update school in schema %s: %w", schemaName, err)
		}

		return nil // Commit transaksi jika tidak ada error
	})
}

// Delete (Menghapus Data Semester berdasarkan ID)
func (r *semesterRepositoryImpl) Delete(ctx context.Context, SemesterID string, schemaName string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		//  Set schema sebelum query
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		//  Hapus data dalam transaksi
		if err := tx.Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), tabelSemester)).
			Where("Semester_id = ?", SemesterID).
			Delete(nil).Error; err != nil {
			return fmt.Errorf("failed to delete school in schema %s: %w", schemaName, err)
		}

		return nil // Commit transaksi jika tidak ada error
	})
}
func (r *semesterRepositoryImpl) FindAll(ctx context.Context, schemaName string, limit, offset int) ([]*models.Semester, error) {
	var SemesterList []*models.Semester

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := r.db.WithContext(ctx).
		Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), tabelSemester)).
		Limit(limit).
		Offset(offset).
		Find(&SemesterList).Error; err != nil {
		return nil, fmt.Errorf("failed to find records in schema %s: %w", schemaName, err)
	}

	return SemesterList, nil
}
