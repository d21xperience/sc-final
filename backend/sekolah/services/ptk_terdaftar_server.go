package services

import (
	"context"
	"fmt"
	"sekolah/config"
	pb "sekolah/generated"
	"sekolah/models"
	"sekolah/repositories"
	"sekolah/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PTKTerdaftarServiceServer struct {
	pb.UnimplementedPTKTerdaftarServiceServer
	repo    repositories.GenericRepository[models.PTKTerdaftar]
	repoPTK repositories.GenericRepository[models.TabelPTK]
}

func NewPTKTerdaftarServiceServer() *PTKTerdaftarServiceServer {
	repoPTKTerdaftar := repositories.NewPTKTerdaftarRepository(config.DB)
	repoPTK := repositories.NewPTKRepository(config.DB)
	return &PTKTerdaftarServiceServer{
		repo:    *repoPTKTerdaftar,
		repoPTK: *repoPTK,
	}
}

// **CreatePTKTerdaftar**
func (s *PTKTerdaftarServiceServer) CreatePTKTerdaftar(ctx context.Context, req *pb.CreatePTKTerdaftarRequest) (*pb.CreatePTKTerdaftarResponse, error) {
	// Debugging: Cek nilai request yang diterima
	// log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "PTKTerdaftar"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	PTKTerdaftar := req.GetPtkTerdaftar()

	PTKTerdaftarModel := utils.ConvertPBToModels(PTKTerdaftar, func(item *pb.PTKTerdaftar) *models.PTKTerdaftar {
		return &models.PTKTerdaftar{
			PtkTerdaftarId: utils.StringToUUID(item.PtkTerdaftarId),
			TahunAjaranId:  item.TahunAjaranId,
			PtkID:          utils.StringToUUID(item.PtkId),
			JenisKeluarId:  &item.JenisKeluarId,
			PTK: models.TabelPTK{
				PtkID:             utils.StringToUUID(item.PtkId),
				Nama:              item.Ptk.Nama,
				NIP:               &item.Ptk.Nip,
				JenisPtkID:        item.Ptk.JenisPtkId,
				JenisKelamin:      item.Ptk.JenisKelamin,
				TempatLahir:       item.Ptk.TempatLahir,
				TanggalLahir:      utils.TimeToPointer(item.Ptk.TanggalLahir),
				NUPTK:             &item.Ptk.Nuptk,
				AlamatJalan:       item.Ptk.AlamatJalan,
				StatusKeaktifanID: item.Ptk.StatusKeaktifanId,
			},
		}
	})

	// var daftarPTK []models.TabelPTK
	// for _, v := range PTKTerdaftarModel {
	// 	daftarPTK = append(daftarPTK, v.PTK)
	// }
	// conflicts, err := s.repoPTK.SaveManyWithConflictCheck(ctx, schemaName, utils.SliceToPointer(daftarPTK), "PtkID", "ptk_id", 100)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "insert failed: %v", err)
	// }

	// conflictProto := repositories.ConvertConflictsToProto(conflicts, "PtkID", "Nama")

	// simpan ke tabel_ptk_terdaftar
	conflicts1, err := s.repo.SaveManyWithConflictCheck(ctx, schemaName, PTKTerdaftarModel, "PtkTerdaftarId", "ptk_terdaftar_id", 5)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "insert failed: %v", err)
	}

	conflictProto2 := repositories.ConvertConflictsToProto(conflicts1, "PtkTerdaftarId", "PtkID")
	// conflictProto2 = append(conflictProto2, conflictProto...)
	fmt.Print(conflictProto2)
	return &pb.CreatePTKTerdaftarResponse{
		Message: "PTKTerdaftar berhasil ditambahkan",
		Status:  true,
		Conflicts: &pb.ConflictResponse{
			Message:       "Sebagian data berhasil disimpan",
			Conflicts:     conflictProto2,
			TotalConflict: int32(len(conflictProto2)),
		},
	}, nil
}

