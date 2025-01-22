package services

import (
	"context"
	"sekolah/models"
	"sekolah/repositories"
)

type RombonganBelajarService interface {
	Save(ctx context.Context, pd *models.RombonganBelajar, schemaName string) error
	FindAll(ctx context.Context, schemaName string, limit, offset int) ([]*models.RombonganBelajar, error)
	FindByID(ctx context.Context, RombonganBelajarID string, schemaName string) (*models.RombonganBelajar, error)
	Update(ctx context.Context, schemaName string, pd *models.RombonganBelajar) error
	Delete(ctx context.Context, RombonganBelajarID string, schemaName string) error
	SaveMany(ctx context.Context, schemaName string, banyakKelas []*models.RombonganBelajar) error
	FindAllByConditions(ctx context.Context, schemaName string, conditions map[string]interface{}, limit, offset int) ([]*models.RombonganBelajar, error)
}

type RombonganBelajarServiceImpl struct {
	repo *repositories.GenericRepository[models.RombonganBelajar]
	// Batch upload
}

func NewRombonganBelajarService(s *repositories.GenericRepository[models.RombonganBelajar]) RombonganBelajarService {
	return &RombonganBelajarServiceImpl{repo: s}
}

func (s *RombonganBelajarServiceImpl) Save(ctx context.Context, pd *models.RombonganBelajar, schemaName string) error {
	return s.repo.Save(ctx, pd, schemaName)
}
func (s *RombonganBelajarServiceImpl) FindAll(ctx context.Context, schemaName string, limit, offset int) ([]*models.RombonganBelajar, error) {
	return s.repo.FindAll(ctx, schemaName, limit, offset)
}
func (s *RombonganBelajarServiceImpl) FindByID(ctx context.Context, RombonganBelajarID string, schemaName string) (*models.RombonganBelajar, error) {
	return s.repo.FindByID(ctx, RombonganBelajarID, schemaName, "rombongan_belajar_id")
}

func (s *RombonganBelajarServiceImpl) Update(ctx context.Context, schemaName string, pd *models.RombonganBelajar) error {
	return s.repo.Update(ctx, pd, schemaName, "rombongan_belajar_id", pd.RombonganBelajarId)
}
func (s *RombonganBelajarServiceImpl) Delete(ctx context.Context, RombonganBelajarID string, schemaName string) error {
	return s.repo.Delete(ctx, RombonganBelajarID, schemaName, "rombongan_belajar_id")
}
func (s *RombonganBelajarServiceImpl) SaveMany(ctx context.Context, schemaName string, banyakKelas []*models.RombonganBelajar) error {
	// Batasi batch size agar operasi lebih efisien (misalnya 100 record per batch)
	batchSize := 100
	return s.repo.SaveMany(ctx, schemaName, banyakKelas, batchSize)
}
func (s *RombonganBelajarServiceImpl) FindAllByConditions(ctx context.Context, schemaName string, conditions map[string]interface{}, limit, offset int) ([]*models.RombonganBelajar, error) {
	return s.repo.FindAllByConditions(ctx, schemaName, conditions, limit, offset)
}
