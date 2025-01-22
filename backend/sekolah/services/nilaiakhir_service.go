package services

import (
	"context"
	"sekolah/models"
	"sekolah/repositories"
)

type NilaiAkhirService interface {
	Save(ctx context.Context, pd *models.NilaiAkhir, schemaName string) error
	FindAll(ctx context.Context, schemaName string, limit, offset int) ([]*models.NilaiAkhir, error)
	FindByID(ctx context.Context, NilaiAkhirID string, schemaName string) (*models.NilaiAkhir, error)
	// Update(ctx context.Context, schemaName string, pd *models.NilaiAkhir) error
	Delete(ctx context.Context, NilaiAkhirID string, schemaName string) error
	SaveMany(ctx context.Context, schemaName string, banyakKelas []*models.NilaiAkhir) error
	FindAllByConditions(ctx context.Context, schemaName string, conditions map[string]interface{}, limit, offset int) ([]*models.NilaiAkhir, error)
}

type NilaiAkhirServiceImpl struct {
	repo *repositories.GenericRepository[models.NilaiAkhir]
	// Batch upload
}

func NewNilaiAkhirService(s *repositories.GenericRepository[models.NilaiAkhir]) NilaiAkhirService {
	return &NilaiAkhirServiceImpl{repo: s}
}

func (s *NilaiAkhirServiceImpl) Save(ctx context.Context, NilaiAkhir *models.NilaiAkhir, schemaName string) error {
	return s.repo.Save(ctx, NilaiAkhir, schemaName)
}
func (s *NilaiAkhirServiceImpl) FindAll(ctx context.Context, schemaName string, limit, offset int) ([]*models.NilaiAkhir, error) {
	return s.repo.FindAll(ctx, schemaName, limit, offset)
}
func (s *NilaiAkhirServiceImpl) FindByID(ctx context.Context, NilaiAkhirId string, schemaName string) (*models.NilaiAkhir, error) {
	return s.repo.FindByID(ctx, NilaiAkhirId, schemaName, "rombongan_belajar_id")
}

//	func (s *NilaiAkhirServiceImpl) Update(ctx context.Context, schemaName string, NilaiAkhir *models.NilaiAkhir) error {
//		return s.repo.Update(ctx, NilaiAkhir, schemaName, "rombongan_belajar_id", NilaiAkhir.ID)
//	}
func (s *NilaiAkhirServiceImpl) Delete(ctx context.Context, NilaiAkhirID string, schemaName string) error {
	return s.repo.Delete(ctx, NilaiAkhirID, schemaName, "rombongan_belajar_id")
}
func (s *NilaiAkhirServiceImpl) SaveMany(ctx context.Context, schemaName string, banyakNilaiAkhir []*models.NilaiAkhir) error {
	// Batasi batch size agar operasi lebih efisien (misalnya 100 record per batch)
	batchSize := 100
	return s.repo.SaveMany(ctx, schemaName, banyakNilaiAkhir, batchSize)
}
func (s *NilaiAkhirServiceImpl) FindAllByConditions(ctx context.Context, schemaName string, conditions map[string]interface{}, limit, offset int) ([]*models.NilaiAkhir, error) {
	return s.repo.FindAllByConditions(ctx, schemaName, conditions, limit, offset)
}
