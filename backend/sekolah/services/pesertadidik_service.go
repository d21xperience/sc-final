package services

import (
	"context"
	"sekolah/models"
	"sekolah/repositories"
)

type PesertaDidikService interface {
	Save(ctx context.Context, pd *models.PesertaDidik, schemaName string) error
	FindAll(ctx context.Context, schemaName string, limit, offset int) ([]*models.PesertaDidik, error)
	FindByID(ctx context.Context, pesertaDidikID string, schemaName string) (*models.PesertaDidik, error)
	Update(ctx context.Context, schemaName string, pd *models.PesertaDidik) error
	Delete(ctx context.Context, pesertaDidikID string, schemaName string) error
	SaveMany(ctx context.Context, schemaName string, banyakSiswa []*models.PesertaDidik) error
}

type pesertaDidikServiceImpl struct {
	repo *repositories.GenericRepository[models.PesertaDidik]
	// Batch upload
}

func NewPesertaDidikService(s *repositories.GenericRepository[models.PesertaDidik]) PesertaDidikService {
	return &pesertaDidikServiceImpl{repo: s}
}

func (s *pesertaDidikServiceImpl) Save(ctx context.Context, pd *models.PesertaDidik, schemaName string) error {
	return s.repo.Save(ctx, pd, schemaName)
}
func (s *pesertaDidikServiceImpl) FindAll(ctx context.Context, schemaName string, limit, offset int) ([]*models.PesertaDidik, error) {
	return s.repo.FindAll(ctx, schemaName, limit, offset)
}
func (s *pesertaDidikServiceImpl) FindByID(ctx context.Context, pesertaDidikID string, schemaName string) (*models.PesertaDidik, error) {
	return s.repo.FindByID(ctx, pesertaDidikID, schemaName, "peserta_didik")
}

func (s *pesertaDidikServiceImpl) Update(ctx context.Context, schemaName string, pd *models.PesertaDidik) error {
	return s.repo.Update(ctx, pd, schemaName, "peserta_didik_id", pd.PesertaDidikID)
}
func (s *pesertaDidikServiceImpl) Delete(ctx context.Context, pesertaDidikID string, schemaName string) error {
	return s.repo.Delete(ctx, pesertaDidikID, schemaName, "peserta_didik_id")
}
func (s *pesertaDidikServiceImpl) SaveMany(ctx context.Context, schemaName string, banyakSiswa []*models.PesertaDidik) error {
	// Batasi batch size agar operasi lebih efisien (misalnya 100 record per batch)
	batchSize := 100
	return s.repo.SaveMany(ctx, schemaName, banyakSiswa, batchSize)
}
