package services

import (
	"context"
	"errors"
	"sekolah/models"
	"sekolah/repositories"
)

type PesertaDidikService interface {
	Save(ctx context.Context, pd *models.PesertaDidik, schemaName string) error
	FindByID(ctx context.Context, pesertaDidikID string, schemaName string) (*models.PesertaDidik, error)
	Update(ctx context.Context, pesertaDidik *models.PesertaDidik, schemaName string) error
	Delete(ctx context.Context, pesertaDidikID string, schemaName string) error
}

type pesertaDidikServiceImpl struct {
	pesertaDidikRepo repositories.PesertaDidikRepository
}

func NewPesertaDidikService(sr repositories.PesertaDidikRepository) PesertaDidikService {
	return &pesertaDidikServiceImpl{pesertaDidikRepo: sr}
}

func (s *pesertaDidikServiceImpl) Save(ctx context.Context, pd *models.PesertaDidik, schemaName string) error {
	err := s.pesertaDidikRepo.Save(ctx, pd, schemaName)
	if err != nil {
		return errors.New("gagal menyimpan siswa")
	}
	return err
}
func (s *pesertaDidikServiceImpl) FindByID(ctx context.Context, pesertaDidikID string, schemaName string) (*models.PesertaDidik, error) {
	return s.pesertaDidikRepo.FindByID(ctx, pesertaDidikID, schemaName)
}
func (s *pesertaDidikServiceImpl) Update(ctx context.Context, pesertaDidik *models.PesertaDidik, schemaName string) error {
	return s.pesertaDidikRepo.Update(ctx, pesertaDidik, schemaName)
}
func (s *pesertaDidikServiceImpl) Delete(ctx context.Context, pesertaDidikID string, schemaName string) error {
	return s.pesertaDidikRepo.Delete(ctx, pesertaDidikID, schemaName)
}
