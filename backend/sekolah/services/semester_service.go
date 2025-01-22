package services

import (
	"context"
	"sekolah/models"
	"sekolah/repositories"
)

type SemesterService interface {
	Save(ctx context.Context, pd *models.Semester, schemaName string) error
	FindByID(ctx context.Context, SemesterID string, schemaName string) (*models.Semester, error)
	FindAll(ctx context.Context, schemaName string, limit, offset int) ([]*models.Semester, error)
	Update(ctx context.Context, Semester *models.Semester, schemaName string) error
	Delete(ctx context.Context, SemesterID string, schemaName string) error
}

type SemesterServiceImpl struct {
	SemesterRepo repositories.SemesterRepository
}

func NewSemesterService(sr repositories.SemesterRepository) SemesterService {
	return &SemesterServiceImpl{SemesterRepo: sr}
}

func (s *SemesterServiceImpl) Save(ctx context.Context, SemesterModel *models.Semester, schemaName string) error {
	err := s.SemesterRepo.Save(ctx, SemesterModel, schemaName)
	if err != nil {
		return err
	}
	return err
}
func (s *SemesterServiceImpl) FindByID(ctx context.Context, SemesterID string, schemaName string) (*models.Semester, error) {
	return s.SemesterRepo.FindByID(ctx, SemesterID, schemaName)
}
func (s *SemesterServiceImpl) Update(ctx context.Context, Semester *models.Semester, schemaName string) error {
	return s.SemesterRepo.Update(ctx, Semester, schemaName)
}
func (s *SemesterServiceImpl) Delete(ctx context.Context, SemesterID string, schemaName string) error {
	return s.SemesterRepo.Delete(ctx, SemesterID, schemaName)
}
func (s *SemesterServiceImpl) FindAll(ctx context.Context, schemaName string, limit, offset int) ([]*models.Semester, error) {
	return s.SemesterRepo.FindAll(ctx, schemaName, limit, offset)
}
