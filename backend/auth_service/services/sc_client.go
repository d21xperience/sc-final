package services

import (
	pb "auth_service/generated/sc"
	"auth_service/models"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SCServiceClient struct {
	client pb.BlockchainAccountServiceClient
}

func NewSCServiceClient() (*SCServiceClient, error) {
	conn, err := grpc.NewClient("localhost:50054", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("gagal terhubung ke sc-service: %w", err)
	}

	client := pb.NewBlockchainAccountServiceClient(conn)
	return &SCServiceClient{client: client}, nil
}

// func (s *SCServiceClient) RegistrasiSC() error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	_, err := s.client.RegistrasiSC(ctx, &pb.TabelSCRequest{
// 		SC: &pb.SC{
// 			SCIdEnkrip: SC.SCIDEnkrip,
// 			SCId:       int32(SC.ID),
// 			NamaSC:     SC.NamaSC,
// 		},
// 	})
// 	if err != nil {
// 		return fmt.Errorf("gagal mendaftarkan SC di SC_service: %w", err)
// 	}

// 	log.Printf("SChema SC %s berhasil dibuat di SC_service", SC.SCIDEnkrip)
// 	return nil
// }
func (s *SCServiceClient) CreateBlockchainAccount(userModel *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.client.CreateBlockchainAccount(ctx, &pb.CreateBlockchainAccountRequest{

	})
	if err != nil {
		return fmt.Errorf("gagal mendaftarkan SC di SC_service: %w", err)
	}
	return nil
}
