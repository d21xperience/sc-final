package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	pb "sekolah/generated"
	"sekolah/models"
	"sekolah/services"

	"gorm.io/gorm"
)

type SekolahServiceServer struct {
	pb.UnimplementedSekolahServiceServer
	// RedisClient    *redis.Client // Tambahkan Redis sebagai field
	schemaService  services.SchemaService
	sekolahService services.SekolahService
}

// // Constructor untuk AuthServiceServer dengan Redis
// func NewAuthServiceServer() *SekolahServiceServer {
// 	return &SekolahServiceServer{}
// }

func (s *SekolahServiceServer) RegistrasiSekolah(ctx context.Context, req *pb.TabelSekolahRequest) (*pb.TabelSekolahResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)

	// Cek apakah req atau req.Sekolah kosong
	if req == nil {
		log.Println("Request is nil")
		return nil, errors.New("invalid request: request is nil")
	}

	if req.Sekolah == nil {
		log.Println("Sekolah is nil in request")
		return nil, errors.New("invalid request: sekolah is nil")
	}
	sekolah := req.GetSekolah()
	namaSchema := sekolah.SekolahIdEnkrip
	existingSchema, err := s.schemaService.GetSchemaBySekolahID(int(sekolah.SekolahId))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Lanjutkan pendaftaran: %v", err)
		// return nil, errors.New("gagal mengecek pendaftaran sekolah")
	}

	if existingSchema != nil {
		return &pb.TabelSekolahResponse{
			Message: "Pendaftaran dibatalkan: Sekolah sudah terdaftar",
			Status:  false,
		}, nil
	}

	cek := s.schemaService.RegistrasiSekolah(ctx, namaSchema)
	if errors.Is(cek, gorm.ErrInvalidData) {
		return nil, errors.New("gagal total")
	}
	// 2 Simpan informasi schema sekolah
	err = s.schemaService.SimpanSchemaSekolah(&models.SekolahTabelTenant{
		SekolahID:  int(sekolah.SekolahId),
		NamaSchema: namaSchema,
		Nama:       sekolah.NamaSekolah,
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

func (s *SekolahServiceServer) GetSekolahTabelTenant(ctx context.Context, req *pb.SekolahTabelTenantRequest) (*pb.SekolahTabelTenantResponse, error) {
	sekolahID := req.GetSekolahId()
	sekolahTerdaftar, err := s.schemaService.GetSchemaBySekolahID(int(sekolahID))
	if err != nil {
		return nil, err
	}

	return &pb.SekolahTabelTenantResponse{
		SekolahId:  int32(sekolahTerdaftar.SekolahID),
		Nama:       sekolahTerdaftar.Nama,
		NamaSchema: sekolahTerdaftar.NamaSchema, // nama schema
	}, err

}

// ================================================================================//
// Tenant table
func (s *SekolahServiceServer) CreateSekolah(ctx context.Context, req *pb.CreateSekolahRequest) (*pb.CreateSekolahResponse, error) {
	schemaName := req.GetSchemaname()
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

func (s *SekolahServiceServer) GetSekolah(ctx context.Context, req *pb.GetSekolahRequest) (*pb.GetSekolahResponse, error) {
	// 🔥 Ambil schema dari request
	schemaName := req.GetSchemaname()
	// sekolahID := req.GetSekolahId()

	// 🔥 Cari sekolah berdasarkan ID dan schema
	sekolah, err := s.sekolahService.Find(ctx, schemaName)
	if err != nil {
		log.Printf("Gagal menemukan sekolah: %v", err)
		return nil, fmt.Errorf("gagal menemukan sekolah: %w", err)
	}

	// 🔥 Return response dalam format protobuf
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
