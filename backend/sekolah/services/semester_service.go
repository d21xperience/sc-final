package services

import (
	"context"
	"sekolah/models"
	"sekolah/repositories"
)

type SemesterService interface {
	Save(ctx context.Context, pd *models.Semester, schemaName string) error
	FindAll(ctx context.Context, schemaName string, limit, offset int) ([]*models.Semester, error)
	FindByID(ctx context.Context, SemesterID string, schemaName, columnName string) (*models.Semester, error)
	Update(ctx context.Context, Semester *models.Semester, schemaName, columnName, IdColumn string) error
	Delete(ctx context.Context, SemesterID string, schemaName string, columnName string) error
	SaveMany(ctx context.Context, schemaName string, banyakKelas []*models.Semester) error
	FindAllByConditions(ctx context.Context, schemaName string, conditions map[string]interface{}, limit, offset int) ([]*models.Semester, error)
}

type semesterServiceImpl struct {
	repo *repositories.GenericRepository[models.Semester]
}

func NewSemesterService(sr *repositories.GenericRepository[models.Semester]) SemesterService {
	return &semesterServiceImpl{repo: sr}
}

func (s *semesterServiceImpl) Save(ctx context.Context, SemesterModel *models.Semester, schemaName string) error {
	err := s.repo.Save(ctx, SemesterModel, schemaName)
	if err != nil {
		return err
	}
	return err
}
func (s *semesterServiceImpl) FindByID(ctx context.Context, SemesterID string, schemaName, columnName string) (*models.Semester, error) {
	return s.repo.FindByID(ctx, SemesterID, schemaName, columnName)
}
func (s *semesterServiceImpl) Update(ctx context.Context, Semester *models.Semester, schemaName, columnName, IdColumn string) error {
	return s.repo.Update(ctx, Semester, schemaName, columnName, IdColumn)
}
func (s *semesterServiceImpl) Delete(ctx context.Context, SemesterID, schemaName, columnName string) error {
	return s.repo.Delete(ctx, SemesterID, schemaName, columnName)
}
func (s *semesterServiceImpl) FindAll(ctx context.Context, schemaName string, limit, offset int) ([]*models.Semester, error) {
	results, err := s.repo.FindAll(ctx, schemaName, limit, offset)
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (s *semesterServiceImpl) FindAllByConditions(ctx context.Context, schemaName string, conditions map[string]interface{}, limit, offset int) ([]*models.Semester, error) {
	results, err := s.repo.FindAllByConditions(ctx, schemaName, conditions, limit, offset)
	if err != nil {
		return nil, err
	}
	return results, nil
}
func (s *semesterServiceImpl) SaveMany(ctx context.Context, schemaName string, banyakNilaiAkhir []*models.Semester) error {
	// Batasi batch size agar operasi lebih efisien (misalnya 100 record per batch)
	batchSize := 100
	return s.repo.SaveMany(ctx, schemaName, banyakNilaiAkhir, batchSize)
}
