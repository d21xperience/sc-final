package server

import (
	pb "sc-service/generated"
	"sc-service/services"

	"google.golang.org/grpc"
)

func RunGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	blockchainService := services.NewBlockchainService()
	pb.RegisterBlockchainServiceServer(grpcServer, blockchainService)
	return grpcServer
}
