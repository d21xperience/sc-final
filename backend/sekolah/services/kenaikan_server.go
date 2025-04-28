package services

import (
	"context"
	"fmt"
	"sekolah/config"
	pb "sekolah/generated"
	"sekolah/models"
	"sekolah/repositories"
	"sekolah/utils"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type KenaikanServiceServer struct {
	pb.UnimplementedKenaikanServiceServer
	repo    repositories.GenericRepository[models.Kenaikan]
	repoPTK repositories.GenericRepository[models.TabelPTK]
}

func NewKenaikanServiceServer() *KenaikanServiceServer {
	repoKenaikan := repositories.NewKenaikanRepository(config.DB)
	repoPTK := repositories.NewPTKRepository(config.DB)
	return &KenaikanServiceServer{
		repo:    *repoKenaikan,
		repoPTK: *repoPTK,
	}
}

// **CreateKenaikan**
func (s *KenaikanServiceServer) CreateKenaikan(ctx context.Context, req *pb.CreateKenaikanRequest) (*pb.CreateKenaikanResponse, error) {
	// Debugging: Cek nilai request yang diterima
	// log.Printf("Received Sekolah data request: %+v\n", req)
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "AnggotaKelas"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	anggotaKelas := req.GetAnggotaKelas()
	// anggotaKelas := s.repo.FindAll()
	KenaikanModel := utils.ConvertPBToModels(anggotaKelas, func(item *pb.AnggotaKelas) *models.Kenaikan {
		return &models.Kenaikan{
			KdKenaikan:      uuid.New(),
			SemesterId:      item.SemesterId,
			AnggotaRombelId: utils.StringToUUID(item.AnggotaRombelId),
			PesertaDidikId:  utils.StringToUUID(item.PesertaDidikId),
			Kenaikan:        req.Kenaikan,
			Tingkat:         req.Tingkat,
		}
	})
	// simpan ke tabel_kenaikan
	err = s.repo.SaveMany(ctx, schemaName, KenaikanModel, 100)
	// conflicts1, err := s.repo.SaveManyWithConflictCheck(ctx, schemaName, KenaikanModel, "kdKenaikan", "peserta_didik_id", 100, []string{"peserta_didik_id"})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "insert failed: %v", err)
	}

	// conflictProto2 := repositories.ConvertConflictsToProto(conflicts1, "kdKenaikan", "PesertaDidikId")
	// conflictProto2 = append(conflictProto2, conflictProto...)
	// fmt.Print(conflictProto2)
	return &pb.CreateKenaikanResponse{
		Message: "Kenaikan berhasil ditambahkan",
		Status:  true,
		// Conflicts: &pb.ConflictResponse{
		// 	Message:       "Sebagian data berhasil disimpan",
		// 	Conflicts:     conflictProto2,
		// 	TotalConflict: int32(len(conflictProto2)),
		// },
	}, nil
}

