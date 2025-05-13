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

type TahunAjaranService struct {
	pb.UnimplementedTahunAjaranServiceServer
	repo repositories.TahunAjaranRepository
}

func NewTahunAjararanService() *TahunAjaranService {
	repoTahunAjaran := repositories.NewTahunAjaranRepository(config.DB)
	return &TahunAjaranService{
		repo: repoTahunAjaran,
	}
}

// **CreateTahunAjaran**
func (s *TahunAjaranService) CreateTahunAjaran(ctx context.Context, req *pb.CreateTahunAjaranRequest) (*pb.CreateTahunAjaranResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"TahunAjaran"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := "ref" //req.GetSchemaname()
	tahunAjaran := req.GetTahunAjaran()
	tahunAjaranModel := &models.TahunAjaran{
		TahunAjaranID:  tahunAjaran.TahunAjaranId,
		Nama:           tahunAjaran.Nama,
		PeriodeAktif:   tahunAjaran.PeriodeAktif,
		TanggalMulai:   tahunAjaran.TanggalMulai,
		TanggalSelesai: tahunAjaran.TanggalSelesai,
	}

	err = s.repo.Save(ctx, tahunAjaranModel, schemaName)
	if err != nil {
		log.Printf("Gagal menyimpan tahun ajaran: %v", err)
		return nil, fmt.Errorf("gagal menyimpan tahun ajaran: %w", err)
	}

	return &pb.CreateTahunAjaranResponse{
		Message: "Tahun ajaran berhasil ditambahkan",
		Status:  true,
	}, nil
}

// **GetTahunAjaran**
func (s *TahunAjaranService) GetTahunAjaran(ctx context.Context, req *pb.GetTahunAjaranRequest) (*pb.GetTahunAjaranResponse, error) {
	schemaName := "ref"
	// Cek apakah harus mengambil semua data atau data spesifik berdasarkan TahunAjaranId
	tahunAjaranID := req.GetTahunAjaranId()
	var err error
	var tahunAjaranModels []*models.TahunAjaran
	if tahunAjaranID == "" {
		// Ambil semua Tahun Ajaran
		tahunAjaranModels, err = s.repo.FindAll(ctx, schemaName, 100, 0)
		if err != nil {
			log.Printf("[ERROR] Gagal menemukan tahun ajaran di schema '%s': %v", schemaName, err)
			return nil, fmt.Errorf("gagal menemukan tahun ajaran di schema '%s': %w", schemaName, err)
		}

		// Konversi hasil ke response protobuf
		tahunAjaranList := utils.ConvertModelsToPB(tahunAjaranModels, func(item *models.TahunAjaran) *pb.TahunAjaran {
			return &pb.TahunAjaran{
				TahunAjaranId:  item.TahunAjaranID,
				Nama:           item.Nama,
				PeriodeAktif:   item.PeriodeAktif,
				TanggalMulai:   item.TanggalMulai,
				TanggalSelesai: item.TanggalSelesai,
			}
		})

		// Return response
		return &pb.GetTahunAjaranResponse{
			TahunAjaran: tahunAjaranList,
		}, nil
	}

	// Ambil data spesifik berdasarkan TahunAjaranId
	tahunAjaranModel, err := s.repo.FindByID(ctx, tahunAjaranID, schemaName)
	if err != nil {
		log.Printf("[ERROR] Gagal menemukan tahun ajaran dengan ID '%s' di schema '%s': %v", tahunAjaranID, schemaName, err)
		return nil, fmt.Errorf("gagal menemukan tahun ajaran dengan ID '%s': %w", tahunAjaranID, err)
	}
	tahunAjaranModels = append(tahunAjaranModels, tahunAjaranModel)
	// Return response untuk satu data
	return &pb.GetTahunAjaranResponse{
		TahunAjaran: utils.ConvertModelsToPB(tahunAjaranModels, func(item *models.TahunAjaran) *pb.TahunAjaran {
			return &pb.TahunAjaran{
				TahunAjaranId:  item.TahunAjaranID,
				Nama:           item.Nama,
				PeriodeAktif:   item.PeriodeAktif,
				TanggalMulai:   item.TanggalMulai,
				TanggalSelesai: item.TanggalSelesai,
			}
		}),
	}, nil
}

// **UpdateTahunAjaran**
func (s *TahunAjaranService) UpdateTahunAjaran(ctx context.Context, req *pb.UpdateTahunAjaranRequest) (*pb.UpdateTahunAjaranResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"TahunAjaran"}
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	tahunAjaranReq := req.GetTahunAjaran()
	tahunAjaranModel := &models.TahunAjaran{
		Nama:           tahunAjaranReq.Nama,
		PeriodeAktif:   tahunAjaranReq.PeriodeAktif,
		TanggalMulai:   tahunAjaranReq.TanggalMulai,
		TanggalSelesai: tahunAjaranReq.TanggalSelesai,
	}
	err = s.repo.Update(ctx, tahunAjaranModel, "ref")
	if err != nil {
		log.Printf("Gagal memperbarui tahun ajaran: %v", err)
		return &pb.UpdateTahunAjaranResponse{
			Message: fmt.Sprintf("Gagal memperbarui tahun ajaran: %v", err),
			Status:  false,
		}, nil
	}
	return &pb.UpdateTahunAjaranResponse{
		Message: "TahunAjaran berhasil diperbarui",
		Status:  true,
	}, nil
}

// // **DeleteTahunAjaran**
func (s *TahunAjaranService) DeleteTahunAjaran(ctx context.Context, req *pb.DeleteTahunAjaranRequest) (*pb.DeleteTahunAjaranResponse, error) {
	tahunAjaranID := req.GetTahunAjaranId()
	err := s.repo.Delete(ctx, tahunAjaranID, "ref")
	if err != nil {
		log.Printf("Gagal menghapus tahun ajaran: %v", err)
		// return nil, fmt.Errorf("gagal menghapus tahun ajaran: %w", err)
		return &pb.DeleteTahunAjaranResponse{
			Message: fmt.Sprintf("Gagal menghapus tahun ajaran: %v", err),
			Status:  false,
		}, nil
	}
	return &pb.DeleteTahunAjaranResponse{
		Message: "Tahun ajaran berhasil dihapus",
		Status:  true,
	}, nil
}

// func convertModelToPB(model *models.TahunAjaran) *pb.TahunAjaran {
// 	return &pb.TahunAjaran{
// 		TahunAjaranId:  model.TahunAjaranID,
// 		Nama:           model.Nama,
// 		PeriodeAktif:   model.PeriodeAktif,
// 		TanggalMulai:   model.TanggalMulai,
// 		TanggalSelesai: model.TanggalSelesai,
// 	}
// }
// func convertModelsToPB(models []*models.TahunAjaran) []*pb.TahunAjaran {
// 	var pbList []*pb.TahunAjaran
// 	for _, model := range models {
// 		pbList = append(pbList, convertModelToPB(model))
// 	}
// 	return pbList
// }
