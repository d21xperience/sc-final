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

	// REGISTER PTKSERVICE
	ptkService := services.NewPTKServiceServer()
	pb.RegisterPTKServiceServer(grpcServer, ptkService)

	// REGISTER PTK TERDAFTAR SERVICE
	ptkTerdaftarService := services.NewPTKTerdaftarServiceServer()
	pb.RegisterPTKTerdaftarServiceServer(grpcServer, ptkTerdaftarService)

	// REGISTER DASHBOARD SERVICE
	dashboardService := services.NewDashboardServiceServer()
	pb.RegisterDashboardServiceServer(grpcServer, dashboardService)

	// REGISTER REFERENSI TABEL SERVICE
	referensiService := services.NewReferensiServiceServer()
	pb.RegisterReferensiServiceServer(grpcServer, referensiService)

	// REGISTER PEMBELAJARAN SERVICE
	pembelajaranService := services.NewPembelajaranServiceServer()
	pb.RegisterPembelajaranServiceServer(grpcServer, pembelajaranService)

	// REGISTER KENAIKAN SERVICE
	kenaikanService := services.NewKenaikanServiceServer()
	pb.RegisterKenaikanServiceServer(grpcServer, kenaikanService)

	// REGISTER PEMBELAJARAN SERVICE
	ijazahService := services.NewIjazahServiceServer()
	pb.RegisterIjazahServiceServer(grpcServer, ijazahService)

	return grpcServer
}
