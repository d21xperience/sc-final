package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sekolah/config"
	pb "sekolah/generated"
	"sekolah/models"
	"sekolah/repositories"
	"sekolah/utils"

	"gorm.io/gorm"
)

type SekolahService struct {
	pb.UnimplementedSekolahServiceServer
	// RedisClient    *redis.Client // Tambahkan Redis sebagai field
	sekolahService repositories.SekolahRepository
	schemaService  SchemaService
}

func NewSekolahService() *SekolahService {
	sekolahRepo := repositories.NewSekolahRepository(config.DB)
	schemaRepo := repositories.NewSchemaRepository(config.DB)
	sekolahTabelTenant := repositories.NewsekolahTenantRepository(config.DB)
	schemaService := NewSchemaService(schemaRepo, sekolahTabelTenant)
	return &SekolahService{
		sekolahService: sekolahRepo,
		schemaService:  schemaService,
	}
}

func (s *SekolahService) RegistrasiSekolah(ctx context.Context, req *pb.TabelSekolahRequest) (*pb.TabelSekolahResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	requiredFields := []string{"Sekolah"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}

	sekolah := req.GetSekolah()
	schemaName := sekolah.SekolahIdEnkrip
	existingSchema, err := s.schemaService.GetSchemaBySekolahID(int(sekolah.SekolahId))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Lanjutkan pendaftaran: %v", err)
	}

	if existingSchema != nil {
		return &pb.TabelSekolahResponse{
			Message: "Pendaftaran dibatalkan: Sekolah sudah terdaftar",
			Status:  false,
		}, nil
	}

	cek := s.schemaService.RegistrasiSekolah(ctx, schemaName)
	if cek != nil {
		// Tangani error spesifik jika error adalah gorm.ErrInvalidData
		if errors.Is(cek, gorm.ErrInvalidData) {
			return nil, errors.New("registrasi sekolah gagal: data tidak valid")
		}
		// Tangani error lainnya
		return nil, fmt.Errorf("registrasi sekolah gagal: %w", cek)
	}
	// 2 Simpan informasi schema sekolah
	err = s.schemaService.SimpanSchemaSekolah(&models.SekolahTabelTenant{
		SekolahID:   int(sekolah.SekolahId),
		SchemaName:  schemaName,
		NamaSekolah: sekolah.NamaSekolah,
	})
	if err != nil {
		log.Printf("Gagal menyimpan schema sekolah: %v", err)
		return nil, errors.New("gagal menyimpan informasi sekolah")
	}

	// 3 Kirim respon sukses
	return &pb.TabelSekolahResponse{
		Message: "Pembuatan database berhasil",
		Status:  true,
	}, nil
}

func (s *SekolahService) GetSekolahTabelTenant(ctx context.Context, req *pb.SekolahTabelTenantRequest) (*pb.SekolahTabelTenantResponse, error) {
	sekolahID := req.GetSekolahId()
	sekolahTerdaftar, err := s.schemaService.GetSchemaBySekolahID(int(sekolahID))
	if err != nil {
		return nil, err
	}

	return &pb.SekolahTabelTenantResponse{
		SekolahId:   int32(sekolahTerdaftar.SekolahID),
		NamaSekolah: sekolahTerdaftar.NamaSekolah,
		SchemaName:  sekolahTerdaftar.SchemaName, // nama schema
	}, err

}

// SCHEMA TABLE SEKOLAH---------------------input informasi sekolah yang telah terdaftar
// ================================================================================//
func (s *SekolahService) CreateSekolah(ctx context.Context, req *pb.CreateSekolahRequest) (*pb.CreateSekolahResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	requiredFields := []string{"SchemaName", "Sekolah"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaName()
	sekolah := req.GetSekolah()
	// sekolahID, _ := uuid.Parse(sekolah.SekolahId)
	sekolahModel := &models.Sekolah{
		SekolahID:           sekolah.SekolahId,
		Nama:                sekolah.Nama,
		Npsn:                sekolah.Npsn,
		Alamat:              sekolah.Alamat,
		KdPos:               sekolah.KdPos,
		Telepon:             sekolah.Telepon,
		Fax:                 sekolah.Fax,
		Kelurahan:           sekolah.Kelurahan,
		Kecamatan:           sekolah.Kecamatan,
		KabKota:             sekolah.KabKota,
		Propinsi:            sekolah.Propinsi,
		Website:             sekolah.Website,
		Email:               sekolah.Email,
		NmKepsek:            sekolah.NmKepsek,
		NipKepsek:           sekolah.NipKepsek,
		NiyKepsek:           sekolah.NipKepsek,
		StatusKepemilikanId: sekolah.StatusKepemilikanId,
		KodeAktivasi:        sekolah.KodeAktivasi,
		Jenjang:             sekolah.Jenjang,
		BentukPendidikanId:  sekolah.BentukPendidikanId,
	}

	sekolahTerdaftar := s.sekolahService.Save(ctx, sekolahModel, schemaName)
	if sekolahTerdaftar != nil {
		log.Printf("Gagal menyimpan sekolah: %v", sekolahTerdaftar.Error())
		return nil, errors.New("gagal menyimpan informasi sekolah")
	}

	return &pb.CreateSekolahResponse{
		Message: "sekolah berhasil ditambahkan",
		Status:  true,
	}, nil

}

func (s *SekolahService) GetSekolah(ctx context.Context, req *pb.GetSekolahRequest) (*pb.GetSekolahResponse, error) {
	//  Ambil schema dari request
	schemaName := req.GetSchemaName()
	sekolahID := req.GetSekolahId()

	//  Cari sekolah berdasarkan ID dan schema
	sekolah, err := s.sekolahService.FindByID(ctx, sekolahID, schemaName)
	if err != nil {
		log.Printf("Gagal menemukan sekolah: %v", err)
		return nil, fmt.Errorf("gagal menemukan sekolah: %w", err)
	}

	//  Return response dalam format protobuf
	return &pb.GetSekolahResponse{
		Sekolah: &pb.SekolahDapo{
			SekolahId:           sekolah.SekolahID,
			Nama:                sekolah.Nama,
			Npsn:                sekolah.Npsn,
			Alamat:              sekolah.Alamat,
			KdPos:               sekolah.KdPos,
			Telepon:             sekolah.Telepon,
			Fax:                 sekolah.Fax,
			Kelurahan:           sekolah.Kelurahan,
			Kecamatan:           sekolah.Kecamatan,
			KabKota:             sekolah.KabKota,
			Propinsi:            sekolah.Propinsi,
			Website:             sekolah.Website,
			Email:               sekolah.Email,
			NmKepsek:            sekolah.NmKepsek,
			NipKepsek:           sekolah.NipKepsek,
			NiyKepsek:           sekolah.NipKepsek,
			StatusKepemilikanId: sekolah.StatusKepemilikanId,
			KodeAktivasi:        sekolah.KodeAktivasi,
			Jenjang:             sekolah.Jenjang,
			BentukPendidikanId:  sekolah.BentukPendidikanId,
		},
	}, nil
}

// Tambahkan fitur tambahan DELET, UPDATE , dan LIST digunakan untuk SUPER ADMIN
