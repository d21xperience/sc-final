package services

import (
	"context"
	"fmt"
	"log"
	"sekolah/config"
	pb "sekolah/generated"
	"sekolah/models"
	"sekolah/repositories"
	"sekolah/utils"
)

type IjazahServiceServer struct {
	pb.UnimplementedIjazahServiceServer
	repo             repositories.GenericRepository[models.Ijazah]
	repoAnggotaKelas repositories.GenericRepository[models.RombelAnggota]
}

func NewIjazahServiceServer() *IjazahServiceServer {
	repoIjazah := repositories.NewIjazahRepository(config.DB)
	repoAnggotaKelas := repositories.NewRombelAnggotaRepository(config.DB)
	return &IjazahServiceServer{
		repo:             *repoIjazah,
		repoAnggotaKelas: *repoAnggotaKelas,
	}
}

// **CreateIjazah**
func (s *IjazahServiceServer) CreateIjazah(ctx context.Context, req *pb.CreateIjazahRequest) (*pb.CreateIjazahResponse, error) {

	return &pb.CreateIjazahResponse{
		Message: "ok",
		Status:  true,
	}, nil
}

// **GetIjazah**
func (s *IjazahServiceServer) GetProsesIjazah(ctx context.Context, req *pb.GetProsesIjazahRequest) (*pb.GetProsesIjazahResponse, error) {
	var err error
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "SemesterId"}
	// Validasi request
	err = utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	if schemaName == "\"\"" {
		return nil, fmt.Errorf("schema name is required")
	}

	// Cek apakah harus mengambil semua data atau data spesifik berdasarkan SemesterId
	// kelasId := req.GetKelasId()
	semesterId := req.GetSemesterId()
	var rombelAnggota []models.RombelAnggota
	var conditions = map[string]any{
		"tabel_anggotakelas.status_keaktifan": 1,
		"tabel_anggotakelas.semester_id":      semesterId,
	}
	if req.GetPesertaDidikId() != "" {
		conditions["tabel_anggotakelas.peserta_didik_id"] = req.GetPesertaDidikId()
	}

	joins := []string{
		// "JOIN tabel_ptk ON tabel_kelas.ptk_id = tabel_ptk.ptk_id",
		// "JOIN tabel_pembelajaran ON tabel_kelas.rombongan_belajar_id = tabel_pembelajaran.rombongan_belajar_id",
		// fmt.Sprintf("JOIN ref.jurusan ON %s.tabel_kelas.jurusan_id = ref.jurusan.jurusan_id", schemaName),
		// fmt.Sprintf("JOIN ref.kurikulum ON %s.tabel_kelas.kurikulum_id = ref.kurikulum.kurikulum_id", schemaName),
		// fmt.Sprintf("JOIN ref.tingkat_pendidikan ON %s.tabel_kelas.tingkat_pendidikan_id = ref.tingkat_pendidikan.tingkat_pendidikan_id", schemaName),
	}
	preloads := []string{"PesertaDidik", "RombonganBelajar"}

	groupByColumns := []string{} // Hindari duplikasi
	rombelAnggota, err = s.repoAnggotaKelas.FindWithPreloadAndJoins(ctx, schemaName, joins, preloads, conditions, groupByColumns)
	if err != nil {
		return nil, err
	}

	banyakKelasList := utils.ConvertModelsToPB(rombelAnggota, func(kelas models.RombelAnggota) *pb.AnggotaKelas {
		// jmlhAnggota, err := s.repoRombelAnggota.CountRows(ctx, schemaName, "rombongan_belajar_id", kelas.RombonganBelajarId.String())
		if err != nil {
			return nil
		}
		return &pb.AnggotaKelas{
			RombonganBelajarId: kelas.RombonganBelajarId.String(),
			SemesterId:         kelas.SemesterId,
			// PesertaDidikId:     kelas.PesertaDidikId.String(),
			AnggotaRombelId: kelas.AnggotaRombelId.String(),
			PesertaDidik: &pb.Siswa{
				PesertaDidikId: kelas.PesertaDidik.PesertaDidikId,
				NmSiswa:        kelas.PesertaDidik.NmSiswa,
				JenisKelamin:   kelas.PesertaDidik.JenisKelamin,
				Nis:            kelas.PesertaDidik.Nis,
				Nisn:           kelas.PesertaDidik.Nisn,
				Nik:            utils.SafeString(kelas.PesertaDidik.Nik),
				TempatLahir:    kelas.PesertaDidik.TempatLahir,
				TanggalLahir:   kelas.PesertaDidik.TanggalLahir.String(),
				Agama:          kelas.PesertaDidik.Agama,
			},
			RombonganBelajar: &pb.Kelas{
				NmKelas: kelas.RombonganBelajar.NmKelas,
			},
		}
	})
	return &pb.GetProsesIjazahResponse{
		Status:       true,
		Message:      "Sukses",
		AnggotaKelas: banyakKelasList,
	}, nil
}

// **UpdateIjazah**
func (s *IjazahServiceServer) UpdateIjazah(ctx context.Context, req *pb.UpdateIjazahRequest) (*pb.UpdateIjazahResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received UpdateUserProfile request: %+v\n", req)
	schemaName := req.GetSchemaname()
	ijazahReq := req.GetIjazah()
	ijazahModel := &models.Ijazah{
		Nama:                        ijazahReq.Nama,
		Nis:                         ijazahReq.Nis,
		NISN:                        ijazahReq.Nisn,
		NPSN:                        ijazahReq.Npsn,
		NoIjazah:                    ijazahReq.NomorIjazah,
		TempatLahir:                 ijazahReq.TempatLahir,
		TanggalLahir:                ijazahReq.TanggalLahir,
		NamaOrtuWali:                ijazahReq.NamaOrtuwali,
		PaketKeahlian:               ijazahReq.PaketKeahlian,
		KabupatenKota:               ijazahReq.Kabupatenkota,
		Provinsi:                    ijazahReq.Provinsi,
		ProgramKeahlian:             ijazahReq.ProgramKeahlian,
		SekolahPenyelenggaraUjianUS: ijazahReq.SekolahPenyelenggaraUjianUs,
		SekolahPenyelenggaraUjianUN: ijazahReq.SekolahPenyelenggaraUjianUn,
		AsalSekolah:                 ijazahReq.AsalSekolah,
		NomorIjazah:                 ijazahReq.NomorIjazah,
		TempatIjazah:                ijazahReq.TempatIjazah,
		TanggalIjazah:               ijazahReq.TanggalIjazah,
	}
	err := s.repo.Update(ctx, ijazahModel, schemaName, "id", ijazahReq.ID)
	if err != nil {
		log.Printf("Gagal memperbarui Ijazah: %v", err)
		return nil, fmt.Errorf("gagal memperbarui Ijazah: %w", err)
	}
	return &pb.UpdateIjazahResponse{
		Message: "Ijazah berhasil diperbarui",
		Status:  true,
	}, nil
}

// **DeleteIjazah**
func (s *IjazahServiceServer) DeleteIjazah(ctx context.Context, req *pb.DeleteIjazahRequest) (*pb.DeleteIjazahResponse, error) {
	schemaName := req.GetSchemaname()
	IjazahID := req.GetIjazahId()

	err := s.repo.Delete(ctx, IjazahID, schemaName, "id")
	if err != nil {
		log.Printf("Gagal menghapus Ijazah: %v", err)
		return nil, fmt.Errorf("gagal menghapus Ijazah: %w", err)
	}

	return &pb.DeleteIjazahResponse{
		Message: "Ijazah berhasil dihapus",
		Status:  true,
	}, nil
}
