package services

import (
	"context"
	"errors"
	"sekolah/models"
	"sekolah/repositories"
)

type SekolahService interface {
	Save(ctx context.Context, pd *models.Sekolah, schemaName string) error
	Find(ctx context.Context, schemaName string) (*models.Sekolah, error)
	FindByID(ctx context.Context, SekolahID string, schemaName string) (*models.Sekolah, error)
	Update(ctx context.Context, Sekolah *models.Sekolah, schemaName string) error
	Delete(ctx context.Context, SekolahID string, schemaName string) error
}

type SekolahServiceImpl struct {
	SekolahRepo repositories.SekolahRepository
}

func NewSekolahService(sr repositories.SekolahRepository) SekolahService {
	return &SekolahServiceImpl{SekolahRepo: sr}
}

func (s *SekolahServiceImpl) Save(ctx context.Context, sekolah *models.Sekolah, schemaName string) error {
	err := s.SekolahRepo.Save(ctx, sekolah, schemaName)
	if err != nil {
		return errors.New("gagal menyimpan Sekolah")
	}
	return err
}
func (s *SekolahServiceImpl) Find(ctx context.Context, schemaName string) (*models.Sekolah, error) {
	return s.SekolahRepo.Find(ctx, schemaName)
}
func (s *SekolahServiceImpl) FindByID(ctx context.Context, SekolahID string, schemaName string) (*models.Sekolah, error) {
	return s.SekolahRepo.FindByID(ctx, SekolahID, schemaName)
}
func (s *SekolahServiceImpl) Update(ctx context.Context, Sekolah *models.Sekolah, schemaName string) error {
	return s.SekolahRepo.Update(ctx, Sekolah, schemaName)
}
func (s *SekolahServiceImpl) Delete(ctx context.Context, SekolahID string, schemaName string) error {
	return s.SekolahRepo.Delete(ctx, SekolahID, schemaName)
}
