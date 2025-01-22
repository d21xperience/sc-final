package services

import (
	"context"
	"sekolah/models"
	"sekolah/repositories"
)

type IjazahService interface {
	Save(ctx context.Context, pd *models.Ijazah, schemaName string) error
	FindAll(ctx context.Context, schemaName string, limit, offset int) ([]*models.Ijazah, error)
	FindByID(ctx context.Context, IjazahID string, schemaName string) (*models.Ijazah, error)
	// Update(ctx context.Context, schemaName string, pd *models.Ijazah) error
	Delete(ctx context.Context, IjazahID string, schemaName string) error
	SaveMany(ctx context.Context, schemaName string, banyakKelas []*models.Ijazah) error
	FindAllByConditions(ctx context.Context, schemaName string, conditions map[string]interface{}, limit, offset int) ([]*models.Ijazah, error)
}

type IjazahServiceImpl struct {
	repo *repositories.GenericRepository[models.Ijazah]
	// Batch upload
}

func NewIjazahService(s *repositories.GenericRepository[models.Ijazah]) IjazahService {
	return &IjazahServiceImpl{repo: s}
}

func (s *IjazahServiceImpl) Save(ctx context.Context, ijazah *models.Ijazah, schemaName string) error {
	return s.repo.Save(ctx, ijazah, schemaName)
}
func (s *IjazahServiceImpl) FindAll(ctx context.Context, schemaName string, limit, offset int) ([]*models.Ijazah, error) {
	return s.repo.FindAll(ctx, schemaName, limit, offset)
}
func (s *IjazahServiceImpl) FindByID(ctx context.Context, ijazahId string, schemaName string) (*models.Ijazah, error) {
	return s.repo.FindByID(ctx, ijazahId, schemaName, "rombongan_belajar_id")
}

//	func (s *IjazahServiceImpl) Update(ctx context.Context, schemaName string, ijazah *models.Ijazah) error {
//		return s.repo.Update(ctx, ijazah, schemaName, "rombongan_belajar_id", ijazah.ID)
//	}
func (s *IjazahServiceImpl) Delete(ctx context.Context, IjazahID string, schemaName string) error {
	return s.repo.Delete(ctx, IjazahID, schemaName, "rombongan_belajar_id")
}
func (s *IjazahServiceImpl) SaveMany(ctx context.Context, schemaName string, banyakijazah []*models.Ijazah) error {
	// Batasi batch size agar operasi lebih efisien (misalnya 100 record per batch)
	batchSize := 100
	return s.repo.SaveMany(ctx, schemaName, banyakijazah, batchSize)
}
func (s *IjazahServiceImpl) FindAllByConditions(ctx context.Context, schemaName string, conditions map[string]interface{}, limit, offset int) ([]*models.Ijazah, error) {
	return s.repo.FindAllByConditions(ctx, schemaName, conditions, limit, offset)
}
