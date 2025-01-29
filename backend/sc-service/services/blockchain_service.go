package services

import (
	"context"
	"errors"

	pb "sc-service/generated"
)

type BlockchainService struct {
	pb.UnimplementedBlockchainServiceServer
	config *Config   // Konfigurasi runtime
	client EthClient // Client yang digunakan (Ethereum/Quorum)
}

// Constructor untuk BlockchainService
func NewBlockchainService() *BlockchainService {
	return &BlockchainService{
		config: &Config{},
	}
}

// SetConfig: Mengatur konfigurasi blockchain
func (s *BlockchainService) SetConfig(ctx context.Context, req *pb.SetConfigRequest) (*pb.SetConfigResponse, error) {
	// Validasi input
	if req.BlockchainType != "ethereum" && req.BlockchainType != "quorum" {
		return nil, errors.New("blockchain_type harus 'ethereum' atau 'quorum'")
	}

	// Update konfigurasi runtime
	s.config.BlockchainType = req.BlockchainType
	s.config.RPCURL = req.RpcUrl

	// Buat client sesuai konfigurasi
	client, err := CreateClientFactory(s.config)
	if err != nil {
		return nil, err
	}
	s.client = client

	return &pb.SetConfigResponse{
		Message: "Konfigurasi blockchain berhasil diperbarui",
	}, nil
}

// GetNetworkID: Mendapatkan Network ID dari blockchain
func (s *BlockchainService) GetNetworkID(ctx context.Context, _ *pb.Empty) (*pb.NetworkIDResponse, error) {
	if s.client == nil {
		return nil, errors.New("client belum dikonfigurasi")
	}

	networkID, err := s.client.NetworkID(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.NetworkIDResponse{
		NetworkId: networkID.String(),
	}, nil
}

// GetConsensusAlgorithm: Mendapatkan algoritma konsensus (hanya untuk Quorum)
func (s *BlockchainService) GetConsensusAlgorithm(ctx context.Context, _ *pb.Empty) (*pb.ConsensusAlgorithmResponse, error) {
	// Periksa apakah client adalah QuorumClient
	quorumClient, ok := s.client.(QuorumClient)
	if !ok {
		return nil, errors.New("fitur ini hanya tersedia untuk Quorum")
	}

	consensus, err := quorumClient.GetConsensusAlgorithm(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.ConsensusAlgorithmResponse{
		ConsensusAlgorithm: consensus,
	}, nil
}
