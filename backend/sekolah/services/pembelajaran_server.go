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

type PembelajaranServiceServer struct {
	pb.UnimplementedPembelajaranServiceServer
	repo repositories.GenericRepository[models.Pembelajaran]
}

func NewPembelajaranServiceServer() *PembelajaranServiceServer {
	repoPembelajaran := repositories.NewPembelajaranRepository(config.DB)
	return &PembelajaranServiceServer{
		repo: *repoPembelajaran,
	}
}

// **CreatePembelajaran**
func (s *PembelajaranServiceServer) CreatePembelajaran(ctx context.Context, req *pb.CreatePembelajaranRequest) (*pb.CreatePembelajaranResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "Pembelajaran"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	pembelajaran := req.GetPembelajaran()

	pembelajaranModel := utils.ConvertPBToModels(pembelajaran, func(item *pb.Pembelajaran) *models.Pembelajaran {
		return &models.Pembelajaran{
			PembelajaranId:     utils.StringToUUID(item.PembelajaranId),
			RombonganBelajarId: utils.StringToUUID(item.RombonganBelajarId),
			MataPelajaranId:    int(item.MataPelajaranId),
			SemesterId:         item.SemesterId,
			PtkTerdaftarId:     utils.PointerToUUID(item.PtkTerdaftarId),
			NamaMataPelajaran:  &item.NamaMataPelajaran,
		}
	})
	for _, v := range pembelajaranModel {
		err = s.repo.Save(ctx, v, schemaName)
		if err != nil {
			log.Printf("Gagal menyimpan Kelas: %v", err)
			return nil, fmt.Errorf("gagal menyimpan Kelas: %w", err)
		}
	}

	return &pb.CreatePembelajaranResponse{
		Message: "Kelas berhasil ditambahkan",
		Status:  true,
	}, nil
}

// func (s *PembelajaranServiceServer) CreateBanyakPembelajaran(ctx context.Context, req *pb.CreateBanyakPembelajaranRequest) (*pb.CreateBanyakPembelajaranResponse, error) {
// 	// Debugging: Cek nilai request yang diterima
// 	log.Printf("Received Sekolah data request: %+v\n", req)
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"Schemaname", "Pembelajaran"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	schemaName := req.GetSchemaname()
// 	Pembelajaran := req.Pembelajaran

// 	anggotaRombel := ConvertPBToModels(Pembelajaran, func(anggota *pb.Pembelajaran) *models.Pembelajaran {
// 		return &models.Pembelajaran{
// 			RombonganBelajarId: utils.StringToUUID(anggota.RombonganBelajarId),
// 			AnggotapembelajaranId:    utils.StringToUUID(anggota.AnggotapembelajaranId),
// 			PesertaDidikId:     utils.StringToUUID(anggota.PesertaDidikId),
// 			SemesterId:         anggota.SemesterId,
// 		}
// 	})
// 	err = s.repo.SaveMany(ctx, schemaName, anggotaRombel, 100)
// 	if err != nil {
// 		log.Printf("Gagal menyimpan Kelas: %v", err)
// 		return nil, fmt.Errorf("gagal menyimpan Kelas: %w", err)
// 	}

// 	return &pb.CreateBanyakPembelajaranResponse{
// 		Message: "Kelas berhasil ditambahkan",
// 		Status:  true,
// 	}, nil
// }

