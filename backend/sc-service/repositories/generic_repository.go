package repositories

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
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
	if entity == nil {
		return fmt.Errorf("entity is nil")
	}

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

func (r *GenericRepository[T]) FindByID(ctx context.Context, id string, schemaName, idColumn string) (*T, error) {
	var entity T
	if err := r.db.WithContext(ctx).Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
		return nil, fmt.Errorf("failed to set schema: %w", err)
	}

	if err := r.db.WithContext(ctx).
		Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), r.tableName)).
		First(&entity, fmt.Sprintf("%s = ?", idColumn), id).Error; err != nil {
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
	// Pastikan koneksi database aktif
	sqlDB, err := r.db.DB()
	if err != nil {
		log.Fatal("Gagal mendapatkan koneksi database:", err)
	}
	if err = sqlDB.Ping(); err != nil {
		log.Fatal("Database tidak terhubung:", err)
	}

	query := r.db.WithContext(ctx).
		Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), r.tableName)).
		Limit(limit).
		Offset(offset)

	// Debug: Lihat kondisi yang dikirim
	fmt.Println("Kondisi Query:", conditions)

	// Tambahkan kondisi WHERE jika ada
	if len(conditions) > 0 {
		query = query.Where(conditions)
	}

	// Debug: Lihat query sebelum eksekusi
	dryRunQuery := query.Session(&gorm.Session{DryRun: true}).Find(&entities)
	fmt.Println("Generated SQL:", dryRunQuery.Statement.SQL.String())

	// Eksekusi query
	if err := query.Find(&entities).Error; err != nil {
		return nil, fmt.Errorf("failed to find records in schema %s: %w", schemaName, err)
	}

	return entities, nil
}

// func (r *GenericRepository[T]) Update(ctx context.Context, entity *T, schemaName, idColumn, id string) error {
// 	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
// 		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
// 			return fmt.Errorf("failed to set schema: %w", err)
// 		}

// 		if err := tx.Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), r.tableName)).
// 			Where(fmt.Sprintf("%s = ?", idColumn), id).
// 			Updates(entity).Error; err != nil {
// 			return fmt.Errorf("failed to update record in schema %s: %w", schemaName, err)
// 		}

//			return nil
//		})
//	}
// func (r *GenericRepository[T]) Update(ctx context.Context, entity *T, schemaName, idColumn, id string) error {
// 	if entity == nil {
// 		return fmt.Errorf("entity cannot be nil or invalid pointer")
// 	}

// 	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
// 		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
// 			return fmt.Errorf("failed to set schema: %w", err)
// 		}

// 		// if err := tx.Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), r.tableName)).
// 		// 	Where(fmt.Sprintf("%s = ?", idColumn), id).
// 		// 	// Updates(entity).Error; err != nil {
// 		// 	UpdateColumns(entity).Error; err != nil {
// 		// 	return fmt.Errorf("failed to update record in schema %s: %w", schemaName, err)
// 		// }

// 		// Debugging: Log entity sebelum update
// 		log.Printf("Updating entity in schema %s: %+v", schemaName, entity)

// 		// Pastikan field yang akan diupdate tidak diabaikan (termasuk false)
// 		if err := tx.Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), r.tableName)).
// 			Where(fmt.Sprintf("%s = ?", idColumn), id).
// 			Select("active"). // Memaksa hanya kolom active yang diperbarui
// 			Updates(entity).Error; err != nil {
// 			return fmt.Errorf("failed to update record in schema %s: %w", schemaName, err)
// 		}

