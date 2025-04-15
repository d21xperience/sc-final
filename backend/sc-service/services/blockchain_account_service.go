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
	config   *Config          // Konfigurasi runtime
	client   BlockchainClient // Client yang digunakan (Ethereum/Quorum)
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
	requiredFields := []string{"AkunParam", "schemaname"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	if req.GetSchemaname() == "\"\"" {
		return nil, errors.New("schemaname tidak boleh kosong")
	}

	var tenantSekolah = TenantSekolah{
		// SekolahId:       req.AkunParam.SekolahId,
		UserId:      req.AkunParam.UserId,
		Password:    req.AkunParam.Password,
		NamaSekolah: req.AkunParam.Organization,
		// SekolahIdEnkrip: req.AkunParam.SekolahIdEnkrip,
		Schemaname: req.GetSchemaname(),
	}
	schemaModel, schemaName, err := s.schema.GetOrCreateSchema(ctx, &tenantSekolah)
	if err != nil {
		return nil, err
	}
	contractAddress, err := s.client.GenerateNewAccount(ctx, tenantSekolah.UserId, tenantSekolah.Password)
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
		UserID:            tenantSekolah.UserId,
		Type:              models.AccountType("KEYSTORE"),
		Address:           address,
		Password:          pass,
		KeystrokeFilename: key,
		NetworkID:         uint32(req.AkunParam.NetworkId),
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
	requiredFields := []string{"Schemaname"}
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
		"schemaname": req.GetSchemaname(),
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

// func (s *BlockchainAccountService) ImportBlockchainAccount(ctx context.Context, req *pb.ImportBlockchainAccountRequest) (*pb.ImportBlockchainAccountResponse, error) {
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"Admin", "Network"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// GetOrCreate(schemaname)
// 	var tenantSekolah = tenantSekolah{
// 		SekolahId:       req.Admin.SekolahId,
// 		UserId:          req.Admin.UserId,
// 		Password:        req.Admin.Password,
// 		NamaSekolah:     req.Admin.NamaSekolah,
// 		SekolahIdEnkrip: req.Admin.SekolahIdEnkrip,
// 		Schemaname:      req.GetSchemaname(),
// 	}
// 	schemaModel, schemaName, err := s.schema.GetOrCreateSchema(ctx, &tenantSekolah)
// 	if !errors.Is(err, ErrSchemaFound) {
// 		return nil, err
// 	}
// 	address, err := ImportPrivateKey(req.GetPrivateKey())
// 	if err != nil {
// 		log.Printf("Gagal membuat akun: %v", err)
// 		return nil, fmt.Errorf("gagal membuat akun: %w", err)
// 	}
// 	err = s.repoAkun.Save(ctx, &models.Account{
// 		Username:     "",
// 		UserID:       tenantSekolah.UserId,
// 		Type:         models.AccountType("import"),
// 		Address:      address.Hex(),
// 		NetworkID:    req.Network.Id,
// 		Organization: schemaModel.NamaSekolah,
// 	}, schemaName)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &pb.ImportBlockchainAccountResponse{
// 		Status:  true,
// 		Message: "Akun berhasi diimport dengan address " + address.Hex(),
// 	}, nil
// }

// IssueDegree mengeluarkan ijazah di Ethereum
// func (e *BlockchainAccountService) IssueDegree(ctx context.Context, contractAddress string, degreeHash [32]byte, sekolah string, issueDate uint64, privateKey string, gasLimit uint64) (string, error) {
// 	//  Load ABI
// 	parsedABI, err := abi.JSON(strings.NewReader(abiJSON))
// 	if err != nil {
// 		return "", fmt.Errorf("error parsing ABI: %v", err)
// 	}

// 	//  Encode data untuk fungsi `issueDegree`
// 	data, err := parsedABI.Pack("issueDegree", degreeHash, sekolah, big.NewInt(int64(issueDate)))
// 	if err != nil {
// 		return "", fmt.Errorf("error packing data: %v", err)
// 	}

// 	//  Kirim transaksi menggunakan SendTransactionToContract
// 	txHash, err := SendTransactionToContract(ctx, e.client, contractAddress, data, privateKey, gasLimit)
// 	if err != nil {
// 		return "", fmt.Errorf("transaction failed: %v", err)
// 	}

// 	return txHash, nil
// }
