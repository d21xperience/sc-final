package services

import (
	"context"
	"errors"
	"sekolah/models"
	"sekolah/repositories"
)

type NilaiAkhirService interface {
	Save(ctx context.Context, pd *models.TabelNilaiAkhir, schemaName string) error
	FindByID(ctx context.Context, NilaiAkhirID string, schemaName string) (*models.TabelNilaiAkhir, error)
	Update(ctx context.Context, NilaiAkhir *models.TabelNilaiAkhir, schemaName string) error
	Delete(ctx context.Context, NilaiAkhirID string, schemaName string) error
}

type NilaiAkhirServiceImpl struct {
	NilaiAkhirRepo repositories.NilaiAkhirRepository
}

func NewNilaiAkhirService(sr repositories.NilaiAkhirRepository) NilaiAkhirService {
	return &NilaiAkhirServiceImpl{NilaiAkhirRepo: sr}
}

func (s *NilaiAkhirServiceImpl) Save(ctx context.Context, nilaiAkhirModel *models.TabelNilaiAkhir, schemaName string) error {
	err := s.NilaiAkhirRepo.Save(ctx, nilaiAkhirModel, schemaName)
	if err != nil {
		return errors.New("gagal menyimpan NilaiAkhir")
	}
	return err
}
func (s *NilaiAkhirServiceImpl) FindByID(ctx context.Context, NilaiAkhirID string, schemaName string) (*models.TabelNilaiAkhir, error) {
	return s.NilaiAkhirRepo.FindByID(ctx, NilaiAkhirID, schemaName)
}
func (s *NilaiAkhirServiceImpl) Update(ctx context.Context, nilaiAkhirModel *models.TabelNilaiAkhir, schemaName string) error {
	return s.NilaiAkhirRepo.Update(ctx, nilaiAkhirModel, schemaName)
}
func (s *NilaiAkhirServiceImpl) Delete(ctx context.Context, NilaiAkhirID string, schemaName string) error {
	return s.NilaiAkhirRepo.Delete(ctx, NilaiAkhirID, schemaName)
}
