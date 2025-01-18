package services

import (
	"context"
	"errors"
	"sekolah/models"
	"sekolah/repositories"
)

type TahunAjaranService interface {
	Save(ctx context.Context, pd *models.TahunAjaran, schemaName string) error
	FindByID(ctx context.Context, TahunAjaranID string, schemaName string) (*models.TahunAjaran, error)
	Update(ctx context.Context, TahunAjaran *models.TahunAjaran, schemaName string) error
	Delete(ctx context.Context, TahunAjaranID string, schemaName string) error
}

type TahunAjaranServiceImpl struct {
	TahunAjaranRepo repositories.TahunAjaranRepository
}

func NewTahunAjaranService(sr repositories.TahunAjaranRepository) TahunAjaranService {
	return &TahunAjaranServiceImpl{TahunAjaranRepo: sr}
}

func (s *TahunAjaranServiceImpl) Save(ctx context.Context, TahunAjaranModel *models.TahunAjaran, schemaName string) error {
	err := s.TahunAjaranRepo.Save(ctx, TahunAjaranModel, schemaName)
	if err != nil {
		return errors.New("gagal menyimpan TahunAjaran")
	}
	return err
}
func (s *TahunAjaranServiceImpl) FindByID(ctx context.Context, TahunAjaranID string, schemaName string) (*models.TahunAjaran, error) {
	return s.TahunAjaranRepo.FindByID(ctx, TahunAjaranID, schemaName)
}
func (s *TahunAjaranServiceImpl) Update(ctx context.Context, TahunAjaran *models.TahunAjaran, schemaName string) error {
	return s.TahunAjaranRepo.Update(ctx, TahunAjaran, schemaName)
}
func (s *TahunAjaranServiceImpl) Delete(ctx context.Context, TahunAjaranID string, schemaName string) error {
	return s.TahunAjaranRepo.Delete(ctx, TahunAjaranID, schemaName)
}
