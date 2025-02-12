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

type IjazahServiceServer struct {
	pb.UnimplementedIjazahServiceServer
	repo repositories.GenericRepository[models.Ijazah]
}

func NewIjazahService() *IjazahServiceServer {
	repoIjazah := repositories.NewIjazahRepository(config.DB)
	return &IjazahServiceServer{
		repo: *repoIjazah,
	}
}

// **CreateIjazah**
func (s *IjazahServiceServer) CreateIjazah(ctx context.Context, req *pb.CreateIjazahRequest) (*pb.CreateIjazahResponse, error) {

	return &pb.CreateIjazahResponse{
		Message: "ok",
		Status:  true,
	}, nil
}

// **GetIjazah**
func (s *IjazahServiceServer) GetIjazah(ctx context.Context, req *pb.GetIjazahRequest) (*pb.GetIjazahResponse, error) {
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

	ijazahId := req.GetIjazahId()
	// semesterId := req.GetSemesterId()
	var conditions = map[string]interface{}{
		"id": ijazahId,
	}

	if ijazahId != "" {
		// Ambil data ijazah berdasarkan RombonganBelajarId
		rombel, err := s.repo.FindByID(ctx, ijazahId, schemaName, "id")
		if err != nil {
			return nil, err
		}
		return &pb.GetIjazahResponse{
			Ijazah: []*pb.Ijazah{
				ConvertModelToPB(rombel, func(model *models.Ijazah) *pb.Ijazah {
					return &pb.Ijazah{
						Nama: model.Nama,
						Nisn: model.NISN,
						Nis:  model.Nis,
						Npsn: model.NPSN,
					}
				}),
			},
		}, nil
	}
	// Ambil semua data ijazah
	// limit := req.GetLimit()
	limit := 100
	if limit == 0 {
		limit = 100
	}
	// offset := req.GetOffset()
	offset := 0
	if offset == 0 {
		offset = 0
	}
	banyakijazah, err := s.repo.FindAllByConditions(ctx, schemaName, conditions, int(limit), int(offset))
	if err != nil {
		log.Printf("[ERROR] Gagal menemukan ijazah di schema '%s': %v", schemaName, err)
		return nil, fmt.Errorf("gagal menemukan ijazah di schema '%s': %w", schemaName, err)
	}
	banyakijazahList := ConvertModelsToPB(banyakijazah, func(model *models.Ijazah) *pb.Ijazah {
		return &pb.Ijazah{
			Nama: model.Nama,
		}
	})
	return &pb.GetIjazahResponse{
		Ijazah: banyakijazahList,
	}, nil
}

// **UpdateIjazah**
func (s *IjazahServiceServer) UpdateIjazah(ctx context.Context, req *pb.UpdateIjazahRequest) (*pb.UpdateIjazahResponse, error) {
	// Debugging: Cek nilai request yang diterima
	log.Printf("Received UpdateUserProfile request: %+v\n", req)
	schemaName := req.GetSchemaName()
	ijazahReq := req.GetIjazah()
	ijazahModel := &models.Ijazah{
		Nama:                        ijazahReq.Nama,
		Nis:                         ijazahReq.Nis,
		NISN:                        ijazahReq.Nisn,
		NPSN:                        ijazahReq.Npsn,
		NoIjazah:                    ijazahReq.NomorIjazah,
		TempatLahir:                 ijazahReq.TempatLahir,
		TanggalLahir:                ijazahReq.TanggalLahir,
		NamaOrtuWali:                ijazahReq.NamaOrtuwali,
		PaketKeahlian:               ijazahReq.PaketKeahlian,
		KabupatenKota:               ijazahReq.Kabupatenkota,
		Provinsi:                    ijazahReq.Provinsi,
		ProgramKeahlian:             ijazahReq.ProgramKeahlian,
		SekolahPenyelenggaraUjianUS: ijazahReq.SekolahPenyelenggaraUjianUs,
		SekolahPenyelenggaraUjianUN: ijazahReq.SekolahPenyelenggaraUjianUn,
		AsalSekolah:                 ijazahReq.AsalSekolah,
		NomorIjazah:                 ijazahReq.NomorIjazah,
		TempatIjazah:                ijazahReq.TempatIjazah,
		TanggalIjazah:               ijazahReq.TanggalIjazah,
	}
	err := s.repo.Update(ctx, ijazahModel, schemaName, "id", ijazahReq.ID)
	if err != nil {
		log.Printf("Gagal memperbarui Ijazah: %v", err)
		return nil, fmt.Errorf("gagal memperbarui Ijazah: %w", err)
	}
	return &pb.UpdateIjazahResponse{
		Message: "Ijazah berhasil diperbarui",
		Status:  true,
	}, nil
}

// **DeleteIjazah**
func (s *IjazahServiceServer) DeleteIjazah(ctx context.Context, req *pb.DeleteIjazahRequest) (*pb.DeleteIjazahResponse, error) {
	schemaName := req.GetSchemaName()
	IjazahID := req.GetIjazahId()

	err := s.repo.Delete(ctx, IjazahID, schemaName, "id")
	if err != nil {
		log.Printf("Gagal menghapus Ijazah: %v", err)
		return nil, fmt.Errorf("gagal menghapus Ijazah: %w", err)
	}

	return &pb.DeleteIjazahResponse{
		Message: "Ijazah berhasil dihapus",
		Status:  true,
	}, nil
}
