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

// AuthServiceServer dengan Redis Client sebagai Dependency Injection
type SekolahServiceServer struct {
	pb.UnimplementedSchoolServiceServer
	// RedisClient    *redis.Client // Tambahkan Redis sebagai field
	schemaService       services.SchemaService
	sekolahService      services.SekolahService
	pesertaDidikService services.PesertaDidikService
	// nilaiAkhirService   services.NilaiAkhirService
}

// Constructor untuk AuthServiceServer dengan Redis
func NewAuthServiceServer() *SekolahServiceServer {
	return &SekolahServiceServer{}
}

func (s *SekolahServiceServer) RegistrasiSekolah(ctx context.Context, req *pb.TabelSekolahRequest) (*pb.TabelSekolahResponse, error) {
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

// **CreateSiswa**
func (s *SekolahServiceServer) CreateSiswa(ctx context.Context, req *pb.CreateSiswaRequest) (*pb.CreateSiswaResponse, error) {
	schemaName := req.GetSchemaname()
	siswa := req.GetSiswa()

	siswaModel := &models.PesertaDidik{
		PesertaDidikID:  siswa.PesertaDidikID,
		NIS:             siswa.NIS,
		NISN:            siswa.NISN,
		NamaSiswa:       siswa.NamaSiswa,
		TempatLahir:     siswa.TempatLahir,
		TanggalLahir:    siswa.TanggalLahir,
		JenisKelamin:    siswa.JenisKelamin,
		Agama:           siswa.Agama,
		AlamatSiswa:     &siswa.AlamatSiswa,
		TeleponSiswa:    siswa.TeleponSiswa,
		DiterimaTanggal: siswa.DiterimaTanggal,
		NamaAyah:        siswa.NamaAyah,
		NamaIbu:         siswa.NamaIbu,
		PekerjaanAyah:   siswa.PekerjaanAyah,
		PekerjaanIbu:    siswa.PekerjaanIbu,
		NamaWali:        &siswa.NamaWali,
		PekerjaanWali:   &siswa.PekerjaanWali,
	}

	err := s.pesertaDidikService.Save(ctx, siswaModel, schemaName)
	if err != nil {
		log.Printf("Gagal menyimpan siswa: %v", err)
		return nil, fmt.Errorf("gagal menyimpan siswa: %w", err)
	}

	return &pb.CreateSiswaResponse{
		Message: "Siswa berhasil ditambahkan",
		Status:  true,
	}, nil
}

// **GetSiswa**
// func (s *SekolahServiceServer) GetSiswa(ctx context.Context, req *pb.GetSiswaRequest) (*pb.GetSiswaResponse, error) {
// 	schemaName := req.GetSchemaname()
// 	siswaID := req.GetSiswaId()

// 	siswa, err := s.siswaRepo.FindByID(ctx, siswaID, schemaName)
// 	if err != nil {
// 		log.Printf("Gagal menemukan siswa: %v", err)
// 		return nil, fmt.Errorf("gagal menemukan siswa: %w", err)
// 	}

// 	return &pb.GetSiswaResponse{
// 		SiswaId: siswa.SiswaID.String(),
// 		Nama:    siswa.Nama,
// 		Kelas:   siswa.Kelas,
// 	}, nil
// }

// // **UpdateSiswa**
// func (s *SekolahServiceServer) UpdateSiswa(ctx context.Context, req *pb.UpdateSiswaRequest) (*pb.UpdateSiswaResponse, error) {
// 	schemaName := fmt.Sprintf("tabel_%s", req.GetSchemaName())
// 	siswaID, err := uuid.Parse(req.GetSiswaId())
// 	if err != nil {
// 		return nil, fmt.Errorf("format UUID tidak valid: %w", err)
// 	}

// 	siswa := &models.PesertaDidik{
// 		SiswaID: siswaID,
// 		Nama:    req.GetNama(),
// 		Kelas:   req.GetKelas(),
// 	}

// 	err = s.siswaRepo.Update(ctx, siswa, schemaName)
// 	if err != nil {
// 		log.Printf("Gagal memperbarui siswa: %v", err)
// 		return nil, fmt.Errorf("gagal memperbarui siswa: %w", err)
// 	}

// 	return &pb.UpdateSiswaResponse{
// 		Message: "Siswa berhasil diperbarui",
// 		Status:  true,
// 	}, nil
// }

// // **DeleteSiswa**
// func (s *SekolahServiceServer) DeleteSiswa(ctx context.Context, req *pb.DeleteSiswaRequest) (*pb.DeleteSiswaResponse, error) {
// 	schemaName := fmt.Sprintf("tabel_%s", req.GetSchemaName())
// 	siswaID := req.GetSiswaId()

// 	err := s.siswaRepo.Delete(ctx, siswaID, schemaName)
// 	if err != nil {
// 		log.Printf("Gagal menghapus siswa: %v", err)
// 		return nil, fmt.Errorf("gagal menghapus siswa: %w", err)
// 	}

// 	return &pb.DeleteSiswaResponse{
// 		Message: "Siswa berhasil dihapus",
// 		Status:  true,
// 	}, nil
// }
