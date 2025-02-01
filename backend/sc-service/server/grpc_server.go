package server

import (
	pb "sc-service/generated"
	"sc-service/services"

	"google.golang.org/grpc"
)
type GRPCServer struct {
	// grpcServer              *grpc.Server
// 	schemaService           services.SchemaService

}

func RunGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer()
	blockchainService := services.NewBlockchainService()
	pb.RegisterBlockchainServiceServer(grpcServer, blockchainService)
	return grpcServer
}
