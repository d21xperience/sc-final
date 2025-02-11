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
	requiredFields := []string{"SchemaName", "TahunAjaran"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := "ref" //req.GetSchemaName()
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
	// Validasi SchemaName
	schemaName := "ref" //req.GetSchemaName()
	if schemaName == "" {
		return nil, fmt.Errorf("schema name is required")
	}

	// Cek apakah harus mengambil semua data atau data spesifik berdasarkan TahunAjaranId
	tahunAjaranID := req.GetTahunAjaranId()
	findAll := tahunAjaranID == ""

	if findAll {
		// Ambil semua Tahun Ajaran
		tahunAjaranModels, err := s.repo.FindAll(ctx, schemaName, int(req.GetLimit()), int(req.GetOffset()))
		if err != nil {
			log.Printf("[ERROR] Gagal menemukan tahun ajaran di schema '%s': %v", schemaName, err)
			return nil, fmt.Errorf("gagal menemukan tahun ajaran di schema '%s': %w", schemaName, err)
		}

		// Konversi hasil ke response protobuf
		tahunAjaranList := convertModelsToPB(tahunAjaranModels)

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

	// Return response untuk satu data
	return &pb.GetTahunAjaranResponse{
		TahunAjaran: []*pb.TahunAjaran{
			convertModelToPB(tahunAjaranModel),
		},
	}, nil
}

// **UpdateTahunAjaran**
func (s *TahunAjaranService) UpdateTahunAjaran(ctx context.Context, req *pb.UpdateTahunAjaranRequest) (*pb.UpdateTahunAjaranResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"schemaname", "tahun_ajaran"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaName()
	tahunAjaranReq := req.GetTahunAjaran()
	tahunAjaranModel := &models.TahunAjaran{
		Nama:           tahunAjaranReq.Nama,
		PeriodeAktif:   tahunAjaranReq.PeriodeAktif,
		TanggalMulai:   tahunAjaranReq.TanggalMulai,
		TanggalSelesai: tahunAjaranReq.TanggalSelesai,
	}
	err = s.repo.Update(ctx, tahunAjaranModel, schemaName)
	if err != nil {
		log.Printf("Gagal memperbarui tahun ajaran: %v", err)
		return nil, fmt.Errorf("gagal memperbarui tahun ajaran: %w", err)
	}
	return &pb.UpdateTahunAjaranResponse{
		Message: "TahunAjaran berhasil diperbarui",
		Status:  true,
	}, nil
}

// // **DeleteTahunAjaran**
func (s *TahunAjaranService) DeleteTahunAjaran(ctx context.Context, req *pb.DeleteTahunAjaranRequest) (*pb.DeleteTahunAjaranResponse, error) {
	schemaName := req.GetSchemaName()
	tahunAjaranID := req.GetTahunAjaranId()

	err := s.repo.Delete(ctx, tahunAjaranID, schemaName)
	if err != nil {
		log.Printf("Gagal menghapus tahun ajaran: %v", err)
		return nil, fmt.Errorf("gagal menghapus tahun ajaran: %w", err)
	}

	return &pb.DeleteTahunAjaranResponse{
		Message: "Tahun ajaran berhasil dihapus",
		Status:  true,
	}, nil
}

func convertModelToPB(model *models.TahunAjaran) *pb.TahunAjaran {
	return &pb.TahunAjaran{
		TahunAjaranId:  model.TahunAjaranID,
		Nama:           model.Nama,
		PeriodeAktif:   model.PeriodeAktif,
		TanggalMulai:   model.TanggalMulai,
		TanggalSelesai: model.TanggalSelesai,
	}
}
func convertModelsToPB(models []*models.TahunAjaran) []*pb.TahunAjaran {
	var pbList []*pb.TahunAjaran
	for _, model := range models {
		pbList = append(pbList, convertModelToPB(model))
	}
	return pbList
}
