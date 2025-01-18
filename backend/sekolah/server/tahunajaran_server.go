package server

import (
	"context"
	"fmt"
	"log"
	pb "sekolah/generated"
	"sekolah/models"
	"sekolah/services"
	"sekolah/utils"
)

type TahunAjaranServiceServer struct {
	pb.UnimplementedTahunAjaranServiceServer
	TahunAjaranService services.TahunAjaranService
}

// **CreateTahunAjaran**
func (s *TahunAjaranServiceServer) CreateTahunAjaran(ctx context.Context, req *pb.CreateTahunAjaranRequest) (*pb.CreateTahunAjaranResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"schemaname", "tahun_ajaran"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	tahunAjaran := req.GetTahunAjaran()
	tahunAjaranModel := &models.TahunAjaran{
		Nama:           tahunAjaran.Nama,
		PeriodeAktif:   tahunAjaran.PeriodeAktif,
		TanggalMulai:   tahunAjaran.TanggalMulai,
		TanggalSelesai: tahunAjaran.TanggalSelesai,
	}

	err = s.TahunAjaranService.Save(ctx, tahunAjaranModel, schemaName)
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
func (s *TahunAjaranServiceServer) GetTahunAjaran(ctx context.Context, req *pb.GetTahunAjaranRequest) (*pb.GetTahunAjaranResponse, error) {
	schemaName := req.GetSchemaname()
	tahunAjaranID := req.GetTahunAjaranId()

	tahunAjaranModel, err := s.TahunAjaranService.FindByID(ctx, tahunAjaranID, schemaName)
	if err != nil {
		log.Printf("Gagal menemukan tahun ajaran: %v", err)
		return nil, fmt.Errorf("gagal menemukan tahun ajaran: %w", err)
	}

	return &pb.GetTahunAjaranResponse{
		TahunAjaran: &pb.TahunAjaran{
			Nama:           tahunAjaranModel.Nama,
			PeriodeAktif:   tahunAjaranModel.PeriodeAktif,
			TanggalMulai:   tahunAjaranModel.TanggalMulai,
			TanggalSelesai: tahunAjaranModel.TanggalSelesai,
		},
	}, nil
}

// **UpdateTahunAjaran**
func (s *TahunAjaranServiceServer) UpdateTahunAjaran(ctx context.Context, req *pb.UpdateTahunAjaranRequest) (*pb.UpdateTahunAjaranResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"schemaname", "tahun_ajaran"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	tahunAjaranReq := req.GetTahunAjaran()
	tahunAjaranModel := &models.TahunAjaran{
		Nama:           tahunAjaranReq.Nama,
		PeriodeAktif:   tahunAjaranReq.PeriodeAktif,
		TanggalMulai:   tahunAjaranReq.TanggalMulai,
		TanggalSelesai: tahunAjaranReq.TanggalSelesai,
	}
	err = s.TahunAjaranService.Update(ctx, tahunAjaranModel, schemaName)
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
func (s *TahunAjaranServiceServer) DeleteTahunAjaran(ctx context.Context, req *pb.DeleteTahunAjaranRequest) (*pb.DeleteTahunAjaranResponse, error) {
	schemaName := req.GetSchemaname()
	tahunAjaranID := req.GetTahunAjaranId()

	err := s.TahunAjaranService.Delete(ctx, tahunAjaranID, schemaName)
	if err != nil {
		log.Printf("Gagal menghapus tahun ajaran: %v", err)
		return nil, fmt.Errorf("gagal menghapus tahun ajaran: %w", err)
	}

	return &pb.DeleteTahunAjaranResponse{
		Message: "Tahun ajaran berhasil dihapus",
		Status:  true,
	}, nil
}
