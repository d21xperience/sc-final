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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RombelServiceServer struct {
	pb.UnimplementedKelasServiceServer
	repo              repositories.GenericRepository[models.RombonganBelajar]
	repoRombelAnggota repositories.GenericRepository[models.RombelAnggota]
	repoSemester      repositories.GenericRepository[models.Semester]
	repoSiswa         repositories.GenericRepository[models.PesertaDidik]
	repoPembelajaran  repositories.GenericRepository[models.Pembelajaran]
}

func NewRombelServiceServer() *RombelServiceServer {
	repoRombel := repositories.NewrombonganBelajarRepository(config.DB)
	repoRombelAnggota := repositories.NewRombelAnggotaRepository(config.DB)
	repoSemester := repositories.NewSemesterRepository(config.DB)
	repoSiswa := repositories.NewSiswaRepository(config.DB)
	repoPembelajaran := repositories.NewPembelajaranRepository(config.DB)
	return &RombelServiceServer{
		repo:              *repoRombel,
		repoRombelAnggota: *repoRombelAnggota,
		repoSemester:      *repoSemester,
		repoSiswa:         *repoSiswa,
		repoPembelajaran:  *repoPembelajaran,
	}
}

// **CreateKelas**
func (s *RombelServiceServer) CreateKelas(ctx context.Context, req *pb.CreateKelasRequest) (*pb.CreateKelasResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "Kelas"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	kelas := req.Kelas
	// rombelId := uuid.New()
	kelasModel := utils.ConvertPBToModels(kelas, func(item *pb.Kelas) *models.RombonganBelajar {
		return &models.RombonganBelajar{
			RombonganBelajarId:  utils.StringToUUID(item.RombonganBelajarId),
			SekolahId:           utils.StringToUUID(item.SekolahId),
			SemesterId:          item.SemesterId,
			JurusanId:           item.JurusanId,
			PtkID:               utils.StringToUUID(item.PtkId),
			NmKelas:             item.NmKelas,
			TingkatPendidikanId: item.TingkatPendidikanId,
			JenisRombel:         item.JenisRombel,
			NamaJurusanSp:       item.NamaJurusanSp,
			KurikulumId:         item.KurikulumId,
		}
	})
	// simpan kelas ke database
	res := s.repo.SaveMany(ctx, schemaName, kelasModel, 100)
	if res != nil {
		return nil, err
	}

	return &pb.CreateKelasResponse{
		Message: "Kelas berhasil ditambahkan",
		Status:  true,
	}, nil
}

// func (s *RombelServiceServer) CreateBanyakKelas(ctx context.Context, req *pb.CreateBanyakKelasRequest) (*pb.CreateBanyakKelasResponse, error) {
// 	// Debugging: Cek nilai request yang diterima
// 	log.Printf("Received Sekolah data request: %+v\n", req)
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"Schemaname", "Kelas"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	schemaName := req.GetSchemaname()
// 	kelas := req.GetKelas()

// 	rombelId := uuid.New()
// 	sekolahId, err := uuid.Parse(kelas.SekolahId)
// 	if err != nil {
// 		fmt.Println("Error parsing UUID:", err)
// 	}
// 	ptkId, err := uuid.Parse(kelas.PtkId)
// 	if err != nil {
// 		fmt.Println("Error parsing UUID:", err)
// 	}
// 	kelasModels := ConvertPBToModels(kelas, func(rom *pb.Kelas) *models.RombonganBelajar {
// 		return &models.RombonganBelajar{
// 			RombonganBelajarId:  rombelId,
// 			SekolahId:           sekolahId,
// 			SemesterId:          rom.SemesterId,
// 			JurusanId:           rom.JurusanId,
// 			PtkID:               ptkId,
// 			NmKelas:             rom.NmKelas,
// 			TingkatPendidikanId: rom.TingkatPendidikanId,
// 			JenisRombel:         rom.JenisRombel,
// 			NamaJurusanSp:       rom.NamaJurusanSp,
// 			// JurusanSpId:         rom.JurusanSpId,
// 			KurikulumId: rom.KurikulumId,
// 		}
// 	})
// 	err = s.repo.SaveMany(ctx, schemaName, kelasModels, 100)
// 	if err != nil {
// 		log.Printf("Gagal menyimpan Kelas: %v", err)
// 		return nil, fmt.Errorf("gagal menyimpan Kelas: %w", err)
// 	}

