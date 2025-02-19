package services

import (
	"context"
	"fmt"
	"sekolah/config"
	pb "sekolah/generated"
	"sekolah/models"
	"sekolah/repositories"
	"sekolah/utils"
)

type PTKTerdaftarServiceServer struct {
	pb.UnimplementedPTKTerdaftarServiceServer
	repo repositories.GenericRepository[models.PTKTerdaftar]
}

func NewPTKTerdaftarServiceServer() *PTKTerdaftarServiceServer {
	repoPTKTerdaftar := repositories.NewPTKTerdaftarRepository(config.DB)
	return &PTKTerdaftarServiceServer{
		repo: *repoPTKTerdaftar,
	}
}

// **CreatePTKTerdaftar**
// func (s *PTKTerdaftarServiceServer) CreatePTKTerdaftar(ctx context.Context, req *pb.CreatePTKTerdaftarRequest) (*pb.CreatePTKTerdaftarResponse, error) {
// 	// Debugging: Cek nilai request yang diterima
// 	log.Printf("Received Sekolah data request: %+v\n", req)
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"SchemaName", "PTKTerdaftar"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	schemaName := req.GetSchemaName()
// 	PTKTerdaftar := req.PTKTerdaftar

// 	PTKTerdaftarModel := &models.RombonganBelajar{
// 		NmPTKTerdaftar:             PTKTerdaftar.NmPTKTerdaftar,
// 		SekolahId:           PTKTerdaftar.SekolahId,
// 		SemesterId:          PTKTerdaftar.SemesterId,
// 		JurusanId:           PTKTerdaftar.JurusanId,
// 		TingkatPendidikanId: PTKTerdaftar.TingkatPendidikanId,
// 		PtkId:               PTKTerdaftar.PtkId,
// 		JenisPTKTerdaftar:         PTKTerdaftar.JenisPTKTerdaftar,
// 		NamaJurusanSp:       PTKTerdaftar.NamaJurusanSp,
// 		JurusanSpId:         PTKTerdaftar.JurusanSpId,
// 		KurikulumId:         PTKTerdaftar.KurikulumId,
// 	}

// 	err = s.repo.Save(ctx, PTKTerdaftarModel, schemaName)
// 	if err != nil {
// 		log.Printf("Gagal menyimpan PTKTerdaftar: %s", err)
// 		return nil, fmt.Errorf("gagal menyimpan PTKTerdaftar: %w", err)
// 	}

// 	return &pb.CreatePTKTerdaftarResponse{
// 		Message: "PTKTerdaftar berhasil ditambahkan",
// 		Status:  true,
// 	}, nil
// }
// func (s *PTKTerdaftarServiceServer) CreateBanyakPTKTerdaftar(ctx context.Context, req *pb.CreateBanyakPTKTerdaftarRequest) (*pb.CreateBanyakPTKTerdaftarResponse, error) {
// 	// Debugging: Cek nilai request yang diterima
// 	log.Printf("Received Sekolah data request: %+v\n", req)
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"SchemaName", "PTKTerdaftar"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	schemaName := req.GetSchemaName()
// 	PTKTerdaftar := req.PTKTerdaftar

// 	PTKTerdaftarModels := ConvertPBToModels(PTKTerdaftar, func(rom *pb.PTKTerdaftar) *models.RombonganBelajar {
// 		return &models.RombonganBelajar{
// 			RombonganBelajarId:  rom.RombonganBelajarId,
// 			SekolahId:           rom.SekolahId,
// 			SemesterId:          rom.SemesterId,
// 			JurusanId:           rom.JurusanId,
// 			PtkId:               rom.PtkId,
// 			NmPTKTerdaftar:             rom.NmPTKTerdaftar,
// 			TingkatPendidikanId: rom.TingkatPendidikanId,
// 			JenisPTKTerdaftar:         rom.JenisPTKTerdaftar,
// 			NamaJurusanSp:       rom.NamaJurusanSp,
// 			JurusanSpId:         rom.JurusanSpId,
// 			KurikulumId:         rom.KurikulumId,
// 		}
// 	})
// 	err = s.repo.SaveMany(ctx, schemaName, PTKTerdaftarModels, 100)
// 	if err != nil {
// 		log.Printf("Gagal menyimpan PTKTerdaftar: %s", err)
// 		return nil, fmt.Errorf("gagal menyimpan PTKTerdaftar: %w", err)
// 	}

// 	return &pb.CreateBanyakPTKTerdaftarResponse{
// 		Message: "PTKTerdaftar berhasil ditambahkan",
// 		Status:  true,
// 	}, nil
// }

