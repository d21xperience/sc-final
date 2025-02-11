package server

import (
	pb "sekolah/generated"
	"sekolah/services"

	"google.golang.org/grpc"
)

func RunGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer()

	// pb.RegisterSekolahServiceServer(grpcServer, &SekolahServiceServer{
	// 	schemaService:  services.schemaService,
	// 	sekolahService: services.sekolahService,
	// })
	tahunAjaranService := services.NewTahunAjararanService()
	pb.RegisterTahunAjaranServiceServer(grpcServer, tahunAjaranService)
	// pb.RegisterSemesterServiceServer(grpcServer, &SemesterServiceServer{
	// 	SemesterService: services.semesterService,
	// })
	// // REGISTER SISWA
	// pb.RegisterSiswaServiceServer(grpcServer, &SiswaServiceServer{
	// 	pesertaDidikService: services.pesertaDidikService,
	// })
	// // REGISTER KELAS
	// pb.RegisterKelasServiceServer(grpcServer, &RombelServiceServer{
	// 	service: services.rombonganBelajarService,
	// })
	// // REGISTER ANGGOTA KELAS
	// pb.RegisterAnggotaKelasServiceServer(grpcServer, &RombelAnggotaServiceServer{
	// 	rombelAnggotaService: services.rombelAnggotaService,
	// })
	// // REGISTER ANGGOTA KELAS
	// pb.RegisterNilaiAkhirServiceServer(grpcServer, &NilaiAkhirServiceServer{
	// 	NilaiAkhirService: services.nilaiAkhirService,
	// })
	// // REGISTER UPLOAD SERVICE
	// pb.RegisterUploadDataSekolahServiceServer(grpcServer, &UploadDataSekolahServiceServer{
	// 	pd: services.pesertaDidikService,
	// })
	return grpcServer
}
