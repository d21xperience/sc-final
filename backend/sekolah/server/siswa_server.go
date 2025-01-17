package server

import (
	"context"
	"fmt"
	"log"
	pb "sekolah/generated"
	"sekolah/models"
	"sekolah/services"
)

type SiswaServiceServer struct {
	pb.UnimplementedSiswaServiceServer
	pesertaDidikService services.PesertaDidikService
}

// **CreateSiswa**
func (s *SiswaServiceServer) CreateSiswa(ctx context.Context, req *pb.CreateSiswaRequest) (*pb.CreateSiswaResponse, error) {
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
func (s *SiswaServiceServer) GetSiswa(ctx context.Context, req *pb.GetSiswaRequest) (*pb.GetSiswaResponse, error) {
	schemaName := req.GetSchemaname()
	siswaID := req.GetSiswaId()

	siswa, err := s.pesertaDidikService.FindByID(ctx, siswaID, schemaName)
	if err != nil {
		log.Printf("Gagal menemukan siswa: %v", err)
		return nil, fmt.Errorf("gagal menemukan siswa: %w", err)
	}

	return &pb.GetSiswaResponse{
		Siswa: &pb.Siswa{
			PesertaDidikID: siswa.PesertaDidikID,
		},
	}, nil
}

// **UpdateSiswa**
func (s *SiswaServiceServer) UpdateSiswa(ctx context.Context, req *pb.UpdateSiswaRequest) (*pb.UpdateSiswaResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received UpdateUserProfile request: %+v\n", req)
	schemaName := req.GetSchemaname()
	siswaReq := req.GetSiswa()
	siswaPelenReq := req.GetSiswaPelengkap()
	siswa := &models.PesertaDidik{
		PesertaDidikID:  siswaReq.PesertaDidikID,
		NIS:             siswaReq.NIS,
		NISN:            siswaReq.NISN,
		NamaSiswa:       siswaReq.NamaSiswa,
		TempatLahir:     siswaReq.TempatLahir,
		TanggalLahir:    siswaReq.TanggalLahir,
		JenisKelamin:    siswaReq.JenisKelamin,
		Agama:           siswaReq.Agama,
		AlamatSiswa:     &siswaReq.AlamatSiswa,
		TeleponSiswa:    siswaReq.TeleponSiswa,
		DiterimaTanggal: siswaReq.DiterimaTanggal,
		NamaAyah:        siswaReq.NamaAyah,
		NamaIbu:         siswaReq.NamaIbu,
		PekerjaanAyah:   siswaReq.PekerjaanAyah,
		PekerjaanIbu:    siswaReq.PekerjaanIbu,
		NamaWali:        &siswaReq.NamaWali,
		PekerjaanWali:   &siswaReq.PekerjaanWali,
	}
	siswaPelenkap := &models.PesertaDidikPelengkap{
		PelengkapSiswaID: siswaPelenReq.PelengkapSiswaID,
		PesertaDidikID: &models.PesertaDidik{
			PesertaDidikID: siswaReq.PesertaDidikID,
		},
		StatusDalamKel: &siswaPelenReq.StatusDalamKel,
		AnakKe:         &siswaPelenReq.AnakKe,
		SekolahAsal:    siswaPelenReq.SekolahAsal,
		DiterimaKelas:  &siswaPelenReq.DiterimaKelas,
		AlamatOrtu:     &siswaPelenReq.AlamatOrtu,
		TeleponOrtu:    &siswaPelenReq.TeleponOrtu,
		AlamatWali:     &siswaPelenReq.AlamatWali,
		TeleponWali:    &siswaPelenReq.TeleponWali,
		FotoSiswa:      &siswaPelenReq.FotoSiswa,
	}
	err := s.pesertaDidikService.Update(ctx, siswa, siswaPelenkap, schemaName)
	if err != nil {
		log.Printf("Gagal memperbarui siswa: %v", err)
		return nil, fmt.Errorf("gagal memperbarui siswa: %w", err)
	}

	return &pb.UpdateSiswaResponse{
		Message: "Siswa berhasil diperbarui",
		Status:  true,
	}, nil
}

// // **DeleteSiswa**
func (s *SiswaServiceServer) DeleteSiswa(ctx context.Context, req *pb.DeleteSiswaRequest) (*pb.DeleteSiswaResponse, error) {
	schemaName := req.GetSchemaname()
	siswaID := req.GetSiswaId()

	err := s.pesertaDidikService.Delete(ctx, siswaID, schemaName)
	if err != nil {
		log.Printf("Gagal menghapus siswa: %v", err)
		return nil, fmt.Errorf("gagal menghapus siswa: %w", err)
	}

	return &pb.DeleteSiswaResponse{
		Message: "Siswa berhasil dihapus",
		Status:  true,
	}, nil
}

// // UploadSiswa mengunggah data siswa dari file Excel
// func (s *SiswaServiceServer) UploadSiswa(ctx context.Context, req *pb.UploadSiswaRequest) (*pb.UploadSiswaResponse, error) {
// 	schemaName := req.GetSchemaname()
// 	fileData := req.GetFile() // File dalam bentuk byte array

// 	// Simpan file ke sementara
// 	tempFile := "/tmp/uploaded_siswa.xlsx"
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
// 	expectedHeaders := []string{"NIS", "NISN", "NamaSiswa", "TempatLahir", "TanggalLahir", "JenisKelamin", "Agama"}
// 	for i, expected := range expectedHeaders {
// 		if rows[0][i] != expected {
// 			return nil, fmt.Errorf("format kolom tidak sesuai, kolom '%s' seharusnya ada di posisi %d", expected, i+1)
// 		}
// 	}

// 	var siswaList []*models.PesertaDidik

// 	// Mulai dari baris kedua karena baris pertama adalah header
// 	for _, row := range rows[1:] {
// 		if len(row) < len(expectedHeaders) {
// 			log.Println("Skipping row due to insufficient data:", row)
// 			continue
// 		}

// 		// Konversi data sesuai dengan model
// 		namaSiswa := row[2]
// 		nis := row[0]
// 		nisn := row[1]
// 		tempatLahir := row[3]
// 		tanggalLahir := row[4]
// 		jenisKelamin := row[5]
// 		agama := row[6]

// 		// Validasi data
// 		if nis == "" || namaSiswa == "" || nisn == "" {
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
// 		siswa := &models.PesertaDidik{
// 			NIS:          strconv.Itoa(nisInt),
// 			NISN:         strconv.Itoa(nisnInt),
// 			NamaSiswa:    namaSiswa,
// 			TempatLahir:  tempatLahir,
// 			TanggalLahir: tanggalLahir,
// 			JenisKelamin: jenisKelamin,
// 			Agama:        agama,
// 		}
// 		siswaList = append(siswaList, siswa)
// 	}

// 	// Simpan ke database
// 	err = s.pesertaDidikService.BatchSave(ctx, siswaList, schemaName)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal menyimpan data siswa ke database: %w", err)
// 	}

// 	return &pb.UploadSiswaResponse{
// 		Message: "Siswa berhasil diunggah",
// 		Total:   int32(len(siswaList)),
// 		Status:  true,
// 	}, nil
// }
