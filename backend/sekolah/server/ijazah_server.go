package server

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	pb "sekolah/generated"
// 	"sekolah/models"
// 	"sekolah/services"
// )

// type IjazahServiceServer struct {
// 	pb.UnimplementedIjazahServiceServer
// 	ijazahService services.IjazahService
// }

// // **CreateIjazah**
// func (s *IjazahServiceServer) CreateIjazah(ctx context.Context, req *pb.CreateIjazahRequest) (*pb.CreateIjazahResponse, error) {

// }

// // **GetIjazah**
// // func (s *IjazahServiceServer) GetIjazah(ctx context.Context, req *pb.GetIjazahRequest) (*pb.GetIjazahResponse, error) {
// // 	schemaName := req.GetSchemaName()
// // 	ijazahID := req.GetIjazahId()

// // 	ijazah, err := s.ijazahService.FindByID(ctx, ijazahID, schemaName)
// // 	if err != nil {
// // 		log.Printf("Gagal menemukan Ijazah: %v", err)
// // 		return nil, fmt.Errorf("gagal menemukan Ijazah: %w", err)
// // 	}

// // 	return &pb.GetIjazahResponse{
// // 		Ijazah: &pb.Ijazah{
// // 			// PesertaDidikID:              ,
// // 			Nama:                        ijazah.Nama,
// // 			Nis:                         ijazah.Nis,
// // 			NISN:                        ijazah.NISN,
// // 			NPSN:                        ijazah.NPSN,
// // 			NoIjazah:                    ijazah.NoIjazah,
// // 			TempatLahir:                 ijazah.TempatLahir,
// // 			TanggalLahir:                ijazah.TanggalLahir,
// // 			NamaOrtuWali:                ijazah.NamaOrtuWali,
// // 			PaketKeahlian:               ijazah.PaketKeahlian,
// // 			KabupatenKota:               ijazah.KabupatenKota,
// // 			Provinsi:                    ijazah.Provinsi,
// // 			ProgramKeahlian:             ijazah.ProgramKeahlian,
// // 			SekolahPenyelenggaraUjianUS: ijazah.SekolahPenyelenggaraUjianUS,
// // 			SekolahPenyelenggaraUjianUN: ijazah.SekolahPenyelenggaraUjianUN,
// // 			AsalSekolah:                 ijazah.AsalSekolah,
// // 			NomorIjazah:                 ijazah.NomorIjazah,
// // 			TempatIjazah:                ijazah.TempatIjazah,
// // 			TanggalIjazah:               ijazah.TanggalIjazah,
// // 		},
// // 	}, nil
// // }

// // **UpdateIjazah**
// // func (s *IjazahServiceServer) UpdateIjazah(ctx context.Context, req *pb.UpdateIjazahRequest) (*pb.UpdateIjazahResponse, error) {
// // 	// Debugging: Cek nilai request yang diterima
// // 	log.Printf("Received UpdateUserProfile request: %+v\n", req)
// // 	schemaName := req.GetSchemaName()
// // 	ijazahReq := req.GetIjazah()
// // 	ijazahModel := &models.Ijazah{
// // 		Nama:                        ijazahReq.Nama,
// // 		Nis:                         ijazahReq.Nis,
// // 		NISN:                        ijazahReq.NISN,
// // 		NPSN:                        ijazahReq.NPSN,
// // 		NoIjazah:                    ijazahReq.NoIjazah,
// // 		TempatLahir:                 ijazahReq.TempatLahir,
// // 		TanggalLahir:                ijazahReq.TanggalLahir,
// // 		NamaOrtuWali:                ijazahReq.NamaOrtuWali,
// // 		PaketKeahlian:               ijazahReq.PaketKeahlian,
// // 		KabupatenKota:               ijazahReq.KabupatenKota,
// // 		Provinsi:                    ijazahReq.Provinsi,
// // 		ProgramKeahlian:             ijazahReq.ProgramKeahlian,
// // 		SekolahPenyelenggaraUjianUS: ijazahReq.SekolahPenyelenggaraUjianUS,
// // 		SekolahPenyelenggaraUjianUN: ijazahReq.SekolahPenyelenggaraUjianUN,
// // 		AsalSekolah:                 ijazahReq.AsalSekolah,
// // 		NomorIjazah:                 ijazahReq.NomorIjazah,
// // 		TempatIjazah:                ijazahReq.TempatIjazah,
// // 		TanggalIjazah:               ijazahReq.TanggalIjazah,
// // 	}
// // 	err := s.ijazahService.Update(ctx, ijazahModel, schemaName)
// // 	if err != nil {
// // 		log.Printf("Gagal memperbarui Ijazah: %v", err)
// // 		return nil, fmt.Errorf("gagal memperbarui Ijazah: %w", err)
// // 	}
// // 	return &pb.UpdateIjazahResponse{
// // 		Message: "Ijazah berhasil diperbarui",
// // 		Status:  true,
// // 	}, nil
// // }

// // // **DeleteIjazah**
// func (s *IjazahServiceServer) DeleteIjazah(ctx context.Context, req *pb.DeleteIjazahRequest) (*pb.DeleteIjazahResponse, error) {
// 	schemaName := req.GetSchemaName()
// 	IjazahID := req.GetIjazahId()

// 	err := s.ijazahService.Delete(ctx, IjazahID, schemaName)
// 	if err != nil {
// 		log.Printf("Gagal menghapus Ijazah: %v", err)
// 		return nil, fmt.Errorf("gagal menghapus Ijazah: %w", err)
// 	}

// 	return &pb.DeleteIjazahResponse{
// 		Message: "Ijazah berhasil dihapus",
// 		Status:  true,
// 	}, nil
// }
