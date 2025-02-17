package server

import (
	pb "auth_service/generated"
	"auth_service/services"

	"google.golang.org/grpc"
)

// var UploadService *services.UploadServiceServer

func RunGRPCServer() *grpc.Server {
	// gRPC Server
	grpcServer := grpc.NewServer()
	// Register gRPC services
	authServiceServer := services.NewAuthServiceServer()
	pb.RegisterAuthServiceServer(grpcServer, authServiceServer)

	userProfileServiceServer := services.NewUserProfileServiceServer()
	pb.RegisterUserProfileServiceServer(grpcServer, userProfileServiceServer)
	sekolahIndonesiaServer := services.NewSekolahIndonesiaServer()
	pb.RegisterSekolahIndonesiaServiceServer(grpcServer, sekolahIndonesiaServer)

	return grpcServer
}
