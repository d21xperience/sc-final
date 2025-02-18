package server

import (
	pb "sekolah/generated"
	"sekolah/services"

	"google.golang.org/grpc"
)

var UploadService *services.UploadServiceServer

func RunGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer()

	sekolahService := services.NewSekolahService()
	pb.RegisterSekolahServiceServer(grpcServer, sekolahService)

	tahunAjaranService := services.NewTahunAjararanService()
	pb.RegisterTahunAjaranServiceServer(grpcServer, tahunAjaranService)

	semesterService := services.NewSemesterService()
	pb.RegisterSemesterServiceServer(grpcServer, semesterService)

	// REGISTER SISWA
	siswaService := services.NewSiswaServiceServer()
	pb.RegisterSiswaServiceServer(grpcServer, siswaService)

	// REGISTER KELAS
	kelasService := services.NewRombelServiceServer()
	pb.RegisterKelasServiceServer(grpcServer, kelasService)

	// REGISTER ANGGOTA KELAS
	anggotaKelasService := services.NewRombelAnggotaService()
	pb.RegisterAnggotaKelasServiceServer(grpcServer, anggotaKelasService)

	// REGISTER ANGGOTA KELAS
	nilaiAkhirService := services.NewNilaiAkhirServiceServer()
	pb.RegisterNilaiAkhirServiceServer(grpcServer, nilaiAkhirService)

	// REGISTER UPLOAD SERVICE
	UploadService := services.NewUploadServiceServer()
	pb.RegisterUploadDataSekolahServiceServer(grpcServer, UploadService)

	// REGISTER PTK TERDAFTAR SERVICE
	ptkTerdaftarService := services.NewPTKTerdaftarServiceServer()
	pb.RegisterPTKTerdaftarServiceServer(grpcServer, ptkTerdaftarService)

	return grpcServer
}
