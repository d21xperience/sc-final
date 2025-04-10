package repositories

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

type GenericRepository[T any] struct {
	db        *gorm.DB
	tableName string
}

// Membuat instance baru dari GenericRepository
func NewGenericRepository[T any](db *gorm.DB, tableName string) *GenericRepository[T] {
	return &GenericRepository[T]{
		db:        db,
		tableName: tableName,
	}
}

// CRUD Operations
func (r *GenericRepository[T]) Save(ctx context.Context, entity *T, schemaName string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		if err := tx.Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), r.tableName)).Create(entity).Error; err != nil {
			return fmt.Errorf("failed to save record in schema %s: %w", schemaName, err)
		}

		return nil
	})
}

// func (r *GenericRepository[T]) FindByID(ctx context.Context, id string, schemaName, idColumn string) (*T, error) {
// 	var entity T
// 	if err := r.db.WithContext(ctx).Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
// 		return nil, fmt.Errorf("failed to set schema: %w", err)
// 	}

// 	if err := r.db.WithContext(ctx).
// 		Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), r.tableName)).
// 		First(&entity, fmt.Sprintf("%s = ?", idColumn), id).Error; err != nil {
// 		return nil, fmt.Errorf("failed to find record in schema %s: %w", schemaName, err)
// 	}

//		return &entity, nil
//	}
func (r *GenericRepository[T]) FindByID(ctx context.Context, id string, schemaName, idColumn string) (*T, error) {
	var entity T

	// Set Schema (Multi-Tenant)
	if err := r.db.WithContext(ctx).Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
		return nil, fmt.Errorf("failed to set schema: %w", err)
	}

	// Query Data
	err := r.db.WithContext(ctx).
		Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), r.tableName)).
		First(&entity, fmt.Sprintf("%s = ?", idColumn), id).Error

	// Handle jika data tidak ditemukan
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // Mengembalikan nil tanpa error jika tidak ditemukan
	} else if err != nil {
		return nil, fmt.Errorf("failed to find record in schema %s: %w", schemaName, err)
	}

	return &entity, nil
}
func (r *GenericRepository[T]) FindAll(ctx context.Context, schemaName string, limit, offset int) ([]*T, error) {
	var entities []*T
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := r.db.WithContext(ctx).
		Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), r.tableName)).
		Limit(limit).
		Offset(offset).
		Find(&entities).Error; err != nil {
		return nil, fmt.Errorf("failed to find records in schema %s: %w", schemaName, err)
	}

	return entities, nil
}
func (r *GenericRepository[T]) FindAllByConditions(
	ctx context.Context,
	schemaName string,
	conditions map[string]any, // Parameter untuk kondisi WHERE
	limit, offset int,
) ([]*T, error) {
	var entities []*T
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Query ke database dengan kondisi WHERE
	query := r.db.WithContext(ctx).
		Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), r.tableName)).
		Limit(limit).
		Offset(offset)

	// Tambahkan kondisi WHERE jika ada
	if len(conditions) > 0 {
		query = query.Where(conditions)
	}

	// Eksekusi query
	if err := query.Find(&entities).Error; err != nil {
		return nil, fmt.Errorf("failed to find records in schema %s: %w", schemaName, err)
	}

	return entities, nil
}

func (r *GenericRepository[T]) Update(ctx context.Context, entity *T, schemaName, idColumn, id string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		if err := tx.Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), r.tableName)).
			Where(fmt.Sprintf("%s = ?", idColumn), id).
			Updates(entity).Error; err != nil {
			return fmt.Errorf("failed to update record in schema %s: %w", schemaName, err)
		}

		return nil
	})
}

func (r *GenericRepository[T]) Delete(ctx context.Context, id string, schemaName, idColumn string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		if err := tx.Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), r.tableName)).
			Where(fmt.Sprintf("%s = ?", idColumn), id).
			Delete(nil).Error; err != nil {
			return fmt.Errorf("failed to delete record in schema %s: %w", schemaName, err)
		}

		return nil
	})
}
func (r *GenericRepository[T]) SaveMany(ctx context.Context, schemaName string, entities []*T, batchSize int) error {
	// fmt.Println("eksekusi di savemany")
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Set schema
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		// Bulk insert menggunakan CreateInBatches
		if err := tx.Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), r.tableName)).
			CreateInBatches(entities, batchSize).Error; err != nil {
			return fmt.Errorf("failed to save records in schema %s: %w", schemaName, err)
		}

		return nil
	})
}

// FindWithJoins melakukan query dengan joins dan kondisi tertentu
func (r *GenericRepository[T]) FindWithJoins(ctx context.Context, schemaName string, joins []string, conditions map[string]interface{}) (*T, error) {
	var result T

	// Gunakan transaksi agar bisa set schema lebih aman
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Set schema
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		// Query dengan joins
		query := tx.Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), r.tableName))

		// Apply joins
		for _, join := range joins {
			query = query.Joins(join)
		}

		// Apply conditions
		if len(conditions) > 0 {
			query = query.Where(conditions)
		}

		// Execute query
		if err := query.First(&result).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &result, nil
}

