package config

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectGRPCServer(cfg Config) *grpc.ClientConn {
	address := fmt.Sprintf("%s:%d", cfg.GRPCHost, cfg.GRPCPort)
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Gagal terhubung ke server: %v", err)
	}
	return conn
}
