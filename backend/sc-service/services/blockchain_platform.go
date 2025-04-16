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

	"github.com/google/uuid"
)

type BCPlatformServiceServer struct {
	pb.UnimplementedBCPlatformServiceServer
	// config *Config // Konfigurasi runtime
	repo *repositories.GenericRepository[models.BCPlatform]
}

// Constructor untuk BCPlatformService
func NewBCPlatformServiceServer() *BCPlatformServiceServer {
	repoBCPlatform := repositories.NewBCPlatformRepository(config.DB)
	return &BCPlatformServiceServer{
		repo: repoBCPlatform,
	}
}

// SetConfig: Mengatur konfigurasi blockchain
func (s *BCPlatformServiceServer) GetBCPlatform(ctx context.Context, req *pb.GetBCPlatformRequest) (*pb.GetBCPlatformResponse, error) {
	// Daftar field yang wajib diisi
	// requiredFields := []string{"Schemaname"}
	// // Validasi request
	// err := utils.ValidateFields(req, requiredFields)
	// if err != nil {
	// 	return nil, err
	// }
	bcPlatformModels, err := s.repo.FindAll(ctx, "ref", 5, 0)
	if err != nil {
		return nil, err
	}
	results := utils.ConvertModelsToPB(bcPlatformModels, func(model *models.BCPlatform) *pb.BCPlatform {
		return &pb.BCPlatform{
			Id:     model.ID.String(),
			Name:   model.NmBlockchain,
			Active: &model.Active,
		}
	})
	return &pb.GetBCPlatformResponse{
		BcPlatform: results,
	}, nil
}
func (s *BCPlatformServiceServer) SetBCPlatform(ctx context.Context, req *pb.SetBCPlatformRequest) (*pb.SetBCPlatformResponse, error) {
	// // Daftar field yang wajib diisi
	// requiredFields := []string{"BcPlatform", "Schemaname"}
	// // Validasi request
	// err := utils.ValidateFields(req, requiredFields)
	// if err != nil {
	// 	return nil, err
	// }
	// schemaname := req.GetSchemaname()
	bcPlatform := req.GetBcPlatform()

	bcPlatformModel := utils.ConvertModelToPB(bcPlatform, func(mod *pb.BCPlatform) *models.BCPlatform {
		parsedUUID, err := uuid.Parse(mod.Id)
		if err != nil {
			log.Printf("error parsing UUID from string: %v", err)
			return nil
		}
		return &models.BCPlatform{
			ID:           parsedUUID,
			NmBlockchain: mod.Name,
			Active:       *mod.Active,
		}
	})
	if bcPlatformModel == nil {
		return nil, fmt.Errorf("failed to convert model, got nil")
	}
	err := s.repo.Update(ctx, bcPlatformModel, "ref", "id", bcPlatform.Id)
	if err != nil {
		return nil, err
	}
	return &pb.SetBCPlatformResponse{
		Message: "sukses update platform",
	}, nil
}
