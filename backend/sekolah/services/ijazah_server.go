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

	"github.com/google/uuid"
)

type IjazahServiceServer struct {
	pb.UnimplementedIjazahServiceServer
	repo             repositories.GenericRepository[models.Ijazah]
	repoKelas        repositories.GenericRepository[models.RombonganBelajar]
	repoAnggotaKelas repositories.GenericRepository[models.RombelAnggota]
	repoSiswa        repositories.GenericRepository[models.PesertaDidik]
	repoSekolah      repositories.SekolahRepository
}

func NewIjazahServiceServer() *IjazahServiceServer {
	repoIjazah := repositories.NewIjazahRepository(config.DB)
	repoKelas := repositories.NewrombonganBelajarRepository(config.DB)
	repoAnggotaKelas := repositories.NewRombelAnggotaRepository(config.DB)
	repoSiswa := repositories.NewSiswaRepository(config.DB)
	repoSekolah := repositories.NewSekolahRepository(config.DB)

	return &IjazahServiceServer{
		repo:             *repoIjazah,
		repoKelas:        *repoKelas,
		repoAnggotaKelas: *repoAnggotaKelas,
		repoSiswa:        *repoSiswa,
		repoSekolah:      repoSekolah,
	}
}

// **CreateIjazah**
func (s *IjazahServiceServer) CreateIjazah(ctx context.Context, req *pb.CreateIjazahRequest) (*pb.CreateIjazahResponse, error) {
	var err error
	// log.Printf("Received Sekolah data request: %+v\n", req)

	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "TahunAjaranId"}
	// Validasi request
	err = utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	if schemaName == "\"\"" {
		return nil, fmt.Errorf("schema name is required")
	}
	anggotaKelashModels := utils.ConvertPBToModels(req.GetAnggotaKelas(), func(item *pb.AnggotaKelas) *models.RombelAnggota {
		tglLahir, err := utils.StringToTime(item.PesertaDidik.TanggalLahir, "2006-01-02")
		if err != nil {
			return nil
		}
		return &models.RombelAnggota{
			// ID:                          uuid.New(),
			PesertaDidikId:     utils.StringToUUID(item.PesertaDidikId),
			AnggotaRombelId:    utils.StringToUUID(item.AnggotaRombelId),
			RombonganBelajarId: utils.StringToUUID(item.RombonganBelajarId),
			PesertaDidik: models.PesertaDidik{
				PesertaDidikId: item.PesertaDidik.PesertaDidikId,
				Nis:            item.PesertaDidik.Nis,
				Nisn:           item.PesertaDidik.Nisn,
				NmSiswa:        item.PesertaDidik.NmSiswa,
				TempatLahir:    item.PesertaDidik.TempatLahir,
				TanggalLahir:   &tglLahir,
				JenisKelamin:   item.PesertaDidik.JenisKelamin,
				Agama:          item.PesertaDidik.Agama,
				AlamatSiswa:    &item.PesertaDidik.AlamatSiswa,
				TeleponSiswa:   item.PesertaDidik.TeleponSiswa,
				// DiterimaTanggal: &tglDiterima,
				NmAyah:        item.PesertaDidik.NmAyah,
				NmIbu:         item.PesertaDidik.NmIbu,
				PekerjaanAyah: item.PesertaDidik.PekerjaanAyah,
				PekerjaanIbu:  item.PesertaDidik.PekerjaanIbu,
				NmWali:        &item.PesertaDidik.NmWali,
				PekerjaanWali: &item.PesertaDidik.PekerjaanWali,
				Nik:           &item.PesertaDidik.Nik,
			},
			// RombonganBelajarId: ,
		}
	})
	sekolahModel, err := s.repoSekolah.FindByID(ctx, req.GetSekolahId(), schemaName)
	if err != nil {
		return nil, err
	}

	ijazahModels := []models.Ijazah{}
	for _, v := range anggotaKelashModels {
		kelasModel, err := s.repoKelas.FindByID(ctx, v.RombonganBelajarId.String(), schemaName, "rombongan_belajar_id")
		if err != nil {
			return nil, err
		}
		ijazahModel := models.Ijazah{
			ID: uuid.New(),

			PesertaDidikId:              v.PesertaDidikId,
			AnggotaRombelId:             v.AnggotaRombelId,
			NomorIjazah:                 utils.GenerateNomorIjazah(sekolahModel.Npsn, utils.ParseInt(req.TahunAjaranId)),
			TempatIjazah:                sekolahModel.KabKota,
			TanggalIjazah:               nil,
			ProgramKeahlian:             kelasModel.NamaJurusanSp,
			PaketKeahlian:               kelasModel.NamaJurusanSp,
			SekolahID:                   kelasModel.SekolahId,
			NamaOrtuWali:                v.PesertaDidik.NmAyah,
			SekolahPenyelenggaraUjianUS: sekolahModel.Nama,
			SekolahPenyelenggaraUjianUN: sekolahModel.Nama,
			Nis:                         v.PesertaDidik.Nis,
			NISN:                        v.PesertaDidik.Nisn,

			// BlockexplorerUrl : req.GetBlockexplorerUrl()
			// CidUrl : ""
			TahunAjaranId: req.GetTahunAjaranId(),
			Status:        0,
		}
		ijazahModels = append(ijazahModels, ijazahModel)
	}
	err = s.repo.SaveMany(ctx, schemaName, utils.SliceToPointer(ijazahModels), 100)
	if err != nil {
		return nil, err
	}
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
	var rombelAnggota []models.Ijazah
	var conditions = map[string]any{
		"ijazah.tahun_ajaran_id": semesterId,
	}
	if req.GetPesertaDidikId() != "" {
		conditions["ijazah.peserta_didik_id"] = req.GetPesertaDidikId()
	}

	// joins := []string{
	// 	// "JOIN tabel_ptk ON tabel_kelas.ptk_id = tabel_ptk.ptk_id",
	// 	// "JOIN tabel_pembelajaran ON tabel_kelas.rombongan_belajar_id = tabel_pembelajaran.rombongan_belajar_id",
	// 	// fmt.Sprintf("JOIN ref.jurusan ON %s.tabel_kelas.jurusan_id = ref.jurusan.jurusan_id", schemaName),
	// 	// fmt.Sprintf("JOIN ref.kurikulum ON %s.tabel_kelas.kurikulum_id = ref.kurikulum.kurikulum_id", schemaName),
	// 	// fmt.Sprintf("JOIN ref.tingkat_pendidikan ON %s.tabel_kelas.tingkat_pendidikan_id = ref.tingkat_pendidikan.tingkat_pendidikan_id", schemaName),
	// }
	preloads := []string{"PesertaDidik", "AnggotaRombel", "AnggotaRombel.RombonganBelajar"}
	// orderBy := []string{"tabel_kelas.nm_kelas ASC"}

	// groupByColumns := []string{} // Hindari duplikasi
	rombelAnggota, err = s.repo.FindWithPreloadAndJoins(ctx, schemaName, nil, preloads, conditions, nil, nil, false)
	if err != nil {
		return nil, err
	}

	banyakKelasList := utils.ConvertModelsToPB(rombelAnggota, func(kelas models.Ijazah) *pb.Ijazah {
		if err != nil {
			return nil
		}
		return &pb.Ijazah{
			AnggotaRombelId: kelas.AnggotaRombelId.String(),
			NamaOrtuWali:    kelas.PesertaDidik.NmAyah,
			ProgramKeahlian: kelas.ProgramKeahlian,
			RombonganBelajar: &pb.Kelas{
				NmKelas: kelas.AnggotaRombel.RombonganBelajar.NmKelas,
			},
			PesertaDidik: &pb.Siswa{
				PesertaDidikId: kelas.PesertaDidik.PesertaDidikId,
				NmSiswa:        kelas.PesertaDidik.NmSiswa,
				JenisKelamin:   kelas.PesertaDidik.JenisKelamin,
				Nis:            kelas.PesertaDidik.Nis,
				Nisn:           kelas.PesertaDidik.Nisn,
				Nik:            utils.SafeString(kelas.PesertaDidik.Nik),
				TempatLahir:    kelas.PesertaDidik.TempatLahir,
				TanggalLahir:   kelas.PesertaDidik.TanggalLahir.Format("2006-01-02"),
				Agama:          kelas.PesertaDidik.Agama,
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
		NamaOrtuWali:                ijazahReq.NamaOrtuWali,
		PaketKeahlian:               ijazahReq.PaketKeahlian,
		ProgramKeahlian:             ijazahReq.ProgramKeahlian,
		SekolahPenyelenggaraUjianUS: ijazahReq.SekolahPenyelenggaraUjianUs,
		SekolahPenyelenggaraUjianUN: ijazahReq.SekolahPenyelenggaraUjianUn,
		NomorIjazah:                 ijazahReq.NomorIjazah,
		TempatIjazah:                ijazahReq.TempatIjazah,
		TanggalIjazah:               nil,
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
