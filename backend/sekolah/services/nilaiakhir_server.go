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

type NilaiAkhirServiceServer struct {
	pb.UnimplementedNilaiAkhirServiceServer
	repo              repositories.GenericRepository[models.NilaiAkhir]
	repoRombelAnggota repositories.GenericRepository[models.RombelAnggota]
}

func NewNilaiAkhirServiceServer() *NilaiAkhirServiceServer {
	repoNilaiAkhir := repositories.NewNilaiAkhirRepository(config.DB)
	repoRombelAnggota := repositories.NewRombelAnggotaRepository(config.DB)
	return &NilaiAkhirServiceServer{
		repo:              *repoNilaiAkhir,
		repoRombelAnggota: *repoRombelAnggota,
	}
}

func (s *NilaiAkhirServiceServer) CreateNilaiAkhir(ctx context.Context, req *pb.CreateNilaiAkhirRequest) (*pb.CreateNilaiAkhirResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "NilaiAkhir"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	nilaiAkhirReq := req.GetNilaiAkhir()

	nilaiAkhir := ConvertPBToModels(nilaiAkhirReq, func(nilai *pb.NilaiAkhir) *models.NilaiAkhir {
		anggotaRombelID, err := uuid.Parse(nilai.AnggotaRombelId)
		if err != nil {
			log.Printf("Invalid UUID for AnggotaRombelID: %v", err)
		}
		pesertaDidikID, err := uuid.Parse(nilai.AnggotaRombelId)
		if err != nil {
			log.Printf("Invalid UUID for AnggotaRombelID: %v", err)
		}
		return &models.NilaiAkhir{
			IdNilaiAkhir:    uuid.New(),
			AnggotaRombelId: anggotaRombelID,
			MataPelajaranId: &nilai.MataPelajaranId,
			SemesterId:      nilai.SemesterId,
			NilaiPeng:       &nilai.NilaiPeng,
			PredikatPeng:    nilai.PredikatPeng,
			NilaiKet:        &nilai.NilaiKet,
			PredikatKet:     nilai.PredikatKet,
			NilaiSik:        &nilai.NilaiSik,
			PredikatSik:     nilai.PredikatSik,
			NilaiSikSos:     &nilai.NilaiSiksos,
			PredikatSikSos:  nilai.PredikatSiksos,
			PesertaDidikId:  pesertaDidikID,
			IDMinat:         nilai.IdMinat,
		}
	})
	err = s.repo.SaveMany(ctx, schemaName, nilaiAkhir, 100)
	if err != nil {
		log.Printf("Gagal menyimpan Nilai akhir: %v", err)
		return nil, fmt.Errorf("gagal menyimpan Nilai akhir: %w", err)
	}

	return &pb.CreateNilaiAkhirResponse{
		Message: "Nilai akhir berhasil ditambahkan",
		Status:  true,
	}, nil
}

