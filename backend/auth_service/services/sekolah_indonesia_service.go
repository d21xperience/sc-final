package services

import (
	"auth_service/models"
	"auth_service/repository"
	"context"
)

type SekolahIndonesiaService interface {
	FindByQuery(ctx context.Context, inColumnName, findWhat string) ([]*models.SekolahIndonesia, error)
	FindByTextPattern(ctx context.Context, inColumnName, findWhat string) ([]*models.SekolahIndonesia, error)
}

type sekolahIndonesiaService struct {
	repo *repository.GenericRepository[models.SekolahIndonesia]
}

func NewSekolahIndonesiaService(repo *repository.GenericRepository[models.SekolahIndonesia]) SekolahIndonesiaService {
	return &sekolahIndonesiaService{repo: repo}

}
func (s *sekolahIndonesiaService) FindByQuery(ctx context.Context, inColumnName string, findWhat string) ([]*models.SekolahIndonesia, error) {
	return s.repo.FindByQuery(ctx, inColumnName, findWhat)
}
func (s *sekolahIndonesiaService) FindByTextPattern(ctx context.Context, inColumnName string, findWhat string) ([]*models.SekolahIndonesia, error) {
	return s.repo.FindByTextPattern(ctx, inColumnName, findWhat)
}
