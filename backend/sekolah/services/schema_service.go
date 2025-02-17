package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sekolah/models"
	"sekolah/repositories"
	"strconv"
	"time"
)

type SchemaService interface {
	RegistrasiSekolah(ctx context.Context, schemaName string) error
	SimpanSchemaSekolah(sekolahTabelTenant *models.SekolahTabelTenant) error
	GetSchemaBySekolahID(int) (*models.SekolahTabelTenant, error)
	GetSchemaBySchemaname(schemaName string) (*models.SekolahTabelTenant, error)
}

type schemaServiceImpl struct {
	schemaRepo         repositories.SchemaRepository
	sekolahTabelTenant repositories.GenericRepository[models.SekolahTabelTenant]
}

func NewSchemaService(s repositories.SchemaRepository, sTT repositories.GenericRepository[models.SekolahTabelTenant]) SchemaService {
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Call the schema repository to save the schema
	schema := s.sekolahTabelTenant.Save(ctx, sekolahTabelTenant, "public")
	if schema != nil {
		return errors.New("gagal menyimpan tabel sekolah tenant")
	}
	return schema
}

func (s *schemaServiceImpl) GetSchemaBySchemaname(schemaName string) (*models.SekolahTabelTenant, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conditions := map[string]interface{}{
		"schema_name": schemaName,
	}
	res, err := s.sekolahTabelTenant.FindAllByConditions(ctx, "public", conditions, 100, 0)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, err
	}
	return res[0], nil
}
func (s *schemaServiceImpl) GetSchemaBySekolahID(sekolahId int) (*models.SekolahTabelTenant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Call the schema repository to get the schema by sekolah ID
	sc, err := s.sekolahTabelTenant.FindByID(ctx, strconv.Itoa(int(sekolahId)), "public", "sekolah_id")
	if err != nil {
		return nil, errors.New("failed to retrieve tenant table: " + err.Error())
	}
	if sc == nil {
		return nil, errors.New("school tenant not found")
	}
	return sc, nil
}
