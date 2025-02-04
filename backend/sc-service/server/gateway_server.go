package server

import (
	"context"
	"log"
	pb "sc-service/generated"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunHTTPGateway(ctx context.Context, mux *runtime.ServeMux, grpcServerEndpoint, httpPort string) {
	// Gunakan insecure.NewCredentials() sebagai pengganti grpc.WithInsecure()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Register gRPC service ke HTTP Gateway
	err := pb.RegisterBlockchainServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register BlockchainService on HTTP gateway: %v", err)
	}
	err = pb.RegisterBlockchainNetworkServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register BlockcahainNetwork on HTTP gateway: %v", err)
	}
	err = pb.RegisterBlockchainAccountServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		log.Fatalf("Failed to register BlockcahainAccount on HTTP gateway: %v", err)
	}
}