// 	return &pb.CreateBanyakKelasResponse{
// 		Message: "Kelas berhasil ditambahkan",
// 		Status:  true,
// 	}, nil
// }

// **GetKelas**
func (s *RombelServiceServer) GetKelas(ctx context.Context, req *pb.GetKelasRequest) (*pb.GetKelasResponse, error) {
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
	var rombelModel []models.RombonganBelajar
	var conditions map[string]any
	if req.KelasId != "" {
		conditions = map[string]any{
			"rombongan_belajar_id": req.KelasId,
		}
	} else {
		conditions = map[string]any{
			"semester_id": semesterId,
		}
	}
	joins := []string{
		"JOIN tabel_ptk ON tabel_kelas.ptk_id = tabel_ptk.ptk_id",
		fmt.Sprintf("JOIN ref.jurusan ON %s.tabel_kelas.jurusan_id = ref.jurusan.jurusan_id", schemaName),
		fmt.Sprintf("JOIN ref.kurikulum ON %s.tabel_kelas.kurikulum_id = ref.kurikulum.kurikulum_id", schemaName),
		fmt.Sprintf("JOIN ref.tingkat_pendidikan ON %s.tabel_kelas.tingkat_pendidikan_id = ref.tingkat_pendidikan.tingkat_pendidikan_id", schemaName),
	}
	preloads := []string{"PTK", "Jurusan", "Kurikulum", "TingkatPendidikan"}

	groupByColumns := []string{"tabel_kelas.rombongan_belajar_id"} // Hindari duplikasi
	rombelModel, err = s.repo.FindWithPreloadAndJoins(ctx, schemaName, joins, preloads, conditions, groupByColumns)
	if err != nil {
		return nil, err
	}

	banyakKelasList := utils.ConvertModelsToPB(rombelModel, func(kelas models.RombonganBelajar) *pb.Kelas {
		// jurusanSPId, err := utils.ConvertUUIDToStringViceVersa(kelas.JurusanSpId)
		// if err != nil {
		// 	return nil
		// }
		jmlhAnggota, err := s.repoRombelAnggota.CountRows(ctx, schemaName, "rombongan_belajar_id", kelas.RombonganBelajarId.String())
		if err != nil {
			return nil
		}
		return &pb.Kelas{
			RombonganBelajarId:  kelas.RombonganBelajarId.String(),
			SekolahId:           kelas.SekolahId.String(),
			SemesterId:          kelas.SemesterId,
			JurusanId:           kelas.JurusanId,
			PtkId:               kelas.PtkID.String(),
			NmKelas:             kelas.NmKelas,
			TingkatPendidikanId: kelas.TingkatPendidikanId,
			JenisRombel:         kelas.JenisRombel,
			NamaJurusanSp:       kelas.NamaJurusanSp,
			// JurusanSpId:         jurusanSPId.(*string),
			KurikulumId: kelas.KurikulumId,
			Ptk: &pb.PTK{
				PtkId:             kelas.PTK.PtkID.String(),
				Nama:              kelas.PTK.Nama,
				JenisKelamin:      kelas.PTK.JenisKelamin,
				JenisPtkId:        kelas.PTK.JenisPtkID,
				TempatLahir:       kelas.PTK.TempatLahir,
				TanggalLahir:      kelas.PTK.TanggalLahir.Format("2006-01-02"),
				AlamatJalan:       kelas.PTK.AlamatJalan,
				StatusKeaktifanId: kelas.PTK.StatusKeaktifanID,
				Nuptk:             utils.SafeString(kelas.PTK.NUPTK),
				Nip:               utils.SafeString(kelas.PTK.NIP),
			},
			Jurusan: &pb.Jurusan{
				JurusanId:           kelas.Jurusan.JurusanID,
				NamaJurusan:         kelas.Jurusan.NamaJurusan,
				JenjangPendidikanId: utils.PointerToUint32(utils.Uint16ToUint32Pointer(kelas.Jurusan.JenjangPendidikanID)),
				UntukSma:            uint32(kelas.Jurusan.UntukSMA),
				UntukSmk:            uint32(kelas.Jurusan.UntukSMK),
				UntukPt:             uint32(kelas.Jurusan.UntukPT),
				UntukSlb:            uint32(kelas.Jurusan.UntukSLB),
				UntukSmklb:          uint32(kelas.Jurusan.UntukSMKLB),
				JurusanInduk:        utils.SafeString(kelas.Jurusan.JurusanInduk),
				LevelBidangId:       kelas.Jurusan.LevelBidangID,
			},
			Kurikulum: &pb.Kurikulum{
				KurikulumId:         uint32(kelas.Kurikulum.KurikulumID),
				NamaKurikulum:       kelas.Kurikulum.NamaKurikulum,
				MulaiBerlaku:        kelas.Kurikulum.MulaiBerlaku.Format("2006-01-02"),
				JenjangPendidikanId: uint32(kelas.Kurikulum.JenjangPendidikanID),
				SistemSks:           uint32(kelas.Kurikulum.SistemSKS),
				JurusanId:           *kelas.Kurikulum.JurusanID,
			},
			TingkatPendidikan: &pb.TingkatPendidikan{
				TingkatPendidikanId: uint32(kelas.TingkatPendidikan.TingkatPendidikanID),
				Kode:                kelas.TingkatPendidikan.Kode,
				Nama:                kelas.TingkatPendidikan.Nama,
				JenjangPendidikanId: uint32(kelas.TingkatPendidikan.JenjangPendidikanID),
			},
			JumlahAnggota: uint32(jmlhAnggota),
		}
	})
	return &pb.GetKelasResponse{
		Kelas: banyakKelasList,
	}, nil
}

