package services

import (
	"context"
	"sekolah/models"
	"sekolah/repositories"
)

type RombelAnggotaService interface {
	Save(ctx context.Context, pd *models.RombelAnggota, schemaName string) error
	FindAll(ctx context.Context, schemaName string, limit, offset int) ([]*models.RombelAnggota, error)
	FindByID(ctx context.Context, RombelAnggotaID string, schemaName string) (*models.RombelAnggota, error)
	Update(ctx context.Context, schemaName string, pd *models.RombelAnggota) error
	Delete(ctx context.Context, RombelAnggotaID string, schemaName string) error
	SaveMany(ctx context.Context, schemaName string, banyakKelas []*models.RombelAnggota) error
	FindAllByConditions(ctx context.Context, schemaName string, conditions map[string]interface{}, limit, offset int) ([]*models.RombelAnggota, error)
}

type RombelAnggotaServiceImpl struct {
	repo *repositories.GenericRepository[models.RombelAnggota]
	// Batch upload
}

func NewRombelAnggotaService(s *repositories.GenericRepository[models.RombelAnggota]) RombelAnggotaService {
	return &RombelAnggotaServiceImpl{repo: s}
}

func (s *RombelAnggotaServiceImpl) Save(ctx context.Context, anggotaRombel *models.RombelAnggota, schemaName string) error {
	return s.repo.Save(ctx, anggotaRombel, schemaName)
}
func (s *RombelAnggotaServiceImpl) FindAll(ctx context.Context, schemaName string, limit, offset int) ([]*models.RombelAnggota, error) {
	return s.repo.FindAll(ctx, schemaName, limit, offset)
}
func (s *RombelAnggotaServiceImpl) FindByID(ctx context.Context, anggotaRombelId string, schemaName string) (*models.RombelAnggota, error) {
	return s.repo.FindByID(ctx, anggotaRombelId, schemaName, "rombongan_belajar_id")
}

func (s *RombelAnggotaServiceImpl) Update(ctx context.Context, schemaName string, anggotaRombel *models.RombelAnggota) error {
	return s.repo.Update(ctx, anggotaRombel, schemaName, "rombongan_belajar_id", anggotaRombel.AnggotaRombelId)
}
func (s *RombelAnggotaServiceImpl) Delete(ctx context.Context, RombelAnggotaID string, schemaName string) error {
	return s.repo.Delete(ctx, RombelAnggotaID, schemaName, "rombongan_belajar_id")
}
func (s *RombelAnggotaServiceImpl) SaveMany(ctx context.Context, schemaName string, banyakAnggotaRombel []*models.RombelAnggota) error {
	// Batasi batch size agar operasi lebih efisien (misalnya 100 record per batch)
	batchSize := 100
	return s.repo.SaveMany(ctx, schemaName, banyakAnggotaRombel, batchSize)
}
func (s *RombelAnggotaServiceImpl) FindAllByConditions(ctx context.Context, schemaName string, conditions map[string]interface{}, limit, offset int) ([]*models.RombelAnggota, error) {
	return s.repo.FindAllByConditions(ctx, schemaName, conditions, limit, offset)
}