// **GetKelas**
func (s *PembelajaranServiceServer) GetPembelajaran(ctx context.Context, req *pb.GetPembelajaranRequest) (*pb.GetPembelajaranResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "SemesterId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	semesterId := req.GetSemesterId()
	var pembelajaranId string
	if req.GetPembelajaranId() != "" {
		pembelajaranId = req.GetPembelajaranId()
	}
	var conditions = map[string]any{
		"tabel_pembelajaran.semester_id": semesterId,
	}
	if pembelajaranId != "" {
		conditions["tabel_pembelajaran.pembelajar_id"] = pembelajaranId
	}
	joins := []string{
		"JOIN tabel_kelas ON tabel_kelas.rombongan_belajar_id = tabel_pembelajaran.rombongan_belajar_id",
		"JOIN tabel_ptk_terdaftar ON tabel_pembelajaran.ptk_terdaftar_id = tabel_ptk_terdaftar.ptk_terdaftar_id",
		// "JOIN tabel_ptk ON tabel_ptk_terdaftar.ptk_id = tabel_ptk.ptk_id",
		// "JOIN ref.mata_pelajaran ON ref.mata_pelajaran.mata_pelajaran_id = tabel_pembelajaran.mata_pelajaran_id",
	}
	preloads := []string{"PTKTerdaftar", "PTKTerdaftar.PTK"}

	orderBy := []string{} // Hindari duplikasi
	pembelajaranModel, err := s.repo.FindWithPreloadAndJoinsOrigin(ctx, schemaName, joins, preloads, conditions, orderBy)
	if err != nil {
		return nil, err
	}
	banyakPembelajaranList := utils.ConvertModelsToPB(pembelajaranModel, func(item models.Pembelajaran) *pb.Pembelajaran {
		return &pb.Pembelajaran{
			RombonganBelajarId: item.RombonganBelajarId.String(),
			MataPelajaranId:    int32(item.MataPelajaranId),
			NamaMataPelajaran:  utils.SafeString(item.NamaMataPelajaran),
			SemesterId:         item.SemesterId,
			PtkTerdaftarId:     item.PtkTerdaftarId.String(),
			PembelajaranId:     item.PembelajaranId.String(),
			PtkTerdaftar: &pb.PTKTerdaftar{
				PtkId: item.PTKTerdaftar.PtkID.String(),
				Ptk: &pb.PTK{
					Nama:              item.PTKTerdaftar.PTK.Nama,
					PtkId:             item.PTKTerdaftar.PtkID.String(),
					Nip:               utils.SafeString(item.PTKTerdaftar.PTK.NIP),
					JenisPtkId:        item.PTKTerdaftar.PTK.JenisPtkID,
					JenisKelamin:      item.PTKTerdaftar.PTK.JenisKelamin,
					TempatLahir:       item.PTKTerdaftar.PTK.TempatLahir,
					StatusKeaktifanId: item.PTKTerdaftar.PTK.StatusKeaktifanID,
					TanggalLahir:      item.PTKTerdaftar.PTK.TanggalLahir.Format("2006-01-02"),
					Nuptk:             utils.SafeString(item.PTKTerdaftar.PTK.NUPTK),
					AlamatJalan:       item.PTKTerdaftar.PTK.AlamatJalan,
				},
			},
			// RombonganBelajar: &pb.Kelas{
			// 	RombonganBelajarId:  item.RombonganBelajarId.String(),
			// 	NmKelas:             item.RombonganBelajar.NmKelas,
			// 	PtkId:               item.RombonganBelajar.PtkID.String(),
			// 	TingkatPendidikanId: item.RombonganBelajar.TingkatPendidikanId,
			// 	JurusanId:           item.RombonganBelajar.JurusanId,
			// 	NamaJurusanSp:       item.RombonganBelajar.NamaJurusanSp,
			// 	JenisRombel:         item.RombonganBelajar.JenisRombel,
			// 	KurikulumId:         item.RombonganBelajar.KurikulumId,

			// },
		}
	})
	return &pb.GetPembelajaranResponse{
		Pembelajaran: banyakPembelajaranList,
	}, nil
}

// func (s *PembelajaranServiceServer) SearchPembelajaran(ctx context.Context, req *pb.SearchPembelajaranRequest) (*pb.SearchPembelajaranResponse, error) {
// 	log.Printf("Received Sekolah data request: %+v\n", req)
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"Schemaname", "SemesterId"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	schemaName := req.GetSchemaname()
// 	switch schemaName {
// 	case "":
// 		return nil, fmt.Errorf("schema name is required")
// 	case "\"\"":
// 		return nil, fmt.Errorf("schema name cannot nul value")
// 	}
// 	semesterId := req.GetSemesterId()
// 	var banyakSiswa []models.Pembelajaran
// 	pesertaDidikId := utils.SafeString(req.PesertaDidikId)
// 	pembelajaranId := utils.SafeString(req.RombonganBelajarId)

// 	joins := []string{
// 		"JOIN tabel_siswa ON tabel_siswa.peserta_didik_id = tabel_pembelajaran.peserta_didik_id",
// 		"JOIN tabel_kelas ON tabel_kelas.rombongan_belajar_id = tabel_pembelajaran.rombongan_belajar_id",
// 		"JOIN tabel_ptk ON tabel_ptk.ptk_id = tabel_kelas.ptk_id",
// 	}
// 	preloads := []string{"PesertaDidik", "RombonganBelajar", "RombonganBelajar.PTK"}
// 	conditions := map[string]any{
// 		"tabel_pembelajaran.semester_id": semesterId,
// 	}
// 	// orderBy := []string{"tabel_kelas.nm_kelas ASC"} // Hindari duplikasi
// 	if pesertaDidikId != "" {
// 		// mod, err := s.repo.FindByID(ctx, pesertaDidikId, schemaName, "peserta_didik_id")
// 		conditions["tabel_pembelajaran.peserta_didik_id"] = pesertaDidikId
// 		mod, err := s.repo.FindWithRelations(ctx, schemaName, joins, preloads, conditions, nil)
// 		if err != nil {
// 			// 2. Jika ada error (selain not found), kembalikan error
// 			if !errors.Is(err, gorm.ErrRecordNotFound) {
// 				log.Printf("[ERROR] Gagal mencari siswa dengan ID %s di schema %s: %v", pesertaDidikId, schemaName, err)
// 				return nil, fmt.Errorf("gagal mencari siswa: %w", err)
// 			}
// 			// 1. Jika record tidak ditemukan, kembalikan nil tanpa error
// 			return nil, nil
// 		}

