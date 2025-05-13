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

type RombelAnggotaService struct {
	pb.UnimplementedAnggotaKelasServiceServer
	repo      repositories.GenericRepository[models.RombelAnggota]
	repoSiswa repositories.GenericRepository[models.PesertaDidik]
}

func NewRombelAnggotaService() *RombelAnggotaService {
	repoRombelAnggota := repositories.NewRombelAnggotaRepository(config.DB)
	repoRombelSiswa := repositories.NewSiswaRepository(config.DB)
	return &RombelAnggotaService{
		repo:      *repoRombelAnggota,
		repoSiswa: *repoRombelSiswa,
	}
}

// **CreateKelas**
// func (s *RombelAnggotaService) CreateKelas(ctx context.Context, req *pb.CreateAnggotaKelasRequest) (*pb.CreateAnggotaKelasResponse, error) {
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
// 	Kelas := req.Kelas

// 	KelasModel := &models.PesertaDidik{
// 		PesertaDidikID:  Kelas.PesertaDidikId,
// 		Nis:             Kelas.Nis,
// 		Nisn:            Kelas.Nisn,
// 		NmKelas:         Kelas.NmKelas,
// 		TempatLahir:     Kelas.TempatLahir,
// 		TanggalLahir:    Kelas.TanggalLahir,
// 		JenisKelamin:    Kelas.JenisKelamin,
// 		Agama:           Kelas.Agama,
// 		AlamatKelas:     &Kelas.AlamatKelas,
// 		TeleponKelas:    Kelas.TeleponKelas,
// 		DiterimaTanggal: Kelas.DiterimaTanggal,
// 		NmAyah:          Kelas.NmAyah,
// 		NmIbu:           Kelas.NmIbu,
// 		PekerjaanAyah:   Kelas.PekerjaanAyah,
// 		PekerjaanIbu:    Kelas.PekerjaanIbu,
// 		NmWali:          &Kelas.NmWali,
// 		PekerjaanWali:   &Kelas.PekerjaanWali,
// 	}

// 	err = s.pesertaDidikService.Save(ctx, KelasModel, schemaName)
// 	if err != nil {
// 		log.Printf("Gagal menyimpan Kelas: %v", err)
// 		return nil, fmt.Errorf("gagal menyimpan Kelas: %w", err)
// 	}

//		return &pb.CreateKelasResponse{
//			Message: "Kelas berhasil ditambahkan",
//			Status:  true,
//		}, nil
//	}
func (s *RombelAnggotaService) CreateBanyakAnggotaKelas(ctx context.Context, req *pb.CreateBanyakAnggotaKelasRequest) (*pb.CreateBanyakAnggotaKelasResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "AnggotaKelas"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	anggotaKelas := req.AnggotaKelas

	anggotaRombel := ConvertPBToModels(anggotaKelas, func(anggota *pb.AnggotaKelas) *models.RombelAnggota {
		tglLahir, err := utils.StringToTime(anggota.PesertaDidik.TanggalLahir, "2006-01-02")
		if err != nil {
			return nil
		}
		return &models.RombelAnggota{
			RombonganBelajarId: utils.StringToUUID(anggota.RombonganBelajarId),
			AnggotaRombelId:    utils.StringToUUID(anggota.AnggotaRombelId),
			PesertaDidikId:     utils.StringToUUID(anggota.PesertaDidikId),
			SemesterId:         anggota.SemesterId,
			PesertaDidik: models.PesertaDidik{
				PesertaDidikId: anggota.PesertaDidik.PesertaDidikId,
				Nis:            anggota.PesertaDidik.Nis,
				Nisn:           anggota.PesertaDidik.Nisn,
				NmSiswa:        anggota.PesertaDidik.NmSiswa,
				TempatLahir:    anggota.PesertaDidik.TempatLahir,
				TanggalLahir:   &tglLahir,
				JenisKelamin:   anggota.PesertaDidik.JenisKelamin,
				Agama:          anggota.PesertaDidik.Agama,
				AlamatSiswa:    &anggota.PesertaDidik.AlamatSiswa,
				TeleponSiswa:   anggota.PesertaDidik.TeleponSiswa,
				// DiterimaTanggal: &tglDiterima,
				NmAyah:        anggota.PesertaDidik.NmAyah,
				NmIbu:         anggota.PesertaDidik.NmIbu,
				PekerjaanAyah: anggota.PesertaDidik.PekerjaanAyah,
				PekerjaanIbu:  anggota.PesertaDidik.PekerjaanIbu,
				NmWali:        &anggota.PesertaDidik.NmWali,
				PekerjaanWali: &anggota.PesertaDidik.PekerjaanWali,
			},
		}
	})
	err = s.repo.SaveMany(ctx, schemaName, anggotaRombel, 100)
	if err != nil {
		log.Printf("Gagal menyimpan Kelas: %v", err)
		return nil, fmt.Errorf("gagal menyimpan Kelas: %w", err)
	}

	var pesertaDidikList []models.PesertaDidik

	for _, anggota := range anggotaRombel {
		if anggota != nil {
			pesertaDidikList = append(pesertaDidikList, anggota.PesertaDidik)
		}
	}
	err = s.repoSiswa.SaveMany(ctx, schemaName, utils.SliceToPointer(pesertaDidikList), 100)
	if err != nil {
		log.Printf("Gagal menyimpan Kelas: %v", err)
		return nil, fmt.Errorf("gagal menyimpan Kelas: %w", err)
	}

	// err = s.repoSiswa.SaveMany(ctx, schemaName,)
	// if err != nil {
	// 	log.Printf("Gagal menyimpan Kelas: %v", err)
	// 	return nil, fmt.Errorf("gagal menyimpan Kelas: %w", err)
	// }

	return &pb.CreateBanyakAnggotaKelasResponse{
		Message: "Kelas berhasil ditambahkan",
		Status:  true,
	}, nil
}

