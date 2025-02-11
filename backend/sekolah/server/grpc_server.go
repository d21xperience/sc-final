package server

import (
	pb "sekolah/generated"
	"sekolah/services"

	"google.golang.org/grpc"
)

func RunGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer()

	// sekolahService := service
	// schemaService := services.NewSchemaService()
	// pb.RegisterSekolahServiceServer(grpcServer, schemaService)
	tahunAjaranService := services.NewTahunAjararanService()
	pb.RegisterTahunAjaranServiceServer(grpcServer, tahunAjaranService)
	// pb.RegisterSemesterServiceServer(grpcServer, &SemesterServiceServer{
	// 	SemesterService: services.semesterService,
	// })

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
	// sekolahService := services.NewSekolahService()
	// pb.RegisterUploadDataSekolahServiceServer(grpcServer, sekolahService)

	return grpcServer
}