// **GetKenaikan**
func (s *KenaikanServiceServer) GetKenaikan(ctx context.Context, req *pb.GetKenaikanRequest) (*pb.GetKenaikanResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "SemesterId", "TipeKenaikan"}
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
	// joins := []string{
	// 	"JOIN tabel_ptk ON tabel_ptk.ptk_id = tabel_ptk_terdaftar.ptk_id",
	// }
	preloads := []string{"AnggotaRombel", "AnggotaRombel.PesertaDidik", "AnggotaRombel.RombonganBelajar", "AnggotaRombel.NilaiAkhir"}

	conditions := map[string]interface{}{
		"tabel_kenaikan.semester_id": req.GetSemesterId(),
		"tabel_kenaikan.tingkat":     req.GetTipeKenaikan(),
	}
	// orderBy := []string{"nama ASC"}
	// exactConditions := []struct {
	// 	Query string
	// 	Args  []interface{}
	// }{
	// 	{"jenis_ptk_id != ?", []any{9}},
	// 	{"jenis_ptk_id != ?", []any{11}},
	// 	{"jenis_ptk_id != ?", []any{42}},
	// 	{"jenis_ptk_id != ?", []any{44}},
	// 	// {"created_at BETWEEN ? AND ?", []interface{}{startDate, endDate}},
	// }
	// groupByColumns := []string{"tabel_ptk_terdaftar.ptk_terdaftar_id"} // Hindari duplikasi
	// KenaikanModel, err := s.repo.FindWithPreloadAndJoinsOrigin(ctx, schemaName, joins, preloads, conditions, orderBy)
	KenaikanModel, err := s.repo.FindWithRelations(ctx, schemaName, nil, preloads, conditions, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	// Konversi ke protobuf
	KenaikanPB := utils.ConvertModelsToPB(KenaikanModel, func(item models.Kenaikan) *pb.Kenaikan {
		return &pb.Kenaikan{
			SemesterId:      item.SemesterId,
			AnggotaRombelId: item.AnggotaRombelId.String(),
			PesertaDidikId:  item.PesertaDidikId.String(),
			Kenaikan:        item.Kenaikan,
			Tingkat:         item.Tingkat,
			AnggotaKelas: &pb.AnggotaKelas{
				NmSiswa: item.AnggotaRombel.PesertaDidik.NmSiswa,
				NmKelas: item.AnggotaRombel.RombonganBelajar.NmKelas,
				PesertaDidik: &pb.Siswa{
					TempatLahir:  item.AnggotaRombel.PesertaDidik.TempatLahir,
					TanggalLahir: item.AnggotaRombel.PesertaDidik.TanggalLahir.Format("2016-02-01"),
					Nis:          item.AnggotaRombel.PesertaDidik.Nis,
					Nisn:         item.AnggotaRombel.PesertaDidik.Nisn,
					JenisKelamin: item.AnggotaRombel.PesertaDidik.JenisKelamin,
				},
				Nilai: utils.ConvertModelsToPB(item.AnggotaRombel.NilaiAkhir, func(nilai models.NilaiAkhir) *pb.NilaiAkhir {
					return &pb.NilaiAkhir{
						IdNilaiAkhir: nilai.IdNilaiAkhir.String(),
					}
				}),
			},
		}
	})

	return &pb.GetKenaikanResponse{
		Kenaikan: KenaikanPB,
	}, nil
}

// // **UpdateKenaikan**
// func (s *KenaikanServiceServer) UpdateKenaikan(ctx context.Context, req *pb.UpdateKenaikanRequest) (*pb.UpdateKenaikanResponse, error) {
// 	// Debugging: Cek nilai request yang diterima
// 	log.Printf("Received Sekolah data request: %+v\n", req)
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"Schemaname", "Kenaikan"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	schemaName := req.GetSchemaname()
// 	Kenaikan := req.Kenaikan

// 	KenaikanModel := &models.RombonganBelajar{
// 		NmKenaikan:             Kenaikan.NmKenaikan,
// 		SekolahId:           Kenaikan.SekolahId,
// 		SemesterId:          Kenaikan.SemesterId,
// 		JurusanId:           Kenaikan.JurusanId,
// 		TingkatPendidikanId: Kenaikan.TingkatPendidikanId,
// 		PtkId:               Kenaikan.PtkId,
// 		JenisKenaikan:         Kenaikan.JenisKenaikan,
// 		NamaJurusanSp:       Kenaikan.NamaJurusanSp,
// 		JurusanSpId:         Kenaikan.JurusanSpId,
// 		KurikulumId:         Kenaikan.KurikulumId,
// 		// RombonganBelajarId:  Kenaikan.RombonganBelajarId,
// 	}
// 	err = s.repo.Update(ctx, KenaikanModel, schemaName, "rombongan_belajar_id", Kenaikan.SemesterId)
// 	if err != nil {
// 		log.Printf("Gagal memperbaharui Kenaikan: %s", err)
// 		return nil, fmt.Errorf("gagal memperbaharui Kenaikan: %w", err)
// 	}
// 	return &pb.UpdateKenaikanResponse{
// 		Message: "Kenaikan berhasil diperbarui",
// 		Status:  true,
// 	}, nil
// }

// // **DeleteKenaikan**
// func (s *KenaikanServiceServer) DeleteKenaikan(ctx context.Context, req *pb.DeleteKenaikanRequest) (*pb.DeleteKenaikanResponse, error) {
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"Schemaname", "KenaikanId"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	schemaName := req.GetSchemaname()
// 	KenaikanID := req.GetKenaikanId()

// 	err = s.repo.Delete(ctx, KenaikanID, schemaName, "rombongan_belajar_id")
// 	if err != nil {
// 		log.Printf("Gagal menghapus Kenaikan: %s", err)
// 		return nil, fmt.Errorf("gagal menghapus Kenaikan: %w", err)
// 	}

// 	return &pb.DeleteKenaikanResponse{
// 		Message: "Kenaikan berhasil dihapus",
// 		Status:  true,
// 	}, nil
// }
