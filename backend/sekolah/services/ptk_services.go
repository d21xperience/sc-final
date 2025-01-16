package services

import (
	"context"
	"errors"
	"sekolah/models"
	"sekolah/repositories"
)

type PTKService interface {
	Save(ctx context.Context, pd *models.TabelPTK, schemaName string) error
	FindByID(ctx context.Context, PTKID string, schemaName string) (*models.TabelPTK, error)
	Update(ctx context.Context, PTK *models.TabelPTK, schemaName string) error
	Delete(ctx context.Context, PTKID string, schemaName string) error
}

type ptkServiceImpl struct {
	PTKRepo repositories.PTKRepository
}

func NewPTKService(sr repositories.PTKRepository) PTKService {
	return &ptkServiceImpl{PTKRepo: sr}
}

func (s *ptkServiceImpl) Save(ctx context.Context, PTKModel *models.TabelPTK, schemaName string) error {
	err := s.PTKRepo.Save(ctx, PTKModel, schemaName)
	if err != nil {
		return errors.New("gagal menyimpan Guru")
	}
	return err
}
func (s *ptkServiceImpl) FindByID(ctx context.Context, PTKID string, schemaName string) (*models.TabelPTK, error) {
	return s.PTKRepo.FindByID(ctx, PTKID, schemaName)
}
func (s *ptkServiceImpl) Update(ctx context.Context, PTK *models.TabelPTK, schemaName string) error {
	return s.PTKRepo.Update(ctx, PTK, schemaName)
}
func (s *ptkServiceImpl) Delete(ctx context.Context, PTKID string, schemaName string) error {
	return s.PTKRepo.Delete(ctx, PTKID, schemaName)
}
