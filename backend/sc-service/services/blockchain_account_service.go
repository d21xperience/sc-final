package services

import (
	"context"
	"errors"
	"fmt"
	"log"

	"sc-service/config"
	pb "sc-service/generated"
	"sc-service/models"
	"sc-service/repositories"
	"sc-service/utils"
)

type BlockchainAccountService struct {
	pb.UnimplementedBlockchainAccountServiceServer
	config   *Config   // Konfigurasi runtime
	client   EthClient // Client yang digunakan (Ethereum/Quorum)
	schema   SchemaService
	repoAkun *repositories.GenericRepository[models.Account]
}

// Constructor untuk BlockchainAccountService
func NewBlockchainAccountService() *BlockchainAccountService {
	schemaRepository := repositories.NewSchemaRepository(config.DB)
	sekolahTenantRepository := repositories.NewsekolahTenantRepository(config.DB)
	schema := NewSchemaService(schemaRepository, sekolahTenantRepository)
	akunRepository := repositories.NewAccountRepository(config.DB)
	return &BlockchainAccountService{
		config:   &Config{},
		schema:   schema,
		repoAkun: akunRepository,
	}
}

func (s *BlockchainAccountService) CreateBlockchainAccount(ctx context.Context, req *pb.CreateBlockchainAccountRequest) (*pb.CreateBlockchainAccountResponse, error) {
	if s.client == nil {
		return nil, errors.New("client belum dikonfigurasi")
	}
	// Daftar field yang wajib diisi
	requiredFields := []string{"Admin", "Network"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	if req.GetSchemaname() == "\"\"" {
		return nil, errors.New("schemaname tidak boleh kosong")
	}

	// GetOrCreate(schemaname)
	// GetOrCreate(schemaname)
	var adminSekolah = AdminSekolah{
		SekolahId:       req.Admin.SekolahId,
		UserId:          req.Admin.UserId,
		Password:        req.Admin.Password,
		NamaSekolah:     req.Admin.NamaSekolah,
		SekolahIdEnkrip: req.Admin.SekolahIdEnkrip,
		Schemaname:      req.GetSchemaname(),
	}
	schemaModel, schemaName, err := s.schema.GetOrCreateSchema(ctx, &adminSekolah)
	if err != nil {
		return nil, err
	}
	contractAddress, err := s.client.GenerateNewAccount(ctx, adminSekolah.UserId, adminSekolah.Password)
	if err != nil {
		log.Printf("Gagal membuat akun: %v", err)
		return nil, fmt.Errorf("gagal membuat akun: %w", err)
	}
	//  type assertion (.(string))
	address, ok := contractAddress["Address"].(string)
	if !ok {
		log.Fatal("Error: Address is not a string")
	}
	pass, ok := contractAddress["Password"].(string)
	if !ok {
		log.Fatal("Error: Address is not a string")
	}
	key, ok := contractAddress["KeystrokeFilename"].(string)
	if !ok {
		log.Fatal("Error: Address is not a string")
	}
	// Load network
	// bcNetwork :=
	// Simpan ke database
	s.repoAkun.Save(ctx, &models.Account{
		Username:          "",
		UserID:            adminSekolah.UserId,
		Type:              models.AccountType("KEYSTORE"),
		Address:           address,
		Password:          pass,
		KeystrokeFilename: key,
		NetworkID:         req.Network.Id,
		Organization:      schemaModel.NamaSekolah,
	}, schemaName)
	// txHash := ""
	return &pb.CreateBlockchainAccountResponse{
		Status:  true,
		Message: address,
	}, nil
}

func (s *BlockchainAccountService) GetBlockchainAccounts(ctx context.Context, req *pb.GetBlockchainAccountsRequest) (*pb.GetBlockchainAccountsResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"Schemaname", "UserId", "NetworkId"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	if req.GetSchemaname() == "\"\"" {
		return nil, errors.New("schemaname tidak boleh kosong")
	}
	// accounts, err := s.client.GetAccounts(ctx, req.GetUserId(), req.GetSchemaname())
	var condition = map[string]interface{}{
		"user_id":    req.GetUserId(),
		"network_id": req.GetNetworkId(),
	}
	accounts, err := s.repoAkun.FindAllByConditions(ctx, req.GetSchemaname(), condition, 100, 0)

	if err != nil {
		log.Printf("Gagal mendapatkan akun: %v", err)
		return nil, fmt.Errorf("gagal mendapatkan akun: %w", err)
	}

	results := utils.ConvertModelsToPB(accounts, func(model *models.Account) *pb.BlockchainAccount {
		return &pb.BlockchainAccount{
			UserId:            model.UserID,
			Address:           model.Address,
			KeystrokeFilename: model.KeystrokeFilename,
		}
	})
	status := false
	if len(results) > 0 {
		status = true
	}
	return &pb.GetBlockchainAccountsResponse{
		Status:             status,
		Blockchainaccounts: results,
	}, nil
}
func (s *BlockchainAccountService) ImportBlockchainAccount(ctx context.Context, req *pb.ImportBlockchainAccountRequest) (*pb.ImportBlockchainAccountResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"Admin", "Network"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	// GetOrCreate(schemaname)
	var adminSekolah = AdminSekolah{
		SekolahId:       req.Admin.SekolahId,
		UserId:          req.Admin.UserId,
		Password:        req.Admin.Password,
		NamaSekolah:     req.Admin.NamaSekolah,
		SekolahIdEnkrip: req.Admin.SekolahIdEnkrip,
		Schemaname:      req.GetSchemaname(),
	}
	schemaModel, schemaName, err := s.schema.GetOrCreateSchema(ctx, &adminSekolah)
	if !errors.Is(err, ErrSchemaFound) {
		return nil, err
	}
	address, err := ImportPrivateKey(req.GetPrivateKey())
	if err != nil {
		log.Printf("Gagal membuat akun: %v", err)
		return nil, fmt.Errorf("gagal membuat akun: %w", err)
	}
	err = s.repoAkun.Save(ctx, &models.Account{
		Username:     "",
		UserID:       adminSekolah.UserId,
		Type:         models.AccountType("import"),
		Address:      address.Hex(),
		NetworkID:    req.Network.Id,
		Organization: schemaModel.NamaSekolah,
	}, schemaName)
	if err != nil {
		return nil, err
	}
	return &pb.ImportBlockchainAccountResponse{
		Status:  true,
		Message: "Akun berhasi diimport dengan address " + address.Hex(),
	}, nil
}


