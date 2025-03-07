package services

import (
	"context"
	"fmt"
	"log"

	"sc-service/config"
	pb "sc-service/generated"
	"sc-service/models"
	"sc-service/repositories"
	"sc-service/utils"
)

type BlockchainNetworkService struct {
	pb.UnimplementedBlockchainNetworkServiceServer
	config *Config // Konfigurasi runtime
	repo *repositories.GenericRepository[models.Network]
}

// Constructor untuk BlockchainNetworkService
func NewBlockchainNetworkService() *BlockchainNetworkService {
	repoNetwork := repositories.NewNetworkRepository(config.DB)
	return &BlockchainNetworkService{
		config: &Config{},
		repo:   repoNetwork,
	}
}

// SetConfig: Mengatur konfigurasi blockchain
func (s *BlockchainNetworkService) CreateBCNetwork(ctx context.Context, req *pb.CreateBCNetworkRequest) (*pb.CreateBCNetworkResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"Network"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	bcNetwork := req.GetNetwork()
	modelNetwork := utils.ConvertPBToModels(bcNetwork, func(entity *pb.BCNetwork) *models.Network {
		// Konversi ENUM dari Protobuf ke GORM
		networkType, err := ConvertProtoToNetworkType(entity.Type)
		if err != nil {
			log.Printf("invalid network type: %v", err)
		}
		return &models.Network{
			Name:        entity.Name,
			ChainID:     entity.ChainId,
			RPCURL:      entity.RPCURL,
			ExplorerURL: entity.ExplorerURL,
			Symbol:      entity.Symbol,
			Type:        networkType,
		}
	})
	err = s.repo.SaveMany(ctx, "ref", modelNetwork, 100)
	if err != nil {
		return nil, err
	}

	return &pb.CreateBCNetworkResponse{
		Status:  true,
		Message: "sukses",
	}, nil
}
func (s *BlockchainNetworkService) GetBCNetwork(ctx context.Context, req *pb.GetBCNetworkRequest) (*pb.GetBCNetworkResponse, error) {
	var modelNetwork []*models.Network
	var message string
	var status bool
	// Daftar field yang wajib diisi
	requiredFields := []string{"NetworkName"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err == nil {
		// return nil, err
		condition := map[string]interface{}{
			"name": req.GetNetworkName(),
		}
		modelNetwork, err = s.repo.FindAllByConditions(ctx, "ref", condition, 100, 0)
		if err != nil {
			return nil, err
		}
		if len(modelNetwork) > 0 {
			message = "Sukses"
			status = true
		} else {
			message = "Gagal"
			status = false
		}

	} else {
		modelNetwork, err = s.repo.FindAll(ctx, "ref", 100, 0)
		if err != nil {
			return nil, err
		}
		if len(modelNetwork) > 0 {
			message = "Sukses"
			status = true
		} else {
			message = "Gagal"
			status = false
		}
	}
	networks := utils.ConvertModelsToPB(modelNetwork, func(model *models.Network) *pb.BCNetwork {
		return &pb.BCNetwork{
			Name:        model.Name,
			ChainId:     model.ChainID,
			RPCURL:      model.RPCURL,
			ExplorerURL: model.ExplorerURL,
			Symbol:      model.Symbol,
			Type:        ConvertNetworkTypeToProto(model.Type),
			Activate:    model.Activate,
			Available:   model.Available,
			Id:          model.ID,
		}
	})

	return &pb.GetBCNetworkResponse{
		Status:  status,
		Message: message,
		Network: networks,
	}, nil
}
func (s *BlockchainNetworkService) UpdateBCNetwork(ctx context.Context, req *pb.UpdateBCNetworkRequest) (*pb.UpdateBCNetworkResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"network"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	coba := req.GetNetwork()
	modelAkun := utils.ConvertModelToPB(coba, func(model *pb.BCNetwork) *models.Network {
		return &models.Network{
			Name:        model.Name,
			RPCURL:      model.RPCURL,
			ExplorerURL: model.ExplorerURL,
			Activate:    model.Activate,
			Available:   model.Available,
			Symbol:      model.Symbol,
		}
	})

	str := utils.ConvertUintToString(req.Network.Id)
	err = s.repo.Update(ctx, modelAkun, "ref", "id", str)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateBCNetworkResponse{
		Status:  true,
		Message: "sukses",
	}, nil
}

func (s *BlockchainNetworkService) DeleteBCNetwork(ctx context.Context, req *pb.DeleteNetworkRequest) (*pb.DeleteNetworkResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"NetworkId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	id_network := utils.ConvertUintToString(req.NetworkId)

	err = s.repo.Delete(ctx, id_network, "ref", "id")
	if err != nil {
		return nil, err
	}
	return &pb.DeleteNetworkResponse{
		Status:  true,
		Message: "suskes",
	}, nil
}

// Konversi dari Protobuf ENUM ke Golang ENUM
func ConvertProtoToNetworkType(protoType pb.NetworkType) (models.NetworkType, error) {
	switch protoType {
	case pb.NetworkType_mainnet:
		return models.Mainnet, nil
	case pb.NetworkType_testnet:
		return models.Testnet, nil
	case pb.NetworkType_private:
		return models.Private, nil
	default:
		return "", fmt.Errorf("invalid NetworkType: %v", protoType)
	}
}

// ConvertNetworkTypeToProto mengonversi dari GORM `NetworkType` ke Protobuf `NetworkType`
func ConvertNetworkTypeToProto(networkType models.NetworkType) pb.NetworkType {
	switch networkType {
	case models.Mainnet:
		return pb.NetworkType_mainnet
	case models.Testnet:
		return pb.NetworkType_testnet
	case models.Private:
		return pb.NetworkType_private
	default:
		return pb.NetworkType_mainnet // Default fallback ke MAINNET jika tidak valid
	}
}