// **GetKelas**
func (s *RombelAnggotaService) GetAnggotaKelas(ctx context.Context, req *pb.GetAnggotaKelasRequest) (*pb.GetAnggotaKelasResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("call rombel_anggota_server: data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "SemesterId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	semesterId := req.GetSemesterId()
	anggotaRombelId := req.GetAnggotaRombelId()
	var conditions = map[string]any{
		"tabel_anggotakelas.semester_id": semesterId,
		"jenis_rombel":                   1,
	}
	if anggotaRombelId != "" {
		conditions["tabel_anggotakelas.anggota_rombel_id"] = anggotaRombelId
	}
	joins := []string{
		"JOIN tabel_siswa ON tabel_siswa.peserta_didik_id = tabel_anggotakelas.peserta_didik_id",
		"JOIN tabel_kelas ON tabel_kelas.rombongan_belajar_id = tabel_anggotakelas.rombongan_belajar_id",
		"JOIN tabel_ptk ON tabel_ptk.ptk_id = tabel_kelas.ptk_id",
	}
	preloads := []string{"PesertaDidik", "RombonganBelajar", "RombonganBelajar.PTK", "NilaiAkhir", "NilaiAkhir.MataPelajaran"}

	orderBy := []string{"tabel_kelas.nm_kelas ASC", "tabel_siswa.nm_siswa ASC"} // Hindari duplikasi
	anggotaRombelModel, err := s.repo.FindWithPreloadAndJoins(ctx, schemaName, joins, preloads, conditions, nil, orderBy, false)
	if err != nil {
		return nil, err
	}

	banyakAnggotaKelasList := utils.ConvertModelsToPB(anggotaRombelModel, func(anggota models.RombelAnggota) *pb.AnggotaKelas1 {

		return &pb.AnggotaKelas1{
			// SemesterId:         anggota.SemesterId,
			RombonganBelajarId:  anggota.RombonganBelajarId.String(),
			NmKelas:             anggota.RombonganBelajar.NmKelas,
			AnggotaRombelId:     anggota.AnggotaRombelId.String(),
			TingkatPendidikanId: anggota.RombonganBelajar.TingkatPendidikanId,
			WaliKelas:           anggota.RombonganBelajar.PTK.Nama,
			PesertaDidikId:      anggota.PesertaDidikId.String(),
			NmSiswa:             anggota.PesertaDidik.NmSiswa,
			Nis:                 anggota.PesertaDidik.Nis,
			Nisn:                anggota.PesertaDidik.Nisn,
			Nik:                 utils.SafeString(anggota.PesertaDidik.Nik),
			TempatLahir:         anggota.PesertaDidik.TempatLahir,
			TanggalLahir:        anggota.PesertaDidik.TanggalLahir.Format("2006-01-02"),
			JenisKelamin:        anggota.PesertaDidik.JenisKelamin,
			NmIbu:               anggota.PesertaDidik.NmIbu,
			Agama:               anggota.PesertaDidik.Agama,
			// PtkId:              anggota.RombonganBelajar.PtkID.String(),
		}
	})
	return &pb.GetAnggotaKelasResponse{
		AnggotaKelas: banyakAnggotaKelasList,
	}, nil
}

func (s *RombelAnggotaService) SearchAnggotaKelas(ctx context.Context, req *pb.SearchAnggotaKelasRequest) (*pb.SearchAnggotaKelasResponse, error) {
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "SemesterId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
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
	semesterId := req.GetSemesterId()
	var banyakSiswa []models.RombelAnggota
	pesertaDidikId := utils.SafeString(req.PesertaDidikId)
	rombelId := utils.SafeString(req.RombonganBelajarId)
	tingkatId := utils.SafeString(req.TingkatId)
	joins := []string{
		"JOIN tabel_siswa ON tabel_siswa.peserta_didik_id = tabel_anggotakelas.peserta_didik_id",
		"JOIN tabel_kelas ON tabel_kelas.rombongan_belajar_id = tabel_anggotakelas.rombongan_belajar_id",
		"JOIN tabel_ptk ON tabel_ptk.ptk_id = tabel_kelas.ptk_id",
	}
	preloads := []string{"PesertaDidik", "RombonganBelajar", "RombonganBelajar.PTK"}
	conditions := map[string]any{
		"tabel_anggotakelas.semester_id": semesterId,
	}
	// customConditions := []struct {
	// 	Query string
	// 	Args  []interface{}
	// }{
	// 	{"age > ?", []interface{}{25}},
	// 	// {"created_at BETWEEN ? AND ?", []interface{}{startDate, endDate}},
	// }
	// orderBy := []string{"tabel_kelas.nm_kelas ASC"} // Hindari duplikasi
	if pesertaDidikId != "" {
		// mod, err := s.repo.FindByID(ctx, pesertaDidikId, schemaName, "peserta_didik_id")
		conditions["tabel_anggotakelas.peserta_didik_id"] = pesertaDidikId
		mod, err := s.repo.FindWithRelations(ctx, schemaName, joins, preloads, conditions, nil, nil, nil)
		if err != nil {
			// 2. Jika ada error (selain not found), kembalikan error
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				log.Printf("[ERROR] Gagal mencari siswa dengan ID %s di schema %s: %v", pesertaDidikId, schemaName, err)
				return nil, fmt.Errorf("gagal mencari siswa: %w", err)
			}
			// 1. Jika record tidak ditemukan, kembalikan nil tanpa error
			return nil, nil
		}

		banyakSiswa = mod
	} else if rombelId != "" {
		conditions := map[string]any{
			"semester_id": semesterId,
		}
		mod, err := s.repo.FindAllByConditions(ctx, schemaName, conditions, 100, 0, nil)
		if err != nil {
			log.Printf("[ERROR] Gagal menemukan siswa di schema '%s': %v", schemaName, err)
			return nil, fmt.Errorf("gagal menemukan siswa di schema '%s': %w", schemaName, err)
		}
		banyakSiswa = utils.PointerToSlice(mod)
	} else if tingkatId != "" {
		conditions := map[string]any{
			"tabel_kelas.tingkat_id": tingkatId,
		}

		mod, err := s.repo.FindWithPreloadAndJoins(ctx, schemaName, joins, []string{"RombonganBelajar"}, conditions, nil, []string{"tabel_kelas.nm_kelas ASC"}, false)
		if err != nil {
			log.Printf("[ERROR] Gagal menemukan siswa di schema '%s': %v", schemaName, err)
			return nil, fmt.Errorf("gagal menemukan siswa di schema '%s': %w", schemaName, err)
		}
		banyakSiswa = mod
	}
	banyakSiswaList := utils.ConvertModelsToPB(banyakSiswa, func(siswa models.RombelAnggota) *pb.AnggotaKelas {
		return &pb.AnggotaKelas{
			AnggotaRombelId:    siswa.AnggotaRombelId.String(),
			PesertaDidikId:     siswa.PesertaDidikId.String(),
			RombonganBelajarId: siswa.RombonganBelajarId.String(),
			RombonganBelajar: &pb.Kelas{
				NmKelas: siswa.RombonganBelajar.NmKelas,
			},
		}
	})
	return &pb.SearchAnggotaKelasResponse{
		AnggotaKelas: banyakSiswaList,
	}, nil
}

