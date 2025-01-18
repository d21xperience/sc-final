package server

import (
	"context"
	"fmt"
	"os"
	pb "sekolah/generated"
	"sekolah/models"
	"sekolah/services"
)

// UploadServiceServer menangani unggahan Excel
type UploadServiceServer struct {
	pb.UnimplementedUploadServiceServer
	services.PesertaDidikService
	services.PTKService
	services.NilaiAkhirService
	services.IjazahService
}

// UploadFile menangani unggahan Excel berdasarkan jenis data
func (s *UploadServiceServer) UploadFile(ctx context.Context, req *pb.UploadFileRequest) (*pb.UploadFileResponse, error) {
	uploadType := req.GetUploadType()
	schemaName := req.GetSchemaname()
	fileData := req.GetFile()

	// Simpan file sementara
	filePath := fmt.Sprintf("/tmp/upload_%s.xlsx", uploadType)
	err := os.WriteFile(filePath, fileData, 0644)
	if err != nil {
		return nil, fmt.Errorf("gagal menyimpan file: %w", err)
	}

	// Proses file sesuai jenis unggahan
	dataList, err := services.UploadData(ctx, filePath, uploadType, schemaName)
	if err != nil {
		return nil, err
	}
	var jumlBaris int
	if v, ok := dataList.([]models.PesertaDidik); ok {
		jumlBaris = len(v)
	} else if v, ok := dataList.([]models.PTKTerdaftar); ok {
		jumlBaris = len(v)
	} else if v, ok := dataList.([]models.RombonganBelajar); ok {
		jumlBaris = len(v)
	}
	return &pb.UploadFileResponse{
		Message: fmt.Sprintf("Berhasil mengunggah data %d %s", jumlBaris, uploadType),
		Status:  true,
	}, nil
}
