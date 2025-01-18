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

type SemesterServiceServer struct {
	pb.UnimplementedSemesterServiceServer
	SemesterService services.SemesterService
}

// **CreateSemester**
func (s *SemesterServiceServer) CreateSemester(ctx context.Context, req *pb.CreateSemesterRequest) (*pb.CreateSemesterResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "Semester"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	semester := req.GetSemester()
	semesterModel := &models.Semester{
		NamaSemester:   semester.NamaSemester,
		PeriodeAktif:   semester.PeriodeAktif,
		TanggalMulai:   semester.TanggalMulai,
		TanggalSelesai: semester.TanggalSelesai,
	}

	err = s.SemesterService.Save(ctx, semesterModel, schemaName)
	if err != nil {
		log.Printf("Gagal menyimpan tahun ajaran: %v", err)
		return nil, fmt.Errorf("gagal menyimpan tahun ajaran: %w", err)
	}

	return &pb.CreateSemesterResponse{
		Message: "Tahun ajaran berhasil ditambahkan",
		Status:  true,
	}, nil
}

// **GetSemester**
func (s *SemesterServiceServer) GetSemester(ctx context.Context, req *pb.GetSemesterRequest) (*pb.GetSemesterResponse, error) {
	schemaName := req.GetSchemaname()
	SemesterID := req.GetSemesterId()

	semesterModel, err := s.SemesterService.FindByID(ctx, SemesterID, schemaName)
	if err != nil {
		log.Printf("Gagal menemukan tahun ajaran: %v", err)
		return nil, fmt.Errorf("gagal menemukan tahun ajaran: %w", err)
	}

	return &pb.GetSemesterResponse{
		Semester: &pb.Semester{
			NamaSemester:   semesterModel.NamaSemester,
			PeriodeAktif:   semesterModel.PeriodeAktif,
			TanggalMulai:   semesterModel.TanggalMulai,
			TanggalSelesai: semesterModel.TanggalSelesai,
		},
	}, nil
}

// **UpdateSemester**
func (s *SemesterServiceServer) UpdateSemester(ctx context.Context, req *pb.UpdateSemesterRequest) (*pb.UpdateSemesterResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"schemaname", "semester"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	schemaName := req.GetSchemaname()
	semesterReq := req.GetSemester()
	SemesterModel := &models.Semester{
		NamaSemester:   semesterReq.NamaSemester,
		PeriodeAktif:   semesterReq.PeriodeAktif,
		TanggalMulai:   semesterReq.TanggalMulai,
		TanggalSelesai: semesterReq.TanggalSelesai,
	}
	err = s.SemesterService.Update(ctx, SemesterModel, schemaName)
	if err != nil {
		log.Printf("Gagal memperbarui tahun ajaran: %v", err)
		return nil, fmt.Errorf("gagal memperbarui tahun ajaran: %w", err)
	}
	return &pb.UpdateSemesterResponse{
		Message: "Semester berhasil diperbarui",
		Status:  true,
	}, nil
}

// // **DeleteSemester**
func (s *SemesterServiceServer) DeleteSemester(ctx context.Context, req *pb.DeleteSemesterRequest) (*pb.DeleteSemesterResponse, error) {
	schemaName := req.GetSchemaname()
	SemesterID := req.GetSemesterId()

	err := s.SemesterService.Delete(ctx, SemesterID, schemaName)
	if err != nil {
		log.Printf("Gagal menghapus tahun ajaran: %v", err)
		return nil, fmt.Errorf("gagal menghapus tahun ajaran: %w", err)
	}

	return &pb.DeleteSemesterResponse{
		Message: "Tahun ajaran berhasil dihapus",
		Status:  true,
	}, nil
}
