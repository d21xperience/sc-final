package repositories

import (
	"context"
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
	conditions map[string]interface{}, // Parameter untuk kondisi WHERE
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
func (r *GenericRepository[T]) FindWithPreloadAndJoins(ctx context.Context, schemaName string, joins []string, preloads []string, conditions map[string]interface{}) ([]T, error) {
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

	// Eksekusi Query dengan kondisi
	if err := tx.Where(conditions).Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}
