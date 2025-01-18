package services

import (
	"context"
	"errors"
	"sekolah/models"
	"sekolah/repositories"
)

type IjazahService interface {
	Save(ctx context.Context, pd *models.Ijazah, schemaName string) error
	FindByID(ctx context.Context, IjazahID string, schemaName string) (*models.Ijazah, error)
	Update(ctx context.Context, Ijazah *models.Ijazah, schemaName string) error
	Delete(ctx context.Context, IjazahID string, schemaName string) error
}

type IjazahServiceImpl struct {
	IjazahRepo repositories.IjazahRepository
}

func NewIjazahService(sr repositories.IjazahRepository) IjazahService {
	return &IjazahServiceImpl{IjazahRepo: sr}
}

func (s *IjazahServiceImpl) Save(ctx context.Context, IjazahModel *models.Ijazah, schemaName string) error {
	err := s.IjazahRepo.Save(ctx, IjazahModel, schemaName)
	if err != nil {
		return errors.New("gagal menyimpan Ijazah")
	}
	return err
}
func (s *IjazahServiceImpl) FindByID(ctx context.Context, IjazahID string, schemaName string) (*models.Ijazah, error) {
	return s.IjazahRepo.FindByID(ctx, IjazahID, schemaName)
}
func (s *IjazahServiceImpl) Update(ctx context.Context, Ijazah *models.Ijazah, schemaName string) error {
	return s.IjazahRepo.Update(ctx, Ijazah, schemaName)
}
func (s *IjazahServiceImpl) Delete(ctx context.Context, IjazahID string, schemaName string) error {
	return s.IjazahRepo.Delete(ctx, IjazahID, schemaName)
}
