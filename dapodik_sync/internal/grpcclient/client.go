package grpcclient

import (
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	Address     string        // e.g. "localhost:50052"
	DialTimeout time.Duration // tidak terpakai di grpc.NewClient tapi tetap bisa digunakan di logic tambahan
}

type ClientWrapper struct {
	Conn *grpc.ClientConn
}

func NewGRPCClient(cfg Config) (*ClientWrapper, error) {
	conn, err := grpc.NewClient(cfg.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("gagal membuat koneksi gRPC: %w", err)
	}

	return &ClientWrapper{Conn: conn}, nil
}

func (c *ClientWrapper) Close() {
	if c.Conn != nil {
		_ = c.Conn.Close()
	}
}
