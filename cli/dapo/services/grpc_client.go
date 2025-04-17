package services

import (
	"context"
	"fmt"
	"time"

	pb "dapo/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GRPCClient represents a gRPC client
type GRPCClient struct {
	conn          *grpc.ClientConn
	timeout       time.Duration
	clientSekolah pb.SekolahServiceClient
	clientSiswa   pb.SiswaServiceClient
	clientPtk     pb.PTKTerdaftarServiceClient
	clientKelas   pb.KelasServiceClient
}

var Schemaname = "tabel_d4da6b98fcfd71c58f5a"

// NewGRPCClient creates a new gRPC client connection
func NewGRPCClient(host, port string, timeout time.Duration) (*GRPCClient, error) {
	// Set up a connection to the server
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("did not connect: %w", err)
	}

	// client := pb.NewStudentServiceClient(conn)
	clientSekolah := pb.NewSekolahServiceClient(conn)
	clientSiswa := pb.NewSiswaServiceClient(conn)
	clientPtk := pb.NewPTKTerdaftarServiceClient(conn)
	clientKelas := pb.NewKelasServiceClient(conn)

	return &GRPCClient{
		conn:          conn,
		timeout:       timeout,
		clientSekolah: clientSekolah,
		clientSiswa:   clientSiswa,
		clientPtk:     clientPtk,
		clientKelas:   clientKelas,
	}, nil
}

// Close closes the gRPC connection
func (c *GRPCClient) Close() error {
	return c.conn.Close()
}

// SendSekolahData sends student data to gRPC server
func (c *GRPCClient) SendSekolahData(data *pb.SekolahDapo) (*pb.UpdateSekolahResponse, error) {
	// Convert SekolahData to protobuf message
	req := &pb.UpdateSekolahRequest{
		Schemaname: Schemaname,
		Sekolah:    data,
	}

	// Contact the server with timeout
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	return c.clientSekolah.UpdateSekolah(ctx, req)
}

// SendStudentData sends student data to gRPC server
func (c *GRPCClient) SendStudentData(data []*pb.Siswa) (*pb.CreateBanyakSiswaResponse, error) {
	// Convert StudentData to protobuf message
	req := &pb.CreateBanyakSiswaRequest{
		Schemaname: Schemaname,
		Siswa:      data,
	}
	// Contact the server with timeout
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	return c.clientSiswa.CreateBanyakSiswa(ctx, req)
}
func (c *GRPCClient) SendPtkTerdaftarData(data []*pb.PTKTerdaftar) (*pb.CreatePTKTerdaftarResponse, error) {
	// Convert StudentData to protobuf message
	req := &pb.CreatePTKTerdaftarRequest{
		Schemaname:   Schemaname,
		PtkTerdaftar: data,
	}
	// Contact the server with timeout
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	return c.clientPtk.CreatePTKTerdaftar(ctx, req)
}

// SendStudentData sends student data to gRPC server
func (c *GRPCClient) SendRombel(data []*pb.Kelas) (*pb.ImportDapodikRombelResponse, error) {
	// Convert StudentData to protobuf message
	req := &pb.ImportDapodikRombelRequest{
		Schemaname: Schemaname,
		Kelas:      data,
	}
	// Contact the server with timeout
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	return c.clientKelas.ImportDapodikRombel(ctx, req)
}
