package server

import (
	pb "sekolah/generated"
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
// func (s *UploadServiceServer) UploadFile(ctx context.Context, req *pb.UploadFileRequest) (*pb.UploadFileResponse, error) {
// 	uploadType := req.GetUploadType()
// 	schemaName := req.GetSchemaname()
// 	fileData := req.GetFile()

// 	// Simpan file sementara
// 	filePath := fmt.Sprintf("/tmp/upload_%s.xlsx", uploadType)
// 	err := os.WriteFile(filePath, fileData, 0644)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal menyimpan file: %w", err)
// 	}

// 	// Proses file sesuai jenis unggahan
// 	dataList, err := services.UploadData(ctx, filePath, uploadType, schemaName)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.UploadFileResponse{
// 		Message: fmt.Sprintf("Berhasil mengunggah %d data %s", len(dataList), uploadType),
// 		Status:  true,
// 	}, nil
// }
