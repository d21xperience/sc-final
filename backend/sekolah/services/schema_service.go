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
	SimpanSchemaSekolah(SekolahTenant *models.SekolahTenant) error
	GetSchemaBySekolahID(int) (*models.SekolahTenant, error)
	GetSchemaBySchemaname(schemaName string) (*models.SekolahTenant, error)
}

type schemaServiceImpl struct {
	schemaRepo         repositories.SchemaRepository
	SekolahTenant repositories.GenericRepository[models.SekolahTenant]
}

func NewSchemaService(s repositories.SchemaRepository, sTT repositories.GenericRepository[models.SekolahTenant]) SchemaService {
	return &schemaServiceImpl{
		schemaRepo:         s,
		SekolahTenant: sTT,
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
	// cty, cancel := context.WithTimeout(context.Background(), 5*time.Minute) // Tambah timeout
	// defer cancel()
	if err := s.schemaRepo.InitializeDatabase(ctx, schemaFile, schemaName); err != nil {
		return fmt.Errorf("gagal membuat schema %s: %w", schemaName, err)
	}
	return nil
}

func (s *schemaServiceImpl) SimpanSchemaSekolah(SekolahTenant *models.SekolahTenant) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Call the schema repository to save the schema
	schema := s.SekolahTenant.Save(ctx, SekolahTenant, "public")
	if schema != nil {
		return errors.New("gagal menyimpan tabel sekolah tenant")
	}
	return schema
}

func (s *schemaServiceImpl) GetSchemaBySchemaname(schemaName string) (*models.SekolahTenant, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conditions := map[string]interface{}{
		"schema_name": schemaName,
	}
	res, err := s.SekolahTenant.FindAllByConditions(ctx, "public", conditions, 100, 0, nil)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, err
	}
	return res[0], nil
}
func (s *schemaServiceImpl) GetSchemaBySekolahID(sekolahId int) (*models.SekolahTenant, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Call the schema repository to get the schema by sekolah ID
	sc, err := s.SekolahTenant.FindByID(ctx, strconv.Itoa(int(sekolahId)), "public", "sekolah_id")
	if err != nil {
		return nil, errors.New("failed to retrieve tenant table: " + err.Error())
	}
	if sc == nil {
		return nil, errors.New("school tenant not found")
	}
	return sc, nil
}
