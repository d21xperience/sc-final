package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sekolah/models"
	"sekolah/repositories"
)

type SchemaService interface {
	RegistrasiSekolah(ctx context.Context, schemaName string) error
	SimpanSchemaSekolah(sekolahTabelTenant *models.SekolahTabelTenant) error
	GetSchemaBySekolahID(int) (*models.SekolahTabelTenant, error)
}

type schemaServiceImpl struct {
	schemaRepo         repositories.SchemaRepository
	sekolahTabelTenant repositories.SekolahTenantRepository
}

func NewSchemaService(s repositories.SchemaRepository, sTT repositories.SekolahTenantRepository) SchemaService {
	return &schemaServiceImpl{
		schemaRepo:         s,
		sekolahTabelTenant: sTT,
	}
}

func (s *schemaServiceImpl) RegistrasiSekolah(ctx context.Context, schemaName string) error {
	// 1️ Ambil path direktori kerja
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan direktori kerja: %w", err)
	}
	// 2️ Tentukan lokasi file SQL untuk schema
	schemaFile := filepath.Join(wd, "migrations", "schema_template.sql")
	if err := s.schemaRepo.InitializeDatabase(ctx, schemaFile, schemaName); err != nil {
		return fmt.Errorf("gagal membuat schema %s: %w", schemaName, err)
	}
	return nil
}

func (s *schemaServiceImpl) SimpanSchemaSekolah(sekolahTabelTenant *models.SekolahTabelTenant) error {
	// Call the schema repository to save the schema
	schema := s.sekolahTabelTenant.Save(sekolahTabelTenant)
	if schema != nil {
		return errors.New("gagal menyimpan tabel sekolah tenant")
	}
	return schema
}

func (s *schemaServiceImpl) GetSchemaBySekolahID(sekolahId int) (*models.SekolahTabelTenant, error) {
	// Call the schema repository to get the schema by sekolah ID
	sc, err := s.sekolahTabelTenant.FindByID(sekolahId)
	if err != nil {
		return nil, errors.New("failed to retrieve tenant table: " + err.Error())
	}
	if sc == nil {
		return nil, errors.New("school tenant not found")
	}
	return sc, nil
}
