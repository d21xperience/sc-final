package repository

import (
	"context"
	"fmt"
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
func (r *GenericRepository[T]) FindByQuery(ctx context.Context, inColumnName, findWhat string) ([]*T, error) {
	var entities []*T
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := r.db.WithContext(ctx).
		Table(r.tableName).
		Find(&entities, fmt.Sprintf("%s = ?", inColumnName), findWhat).Error; err != nil {
		return nil, fmt.Errorf("failed to find record in schema %s", err)
	}
	return entities, nil
}

func (r *GenericRepository[T]) FindByTextPattern(ctx context.Context, inColumnName, findWhat string) ([]*T, error) {
	var entities []*T
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	textPattern := fmt.Sprintf("%%%s%%", findWhat)
	textQuery := fmt.Sprintf("%s LIKE ?", inColumnName)
	if err := r.db.WithContext(ctx).
		Table(r.tableName).
		Where(textQuery, textPattern).
		Find(&entities).Error; err != nil {
		return nil, fmt.Errorf("failed to find records in schema %s", err)
	}

	return entities, nil
}
