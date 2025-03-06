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

type DashboardServiceServer struct {
	pb.UnimplementedDashboardServiceServer
	repoSiswa repositories.GenericRepository[models.RombelAnggota]
	repoGuru  repositories.GenericRepository[models.PTKTerdaftar]
	repoKelas repositories.GenericRepository[models.RombonganBelajar]
}

func NewDashboardServiceServer() *DashboardServiceServer {
	repoSiswa := repositories.NewRombelAnggotaRepository(config.DB)
	repoGuru := repositories.NewPTKTerdaftarRepository(config.DB)
	repoKelas := repositories.NewrombonganBelajarRepository(config.DB)
	return &DashboardServiceServer{
		repoSiswa: *repoSiswa,
		repoGuru:  *repoGuru,
		repoKelas: *repoKelas,
	}
}

func (s *DashboardServiceServer) GetCountSiswa(ctx context.Context, req *pb.GetCountSiswaRequest) (*pb.GetCountSiswaResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	semesterId := req.GetSemesterId()
	if schemaName == "\"\"" {
		return nil, fmt.Errorf("schema name is required")
	}
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	count, err := s.repoSiswa.CountRows(ctx, schemaName, "semester_id", semesterId)
	if err != nil {
		return nil, err
	}

	return &pb.GetCountSiswaResponse{
		CountSiswa: count,
	}, nil
}
func (s *DashboardServiceServer) GetCountGuru(ctx context.Context, req *pb.GetCountGuruRequest) (*pb.GetCountGuruResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	tahunAjaranId := req.GetTahunAjaranId()
	if schemaName == "\"\"" {
		return nil, fmt.Errorf("schema name is required")
	}
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	count, err := s.repoGuru.CountRows(ctx, schemaName, "tahun_ajaran_id", tahunAjaranId)
	if err != nil {
		return nil, err
	}
	return &pb.GetCountGuruResponse{
		CountGuru: count,
	}, nil

}
func (s *DashboardServiceServer) GetCountKelas(ctx context.Context, req *pb.GetCountKelasRequest) (*pb.GetCountKelasResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	semesterId := req.GetSemesterId()
	if schemaName == "\"\"" {
		return nil, fmt.Errorf("schema name is required")
	}
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	count, err := s.repoKelas.CountRows(ctx, schemaName, "semester_id", semesterId)
	if err != nil {
		return nil, err
	}
	return &pb.GetCountKelasResponse{
		CountKelas: count,
	}, nil
}