// FindWithPreloadAndJoins - Fungsi generic untuk mendapatkan data dengan Preload dan Joins
func (r *GenericRepository[T]) FindWithPreloadAndJoinsOrigin(
	ctx context.Context,
	schemaName string,
	joins []string,
	preloads []string,
	conditions map[string]interface{},
	orderBy []string, // Tambahkan parameter orderBy
) ([]T, error) {
	var results []T
	tx := r.db.WithContext(ctx)

	// Set Schema (Multi-Tenant)
	if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", schemaName)).Error; err != nil {
		return nil, fmt.Errorf("failed to set schema: %w", err)
	}

	// Tambahkan Joins jika ada
	for _, join := range joins {
		tx = tx.Joins(join)
	}

	// Tambahkan Preload jika ada
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	// Tambahkan ORDER BY jika ada
	if len(orderBy) > 0 {
		tx = tx.Order(strings.Join(orderBy, ", ")) // Gabungkan semua kolom ORDER BY
	}

	// Eksekusi Query dengan kondisi
	if err := tx.Where(conditions).Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

// Fungsi Generic untuk Preload One-To-Many
func (r *GenericRepository[T]) FindWithPreloadAndJoins(
	ctx context.Context,
	schemaName string,
	joins []string,
	preloads []string,
	conditions map[string]interface{},
	groupByColumns []string,
) ([]T, error) {
	var results []T
	tx := r.db.WithContext(ctx)

	// Set Schema (Multi-Tenant)
	if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", schemaName)).Error; err != nil {
		return nil, fmt.Errorf("failed to set schema: %w", err)
	}

	// Tambahkan DISTINCT untuk menghindari duplikasi
	tx = tx.Distinct()

	// Tambahkan Joins jika ada
	for _, join := range joins {
		tx = tx.Joins(join)
	}

	// Tambahkan Preload untuk relasi One-To-Many
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	// Tambahkan GROUP BY jika diperlukan
	if len(groupByColumns) > 0 {
		tx = tx.Group(strings.Join(groupByColumns, ", "))
	}

	// Eksekusi Query dengan kondisi
	if err := tx.Where(conditions).Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

//
// //

// FindWithRelations mengambil data dengan relasi tertentu berdasarkan tipe relasinya
//
// Fungsi Generic untuk Relasi dalam GORM.
//  1. Jenis Relasi dalam GORM
//  2. BelongsTo (One-To-One, Foreign Key di tabel anak)
//  3. HasOne (One-To-One, Foreign Key di tabel induk)
//  4. HasMany (One-To-Many)
//  5. ManyToMany (Many-To-Many, menggunakan tabel pivot)
//
// //
//
//	// 1. One-To-One (BelongsTo & HasOne)
//	preloads := []string{"PTK"}  // Preload PTK dalam PTKTerdaftar
//	conditions := map[string]interface{}{"ptk_terdaftar_id": "some-uuid"}
//	data, err := repo.FindWithRelations(ctx, schemaName, nil, preloads, conditions, nil)
//	// 2. One-To-Many (HasMany)
//	preloads := []string{"Pembelajaran"}  // Preload Pembelajaran dalam PTKTerdaftar
//	conditions := map[string]interface{}{"tahun_ajaran_id": "2023"}
//	data, err := repo.FindWithRelations(ctx, schemaName, nil, preloads, conditions, nil)
//	// 3. Many-To-Many
//	preloads := []string{"Kelas"}  // Preload Kelas dalam Guru
//	conditions := map[string]interface{}{"guru_id": "some-uuid"}
//	data, err := repo.FindWithRelations(ctx, schemaName, nil, preloads, conditions, nil)
//
// [docs]: https://gorm.io/docs/query.html#Conditions
func (r *GenericRepository[T]) FindWithRelations(
	ctx context.Context,
	schemaName string,
	joins []string,
	preloads []string,
	conditions map[string]interface{},
	groupByColumns []string,
) ([]T, error) {
	var results []T
	tx := r.db.WithContext(ctx)

	// Set Schema (Multi-Tenant)
	if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", schemaName)).Error; err != nil {
		return nil, fmt.Errorf("failed to set schema: %w", err)
	}

	// Tambahkan DISTINCT untuk menghindari duplikasi
	tx = tx.Distinct()

	// Tambahkan Joins jika ada (untuk hubungan complex)
	for _, join := range joins {
		tx = tx.Joins(join)
	}

	// Tambahkan Preload untuk relasi One-To-One, One-To-Many, dan Many-To-Many
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	// Tambahkan GROUP BY jika diperlukan
	if len(groupByColumns) > 0 {
		tx = tx.Group(strings.Join(groupByColumns, ", "))
	}

	// Eksekusi Query dengan kondisi
	if err := tx.Where(conditions).Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

