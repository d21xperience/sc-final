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
		log.Fatalf("Failed to register HTTP gateway: %v", err)
	}

	// // log.Printf("HTTP gateway running on port %s", httpPort)
	// if err := http.ListenAndServe(":"+httpPort, mux); err != nil {
	// 	log.Fatalf("Failed to serve HTTP gateway: %v", err)
	// }
}