// 		banyakSiswa = mod
// 	} else if pembelajaranId != "" {
// 		conditions := map[string]any{
// 			"semester_id": semesterId,
// 		}
// 		mod, err := s.repo.FindAllByConditions(ctx, schemaName, conditions, 100, 0)
// 		if err != nil {
// 			log.Printf("[ERROR] Gagal menemukan siswa di schema '%s': %v", schemaName, err)
// 			return nil, fmt.Errorf("gagal menemukan siswa di schema '%s': %w", schemaName, err)
// 		}
// 		banyakSiswa = utils.PointerToSlice(mod)
// 	}
// 	banyakSiswaList := utils.ConvertModelsToPB(banyakSiswa, func(siswa models.Pembelajaran) *pb.Pembelajaran {
// 		return &pb.Pembelajaran{
// 			AnggotapembelajaranId:    siswa.AnggotapembelajaranId.String(),
// 			PesertaDidikId:     siswa.PesertaDidikId.String(),
// 			RombonganBelajarId: siswa.RombonganBelajarId.String(),
// 			RombonganBelajar: &pb.Kelas{
// 				NmKelas: siswa.RombonganBelajar.NmKelas,
// 			},
// 		}
// 	})
// 	return &pb.SearchPembelajaranResponse{
// 		Pembelajaran: banyakSiswaList,
// 	}, nil
// }

// **UpdateKelas**
// func (s *PembelajaranServiceServerServer) UpdateKelas(ctx context.Context, req *pb.UpdateKelasRequest) (*pb.UpdateKelasResponse, error) {
// 	// Debugging: Cek nilai request yang diterima
// 	log.Printf("Received UpdateUserProfile request: %+v\n", req)
// 	schemaName := req.GetSchemaname()
// 	KelasReq := req.GetKelas()
// 	KelasPelenReq := req.GetKelasPelengkap()
// 	Kelas := &models.PesertaDidik{
// 		PesertaDidikID:  KelasReq.PesertaDidikID,
// 		NIS:             KelasReq.NIS,
// 		NISN:            KelasReq.NISN,
// 		NamaKelas:       KelasReq.NamaKelas,
// 		TempatLahir:     KelasReq.TempatLahir,
// 		TanggalLahir:    KelasReq.TanggalLahir,
// 		JenisKelamin:    KelasReq.JenisKelamin,
// 		Agama:           KelasReq.Agama,
// 		AlamatKelas:     &KelasReq.AlamatKelas,
// 		TeleponKelas:    KelasReq.TeleponKelas,
// 		DiterimaTanggal: KelasReq.DiterimaTanggal,
// 		NamaAyah:        KelasReq.NamaAyah,
// 		NamaIbu:         KelasReq.NamaIbu,
// 		PekerjaanAyah:   KelasReq.PekerjaanAyah,
// 		PekerjaanIbu:    KelasReq.PekerjaanIbu,
// 		NamaWali:        &KelasReq.NamaWali,
// 		PekerjaanWali:   &KelasReq.PekerjaanWali,
// 	}
// 	KelasPelenkap := &models.PesertaDidikPelengkap{
// 		PelengkapKelasID: KelasPelenReq.PelengkapKelasID,
// 		PesertaDidikID:   &KelasPelenReq.PesertaDidikID,
// 		StatusDalamKel:   &KelasPelenReq.StatusDalamKel,
// 		AnakKe:           &KelasPelenReq.AnakKe,
// 		SekolahAsal:      KelasPelenReq.SekolahAsal,
// 		DiterimaKelas:    &KelasPelenReq.DiterimaKelas,
// 		AlamatOrtu:       &KelasPelenReq.AlamatOrtu,
// 		TeleponOrtu:      &KelasPelenReq.TeleponOrtu,
// 		AlamatWali:       &KelasPelenReq.AlamatWali,
// 		TeleponWali:      &KelasPelenReq.TeleponWali,
// 		FotoKelas:        &KelasPelenReq.FotoKelas,
// 	}
// 	err := s.pesertaDidikService.Update(ctx, Kelas, KelasPelenkap, schemaName)
// 	if err != nil {
// 		log.Printf("Gagal memperbarui Kelas: %v", err)
// 		return nil, fmt.Errorf("gagal memperbarui Kelas: %w", err)
// 	}

