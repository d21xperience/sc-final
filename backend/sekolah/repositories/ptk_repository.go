package repositories

import (
	"sekolah/models"

	"gorm.io/gorm"
)

func NewPTKRepository(db *gorm.DB) *GenericRepository[models.TabelPTK] {
	return NewGenericRepository[models.TabelPTK](db, "tabel_ptk")
}

// type PTKRepository interface {
// 	Save(ctx context.Context, PTK *models.TabelPTK, schemaName string) error
// 	FindByID(ctx context.Context, PTKID string, schemaName string) (*models.TabelPTK, error)
// 	Update(ctx context.Context, PTK *models.TabelPTK, schemaName string) error
// 	Delete(ctx context.Context, PTKID string, schemaName string) error
// }

// type PTKRepositoryImpl struct {
// 	// schemaRepository SchemaRepository
// 	db *gorm.DB
// }

// // NewPTKRepository membuat instance baru dari PTKRepository
// func NewPTKRepository(dB *gorm.DB) PTKRepository {
// 	return &PTKRepositoryImpl{
// 		db: dB,
// 		// schemaRepository: NewSchemaRepository(dB),
// 	}
// }

// var tabelPTK = "tabel_ptk"

// func (r *PTKRepositoryImpl) Save(ctx context.Context, PTK *models.TabelPTK, schemaName string) error {
// 	// Gunakan transaksi agar atomic
// 	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
// 		// Pastikan schema diubah dalam transaksi
// 		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
// 			return fmt.Errorf("failed to set schema: %w", err)
// 		}

// 		// Gunakan `tx.Table(schemaName + ".PTKs")` agar GORM tahu schema yang benar
// 		if err := tx.Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), tabelPTK)).Create(PTK).Error; err != nil {
// 			return fmt.Errorf("failed to save school in schema %s: %w", schemaName, err)
// 		}

// 		return nil
// 	})
// }

// func (r *PTKRepositoryImpl) FindByID(ctx context.Context, PTKID string, schemaName string) (*models.TabelPTK, error) {
// 	var PTK models.TabelPTK

// 	// Pastikan schema diubah sebelum query
// 	if err := r.db.WithContext(ctx).Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
// 		return nil, fmt.Errorf("failed to set schema: %w", err)
// 	}

// 	// Gunakan `tx.Table(schemaName + ".tabel_PTK")` agar GORM tahu schema yang benar
// 	if err := r.db.WithContext(ctx).
// 		Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), tabelPTK)).
// 		First(&PTK, "ptk_id = ?", PTKID).Error; err != nil {
// 		return nil, fmt.Errorf("failed to find school in schema %s: %w", schemaName, err)
// 	}

// 	return &PTK, nil
// }

// // Update (Memperbarui Data PTK)
// func (r *PTKRepositoryImpl) Update(ctx context.Context, PTK *models.TabelPTK, schemaName string) error {
// 	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
// 		// Set schema sebelum query
// 		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
// 			return fmt.Errorf("failed to set schema: %w", err)
// 		}

// 		// Lakukan update dalam transaksi
// 		if err := tx.Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), tabelPTK)).
// 			Where("ptk_id = ?", PTK.PTKID).
// 			Updates(PTK).Error; err != nil {
// 			return fmt.Errorf("failed to update school in schema %s: %w", schemaName, err)
// 		}

// 		return nil // Commit transaksi jika tidak ada error
// 	})
// }

// Delete (Menghapus Data PTK berdasarkan ID)
// func (r *PTKRepositoryImpl) Delete(ctx context.Context, PTKID string, schemaName string) error {
// 	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
// 		// Set schema sebelum query
// 		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
// 			return fmt.Errorf("failed to set schema: %w", err)
// 		}

// 		// Hapus data dalam transaksi
// 		if err := tx.Table(fmt.Sprintf("%s.%v", strings.ToLower(schemaName), tabelPTK)).
// 			Where("PTK_id = ?", PTKID).
// 			Delete(nil).Error; err != nil {
// 			return fmt.Errorf("failed to delete school in schema %s: %w", schemaName, err)
// 		}

// 		return nil // Commit transaksi jika tidak ada error
// 	})
// }