// **UpdateKelas**
// func (s *RombelAnggotaServiceServer) UpdateKelas(ctx context.Context, req *pb.UpdateKelasRequest) (*pb.UpdateKelasResponse, error) {
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
func (s *RombelAnggotaService) DeleteAnggotaKelas(ctx context.Context, req *pb.DeleteAnggotaKelasRequest) (*pb.DeleteAnggotaKelasResponse, error) {
	var err error
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "AnggotaRombelId"}
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
	anggotaRombelId := req.GetAnggotaRombelId()

	err = s.repo.Delete(ctx, anggotaRombelId, schemaName, "anggota_rombel_id")
	if err != nil {
		log.Printf("Gagal menghapus Kelas: %v", err)
		return &pb.DeleteAnggotaKelasResponse{
			Message: fmt.Sprintf("Gagal menghapus anggota kelas %s", err),
			Status:  false,
		}, nil
	}

	return &pb.DeleteAnggotaKelasResponse{
		Message: "Anggota Kelas berhasil dihapus",
		Status:  true,
	}, nil
}

func (s *RombelAnggotaService) DeleteBatchAnggotaKelas(ctx context.Context, req *pb.DeleteBatchAnggotaKelasRequest) (*pb.DeleteBatchAnggotaKelasResponse, error) {
	var err error
	log.Printf("DeleteBatchAnggotaKelas Receiveddata request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "AnggotaRombelId"}
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
	anggotaRombelId := req.GetAnggotaRombelId()

	err = s.repo.DeleteBatch(ctx, anggotaRombelId, schemaName, "anggota_rombel_id")
	if err != nil {
		log.Printf("Gagal menghapus anggota kelas: %v", err)
		// return nil, fmt.Errorf("gagal menghapus Kelas: %w", err)
		return &pb.DeleteBatchAnggotaKelasResponse{
			Message: fmt.Sprintf("Gagal menghapus banyak anggota kelas %s", err),
			Status:  false,
		}, nil
	}

	return &pb.DeleteBatchAnggotaKelasResponse{
		Message: "Anggota Kelas berhasil dihapus",
		Status:  true,
	}, nil

}

// // UploadKelas mengunggah data Kelas dari file Excel
// func (s *RombelAnggotaServiceServer) UploadKelas(ctx context.Context, req *pb.UploadKelasRequest) (*pb.UploadKelasResponse, error) {
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

func (s *RombelAnggotaService) FilterAnggotaKelas(ctx context.Context, req *pb.FilterAnggotaKelasRequest) (*pb.FilterAnggotaKelasResponse, error) {
	var err error
	log.Printf("rombel_anggota_server->FilterAnggotaKelas with data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "SemesterId"}
	// Validasi request
	err = utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaname := req.GetSchemaname()
	semesterId := req.GetSemesterId()

	tingkatPendId := req.GetTingkatPendidikanId()
	rombelId := req.GetRombonganBelajarId()
	var anggotaRombelModels []models.RombelAnggota
	joins := []string{
		"JOIN tabel_kelas ON tabel_anggotakelas.rombongan_belajar_id = tabel_kelas.rombongan_belajar_id",
		"JOIN tabel_siswa ON tabel_anggotakelas.peserta_didik_id = tabel_siswa.peserta_didik_id",
	}
	preloads := []string{"RombonganBelajar", "PesertaDidik"}
	conditions := map[string]any{"tabel_anggotakelas.semester_id": semesterId}
	orderBy := []string{"tabel_kelas.nm_kelas", "tabel_siswa.nm_siswa"}

	if tingkatPendId != 0 {
		conditions["tabel_kelas.tingkat_pendidikan_id"] = tingkatPendId
		// orderBy = append(orderBy, "tabel_siswa.nm_siswa")
		anggotaRombelModels, err = s.repo.FindWithPreloadAndJoins(ctx, schemaname, joins, preloads, conditions, nil, orderBy, false)
		if err != nil {
			return nil, err
		}
	} else if rombelId != "" {
		conditions["tabel_anggotakelas.rombongan_belajar_id"] = rombelId
		anggotaRombelModels, err = s.repo.FindWithPreloadAndJoins(ctx, schemaname, joins, preloads, conditions, nil, orderBy, false)
		if err != nil {
			return nil, err
		}
	}
	if anggotaRombelModels == nil {
		return nil, fmt.Errorf("parameter kosong, pilih salah satu %v atau %v", tingkatPendId, rombelId)
	}
	results := utils.ConvertModelsToPB(anggotaRombelModels, func(item models.RombelAnggota) *pb.AnggotaKelas {
		return &pb.AnggotaKelas{
			AnggotaRombelId:    item.AnggotaRombelId.String(),
			PesertaDidikId:     item.PesertaDidikId.String(),
			NmSiswa:            item.PesertaDidik.NmSiswa,
			NmKelas:            item.RombonganBelajar.NmKelas,
			RombonganBelajarId: item.RombonganBelajarId.String(),
			PesertaDidik: &pb.Siswa{
				Nis:          item.PesertaDidik.Nis,
				Nisn:         item.PesertaDidik.Nisn,
				TempatLahir:  item.PesertaDidik.TempatLahir,
				TanggalLahir: item.PesertaDidik.TanggalLahir.Format("2006-01-02"),
			},
		}
	})

	return &pb.FilterAnggotaKelasResponse{
		AnggotaKelas: results,
	}, nil
}
