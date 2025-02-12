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

type RombelServiceServer struct {
	pb.UnimplementedKelasServiceServer
	repo repositories.GenericRepository[models.RombonganBelajar]
}

func NewRombelServiceServer() *RombelServiceServer {
	repoRombel := repositories.NewrombonganBelajarRepository(config.DB)
	return &RombelServiceServer{
		repo: *repoRombel,
	}
}

// **CreateKelas**
func (s *RombelServiceServer) CreateKelas(ctx context.Context, req *pb.CreateKelasRequest) (*pb.CreateKelasResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"SchemaName", "Kelas"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaName()
	Kelas := req.Kelas

	KelasModel := &models.RombonganBelajar{
		NmKelas:             Kelas.NmKelas,
		SekolahId:           Kelas.SekolahId,
		SemesterId:          Kelas.SemesterId,
		JurusanId:           Kelas.JurusanId,
		TingkatPendidikanId: Kelas.TingkatPendidikanId,
		PtkId:               Kelas.PtkId,
		JenisRombel:         Kelas.JenisRombel,
		NamaJurusanSp:       Kelas.NamaJurusanSp,
		JurusanSpId:         Kelas.JurusanSpId,
		KurikulumId:         Kelas.KurikulumId,
	}

	err = s.repo.Save(ctx, KelasModel, schemaName)
	if err != nil {
		log.Printf("Gagal menyimpan Kelas: %v", err)
		return nil, fmt.Errorf("gagal menyimpan Kelas: %w", err)
	}

	return &pb.CreateKelasResponse{
		Message: "Kelas berhasil ditambahkan",
		Status:  true,
	}, nil
}
func (s *RombelServiceServer) CreateBanyakKelas(ctx context.Context, req *pb.CreateBanyakKelasRequest) (*pb.CreateBanyakKelasResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"SchemaName", "Kelas"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaName()
	kelas := req.Kelas

	kelasModels := ConvertPBToModels(kelas, func(rom *pb.Kelas) *models.RombonganBelajar {
		return &models.RombonganBelajar{
			RombonganBelajarId:  rom.RombonganBelajarId,
			SekolahId:           rom.SekolahId,
			SemesterId:          rom.SemesterId,
			JurusanId:           rom.JurusanId,
			PtkId:               rom.PtkId,
			NmKelas:             rom.NmKelas,
			TingkatPendidikanId: rom.TingkatPendidikanId,
			JenisRombel:         rom.JenisRombel,
			NamaJurusanSp:       rom.NamaJurusanSp,
			JurusanSpId:         rom.JurusanSpId,
			KurikulumId:         rom.KurikulumId,
		}
	})
	err = s.repo.SaveMany(ctx, schemaName, kelasModels, 100)
	if err != nil {
		log.Printf("Gagal menyimpan Kelas: %v", err)
		return nil, fmt.Errorf("gagal menyimpan Kelas: %w", err)
	}

	return &pb.CreateBanyakKelasResponse{
		Message: "Kelas berhasil ditambahkan",
		Status:  true,
	}, nil
}

