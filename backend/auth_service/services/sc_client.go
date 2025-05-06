package services

import (
	"auth_service/models"
	"context"
	"fmt"
	"time"

	pb "auth_service/generated/sc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SCServiceClient struct {
	client   pb.TenantServiceClient
	clientBC pb.BlockchainAccountServiceClient
	// client pb.BlockchainAccountServiceClient
}
type AdminSekolah struct {
	SekolahId   int32
	UserId      int32
	Password    string
	NamaSekolah string
	EnkripID    string
}

type BCNetwork struct {
	Name        string
	ChainId     int64
	RPCURL      string
	ExplorerURL string
	Symbol      string
	// Type        NetworkType
	Activate     bool
	Available    bool
	Id           uint32
	Architecture string
}

func NewSCServiceClient() (*SCServiceClient, error) {
	conn, err := grpc.NewClient("localhost:50054", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("gagal terhubung ke sc-service: %w", err)
	}

	// client := pb.NewBlockchainAccountServiceClient(conn)
	client := pb.NewTenantServiceClient(conn)
	clientBC := pb.NewBlockchainAccountServiceClient(conn)
	return &SCServiceClient{client: client, clientBC: clientBC}, nil
}

func (s *SCServiceClient) RegistrasiSekolahTenant(sekolah *models.SekolahTenant, userId int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.client.RegistrasiSekolahTenant(ctx, &pb.RegistrasiSekolahTenantRequest{
		SekolahTenant: &pb.SekolahTenant{
			NamaSekolah:     sekolah.NamaSekolah,
			Schemaname:      fmt.Sprintf("tabel_%s", sekolah.EnkripID),
			SekolahTenantId: sekolah.ID,
			UserId:          userId,
		},
	})
	if err != nil {
		return fmt.Errorf("gagal mendaftarkan SC di SC_service: %w", err)
	}

	// log.Printf("SChema SC %s berhasil dibuat di SC_service", SC.SCIDEnkrip)
	return nil
}

// func (s *SCServiceClient) CreateDefaultBlockchainAccount(admin AdminSekolah, bcNetowrk BCNetwork, schemaname string) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	// Buat akun ethereum
// 	_, err := s.clientBC.CreateBlockchainAccount(ctx, &pb.CreateBlockchainAccountRequest{

// 		Admin:      &pb.AdminSekolah{
// 			SekolahId: admin.SekolahId,
// 			UserId: admin.UserId,
// 			Password: admin.Password,
// 			NamaSekolah: admin.NamaSekolah,
// 			EnkripID: admin.EnkripID,
// 		},
// 		Network:    &pb.BCNetwork{
// 			Name: bcNetowrk.Name,
// 		},
// 		Schemaname: "tes",
// 	})
// 	if err != nil {
// 		return fmt.Errorf("gagal mendaftarkan SC di SC_service: %w", err)
// 	}
// 	return nil
// }
