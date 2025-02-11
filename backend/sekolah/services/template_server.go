package services

// import (
// 	"context"
// 	"fmt"
// 	"os"

// 	pb "sekolah/generated"
// 	"sekolah/services"
// )

// // TemplateServiceServer menangani template Excel
// type TemplateServiceServer struct {
// 	pb.UnimplementedTemplateServiceServer
// }

// // GetTemplate menyediakan template Excel berdasarkan jenis data
// func (s *TemplateServiceServer) GetTemplate(ctx context.Context, req *pb.GetTemplateRequest) (*pb.GetTemplateResponse, error) {
// 	templateType := req.GetTemplateType()
// 	templatePath := fmt.Sprintf("/tmp/template_%s.xlsx", templateType)

// 	// Buat file template jika belum ada
// 	if _, err := os.Stat(templatePath); os.IsNotExist(err) {
// 		err := services.GenerateTemplate(templateType, templatePath)
// 		if err != nil {
// 			return nil, fmt.Errorf("gagal membuat template %s: %w", templateType, err)
// 		}
// 	}

// 	// Baca file template
// 	data, err := os.ReadFile(templatePath)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal membaca template %s: %w", templateType, err)
// 	}

// 	return &pb.GetTemplateResponse{
// 		FileName: fmt.Sprintf("template_%s.xlsx", templateType),
// 		FileData: data,
// 	}, nil
// }