func (r *GenericRepository[T]) CountRows(ctx context.Context, schemaName, semesterIdColumn, semesterId string) (int64, error) {
	var count int64

	//  Set schema terlebih dahulu
	if err := r.db.WithContext(ctx).Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
		return 0, fmt.Errorf("failed to set schema: %w", err)
	}

	//  Query COUNT dengan filter berdasarkan `semester_id`
	if err := r.db.WithContext(ctx).
		Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), r.tableName)).
		Where(fmt.Sprintf("%s = ?", semesterIdColumn), semesterId).
		Count(&count).Error; err != nil {
		return 0, fmt.Errorf("failed to count rows in schema %s with semester_id %s: %w", schemaName, semesterId, err)
	}

	return count, nil
}

// Fungsi untuk pencarian data berdasarkan kolom, seperti nama
// Fungsi Generic untuk Preload One-To-Many dengan Pagination dan Pencarian Nama
func (r *GenericRepository[T]) SearchByColumnNamePreloadAndJoins(
	ctx context.Context,
	schemaName string,
	joins []string,
	preloads []string,
	conditions map[string]interface{},
	groupByColumns []string,
	limit int,
	offset int,
	nameColumn string, // Kolom yang akan dicari berdasarkan nama
	nameValue string, // Nilai yang akan dicari dalam kolom nama
) ([]T, int64, error) {
	var results []T
	var totalCount int64
	tx := r.db.WithContext(ctx)

	// Set Schema (Multi-Tenant)
	if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", schemaName)).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to set schema: %w", err)
	}

	// Hitung total data sebelum paginasi
	countTx := tx.Model(new(T))
	for _, join := range joins {
		countTx = countTx.Joins(join)
	}

	// Tambahkan kondisi pencarian berdasarkan nama (LIKE)
	if nameColumn != "" && nameValue != "" {
		searchPattern := fmt.Sprintf("%%%s%%", nameValue) // LIKE '%nameValue%'
		countTx = countTx.Where(fmt.Sprintf("%s ILIKE ?", nameColumn), searchPattern)
	}

	if err := countTx.Where(conditions).Count(&totalCount).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count records: %w", err)
	}

	// Tambahkan DISTINCT untuk menghindari duplikasi
	tx = tx.Distinct()

	// Tambahkan Joins jika ada
	for _, join := range joins {
		tx = tx.Joins(join)
	}

	// Tambahkan Preload untuk relasi One-To-Many
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	// Tambahkan GROUP BY jika diperlukan
	if len(groupByColumns) > 0 {
		tx = tx.Group(strings.Join(groupByColumns, ", "))
	}

	// Tambahkan kondisi pencarian berdasarkan nama
	if nameColumn != "" && nameValue != "" {
		searchPattern := fmt.Sprintf("%%%s%%", nameValue)
		tx = tx.Where(fmt.Sprintf("%s ILIKE ?", nameColumn), searchPattern)
	}

	// Tambahkan Pagination
	if limit > 0 {
		tx = tx.Limit(limit)
	}
	if offset >= 0 {
		tx = tx.Offset(offset)
	}

	// Eksekusi Query dengan kondisi
	if err := tx.Where(conditions).Find(&results).Error; err != nil {
		return nil, 0, err
	}

	return results, totalCount, nil
}

// Fungsi untuk mengambil semua data namun di batasi oleh menggunakan pagination
// Fungsi Generic untuk Preload One-To-Many dengan Pagination
func (r *GenericRepository[T]) FindAllWithPagination(
	ctx context.Context,
	schemaName string,
	joins []string,
	preloads []string,
	conditions map[string]any,
	groupByColumns []string,
	limit int, // Menentukan jumlah data per halaman
	offset int, // Menentukan posisi mulai data
) ([]T, int64, error) { // Tambahkan return totalCount untuk mengetahui jumlah total data
	var results []T
	var totalCount int64
	tx := r.db.WithContext(ctx)

	// Set Schema (Multi-Tenant)
	if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", schemaName)).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to set schema: %w", err)
	}

	// Hitung total data sebelum paginasi
	countTx := tx.Model(new(T))
	for _, join := range joins {
		countTx = countTx.Joins(join)
	}
	if err := countTx.Where(conditions).Count(&totalCount).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count records: %w", err)
	}

	// Tambahkan DISTINCT untuk menghindari duplikasi
	tx = tx.Distinct()

	// Tambahkan Joins jika ada
	for _, join := range joins {
		tx = tx.Joins(join)
	}

	// Tambahkan Preload untuk relasi One-To-Many
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	// Tambahkan GROUP BY jika diperlukan
	if len(groupByColumns) > 0 {
		tx = tx.Group(strings.Join(groupByColumns, ", "))
	}

	// Tambahkan Pagination
	if limit > 0 {
		tx = tx.Limit(limit)
	}
	if offset >= 0 {
		tx = tx.Offset(offset)
	}

	// Eksekusi Query dengan kondisi
	if err := tx.Where(conditions).Find(&results).Error; err != nil {
		return nil, 0, err
	}

	return results, totalCount, nil
}

