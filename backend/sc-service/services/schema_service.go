package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sc-service/models"
	"sc-service/repositories"
)

type SchemaService interface {
	RegistrasiSekolah(ctx context.Context, schemaName string) error
	SimpanSchemaSekolah(SekolahTenant *models.SekolahTenant) error
	GetSchemaBySekolahID(int) (*models.SekolahTenant, error)
	GetSchemaByName(string) (*models.SekolahTenant, error)
	GetOrCreateSchema(ctx context.Context, TenantSekolah *TenantSekolah) (*models.SekolahTenant, string, error)
}
type TenantSekolah struct {
	SekolahTenantId uint32
	UserId          int32
	Password        string
	NamaSekolah     string
	SekolahIdEnkrip string
	Schemaname      string
}
type schemaServiceImpl struct {
	schemaRepo    repositories.SchemaRepository
	SekolahTenant repositories.SekolahTenantRepository
}

func NewSchemaService(s repositories.SchemaRepository, sTT repositories.SekolahTenantRepository) SchemaService {
	return &schemaServiceImpl{
		schemaRepo:    s,
		SekolahTenant: sTT,
	}
}

var ErrSchemaNotFound = errors.New("schema tidak ditemukan")
var ErrSchemaFound = errors.New("schema ditemukan")

func (s *schemaServiceImpl) RegistrasiSekolah(ctx context.Context, schemaName string) error {
	// 1 Ambil path direktori kerja
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("gagal mendapatkan direktori kerja: %w", err)
	}
	// 2 Tentukan lokasi file SQL untuk schema
	schemaFile := filepath.Join(wd, "migrations", "schema_template.sql")
	if err := s.schemaRepo.InitializeDatabase(ctx, schemaFile, schemaName); err != nil {
		return fmt.Errorf("gagal membuat schema %s: %w", schemaName, err)
	}

	return nil
}

func (s *schemaServiceImpl) SimpanSchemaSekolah(SekolahTenant *models.SekolahTenant) error {
	// Call the schema repository to save the schema
	schema := s.SekolahTenant.Save(SekolahTenant)
	if schema != nil {
		return errors.New("gagal menyimpan tabel sekolah tenant")
	}
	return schema
}

func (s *schemaServiceImpl) GetSchemaBySekolahID(sekolahId int) (*models.SekolahTenant, error) {
	// Call the schema repository to get the schema by sekolah ID
	sc, err := s.SekolahTenant.FindByID(sekolahId)
	if err != nil {
		return nil, errors.New("failed to retrieve tenant table: " + err.Error())
	}
	if sc == nil {
		return nil, errors.New("school tenant not found")
	}
	return sc, nil
}
func (s *schemaServiceImpl) GetSchemaByName(schemaname string) (*models.SekolahTenant, error) {
	// Call the schema repository to get the schema by sekolah ID
	sc, err := s.SekolahTenant.FindByName(schemaname)
	if err != nil {
		return nil, ErrSchemaNotFound
	}
	if sc == nil {
		return nil, ErrSchemaFound
	}
	return sc, nil
}
func (s *schemaServiceImpl) GetOrCreateSchema(ctx context.Context, TenantSekolah *TenantSekolah) (*models.SekolahTenant, string, error) {
	var schema *models.SekolahTenant
	var err error

	schema, err = s.GetSchemaByName(TenantSekolah.Schemaname)
	// Jika error selain "schema tidak ditemukan", return error
	if err != nil && !errors.Is(err, ErrSchemaNotFound) {
		return nil, "", fmt.Errorf("gagal mengambil schema: %w", err)
	}
	// Jika schema sudah ada, kembalikan error
	if schema != nil {
		return schema, schema.SchemaName, ErrSchemaFound
		// return schema, schema.SchemaName, fmt.Errorf("schema '%s' sudah ada", TenantSekolah.Schemaname)
	}

	// ==================================CEK ULANG DARI DATA admin sekolah============
	schema, err = s.GetSchemaByName(TenantSekolah.SekolahIdEnkrip)
	// Jika error selain "schema tidak ditemukan", return error
	if err != nil && !errors.Is(err, ErrSchemaNotFound) {
		return nil, "", fmt.Errorf("gagal mengambil schema: %w", err)
	}
	// Jika schema sudah ada, kembalikan error
	if schema != nil {
		// return schema, schema.SchemaName, fmt.Errorf("schema '%s' sudah ada", TenantSekolah.Schemaname)
		return schema, schema.SchemaName, ErrSchemaFound
	}

	// Registrasi schema baru
	if err := s.RegistrasiSekolah(ctx, TenantSekolah.SekolahIdEnkrip); err != nil {
		return nil, "", fmt.Errorf("gagal registrasi schema '%s': %w", TenantSekolah.Schemaname, err)
	}
	// Simpan informasi schema sekolah
	err = s.SimpanSchemaSekolah(&models.SekolahTenant{
		NamaSekolah:     TenantSekolah.NamaSekolah,
		UserId:          TenantSekolah.UserId,
		SekolahTenantId: TenantSekolah.SekolahTenantId,
		SchemaName:      TenantSekolah.SekolahIdEnkrip,
	})
	if err != nil {
		return nil, "", err
	}
	// Berhasil, return schema atau nilai default jika nil
	return schema, func() string {
		if schema != nil {
			return schema.SchemaName
		}
		return ""
	}(), nil
}
