package services

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	pb "sekolah/generated"
// 	"sekolah/models"
// )

// type IjazahService struct {
// 	pb.UnimplementedIjazahServiceServer
// 	ijazahService services.IjazahService
// }

// func NewIjazahService() *IjazahService {
// 	return &IjazahService{}
// }

// func (s *IjazahService) CreateIjazah(ctx context.Context, req *pb.CreateIjazahRequest) (*pb.CreateIjazahResponse, error) {
// 	schemaName := req.GetSchemaName()
// 	ijazah := req.GetIjazah()
// 	// tanggalIjazah, err := time.Parse("2006-01-02", ijazah.TanggalIjazah)
// 	IjazahModel := &models.Ijazah{
// 		Nama:                        ijazah.Nama,
// 		Nis:                         ijazah.Nis,
// 		NISN:                        ijazah.Nisn,
// 		NPSN:                        ijazah.Npsn,
// 		NoIjazah:                    ijazah.NomorIjazah,
// 		TempatLahir:                 ijazah.TempatLahir,
// 		TanggalLahir:                ijazah.TanggalLahir,
// 		NamaOrtuWali:                ijazah.NamaOrtuwali,
// 		PaketKeahlian:               ijazah.PaketKeahlian,
// 		KabupatenKota:               ijazah.Kabupatenkota,
// 		Provinsi:                    ijazah.Provinsi,
// 		ProgramKeahlian:             ijazah.ProgramKeahlian,
// 		SekolahPenyelenggaraUjianUS: ijazah.SekolahPenyelenggaraUjianUs,
// 		SekolahPenyelenggaraUjianUN: ijazah.SekolahPenyelenggaraUjianUn,
// 		AsalSekolah:                 ijazah.AsalSekolah,
// 		NomorIjazah:                 ijazah.NomorIjazah,
// 		TempatIjazah:                ijazah.TempatIjazah,
// 		TanggalIjazah:               ijazah.TanggalIjazah,
// 	}

// 	err := s.ijazahService.Save(ctx, IjazahModel, schemaName)
// 	if err != nil {
// 		log.Printf("Gagal menyimpan Ijazah: %v", err)
// 		return nil, fmt.Errorf("gagal menyimpan Ijazah: %w", err)
// 	}

// 	return &pb.CreateIjazahResponse{
// 		Message: "Ijazah berhasil ditambahkan",
// 		Status:  true,
// 	}, nil
// }
// func (s *IjazahService) GetIjazah(ctx context.Context, req *pb.GetIjazahRequest) (*pb.GetIjazahResponse, error) {

// 	return &pb.GetIjazahResponse{

// 		Status:  true,
// 		Message: "sukses",
// 	}, nil

// }
// func (s *IjazahService) UpdateIjazah(ctx context.Context, req *pb.UpdateIjazahRequest) (*pb.UpdateIjazahResponse, error) {

// 	return &pb.UpdateIjazahResponse{
// 		Status:  true,
// 		Message: "sukses",
// 	}, nil

// }
// func (s *IjazahService) DeleteIjazah(ctx context.Context, req *pb.DeleteIjazahRequest) (*pb.DeleteIjazahResponse, error) {
// 	return &pb.DeleteIjazahResponse{
// 		Status:  true,
// 		Message: "sukses",
// 	}, nil

// }