// **UpdateKelas**
func (s *RombelServiceServer) UpdateKelas(ctx context.Context, req *pb.UpdateKelasRequest) (*pb.UpdateKelasResponse, error) {
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "Kelas"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	Kelas := req.Kelas
	// rombelId := uuid.New()
	sekolahId, err := uuid.Parse(Kelas.SekolahId)
	if err != nil {
		fmt.Println("Error parsing UUID:", err)
	}
	ptkId, err := uuid.Parse(Kelas.PtkId)
	if err != nil {
		fmt.Println("Error parsing UUID:", err)
	}
	rombelId, err := uuid.Parse(Kelas.RombonganBelajarId)
	if err != nil {
		fmt.Println("Error parsing UUID:", err)
	}
	KelasModel := &models.RombonganBelajar{
		RombonganBelajarId:  rombelId,
		NmKelas:             Kelas.NmKelas,
		SekolahId:           sekolahId,
		SemesterId:          Kelas.SemesterId,
		JurusanId:           Kelas.JurusanId,
		TingkatPendidikanId: Kelas.TingkatPendidikanId,
		PtkID:               ptkId,
		JenisRombel:         Kelas.JenisRombel,
		NamaJurusanSp:       Kelas.NamaJurusanSp,
		JurusanSpId:         &uuid.Nil,
		KurikulumId:         Kelas.KurikulumId,
		// RombonganBelajarId:  kelas.RombonganBelajarId,
	}
	err = s.repo.Update(ctx, KelasModel, schemaName, "rombongan_belajar_id", Kelas.RombonganBelajarId)
	if err != nil {
		log.Printf("Gagal memperbaharui Kelas: %v", err)
		return nil, fmt.Errorf("gagal memperbaharui Kelas: %w", err)
	}

	// cek apakah anggota kelas berisi nilai
	anggotaKelas := req.GetAnggotaKelas()
	if len(anggotaKelas) > 0 {
		// var modelAnggotaKelas []models.RombelAnggota // Inisialisasi slice kosong

		// for _, v := range anggotaKelas {
		// newUUID := uuid.New()
		// ced := models.RombelAnggota{ // Tidak perlu pakai pointer di sini
		// 	AnggotaRombelId:    utils.StringToUUID(v.AnggotaRombelId),
		// 	RombonganBelajarId: utils.StringToUUID(v.RombonganBelajarId),
		// 	PesertaDidikId:     utils.StringToUUID(v.PesertaDidikId),
		// 	SemesterId:         KelasModel.SemesterId,
		// }
		// modelAnggota, err := s.repoRombelAnggota.FindByID(ctx, ced.PesertaDidikId.String(), schemaName, "peserta_didik_id")
		// if err != nil {
		// 	return nil, err
		// }

		// if modelAnggota == nil {
		// 	// cek di tabel_siswa
		// 	// jika di tabel_siswa tidak ada buat di tabel_siswa
		// 	pdModel, err := s.repoSiswa.FindOrCreateByID(ctx, v.PesertaDidikId, schemaName, "peserta_didik_id", func(s string) *models.PesertaDidik {
		// 		return &models.PesertaDidik{
		// 			PesertaDidikId: v.PesertaDidik.Nis,
		// 			Nis:            v.PesertaDidik.Nis,
		// 			Nisn:           v.PesertaDidik.Nisn,
		// 		}
		// 	})
		// 	if err != nil {
		// 		return nil, err
		// 	}
		// 	// daftarkan ke tabel_anggota
		// 	if pdModel == nil {
		// 		fmt.Printf("User ditemukan/ditambahkan: %+v\n", pdModel)
		// 	}
		// 	// jika di tabel_siswa ada
		// 	// daftarkan ke tabel_anggota
		// 	// results := s.repoRombelAnggota.Save(ctx, &ced, schemaName)
		// 	// if results != nil {
		// 	// 	return nil, err
		// 	// }
		// } else {
		// 	results := s.repoRombelAnggota.Update(ctx, &ced, schemaName, "peserta_didik_id", ced.PesertaDidikId.String())
		// 	if results != nil {
		// 		return nil, err
		// 	}
		// }
		// modelAnggotaKelas = append(modelAnggotaKelas, ced) // Tambahkan ke slice
		// modelAnggota, err := s.repoRombelAnggota.Update(ctx, ced,schemaName,"rombel_id",ced.AnggotaRombelId)
		// }
	}
	return &pb.UpdateKelasResponse{
		Message: "Kelas berhasil diperbarui",
		Status:  true,
	}, nil
}