// **GetNilai akhir**
func (s *NilaiAkhirServiceServer) GetNilaiAkhir(ctx context.Context, req *pb.GetNilaiAkhirRequest) (*pb.GetNilaiAkhirResponse, error) {
	// Debugging: Cek nilai request yang diterima
	// log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "SemesterId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	semesterId := req.GetSemesterId()

	var conditions = map[string]interface{}{
		"tabel_anggotakelas.semester_id": semesterId,
		"tabel_kelas.jenis_rombel":       1,
	}
	if req.GetPesertaDidikId() != "" {
		conditions["tabel_anggotakelas.peserta_didik_id"] = req.GetPesertaDidikId()
	}
	joins := []string{
		"JOIN tabel_siswa ON tabel_siswa.peserta_didik_id = tabel_anggotakelas.peserta_didik_id",
		"JOIN tabel_kelas ON tabel_kelas.rombongan_belajar_id = tabel_anggotakelas.rombongan_belajar_id",
		"JOIN tabel_ptk ON tabel_ptk.ptk_id = tabel_kelas.ptk_id",
	}
	// preloads := []string{"PesertaDidik", "RombonganBelajar", "RombonganBelajar.PTK", "NilaiAkhir", "NilaiAkhir.MataPelajaran"}
	preloads := []string{"PesertaDidik", "RombonganBelajar", "RombonganBelajar.PTK"}

	orderBy := []string{"tabel_kelas.nm_kelas ASC", "tabel_siswa.nm_siswa ASC"} // Hindari duplikasi
	anggotaRombelModel, err := s.repoRombelAnggota.FindWithPreloadAndJoinsOrigin(ctx, schemaName, joins, preloads, conditions, orderBy)
	if err != nil {
		return nil, err
	}

	NilaiAkhirList := utils.ConvertModelsToPB(anggotaRombelModel, func(item models.RombelAnggota) *pb.NilaiSiswa {
		// nilaiAkhir, err := s.repo.FindWithRelations(ctx, schemaName, nil, []string{"MataPelajaran"}, map[string]any{"peserta_didik_id": item.PesertaDidikId.String()}, []struct {
		// 	Query string
		// 	Args  []interface{}
		// }{
		// 	{"semester_id <= ?", []interface{}{item.SemesterId}},
		// 	// {"created_at BETWEEN ? AND ?", []interface{}{startDate, endDate}},
		// }, nil, nil)
		// if err != nil {
		// 	return nil
		// }

		return &pb.NilaiSiswa{
			PesertaDidikId:      item.PesertaDidikId.String(),
			NmSiswa:             item.PesertaDidik.NmSiswa,
			RombonganBelajarId:  item.AnggotaRombelId.String(),
			NmKelas:             item.RombonganBelajar.NmKelas,
			TingkatPendidikanId: item.RombonganBelajar.TingkatPendidikanId,
			// NilaiAkhir: utils.ConvertModelsToPB(nilaiAkhir, func(item models.NilaiAkhir) *pb.NilaiAkhir {
			// 	return &pb.NilaiAkhir{
			// 		IdNilaiAkhir:    item.IdNilaiAkhir.String(),
			// 		NilaiPeng:       utils.PointerToUint32(item.NilaiPeng),
			// 		MataPelajaranId: utils.PointerToUint32(item.MataPelajaranId),
			// 		SemesterId:      item.SemesterId,
			// 		Semester:        utils.PointerToUint32(item.Semester),
			// 		MataPelajaran: &pb.Mapel{
			// 			Nama: item.MataPelajaran.Nama,
			// 		},
			// 	}
			// }),
		}
	})
	return &pb.GetNilaiAkhirResponse{
		NilaiSiswa: NilaiAkhirList,
	}, nil
}

// **UpdateNilai akhir**
// func (s *NilaiAkhirServiceServer) UpdateNilai akhir(ctx context.Context, req *pb.UpdateNilai akhirRequest) (*pb.UpdateNilai akhirResponse, error) {
// 	// Debugging: Cek nilai request yang diterima
// 	log.Printf("Received UpdateUserProfile request: %+v\n", req)
// 	schemaName := req.GetSchemaname()
// 	Nilai akhirReq := req.GetNilai akhir()
// 	Nilai akhirPelenReq := req.GetNilai akhirPelengkap()
// 	Nilai akhir := &models.PesertaDidik{
// 		PesertaDidikID:  Nilai akhirReq.PesertaDidikID,
// 		NIS:             Nilai akhirReq.NIS,
// 		NISN:            Nilai akhirReq.NISN,
// 		NamaNilai akhir:       Nilai akhirReq.NamaNilai akhir,
// 		TempatLahir:     Nilai akhirReq.TempatLahir,
// 		TanggalLahir:    Nilai akhirReq.TanggalLahir,
// 		JenisKelamin:    Nilai akhirReq.JenisKelamin,
// 		Agama:           Nilai akhirReq.Agama,
// 		AlamatNilai akhir:     &Nilai akhirReq.AlamatNilai akhir,
// 		TeleponNilai akhir:    Nilai akhirReq.TeleponNilai akhir,
// 		DiterimaTanggal: Nilai akhirReq.DiterimaTanggal,
// 		NamaAyah:        Nilai akhirReq.NamaAyah,
// 		NamaIbu:         Nilai akhirReq.NamaIbu,
// 		PekerjaanAyah:   Nilai akhirReq.PekerjaanAyah,
// 		PekerjaanIbu:    Nilai akhirReq.PekerjaanIbu,
// 		NamaWali:        &Nilai akhirReq.NamaWali,
// 		PekerjaanWali:   &Nilai akhirReq.PekerjaanWali,
// 	}
// 	Nilai akhirPelenkap := &models.PesertaDidikPelengkap{
// 		PelengkapNilai akhirID: Nilai akhirPelenReq.PelengkapNilai akhirID,
// 		PesertaDidikID:   &Nilai akhirPelenReq.PesertaDidikID,
// 		StatusDalamKel:   &Nilai akhirPelenReq.StatusDalamKel,
// 		AnakKe:           &Nilai akhirPelenReq.AnakKe,
// 		SekolahAsal:      Nilai akhirPelenReq.SekolahAsal,
// 		DiterimaNilai akhir:    &Nilai akhirPelenReq.DiterimaNilai akhir,
// 		AlamatOrtu:       &Nilai akhirPelenReq.AlamatOrtu,
// 		TeleponOrtu:      &Nilai akhirPelenReq.TeleponOrtu,
// 		AlamatWali:       &Nilai akhirPelenReq.AlamatWali,
// 		TeleponWali:      &Nilai akhirPelenReq.TeleponWali,
// 		FotoNilai akhir:        &Nilai akhirPelenReq.FotoNilai akhir,
// 	}
// 	err := s.pesertaDidikService.Update(ctx, Nilai akhir, Nilai akhirPelenkap, schemaName)
// 	if err != nil {
// 		log.Printf("Gagal memperbarui Nilai akhir: %v", err)
// 		return nil, fmt.Errorf("gagal memperbarui Nilai akhir: %w", err)
// 	}

