package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sc-service/models"
	"sc-service/utils"

	pb "sc-service/generated"

	"gorm.io/gorm"
)

type TenantServiceServer struct {
	pb.UnimplementedTenantServiceServer
	schemaService SchemaService
	// RedisClient    *redis.Client // Tambahkan Redis sebagai field
}

func (s *TenantServiceServer) RegistrasiSekolahTenant(ctx context.Context, req *pb.RegistrasiSekolahTenantRequest) (*pb.RegistrasiSekolahTenantResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	requiredFields := []string{"Sekolah"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}

	sekolah := req.GetSekolahTenant()
	schemaName := sekolah.Schemaname
	existingSchema, err := s.schemaService.GetSchemaBySekolahID(int(sekolah.SekolahId))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Lanjutkan pendaftaran: %v", err)
	}

	if existingSchema != nil {
		return &pb.RegistrasiSekolahTenantResponse{
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
	err = s.schemaService.SimpanSchemaSekolah(&models.SekolahTenant{
		UserId:      sekolah.UserId,
		SchemaName:  sekolah.Schemaname,
		NamaSekolah: sekolah.NamaSekolah,
	})
	if err != nil {
		log.Printf("Gagal menyimpan schema sekolah: %v", err)
		return nil, errors.New("gagal menyimpan informasi sekolah")
	}

	// 3 Kirim respon sukses
	return &pb.RegistrasiSekolahTenantResponse{
		Message: "Pembuatan database berhasil",
		Status:  true,
	}, nil
}

func (s *TenantServiceServer) GetSekolahTenant(ctx context.Context, req *pb.GetSekolahTenantRequest) (*pb.GetSekolahTenantResponse, error) {
	sekolahID := req.GetSekolahId()
	sekolahTerdaftar, err := s.schemaService.GetSchemaBySekolahID(int(sekolahID))
	if err != nil {
		return nil, err
	}

	return &pb.GetSekolahTenantResponse{
		SekolahTenant: &pb.SekolahTenant{
			NamaSekolah: sekolahTerdaftar.NamaSekolah,
			UserId:      sekolahTerdaftar.UserId,
			SekolahId:   sekolahTerdaftar.SekolahId,
			Schemaname:  sekolahTerdaftar.SchemaName,
		},
	}, err

}