//			return nil
//		})
//	}
func (r *GenericRepository[T]) Update(ctx context.Context, entity *T, schemaName, idColumn, id string) error {
	if entity == nil {
		return fmt.Errorf("entity cannot be nil or invalid pointer")
	}

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Pastikan schema di-set dengan benar
		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
			return fmt.Errorf("failed to set schema: %w", err)
		}

		// Debugging: Log entity sebelum update
		log.Printf("Updating entity in schema %s: %+v", schemaName, entity)

		// Gunakan refleksi untuk mengecek apakah field `active` ada
		val := reflect.ValueOf(entity).Elem()
		hasActiveField := val.FieldByName("Active").IsValid()

		query := tx.Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), r.tableName)).
			Where(fmt.Sprintf("%s = ?", idColumn), id)

		// Jika ada field `active`, gunakan Select agar tetap diperbarui
		if hasActiveField {
			query = query.Select("active")
		}

		// Eksekusi update
		if err := query.Updates(entity).Error; err != nil {
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

// func (r *GenericRepository[T]) SaveMany(ctx context.Context, schemaName string, entities []*T, batchSize int) error {
// 	// fmt.Println("eksekusi di savemany")
// 	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
// 		// Set schema
// 		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
// 			return fmt.Errorf("failed to set schema: %w", err)
// 		}

// 		// Bulk insert menggunakan CreateInBatches
// 		if err := tx.Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), r.tableName)).
// 			CreateInBatches(entities, batchSize).Error; err != nil {
// 			return fmt.Errorf("failed to save records in schema %s: %w", schemaName, err)
// 		}

// 		return nil
// 	})
// }

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
	orderBy []string,
	useDistinct bool, // NEW
) ([]T, error) {
	var results []T
	tx := r.db.WithContext(ctx)

	// Set Schema
	if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", schemaName)).Error; err != nil {
		return nil, fmt.Errorf("failed to set schema: %w", err)
	}

	// Optional DISTINCT
	if useDistinct {
		tx = tx.Distinct()
	}

	// JOINs
	for _, join := range joins {
		tx = tx.Joins(join)
	}

	// Preloads
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	// GROUP BY
	if len(groupByColumns) > 0 {
		tx = tx.Group(strings.Join(groupByColumns, ", "))
	}

	// ORDER BY
	if len(orderBy) > 0 {
		tx = tx.Order(strings.Join(orderBy, ", "))
	}

	// WHERE
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
	exactConditions map[string]interface{},
	customConditions []struct {
		Query string
		Args  []interface{}
	},
	groupByColumns []string,
	orderBy []string,
) ([]T, error) {
	var results []T
	tx := r.db.WithContext(ctx)

	// Set schema ke schemaName
	if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", schemaName)).Error; err != nil {
		return nil, fmt.Errorf("failed to set schema: %w", err)
	}

	// Hapus tx.Distinct() → tidak digunakan

	// Join relasi jika ada
	for _, join := range joins {
		tx = tx.Joins(join)
	}

	// Preload relasi jika ada
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	// GROUP BY jika ada
	if len(groupByColumns) > 0 {
		tx = tx.Group(strings.Join(groupByColumns, ", "))
	}

	// ORDER BY jika ada
	if len(orderBy) > 0 {
		tx = tx.Order(strings.Join(orderBy, ", "))
	}

	// Kondisi exact
	if len(exactConditions) > 0 {
		tx = tx.Where(exactConditions)
	}

	// Kondisi custom
	for _, cond := range customConditions {
		tx = tx.Where(cond.Query, cond.Args...)
	}

	// Eksekusi query
	if err := tx.Find(&results).Error; err != nil {
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

// func getFieldValue[T any](entity *T, fieldName string) interface{} {
// 	val := reflect.ValueOf(entity).Elem()
// 	field := val.FieldByName(fieldName)
// 	if !field.IsValid() || !field.CanInterface() {
// 		return nil
// 	}
// 	return field.Interface()
// }

type ConflictRow[T any] struct {
	Index  int
	Entity *T
	Reason string
}

// func (r *GenericRepository[T]) SaveManyWithConflictCheck(
// 	ctx context.Context,
// 	schemaName string,
// 	entities []*T,
// 	uniqueFieldName string,
// 	uniqueColumnName string,
// 	batchSize int,
// ) ([]ConflictRow[T], error) {
// 	conflictRows := []ConflictRow[T]{}

// 	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
// 		if err := tx.Exec(fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))).Error; err != nil {
// 			return fmt.Errorf("failed to set schema: %w", err)
// 		}

// 		// Ambil nilai unik dari entity
// 		uniqueValues := make([]interface{}, 0, len(entities))
// 		valueToEntity := make(map[interface{}]*T)
// 		for _, e := range entities {
// 			val := getFieldValue(e, uniqueFieldName)
// 			if val != nil {
// 				uniqueValues = append(uniqueValues, val)
// 				valueToEntity[val] = e
// 			}
// 		}

// 		// Ambil data yang sudah ada di DB
// 		var existing []map[string]interface{}
// 		err := tx.Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), r.tableName)).
// 			Select(uniqueColumnName).
// 			Where(fmt.Sprintf("%s IN ?", uniqueColumnName), uniqueValues).
// 			Find(&existing).Error
// 		if err != nil {
// 			return fmt.Errorf("failed to check existing records: %w", err)
// 		}

// 		// Tandai data yang sudah ada
// 		existingMap := make(map[interface{}]bool)
// 		for _, row := range existing {
// 			if val, ok := row[uniqueColumnName]; ok {
// 				existingMap[val] = true
// 			}
// 		}

// 		// Pisahkan data baru dan konflik
// 		newEntities := []*T{}
// 		for idx, e := range entities {
// 			val := getFieldValue(e, uniqueFieldName)
// 			if existingMap[val] {
// 				conflictRows = append(conflictRows, ConflictRow[T]{
// 					Index:  idx,
// 					Entity: e,
// 					Reason: fmt.Sprintf("Conflict on %s = %v", uniqueFieldName, val),
// 				})
// 			} else {
// 				newEntities = append(newEntities, e)
// 			}
// 		}

// 		// Insert data baru
// 		if len(newEntities) > 0 {
// 			if err := tx.Table(fmt.Sprintf("%s.%s", strings.ToLower(schemaName), r.tableName)).
// 				CreateInBatches(newEntities, batchSize).Error; err != nil {
// 				return fmt.Errorf("failed to insert new records: %w", err)
// 			}
// 		}

// 		return nil
// 	})

// 		return conflictRows, err
// 	}

func getFieldValue[T any](entity *T, fieldName string) interface{} {
	val := reflect.ValueOf(entity).Elem()
	field := val.FieldByName(fieldName)
	if !field.IsValid() {
		return nil
	}
	return field.Interface()
}

// versi 1
// func (r *GenericRepository[T]) SaveManyWithConflictCheck(
// 	ctx context.Context,
// 	schemaName string,
// 	entities []*T,
// 	uniqueFieldName string,
// 	uniqueColumnName string,
// 	batchSize int,
// ) ([]ConflictRow[T], error) {
// 	fmt.Printf("Starting SaveManyWithConflictCheck for %d entities in schema %s\n", len(entities), schemaName)
// 	conflictRows := []ConflictRow[T]{}

// 	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
// 		// Set schema
// 		schemaQuery := fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))
// 		fmt.Printf("Executing schema query: %s\n", schemaQuery)
// 		if err := tx.Exec(schemaQuery).Error; err != nil {
// 			fmt.Printf("Error setting schema: %v\n", err)
// 			return fmt.Errorf("failed to set schema: %w", err)
// 		}

// 		// Get unique values from entities
// 		fmt.Printf("Processing %d entities for unique field '%s'\n", len(entities), uniqueFieldName)
// 		uniqueValues := make([]interface{}, 0, len(entities))
// 		valueToEntity := make(map[interface{}]*T)
// 		for i, e := range entities {
// 			val := getFieldValue(e, uniqueFieldName)
// 			if val != nil {
// 				uniqueValues = append(uniqueValues, val)
// 				valueToEntity[val] = e
// 			} else {
// 				fmt.Printf("Warning: Entity at index %d has nil value for field '%s'\n", i, uniqueFieldName)
// 			}
// 		}
// 		fmt.Printf("Found %d unique values to check\n", len(uniqueValues))

// 		// Check existing records in DB
// 		tableName := fmt.Sprintf("%s.%s", strings.ToLower(schemaName), r.tableName)
// 		query := tx.Table(tableName).Select(uniqueColumnName).Where(fmt.Sprintf("%s IN ?", uniqueColumnName), uniqueValues)
// 		fmt.Printf("Checking existing records with query: %v\n", query.Statement.SQL.String())

// 		var existing []map[string]interface{}
// 		err := query.Find(&existing).Error
// 		if err != nil {
// 			fmt.Printf("Error checking existing records: %v\n", err)
// 			return fmt.Errorf("failed to check existing records: %w", err)
// 		}
// 		fmt.Printf("Found %d existing records that may cause conflicts\n", len(existing))

// 		// Mark existing data
// 		existingMap := make(map[interface{}]bool)
// 		for _, row := range existing {
// 			if val, ok := row[uniqueColumnName]; ok {
// 				existingMap[val] = true
// 			}
// 		}

// 		// Separate new and conflicting entities
// 		newEntities := []*T{}
// 		for idx, e := range entities {
// 			val := getFieldValue(e, uniqueFieldName)
// 			if val == nil {
// 				fmt.Printf("Skipping entity at index %d - nil unique value\n", idx)
// 				continue
// 			}

// 			if existingMap[val] {
// 				conflictReason := fmt.Sprintf("Conflict on %s = %v", uniqueFieldName, val)
// 				fmt.Printf("Conflict detected at index %d: %s\n", idx, conflictReason)
// 				conflictRows = append(conflictRows, ConflictRow[T]{
// 					Index:  idx,
// 					Entity: e,
// 					Reason: conflictReason,
// 				})
// 			} else {
// 				newEntities = append(newEntities, e)
// 			}
// 		}
// 		fmt.Printf("Identified %d new entities and %d conflicts\n", len(newEntities), len(conflictRows))

// 		// Insert new entities
// 		if len(newEntities) > 0 {
// 			fmt.Printf("Inserting %d new entities in batches of %d\n", len(newEntities), batchSize)
// 			if err := tx.Table(tableName).CreateInBatches(newEntities, batchSize).Error; err != nil {
// 				fmt.Printf("Error inserting new records: %v\n", err)
// 				return fmt.Errorf("failed to insert new records: %w", err)
// 			}
// 			fmt.Printf("Successfully inserted %d new entities\n", len(newEntities))
// 		} else {
// 			fmt.Println("No new entities to insert")
// 		}

// 		return nil
// 	})

// 	if err != nil {
// 		fmt.Printf("Transaction failed: %v\n", err)
// 	} else {
// 		fmt.Printf("Transaction completed successfully with %d conflicts\n", len(conflictRows))
// 	}

// 	return conflictRows, err
// }

// versi 2
// func (r *GenericRepository[T]) SaveManyWithConflictCheck(
// 	ctx context.Context,
// 	schemaName string,
// 	entities []*T,
// 	uniqueFieldName string,
// 	uniqueColumnName string,
// 	batchSize int,
// ) ([]ConflictRow[T], error) {
// 	fmt.Printf("Starting SaveManyWithConflictCheck for %d entities in schema %s\n", len(entities), schemaName)
// 	conflictRows := []ConflictRow[T]{}

// 	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
// 		// Set schema
// 		schemaQuery := fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))
// 		fmt.Printf("Executing schema query: %s\n", schemaQuery)
// 		if err := tx.Exec(schemaQuery).Error; err != nil {
// 			fmt.Printf("Error setting schema: %v\n", err)
// 			return fmt.Errorf("failed to set schema: %w", err)
// 		}

// 		// Get unique values from entities
// 		fmt.Printf("Processing %d entities for unique field '%s'\n", len(entities), uniqueFieldName)
// 		uniqueValues := make([]interface{}, 0, len(entities))
// 		valueToEntity := make(map[interface{}]*T)
// 		for i, e := range entities {
// 			val := getFieldValue(e, uniqueFieldName)
// 			if val != nil {
// 				uniqueValues = append(uniqueValues, val)
// 				valueToEntity[val] = e
// 			} else {
// 				fmt.Printf("Warning: Entity at index %d has nil value for field '%s'\n", i, uniqueFieldName)
// 			}
// 		}
// 		fmt.Printf("Found %d unique values to check\n", len(uniqueValues))

// 		// Check existing records in DB
// 		tableName := fmt.Sprintf("%s.%s", strings.ToLower(schemaName), r.tableName)
// 		query := tx.Table(tableName).Select(uniqueColumnName).Where(fmt.Sprintf("%s IN ?", uniqueColumnName), uniqueValues)
// 		fmt.Printf("Checking existing records with query: %v\n", query.Statement.SQL.String())

// 		var existing []map[string]interface{}
// 		err := query.Find(&existing).Error
// 		if err != nil {
// 			fmt.Printf("Error checking existing records: %v\n", err)
// 			return fmt.Errorf("failed to check existing records: %w", err)
// 		}
// 		fmt.Printf("Found %d existing records that may cause conflicts\n", len(existing))

// 		// Mark existing data
// 		existingMap := make(map[interface{}]bool)
// 		for _, row := range existing {
// 			if val, ok := row[uniqueColumnName]; ok {
// 				existingMap[val] = true
// 			}
// 		}

// 		// Separate new and conflicting entities
// 		newEntities := []*T{}
// 		for idx, e := range entities {
// 			val := getFieldValue(e, uniqueFieldName)
// 			if val == nil {
// 				fmt.Printf("Skipping entity at index %d - nil unique value\n", idx)
// 				continue
// 			}

// 			if existingMap[val] {
// 				conflictReason := fmt.Sprintf("Conflict on %s = %v", uniqueFieldName, val)
// 				fmt.Printf("Conflict detected at index %d: %s\n", idx, conflictReason)
// 				conflictRows = append(conflictRows, ConflictRow[T]{
// 					Index:  idx,
// 					Entity: e,
// 					Reason: conflictReason,
// 				})
// 			} else {
// 				newEntities = append(newEntities, e)
// 			}
// 		}
// 		fmt.Printf("Identified %d new entities and %d conflicts\n", len(newEntities), len(conflictRows))

// 		// Insert new entities
// 		if len(newEntities) > 0 {
// 			fmt.Printf("Inserting %d new entities in batches of %d\n", len(newEntities), batchSize)
// 			if err := tx.Table(tableName).CreateInBatches(newEntities, batchSize).Error; err != nil {
// 				if isForeignKeyViolation(err) {
// 					// Handle foreign key violation specifically
// 					fmt.Printf("Foreign key violation detected: %v\n", err)

// 					// Get the IDs of the referenced table that don't exist
// 					missingRefs := findMissingReferences(tx, newEntities)

// 					// Add these to conflict rows
// 					for idx, e := range newEntities {
// 						refVal := getReferenceFieldValue(e, "ptk_id") // Adjust field name as needed
// 						if _, exists := missingRefs[refVal]; exists {
// 							conflictReason := fmt.Sprintf("Referenced PTK with ID %v does not exist", refVal)
// 							fmt.Printf("Reference conflict at index %d: %s\n", idx, conflictReason)
// 							conflictRows = append(conflictRows, ConflictRow[T]{
// 								Index:  idx,
// 								Entity: e,
// 								Reason: conflictReason,
// 							})
// 						}
// 					}

// 					return nil // Return nil to commit the transaction with the conflict info
// 				}
// 				fmt.Printf("Error inserting new records: %v\n", err)
// 				return fmt.Errorf("failed to insert new records: %w", err)
// 			}
// 			fmt.Printf("Successfully inserted %d new entities\n", len(newEntities))
// 		} else {
// 			fmt.Println("No new entities to insert")
// 		}

// 		return nil
// 	})

// 	if err != nil {
// 		fmt.Printf("Transaction failed: %v\n", err)
// 	} else {
// 		fmt.Printf("Transaction completed successfully with %d conflicts\n", len(conflictRows))
// 	}

// 	return conflictRows, err
// }

// ============versi 3===============
func (r *GenericRepository[T]) SaveManyWithConflictCheck(
	ctx context.Context,
	schemaName string,
	entities []*T,
	uniqueFieldName string,
	uniqueColumnName string,
	batchSize int,
	foreignKeyFields []string, // Tambahan parameter FK fleksibel
) ([]ConflictRow[T], error) {
	fmt.Printf("Starting SaveManyWithConflictCheck for %d entities in schema %s\n", len(entities), schemaName)
	conflictRows := []ConflictRow[T]{}

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Set schema
		schemaQuery := fmt.Sprintf("SET search_path TO %s", strings.ToLower(schemaName))
		fmt.Printf("Executing schema query: %s\n", schemaQuery)
		if err := tx.Exec(schemaQuery).Error; err != nil {
			fmt.Printf("Error setting schema: %v\n", err)
			return fmt.Errorf("failed to set schema: %w", err)
		}

		// Ambil semua nilai unik untuk pengecekan duplikat
		fmt.Printf("Processing %d entities for unique field '%s'\n", len(entities), uniqueFieldName)
		uniqueValues := make([]interface{}, 0, len(entities))
		valueToEntity := make(map[interface{}]*T)
		for i, e := range entities {
			val := getFieldValue(e, uniqueFieldName)
			if val != nil {
				uniqueValues = append(uniqueValues, val)
				valueToEntity[val] = e
			} else {
				fmt.Printf("Warning: Entity at index %d has nil value for field '%s'\n", i, uniqueFieldName)
			}
		}
		fmt.Printf("Found %d unique values to check\n", len(uniqueValues))

		// Cek apakah sudah ada di DB
		tableName := fmt.Sprintf("%s.%s", strings.ToLower(schemaName), r.tableName)
		query := tx.Table(tableName).Select(uniqueColumnName).Where(fmt.Sprintf("%s IN ?", uniqueColumnName), uniqueValues)
		fmt.Printf("Checking existing records with query: %v\n", query.Statement.SQL.String())

		var existing []map[string]interface{}
		err := query.Find(&existing).Error
		if err != nil {
			fmt.Printf("Error checking existing records: %v\n", err)
			return fmt.Errorf("failed to check existing records: %w", err)
		}
		fmt.Printf("Found %d existing records that may cause conflicts\n", len(existing))

		// Tandai konflik
		existingMap := make(map[interface{}]bool)
		for _, row := range existing {
			if val, ok := row[uniqueColumnName]; ok {
				existingMap[val] = true
			}
		}

		newEntities := []*T{}
		for idx, e := range entities {
			val := getFieldValue(e, uniqueFieldName)
			if val == nil {
				fmt.Printf("Skipping entity at index %d - nil unique value\n", idx)
				continue
			}
			if existingMap[val] {
				conflictReason := fmt.Sprintf("Conflict on %s = %v", uniqueFieldName, val)
				fmt.Printf("Conflict detected at index %d: %s\n", idx, conflictReason)
				conflictRows = append(conflictRows, ConflictRow[T]{
					Index:  idx,
					Entity: e,
					Reason: conflictReason,
				})
			} else {
				newEntities = append(newEntities, e)
			}
		}
		fmt.Printf("Identified %d new entities and %d conflicts\n", len(newEntities), len(conflictRows))

		// Insert ke DB
		// Insert new entities
		if len(newEntities) > 0 {
			fmt.Printf("Inserting %d new entities in batches of %d\n", len(newEntities), batchSize)
			insertErr := tx.Table(tableName).CreateInBatches(newEntities, batchSize).Error
			if insertErr != nil {
				insertErrMsg := insertErr.Error()

				if strings.Contains(insertErrMsg, "duplicate key value violates unique constraint") {
					// Tangani duplicate key
					fmt.Printf("[DUPLICATE] Duplikat terdeteksi saat insert: %v\n", insertErrMsg)

					for idx, e := range newEntities {
						val := getFieldValue(e, uniqueFieldName)
						conflictReason := fmt.Sprintf("Duplicate on %s = %v", uniqueFieldName, val)
						conflictRows = append(conflictRows, ConflictRow[T]{
							Index:  idx,
							Entity: e,
							Reason: conflictReason,
						})
					}
					// Lanjutkan transaksi (return nil untuk melanjutkan)
					return nil
				}

				// Tangani foreign key violation
				if isForeignKeyViolation(insertErr) {
					fmt.Printf("Foreign key violation detected: %v\n", insertErr)
					missingRefs := findMissingReferences(tx, newEntities)
					for idx, e := range newEntities {
						for _, fkField := range foreignKeyFields {
							refVal := getReferenceFieldValue(e, fkField)
							if _, exists := missingRefs[refVal]; exists {
								conflictReason := fmt.Sprintf("Missing reference for %s: %v", fkField, refVal)
								conflictRows = append(conflictRows, ConflictRow[T]{
									Index:  idx,
									Entity: e,
									Reason: conflictReason,
								})
							}
						}
					}
					return nil
				}

				// Error lain, hentikan transaksi
				fmt.Printf("Error inserting new records: %v\n", insertErr)
				return fmt.Errorf("failed to insert new records: %w", insertErr)
			}
			fmt.Printf("Successfully inserted %d new entities\n", len(newEntities))
		} else {
			fmt.Println("No new entities to insert")
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Transaction failed: %v\n", err)
	} else {
		fmt.Printf("Transaction completed successfully with %d conflicts\n", len(conflictRows))
	}

	return conflictRows, err
}

// Helper function to check for foreign key violation
func isForeignKeyViolation(err error) bool {
	if pgErr, ok := err.(*pgconn.PgError); ok {
		return pgErr.Code == "23503" // SQLSTATE for foreign key violation
	}
	return false
}

// Helper function to find missing references
func findMissingReferences[T any](tx *gorm.DB, entities []*T) map[interface{}]bool {
	// Collect all reference IDs from the entities
	refValues := make(map[interface{}]bool)
	for _, e := range entities {
		val := getReferenceFieldValue(e, "ptk_id") // Adjust field name as needed
		if val != nil {
			refValues[val] = true
		}
	}

	// Check which ones don't exist in the referenced table
	missing := make(map[interface{}]bool)
	if len(refValues) > 0 {
		var existingRefs []map[string]interface{}
		tx.Table("tabel_ptk").Select("id").Where("id IN ?", getKeys(refValues)).Find(&existingRefs)

		existingMap := make(map[interface{}]bool)
		for _, row := range existingRefs {
			if id, ok := row["id"]; ok {
				existingMap[id] = true
			}
		}

		for ref := range refValues {
			if !existingMap[ref] {
				missing[ref] = true
			}
		}
	}

	return missing
}

func getKeys(m map[interface{}]bool) []interface{} {
	keys := make([]interface{}, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func getReferenceFieldValue[T any](entity *T, fieldName string) interface{} {
	val := reflect.ValueOf(entity).Elem()
	field := val.FieldByName(fieldName)
	if !field.IsValid() {
		fmt.Printf("Reference field '%s' not found in entity\n", fieldName)
		return nil
	}
	return field.Interface()
}

// func getFieldValue[T any](entity *T, fieldName string) interface{} {
// 	val := reflect.ValueOf(entity).Elem()
// 	field := val.FieldByName(fieldName)
// 	if !field.IsValid() {
// 		fmt.Printf("Field '%s' not found in entity\n", fieldName)
// 		return nil
// 	}
// 	return field.Interface()
// }

// func ConvertConflictsToProto[T any](conflicts []ConflictRow[T], uniqueField string, nameField string) []*pb.ConflictRow {
// 	result := make([]*pb.ConflictRow, 0, len(conflicts))
// 	for _, c := range conflicts {
// 		id := fmt.Sprintf("%v", getFieldValue(c.Entity, uniqueField))
// 		name := fmt.Sprintf("%v", getFieldValue(c.Entity, nameField))

// 		result = append(result, &pb.ConflictRow{
// 			Index:  int32(c.Index),
// 			Id:     id,
// 			Nama:   name,
// 			Reason: c.Reason,
// 		})
// 	}
// 	return result
// }