// 	return &pb.UpdateNilai akhirResponse{
// 		Message: "Nilai akhir berhasil diperbarui",
// 		Status:  true,
// 	}, nil
// }

// // // **DeleteNilai akhir**
// func (s *NilaiAkhirServiceServer) DeleteNilai akhir(ctx context.Context, req *pb.DeleteNilai akhirRequest) (*pb.DeleteNilai akhirResponse, error) {
// 	schemaName := req.GetSchemaname()
// 	Nilai akhirID := req.GetNilai akhirId()

// 	err := s.pesertaDidikService.Delete(ctx, Nilai akhirID, schemaName)
// 	if err != nil {
// 		log.Printf("Gagal menghapus Nilai akhir: %v", err)
// 		return nil, fmt.Errorf("gagal menghapus Nilai akhir: %w", err)
// 	}

// 	return &pb.DeleteNilai akhirResponse{
// 		Message: "Nilai akhir berhasil dihapus",
// 		Status:  true,
// 	}, nil
// }

// // UploadNilai akhir mengunggah data Nilai akhir dari file Excel
// func (s *NilaiAkhirServiceServer) UploadNilai akhir(ctx context.Context, req *pb.UploadNilai akhirRequest) (*pb.UploadNilai akhirResponse, error) {
// 	schemaName := req.GetSchemaname()
// 	fileData := req.GetFile() // File dalam bentuk byte array

// 	// Simpan file ke sementara
// 	tempFile := "/tmp/uploaded_Nilai akhir.xlsx"
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
// 	expectedHeaders := []string{"NIS", "NISN", "NamaNilai akhir", "TempatLahir", "TanggalLahir", "JenisKelamin", "Agama"}
// 	for i, expected := range expectedHeaders {
// 		if rows[0][i] != expected {
// 			return nil, fmt.Errorf("format kolom tidak sesuai, kolom '%s' seharusnya ada di posisi %d", expected, i+1)
// 		}
// 	}

// 	var Nilai akhirList []*models.PesertaDidik

// 	// Mulai dari baris kedua karena baris pertama adalah header
// 	for _, row := range rows[1:] {
// 		if len(row) < len(expectedHeaders) {
// 			log.Println("Skipping row due to insufficient data:", row)
// 			continue
// 		}

// 		// Konversi data sesuai dengan model
// 		namaNilai akhir := row[2]
// 		nis := row[0]
// 		nisn := row[1]
// 		tempatLahir := row[3]
// 		tanggalLahir := row[4]
// 		jenisKelamin := row[5]
// 		agama := row[6]

// 		// Validasi data
// 		if nis == "" || namaNilai akhir == "" || nisn == "" {
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
// 		Nilai akhir := &models.PesertaDidik{
// 			NIS:          strconv.Itoa(nisInt),
// 			NISN:         strconv.Itoa(nisnInt),
// 			NamaNilai akhir:    namaNilai akhir,
// 			TempatLahir:  tempatLahir,
// 			TanggalLahir: tanggalLahir,
// 			JenisKelamin: jenisKelamin,
// 			Agama:        agama,
// 		}
// 		Nilai akhirList = append(Nilai akhirList, Nilai akhir)
// 	}

// 	// Simpan ke database
// 	err = s.pesertaDidikService.BatchSave(ctx, Nilai akhirList, schemaName)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal menyimpan data Nilai akhir ke database: %w", err)
// 	}

// 	return &pb.UploadNilai akhirResponse{
// 		Message: "Nilai akhir berhasil diunggah",
// 		Total:   int32(len(Nilai akhirList)),
// 		Status:  true,
// 	}, nil
// }
