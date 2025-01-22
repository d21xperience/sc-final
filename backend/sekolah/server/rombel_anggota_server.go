package server

import (
	"context"
	"fmt"
	"log"
	pb "sekolah/generated"
	"sekolah/models"
	"sekolah/services"
	"sekolah/utils"
)

type RombelAnggotaServiceServer struct {
	pb.UnimplementedAnggotaKelasServiceServer
	rombelAnggotaService services.RombelAnggotaService
}

// **CreateKelas**
// func (s *RombelAnggotaServiceServer) CreateKelas(ctx context.Context, req *pb.CreateKelasRequest) (*pb.CreateKelasResponse, error) {
// 	// Debugging: Cek nilai request yang diterima
// 	log.Printf("Received Sekolah data request: %+v\n", req)
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"SchemaName", "Kelas"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	schemaName := req.GetSchemaName()
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
func (s *RombelAnggotaServiceServer) CreateBanyakAnggotaKelas(ctx context.Context, req *pb.CreateBanyakAnggotaKelasRequest) (*pb.CreateBanyakAnggotaKelasResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"SchemaName", "AnggotaKelas"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaName()
	anggotaKelas := req.AnggotaKelas

	anggotaRombel := ConvertPBToModels(anggotaKelas, func(anggota *pb.AnggotaKelas) *models.RombelAnggota {
		return &models.RombelAnggota{
			RombonganBelajarId: anggota.RombonganBelajarId,
			AnggotaRombelId:    anggota.AnggotaRombelId,
			PesertaDidikId:     anggota.PesertaDidikId,
			SemesterId:         anggota.SemesterId,
		}
	})
	err = s.rombelAnggotaService.SaveMany(ctx, schemaName, anggotaRombel)
	if err != nil {
		log.Printf("Gagal menyimpan Kelas: %v", err)
		return nil, fmt.Errorf("gagal menyimpan Kelas: %w", err)
	}

	return &pb.CreateBanyakAnggotaKelasResponse{
		Message: "Kelas berhasil ditambahkan",
		Status:  true,
	}, nil
}

// **GetKelas**
func (s *RombelAnggotaServiceServer) GetAnggotaKelas(ctx context.Context, req *pb.GetAnggotaKelasRequest) (*pb.GetAnggotaKelasResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"SchemaName", "SemesterId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaName()
	kelasId := req.GetKelasId()
	semesterId := req.GetSemesterId()
	var conditions = map[string]interface{}{
		"semester_id": semesterId,
	}

	if kelasId != "" {
		// Ambil data Kelas berdasarkan PesertaDidikId
		RombelAnggota, err := s.rombelAnggotaService.FindByID(ctx, kelasId, schemaName)
		if err != nil {
			return nil, err
		}
		return &pb.GetAnggotaKelasResponse{
			AnggotaKelas: []*pb.AnggotaKelas{
				ConvertModelToPB(RombelAnggota, func(anggota *models.RombelAnggota) *pb.AnggotaKelas {
					return &pb.AnggotaKelas{
						RombonganBelajarId: anggota.RombonganBelajarId,
						AnggotaRombelId:    anggota.AnggotaRombelId,
						PesertaDidikId:     anggota.PesertaDidikId,
						SemesterId:         anggota.SemesterId,
					}
				}),
			},
		}, nil
	}
	// Ambil semua data Kelas
	banyakAnggotaKelas, err := s.rombelAnggotaService.FindAllByConditions(ctx, schemaName, conditions, int(req.GetLimit()), int(req.GetOffset()))
	if err != nil {
		log.Printf("[ERROR] Gagal menemukan Kelas di schema '%s': %v", schemaName, err)
		return nil, fmt.Errorf("gagal menemukan Kelas di schema '%s': %w", schemaName, err)
	}
	banyakAnggotaKelasList := ConvertModelsToPB(banyakAnggotaKelas, func(anggota *models.RombelAnggota) *pb.AnggotaKelas {
		return &pb.AnggotaKelas{
			RombonganBelajarId: anggota.RombonganBelajarId,
			AnggotaRombelId:    anggota.AnggotaRombelId,
			PesertaDidikId:     anggota.PesertaDidikId,
			SemesterId:         anggota.SemesterId,
		}
	})
	return &pb.GetAnggotaKelasResponse{
		AnggotaKelas: banyakAnggotaKelasList,
	}, nil
}

// **UpdateKelas**
// func (s *RombelAnggotaServiceServer) UpdateKelas(ctx context.Context, req *pb.UpdateKelasRequest) (*pb.UpdateKelasResponse, error) {
// 	// Debugging: Cek nilai request yang diterima
// 	log.Printf("Received UpdateUserProfile request: %+v\n", req)
// 	schemaName := req.GetSchemaName()
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

// // // **DeleteKelas**
// func (s *RombelAnggotaServiceServer) DeleteKelas(ctx context.Context, req *pb.DeleteKelasRequest) (*pb.DeleteKelasResponse, error) {
// 	schemaName := req.GetSchemaName()
// 	KelasID := req.GetKelasId()

// 	err := s.pesertaDidikService.Delete(ctx, KelasID, schemaName)
// 	if err != nil {
// 		log.Printf("Gagal menghapus Kelas: %v", err)
// 		return nil, fmt.Errorf("gagal menghapus Kelas: %w", err)
// 	}

// 	return &pb.DeleteKelasResponse{
// 		Message: "Kelas berhasil dihapus",
// 		Status:  true,
// 	}, nil
// }

// // UploadKelas mengunggah data Kelas dari file Excel
// func (s *RombelAnggotaServiceServer) UploadKelas(ctx context.Context, req *pb.UploadKelasRequest) (*pb.UploadKelasResponse, error) {
// 	schemaName := req.GetSchemaName()
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