// **DeleteKelas**
func (s *RombelServiceServer) DeleteKelas(ctx context.Context, req *pb.DeleteKelasRequest) (*pb.DeleteKelasResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "KelasId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	KelasID := req.GetKelasId()

	err = s.repo.Delete(ctx, KelasID, schemaName, "rombongan_belajar_id")
	if err != nil {
		log.Printf("Gagal menghapus Kelas: %v", err)
		return nil, fmt.Errorf("gagal menghapus Kelas: %w", err)
	}

	return &pb.DeleteKelasResponse{
		Message: "Kelas berhasil dihapus",
		Status:  true,
	}, nil
}

func (s *RombelServiceServer) ImportDapodikRombel(ctx context.Context, req *pb.ImportDapodikRombelRequest) (*pb.ImportDapodikRombelResponse, error) {
	// Debugging: Cek nilai request yang diterima
	// log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "Kelas"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	kelas := req.Kelas
	// rombelId := uuid.New()
	// type rombelAnggota struct {
	// 	AnggotaRombelId    uuid.UUID // UUID
	// 	PesertaDidikId     uuid.UUID // UUID
	// 	RombonganBelajarId uuid.UUID
	// 	SemesterId         string
	// }
	simpanRombelAnggota := []models.RombelAnggota{}
	simpanRombelPembelajaran := []models.Pembelajaran{}
	kelasModel := utils.ConvertPBToModels(kelas, func(item *pb.Kelas) *models.RombonganBelajar {
		for _, v := range item.AnggotaKelas {
			simpanRombelAnggota = append(simpanRombelAnggota, models.RombelAnggota{
				AnggotaRombelId:    utils.StringToUUID(v.AnggotaRombelId),
				PesertaDidikId:     utils.StringToUUID(v.PesertaDidikId),
				RombonganBelajarId: utils.StringToUUID(v.RombonganBelajarId),
				SemesterId:         item.SemesterId,
			})
		}
		for _, v := range item.Pembelajaran {
			simpanRombelPembelajaran = append(simpanRombelPembelajaran, models.Pembelajaran{
				PembelajaranId:     utils.StringToUUID(v.PembelajaranId),
				RombonganBelajarId: utils.StringToUUID(v.RombonganBelajarId),
				MataPelajaranId:    int(v.MataPelajaranId),
				SemesterId:         v.SemesterId,
				PtkTerdaftarId:     utils.PointerToUUID(v.PtkTerdaftarId),
				NamaMataPelajaran:  &v.NamaMataPelajaran,
			})
		}

		return &models.RombonganBelajar{
			RombonganBelajarId:  utils.StringToUUID(item.RombonganBelajarId),
			SekolahId:           utils.StringToUUID(item.SekolahId),
			SemesterId:          item.SemesterId,
			JurusanId:           item.JurusanId,
			PtkID:               utils.StringToUUID(item.PtkId),
			NmKelas:             item.NmKelas,
			TingkatPendidikanId: item.TingkatPendidikanId,
			JenisRombel:         item.JenisRombel,
			NamaJurusanSp:       item.NamaJurusanSp,
			KurikulumId:         item.KurikulumId,
		}
	})
	// simpan kelas ke database
	conflicts, err := s.repo.SaveManyWithConflictCheck(ctx, schemaName, kelasModel, "RombonganBelajarId", "rombongan_belajar_id", 100, []string{})
	if err != nil {
		// Jika error karena *database failure* (misal connection, syntax error, dll), hentikan proses
		st, ok := status.FromError(err)
		if !ok || st.Code() != codes.AlreadyExists {
			return nil, status.Errorf(codes.Internal, "insert failed: %v", err)
		}
		// Kalau errornya karena duplicate, lanjut aja
		log.Printf("[WARNING] Terjadi duplikat data rombel, melanjutkan proses: %v", err)
	}
	conflictProto := repositories.ConvertConflictsToProto(conflicts, "PtkTerdaftarId", "NmKelas")

	// simpan anggota rombel
	conflicts2, err := s.repoRombelAnggota.SaveManyWithConflictCheck(ctx, schemaName, utils.SliceToPointer(simpanRombelAnggota), "AnggotaRombelId", "anggota_rombel_id", 100, []string{"peserta_didik_id", "rombongan_belajar_id"})
	if err != nil {
		// Jika error karena *database failure* (misal connection, syntax error, dll), hentikan proses
		st, ok := status.FromError(err)
		if !ok || st.Code() != codes.AlreadyExists {
			return nil, status.Errorf(codes.Internal, "insert failed: %v", err)
		}
		// Kalau errornya karena duplicate, lanjut aja
		log.Printf("[WARNING] Terjadi duplikat data rombel, melanjutkan proses: %v", err)
	}
	conflictProto2 := repositories.ConvertConflictsToProto(conflicts2, "AnggotaRombelId", "PesertaDidikId")

	// Simpan pembelajaran
	conflicts3, err := s.repoPembelajaran.SaveManyWithConflictCheck(ctx, schemaName, utils.SliceToPointer(simpanRombelPembelajaran), "PembelajaranId", "pembelajaran_id", 100, []string{"ptk_terdaftar_id", "rombongan_belajar_id"})
	if err != nil {
		// Jika error karena *database failure* (misal connection, syntax error, dll), hentikan proses
		st, ok := status.FromError(err)
		if !ok || st.Code() != codes.AlreadyExists {
			return nil, status.Errorf(codes.Internal, "insert failed: %v", err)
		}
		// Kalau errornya karena duplicate, lanjut aja
		log.Printf("[WARNING] Terjadi duplikat data rombel, melanjutkan proses: %v", err)
	}
	conflictProto3 := repositories.ConvertConflictsToProto(conflicts3, "PembelajaranId", "NamaMataPelajaran")

	var resultsConflict []*pb.ConflictRow
	resultsConflict = append(resultsConflict, conflictProto...)
	resultsConflict = append(resultsConflict, conflictProto2...)
	resultsConflict = append(resultsConflict, conflictProto3...)

	fmt.Print(resultsConflict)

	return &pb.ImportDapodikRombelResponse{
		Message: "Kelas berhasil ditambahkan",
		Status:  true,
		Conflicts: &pb.ConflictResponse{
			Message:       "Sebagian data berhasil disimpan",
			Conflicts:     resultsConflict,
			TotalConflict: int32(len(resultsConflict)),
		},
	}, nil
}