// func (s *PTKTerdaftarServiceServer) CreateBanyakPTKTerdaftar(ctx context.Context, req *pb.CreateBanyakPTKTerdaftarRequest) (*pb.CreateBanyakPTKTerdaftarResponse, error) {
// 	// Debugging: Cek nilai request yang diterima
// 	log.Printf("Received Sekolah data request: %+v\n", req)
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"Schemaname", "PTKTerdaftar"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	schemaName := req.GetSchemaname()
// 	PTKTerdaftar := req.PTKTerdaftar

// 	PTKTerdaftarModels := ConvertPBToModels(PTKTerdaftar, func(rom *pb.PTKTerdaftar) *models.RombonganBelajar {
// 		return &models.RombonganBelajar{
// 			RombonganBelajarId:  rom.RombonganBelajarId,
// 			SekolahId:           rom.SekolahId,
// 			SemesterId:          rom.SemesterId,
// 			JurusanId:           rom.JurusanId,
// 			PtkId:               rom.PtkId,
// 			NmPTKTerdaftar:      rom.NmPTKTerdaftar,
// 			TingkatPendidikanId: rom.TingkatPendidikanId,
// 			JenisPTKTerdaftar:   rom.JenisPTKTerdaftar,
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
	requiredFields := []string{"Schemaname", "TahunAjaranId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	if schemaName == "\"\"" {
		return nil, fmt.Errorf("schema name is required")
	}

	// Cek apakah harus mengambil semua data atau data spesifik berdasarkan SemesterId
	joins := []string{
		"JOIN tabel_ptk ON tabel_ptk.ptk_id = tabel_ptk_terdaftar.ptk_id",
	}
	preloads := []string{"PTK"}

	conditions := map[string]interface{}{
		"tahun_ajaran_id": req.GetTahunAjaranId(),
	}
	orderBy := []string{"nama ASC"}
	// groupByColumns := []string{"tabel_ptk_terdaftar.ptk_terdaftar_id"} // Hindari duplikasi
	PTKTerdaftarModel, err := s.repo.FindWithPreloadAndJoinsOrigin(ctx, schemaName, joins, preloads, conditions, orderBy)
	if err != nil {
		return nil, err
	}
	// Konversi ke protobuf
	ptkTerdaftarPB := utils.ConvertModelsToPB(PTKTerdaftarModel, func(ptk models.PTKTerdaftar) *pb.PTKTerdaftar {
		ptkTerdaftarId, err := utils.ConvertUUIDToStringViceVersa(ptk.PtkTerdaftarId)
		if err != nil {
			return nil
		}
		// tglLahir, err := ptk.PTK.TanggalLahir, "2006-01-02")
		// if err != nil {
		// 	return nil
		// }
		return &pb.PTKTerdaftar{
			PtkTerdaftarId: ptkTerdaftarId.(string),
			TahunAjaranId:  ptk.TahunAjaranId,
			Ptk: &pb.PTK{
				PtkId:             ptk.PTK.PtkID.String(),
				Nama:              ptk.PTK.Nama,
				JenisKelamin:      ptk.PTK.JenisKelamin,
				JenisPtkId:        ptk.PTK.JenisPtkID,
				TempatLahir:       ptk.PTK.TempatLahir,
				TanggalLahir:      ptk.PTK.TanggalLahir.Format("2006-01-02"),
				AlamatJalan:       ptk.PTK.AlamatJalan,
				StatusKeaktifanId: ptk.PTK.StatusKeaktifanID,
				Nuptk:             utils.SafeString(ptk.PTK.NUPTK),
				Nip:               utils.SafeString(ptk.PTK.NIP),
			},
			// PtkPelengkap: &pb.PTKPelengkap{
			// 	PtkPelengkapId: ptk.PTKPelengkap.PTKID,
			// 	GelarDepan: ptk.PTK.,
			// },
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
// 	requiredFields := []string{"Schemaname", "PTKTerdaftar"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	schemaName := req.GetSchemaname()
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
// 	schemaName := req.GetSchemaname()
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