// 	return &pb.UpdateKelasResponse{
// 		Message: "Kelas berhasil diperbarui",
// 		Status:  true,
// 	}, nil
// }

// // **DeleteKelas**
func (s *PembelajaranServiceServer) DeletePembelajaran(ctx context.Context, req *pb.DeletePembelajaranRequest) (*pb.DeletePembelajaranResponse, error) {
	var err error
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "PembelajaranId"}
	// Validasi request
	err = utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	switch schemaName {
	case "":
		return nil, fmt.Errorf("schema name is required")
	case "\"\"":
		return nil, fmt.Errorf("schema name cannot nul value")
	}
	anggotapembelajaranId := req.GetPembelajaranId()

	err = s.repo.Delete(ctx, anggotapembelajaranId, schemaName, "anggota_rombel_id")
	if err != nil {
		log.Printf("Gagal menghapus Kelas: %v", err)
		return nil, fmt.Errorf("gagal menghapus Kelas: %w", err)
	}

	return &pb.DeletePembelajaranResponse{
		Message: "Anggota Kelas berhasil dihapus",
		Status:  true,
	}, nil
}

// // UploadKelas mengunggah data Kelas dari file Excel
// func (s *PembelajaranServiceServerServer) UploadKelas(ctx context.Context, req *pb.UploadKelasRequest) (*pb.UploadKelasResponse, error) {
// 	schemaName := req.GetSchemaname()
// 	fileData := req.GetFile() // File dalam bentuk byte array

// 	// Simpan file ke sementara
// 	tempFile := "/tmp/uploaded_Kelas.xlsx"
// 	err := saveFile(tempFile, fileData)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal menyimpan file sementara: %w", err)
// 	}

// 	// Baca file Excel
// 	f, err := excelize.OpenFile(tempFile)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal membaca file Excel: %w", err)
// 	}
// 	defer f.Close()

// 	// Ambil semua data dari sheet pertama
// 	rows, err := f.GetRows(f.GetSheetName(0))
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal mengambil data dari sheet: %w", err)
// 	}

// 	// Pastikan ada data
// 	if len(rows) < 2 {
// 		return nil, fmt.Errorf("file Excel kosong atau tidak memiliki data yang valid")
// 	}

// 	// Validasi header
// 	expectedHeaders := []string{"NIS", "NISN", "NamaKelas", "TempatLahir", "TanggalLahir", "JenisKelamin", "Agama"}
// 	for i, expected := range expectedHeaders {
// 		if rows[0][i] != expected {
// 			return nil, fmt.Errorf("format kolom tidak sesuai, kolom '%s' seharusnya ada di posisi %d", expected, i+1)
// 		}
// 	}

// 	var KelasList []*models.PesertaDidik

// 	// Mulai dari baris kedua karena baris pertama adalah header
// 	for _, row := range rows[1:] {
// 		if len(row) < len(expectedHeaders) {
// 			log.Println("Skipping row due to insufficient data:", row)
// 			continue
// 		}

// 		// Konversi data sesuai dengan model
// 		namaKelas := row[2]
// 		nis := row[0]
// 		nisn := row[1]
// 		tempatLahir := row[3]
// 		tanggalLahir := row[4]
// 		jenisKelamin := row[5]
// 		agama := row[6]

// 		// Validasi data
// 		if nis == "" || namaKelas == "" || nisn == "" {
// 			log.Println("Skipping row due to missing required fields:", row)
// 			continue
// 		}

// 		// Konversi angka
// 		nisInt, err := strconv.Atoi(nis)
// 		if err != nil {
// 			log.Printf("Format NIS tidak valid: %s", nis)
// 			continue
// 		}

// 		nisnInt, err := strconv.Atoi(nisn)
// 		if err != nil {
// 			log.Printf("Format NISN tidak valid: %s", nisn)
// 			continue
// 		}

// 		// Masukkan ke dalam list
// 		Kelas := &models.PesertaDidik{
// 			NIS:          strconv.Itoa(nisInt),
// 			NISN:         strconv.Itoa(nisnInt),
// 			NamaKelas:    namaKelas,
// 			TempatLahir:  tempatLahir,
// 			TanggalLahir: tanggalLahir,
// 			JenisKelamin: jenisKelamin,
// 			Agama:        agama,
// 		}
// 		KelasList = append(KelasList, Kelas)
// 	}

// 	// Simpan ke database
// 	err = s.pesertaDidikService.BatchSave(ctx, KelasList, schemaName)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal menyimpan data Kelas ke database: %w", err)
// 	}

// 	return &pb.UploadKelasResponse{
// 		Message: "Kelas berhasil diunggah",
// 		Total:   int32(len(KelasList)),
// 		Status:  true,
// 	}, nil
// }