// **GetPTKTerdaftar**
func (s *PTKTerdaftarServiceServer) GetPTKTerdaftar(ctx context.Context, req *pb.GetPTKTerdaftarRequest) (*pb.GetPTKTerdaftarResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"SchemaName", "TahunAjaranId"}
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
	// PTKTerdaftarId := req.GetPtkTerdaftarId()
	joins := []string{
		"JOIN tabel_ptk ON tabel_ptk.ptk_id = tabel_ptk_terdaftar.ptk_id",
		// "JOIN tabel_ptk ON tabel_kelas.ptk_id = tabel_ptk.ptk_id",
	}
	preloads := []string{"PTK"}

	conditions := map[string]interface{}{
		"tahun_ajaran_id": req.GetTahunAjaranId(),
	}

	PTKTerdaftarModel, err := s.repo.FindWithPreloadAndJoins(ctx, schemaName, joins, preloads, conditions)
	if err != nil {
		return nil, err
	}
	// Konversi ke protobuf
	ptkTerdaftarPB := utils.ConvertModelsToPB(PTKTerdaftarModel, func(ptk models.PTKTerdaftar) *pb.PTKTerdaftar {
		ptkTerdaftarId, err := utils.ConvertUUIDToStringViceVersa(ptk.PtkTerdaftarID)
		if err != nil {
			return nil
		}
		return &pb.PTKTerdaftar{
			PtkTerdaftarId: ptkTerdaftarId.(string),
			TahunAjaranId:  ptk.TahunAjaranID,
			Ptk: &pb.PTK{
				PtkId:             ptk.PTK.PtkID,
				Nama:              ptk.PTK.Nama,
				JenisKelamin:      ptk.PTK.JenisKelamin,
				JenisPtkId:        ptk.PTK.JenisPtkID,
				TempatLahir:       ptk.PTK.TempatLahir,
				TanggalLahir:      ptk.PTK.TanggalLahir,
				AlamatJalan:       ptk.PTK.AlamatJalan,
				StatusKeaktifanId: ptk.PTK.StatusKeaktifanID,
				Nuptk:             utils.SafeString(ptk.PTK.NUPTK),
				Nip:               utils.SafeString(ptk.PTK.NIP),
			},
			// Isi field sesuai kebutuhan
		}
	})

	return &pb.GetPTKTerdaftarResponse{
		PtkTerdaftar: ptkTerdaftarPB,
		Message:      "Sukses",
	}, nil
}

// // **UpdatePTKTerdaftar**
// func (s *PTKTerdaftarServiceServer) UpdatePTKTerdaftar(ctx context.Context, req *pb.UpdatePTKTerdaftarRequest) (*pb.UpdatePTKTerdaftarResponse, error) {
// 	// Debugging: Cek nilai request yang diterima
// 	log.Printf("Received Sekolah data request: %+v\n", req)
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"SchemaName", "PTKTerdaftar"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	schemaName := req.GetSchemaName()
// 	PTKTerdaftar := req.PTKTerdaftar

// 	PTKTerdaftarModel := &models.RombonganBelajar{
// 		NmPTKTerdaftar:             PTKTerdaftar.NmPTKTerdaftar,
// 		SekolahId:           PTKTerdaftar.SekolahId,
// 		SemesterId:          PTKTerdaftar.SemesterId,
// 		JurusanId:           PTKTerdaftar.JurusanId,
// 		TingkatPendidikanId: PTKTerdaftar.TingkatPendidikanId,
// 		PtkId:               PTKTerdaftar.PtkId,
// 		JenisPTKTerdaftar:         PTKTerdaftar.JenisPTKTerdaftar,
// 		NamaJurusanSp:       PTKTerdaftar.NamaJurusanSp,
// 		JurusanSpId:         PTKTerdaftar.JurusanSpId,
// 		KurikulumId:         PTKTerdaftar.KurikulumId,
// 		// RombonganBelajarId:  PTKTerdaftar.RombonganBelajarId,
// 	}
// 	err = s.repo.Update(ctx, PTKTerdaftarModel, schemaName, "rombongan_belajar_id", PTKTerdaftar.SemesterId)
// 	if err != nil {
// 		log.Printf("Gagal memperbaharui PTKTerdaftar: %s", err)
// 		return nil, fmt.Errorf("gagal memperbaharui PTKTerdaftar: %w", err)
// 	}
// 	return &pb.UpdatePTKTerdaftarResponse{
// 		Message: "PTKTerdaftar berhasil diperbarui",
// 		Status:  true,
// 	}, nil
// }

// // **DeletePTKTerdaftar**
// func (s *PTKTerdaftarServiceServer) DeletePTKTerdaftar(ctx context.Context, req *pb.DeletePTKTerdaftarRequest) (*pb.DeletePTKTerdaftarResponse, error) {
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"Schemaname", "PTKTerdaftarId"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	schemaName := req.GetSchemaName()
// 	PTKTerdaftarID := req.GetPTKTerdaftarId()

// 	err = s.repo.Delete(ctx, PTKTerdaftarID, schemaName, "rombongan_belajar_id")
// 	if err != nil {
// 		log.Printf("Gagal menghapus PTKTerdaftar: %s", err)
// 		return nil, fmt.Errorf("gagal menghapus PTKTerdaftar: %w", err)
// 	}

// 	return &pb.DeletePTKTerdaftarResponse{
// 		Message: "PTKTerdaftar berhasil dihapus",
// 		Status:  true,
// 	}, nil
// }
