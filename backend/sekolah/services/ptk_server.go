package services

import (
	"sekolah/config"
	pb "sekolah/generated"
	"sekolah/models"
	"sekolah/repositories"
)

type PTKServiceServer struct {
	pb.UnimplementedPTKServiceServer
	repo repositories.GenericRepository[models.TabelPTK]
}

func NewPTKServiceServer() *PTKServiceServer {
	repoPTK := repositories.NewPTKRepository(config.DB)
	return &PTKServiceServer{
		repo: *repoPTK,
	}
}

// **CreatePTK**
// func (s *PTKServiceServer) CreatePTK(ctx context.Context, req *pb.CreatePTKRequest) (*pb.CreatePTKResponse, error) {
// 	// Debugging: Cek nilai request yang diterima
// 	log.Printf("Received Sekolah data request: %+v\n", req)
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"SchemaName", "AnggotaKelas"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	schemaName := req.GetSchemaName()
// 	PTKReq := req.GetPTK()

// 	PTKModel := &models.TabelPTK{}

// 	err = s.repo.Save(ctx, PTKModel, schemaName)
// 	if err != nil {
// 		log.Printf("Gagal menyimpan PTK: %v", err)
// 		return nil, fmt.Errorf("gagal menyimpan PTK: %w", err)
// 	}

// 	return &pb.CreatePTKResponse{
// 		Message: "PTK berhasil ditambahkan",
// 		Status:  true,
// 	}, nil
// }

// **GetPTK**
// func (s *PTKServiceServer) GetPTK(ctx context.Context, req *pb.GetPTKRequest) (*pb.GetPTKResponse, error) {
// 	schemaName := req.GetSchemaname()
// 	PTKID := req.GetPTKId()

// 	PTK, err := s.PTKService.FindByID(ctx, PTKID, schemaName)
// 	if err != nil {
// 		log.Printf("Gagal menemukan PTK: %v", err)
// 		return nil, fmt.Errorf("gagal menemukan PTK: %w", err)
// 	}

// 	return &pb.GetPTKResponse{
// 		PTK: &pb.PTK{
// 			PTKID: PTK.PTKID,
// 		},
// 	}, nil
// }

// **UpdatePTK**
// func (s *PTKServiceServer) UpdatePTK(ctx context.Context, req *pb.UpdatePTKRequest) (*pb.UpdatePTKResponse, error) {
// 	// Debugging: Cek nilai request yang diterima
// 	log.Printf("Received UpdateUserProfile request: %+v\n", req)
// 	schemaName := req.GetSchemaname()
// 	PTKReq := req.GetPTK()
// 	PTKPelenReq := req.GetPTKPelengkap()
// 	PTK := &models.PTK{}
// 	PTKPelenkap := &models.PTKPelengkap{}
// 	err := s.PTKService.Update(ctx, PTK, PTKPelenkap, schemaName)
// 	if err != nil {
// 		log.Printf("Gagal memperbarui PTK: %v", err)
// 		return nil, fmt.Errorf("gagal memperbarui PTK: %w", err)
// 	}

// 	return &pb.UpdatePTKResponse{
// 		Message: "PTK berhasil diperbarui",
// 		Status:  true,
// 	}, nil
// }

// // // **DeletePTK**
// func (s *PTKServiceServer) DeletePTK(ctx context.Context, req *pb.DeletePTKRequest) (*pb.DeletePTKResponse, error) {
// 	schemaName := req.GetSchemaname()
// 	PTKID := req.GetPTKId()

// 	err := s.PTKService.Delete(ctx, PTKID, schemaName)
// 	if err != nil {
// 		log.Printf("Gagal menghapus PTK: %v", err)
// 		return nil, fmt.Errorf("gagal menghapus PTK: %w", err)
// 	}

// 	return &pb.DeletePTKResponse{
// 		Message: "PTK berhasil dihapus",
// 		Status:  true,
// 	}, nil
// }

// UploadPTK mengunggah data PTK dari file Excel
// func (s *PTKServiceServer) UploadPTK(ctx context.Context, req *pb.UploadPTKRequest) (*pb.UploadPTKResponse, error) {
// 	schemaName := req.GetSchemaname()
// 	fileData := req.GetFile() // File dalam bentuk byte array

// 	// Simpan file ke sementara
// 	tempFile := "/tmp/uploaded_PTK.xlsx"
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
// 	expectedHeaders := []string{"NIS", "NISN", "NamaPTK", "TempatLahir", "TanggalLahir", "JenisKelamin", "Agama"}
// 	for i, expected := range expectedHeaders {
// 		if rows[0][i] != expected {
// 			return nil, fmt.Errorf("format kolom tidak sesuai, kolom '%s' seharusnya ada di posisi %d", expected, i+1)
// 		}
// 	}

// 	var PTKList []*models.PTK

// 	// Mulai dari baris kedua karena baris pertama adalah header
// 	for _, row := range rows[1:] {
// 		if len(row) < len(expectedHeaders) {
// 			log.Println("Skipping row due to insufficient data:", row)
// 			continue
// 		}

// 		// Konversi data sesuai dengan model
// 		namaPTK := row[2]
// 		nis := row[0]
// 		nisn := row[1]
// 		tempatLahir := row[3]
// 		tanggalLahir := row[4]
// 		jenisKelamin := row[5]
// 		agama := row[6]

// 		// Validasi data
// 		if nis == "" || namaPTK == "" || nisn == "" {
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
// 		PTK := &models.PTK{
// 			NIS:          strconv.Itoa(nisInt),
// 			NISN:         strconv.Itoa(nisnInt),
// 			NamaPTK:    namaPTK,
// 			TempatLahir:  tempatLahir,
// 			TanggalLahir: tanggalLahir,
// 			JenisKelamin: jenisKelamin,
// 			Agama:        agama,
// 		}
// 		PTKList = append(PTKList, PTK)
// 	}

// 	// Simpan ke database
// 	err = s.PTKService.BatchSave(ctx, PTKList, schemaName)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal menyimpan data PTK ke database: %w", err)
// 	}

// 	return &pb.UploadPTKResponse{
// 		Message: "PTK berhasil diunggah",
// 		Total:   int32(len(PTKList)),
// 		Status:  true,
// 	}, nil
// }