// **GetKelas**
func (s *RombelServiceServer) GetKelas(ctx context.Context, req *pb.GetKelasRequest) (*pb.GetKelasResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaName()
	if schemaName == "\"\"" {
		return nil, fmt.Errorf("schema name is required")
	}

	// Cek apakah harus mengambil semua data atau data spesifik berdasarkan SemesterId
	kelasId := req.GetKelasId()
	semesterId := req.GetSemesterId()
	var conditions = map[string]interface{}{
		"semester_id": semesterId,
	}

	if kelasId != "" {
		// Ambil data Kelas berdasarkan RombonganBelajarId
		rombel, err := s.repo.FindByID(ctx, kelasId, schemaName, "semester_id")
		if err != nil {
			return nil, err
		}
		return &pb.GetKelasResponse{
			Kelas: []*pb.Kelas{
				ConvertModelToPB(rombel, func(rom *models.RombonganBelajar) *pb.Kelas {
					return &pb.Kelas{
						RombonganBelajarId:  rom.RombonganBelajarId,
						SekolahId:           rom.SekolahId,
						SemesterId:          rom.SemesterId,
						JurusanId:           rom.JurusanId,
						PtkId:               rom.PtkId,
						NmKelas:             rom.NmKelas,
						TingkatPendidikanId: rom.TingkatPendidikanId,
						JenisRombel:         rom.JenisRombel,
						NamaJurusanSp:       rom.NamaJurusanSp,
						JurusanSpId:         rom.JurusanSpId,
						KurikulumId:         rom.KurikulumId,
					}
				}),
			},
		}, nil
	}
	// Ambil semua data Kelas
	limit := req.GetLimit()
	if limit == 0 {
		limit = 100
	}
	offset := req.GetOffset()
	if offset == 0 {
		offset = 0
	}
	banyakKelas, err := s.repo.FindAllByConditions(ctx, schemaName, conditions, int(limit), int(req.GetOffset()))
	if err != nil {
		log.Printf("[ERROR] Gagal menemukan Kelas di schema '%s': %v", schemaName, err)
		return nil, fmt.Errorf("gagal menemukan Kelas di schema '%s': %w", schemaName, err)
	}
	banyakKelasList := ConvertModelsToPB(banyakKelas, func(kelas *models.RombonganBelajar) *pb.Kelas {
		return &pb.Kelas{
			RombonganBelajarId:  kelas.RombonganBelajarId,
			SekolahId:           kelas.SekolahId,
			SemesterId:          kelas.SemesterId,
			JurusanId:           kelas.JurusanId,
			PtkId:               kelas.PtkId,
			NmKelas:             kelas.NmKelas,
			TingkatPendidikanId: kelas.TingkatPendidikanId,
			JenisRombel:         kelas.JenisRombel,
			NamaJurusanSp:       kelas.NamaJurusanSp,
			JurusanSpId:         kelas.JurusanSpId,
			KurikulumId:         kelas.KurikulumId,
		}
	})
	return &pb.GetKelasResponse{
		Kelas: banyakKelasList,
	}, nil
}

// **UpdateKelas**
func (s *RombelServiceServer) UpdateKelas(ctx context.Context, req *pb.UpdateKelasRequest) (*pb.UpdateKelasResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"SchemaName", "Kelas"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaName()
	Kelas := req.Kelas

	KelasModel := &models.RombonganBelajar{
		NmKelas:             Kelas.NmKelas,
		SekolahId:           Kelas.SekolahId,
		SemesterId:          Kelas.SemesterId,
		JurusanId:           Kelas.JurusanId,
		TingkatPendidikanId: Kelas.TingkatPendidikanId,
		PtkId:               Kelas.PtkId,
		JenisRombel:         Kelas.JenisRombel,
		NamaJurusanSp:       Kelas.NamaJurusanSp,
		JurusanSpId:         Kelas.JurusanSpId,
		KurikulumId:         Kelas.KurikulumId,
		// RombonganBelajarId:  kelas.RombonganBelajarId,
	}
	err = s.repo.Update(ctx, KelasModel, schemaName, "rombongan_belajar_id", Kelas.SemesterId)
	if err != nil {
		log.Printf("Gagal memperbaharui Kelas: %v", err)
		return nil, fmt.Errorf("gagal memperbaharui Kelas: %w", err)
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
	schemaName := req.GetSchemaName()
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

// UploadKelas mengunggah data Kelas dari file Excel
// func (s *RombelServiceServer) UploadKelas(ctx context.Context, req *pb.UploadKelasRequest) (*pb.UploadKelasResponse, error) {
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

// 	var KelasList []*models.RombonganBelajar

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
// 		Kelas := &models.RombonganBelajar{
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
// 	err = s.repo.BatchSave(ctx, KelasList, schemaName)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal menyimpan data Kelas ke database: %w", err)
// 	}

// 	return &pb.UploadKelasResponse{
// 		Message: "Kelas berhasil diunggah",
// 		Total:   int32(len(KelasList)),
// 		Status:  true,
// 	}, nil
// }
