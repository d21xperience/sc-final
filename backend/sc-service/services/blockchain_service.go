package services

import (
	"context"
	"errors"
	"fmt"
	"log"

	pb "sc-service/generated"

	"sc-service/utils"
)

type BlockchainService struct {
	pb.UnimplementedBlockchainServiceServer
	config *Config // Konfigurasi runtime
	client BlockchainClient
}

// Constructor untuk BlockchainService
func NewBlockchainService() *BlockchainService {
	return &BlockchainService{
		config: &Config{},
		client: nil,
	}
}

// SetConfig: Mengatur konfigurasi blockchain
func (s *BlockchainService) SetConfig(ctx context.Context, req *pb.SetConfigRequest) (*pb.SetConfigResponse, error) {
	// Daftar field yang wajib diisi
	requiredFields := []string{"Network"}
	// Validasi request
	err := utils.ValidateFields(req, requiredFields)
	if err != nil {
		return nil, err
	}
	// Load konfigurasi dari environment variables
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Gagal memuat konfigurasi: %v", err)
	}
	// Overwrite dengan nilai dari request
	network := req.Network
	if network.Architecture == "EVM" {
		cfg.BlockchainType = "ethereum"
	} else if network.Architecture == "NONEVM" {
		cfg.BlockchainType = "hyperledger"
	} else {
		return nil, errors.New("blockchain_type harus 'ethereum' atau 'quorum'")
	}
	// Overwrite RPCURL jika ada dalam request
	if network.RPCURL != "" {
		cfg.RPCURL = network.RPCURL
	}

	// Buat blockchain client sesuai config
	client, err := CreateClientFactory(cfg)
	if err != nil {
		log.Fatalf("Gagal membuat klien: %v", err)
	}
	// Connect ke blockchain
	if err := client.Connect(); err != nil {
		return nil, errors.New("gagal terhubung ke blockhain")
		// log.Fatalf("Gagal terhubung ke blockchain: %v", err)
	}
	s.client = client
	return &pb.SetConfigResponse{
		Message: fmt.Sprintf("berhasil terhubung ke blockchain:%s ", cfg.BlockchainType),
		// Message: "Konfigurasi blockchain berhasil diperbarui",
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
		NetworkId: uint32(networkID.Uint64()),
	}, nil
}

// GetContractEvents: Mendapatkan daftar event dari smart contract
// func (s *BlockchainService) GetContractEvents(ctx context.Context, req *pb.GetContractEventsRequest) (*pb.GetContractEventsResponse, error) {
// 	if s.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}

// 	// Panggil client untuk mendapatkan event logs
// 	events, err := s.client.GetContractEvents(ctx, req.ContractAddress, req.Abi, req.EventName, req.FromBlock, req.ToBlock)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.GetContractEventsResponse{
// 		Events: events,
// 	}, nil
// }

// // TransferToken: Mengirim token ERC20 dari satu alamat ke alamat lain
// func (s *BlockchainService) TransferToken(ctx context.Context, req *pb.TransferTokenRequest) (*pb.TransferTokenResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}

// 	// Panggil fungsi transfer pada smart contract ERC20
// 	txHash, err := s.client.TransferToken(ctx, req.TokenAddress, req.From, req.To, req.Amount, req.PrivateKey, req.GasLimit)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.TransferTokenResponse{
// 		TxHash: txHash,
// 	}, nil
// }

// // SendETH: Mengirim ETH dari satu alamat ke alamat lain
// func (s *BlockchainService) SendETH(ctx context.Context, req *pb.SendETHRequest) (*pb.SendETHResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"From", "To", "Amount", "PrivateKey"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}

// 	amount := new(big.Int)
// 	amount, ok := amount.SetString(req.Amount, 10)
// 	if !ok {
// 		return nil, errors.New("gagal mengonversi amount ke *big.Int")
// 	}
// 	// Kirim transaksi ETH
// 	txHash, err := s.client.SendETH(ctx, req.PrivateKey, req.To, amount)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.SendETHResponse{
// 		TxHash: txHash,
// 	}, nil
// }

// // GetTokenBalance: Mendapatkan saldo token ERC20 dari smart contract
// func (s *BlockchainService) GetTokenBalance(ctx context.Context, req *pb.GetTokenBalanceRequest) (*pb.GetTokenBalanceResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"TokenAddress", "OwnerAddress"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if req.TokenAddress == "\"\"" || req.OwnerAddress == "\"\"" {
// 		return nil, errors.New("token dan owner address tidak boleh kosong")
// 	}

// 	// Panggil fungsi "balanceOf" dari kontrak ERC20
// 	balance, err := s.client.GetTokenBalance(ctx, req.TokenAddress, req.OwnerAddress)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.GetTokenBalanceResponse{
// 		Balance: balance.String(),
// 	}, nil
// }

// // CallContractMethod: Memanggil fungsi read-only pada smart contract
// func (s *BlockchainService) CallContractMethod(ctx context.Context, req *pb.CallContractMethodRequest) (*pb.CallContractMethodResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}

// 	// Panggil client untuk membaca data dari smart contract
// 	result, err := s.client.CallContractMethod(ctx, req.ContractAddress, req.Abi, req.Method, req.Params)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.CallContractMethodResponse{
// 		Result: result,
// 	}, nil
// }

// // GetContractOwner: Mendapatkan alamat pemilik dari smart contract (jika ada)
// func (s *BlockchainService) GetContractOwner(ctx context.Context, req *pb.GetContractOwnerRequest) (*pb.GetContractOwnerResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}

// 	// Panggil fungsi "owner()" pada smart contract
// 	owner, err := s.client.CallContractMethod(ctx, req.ContractAddress, req.Abi, "owner", nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.GetContractOwnerResponse{
// 		OwnerAddress: owner,
// 	}, nil
// }

// // GetContract: Mendapatkan informasi contract dari blockchain
// func (s *BlockchainService) GetContract(ctx context.Context, req *pb.GetContractRequest) (*pb.GetContractResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"ContractAddress"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// Panggil client untuk mendapatkan informasi contract
// 	bytecode, abi, err := s.client.GetContract(ctx, req.ContractAddress)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.GetContractResponse{
// 		ContractAddress: req.ContractAddress,
// 		Bytecode:        bytecode,
// 		Abi:             abi,
// 	}, nil
// }

// func (s *BlockchainService) DeployIjazahContract(ctx context.Context, req *pb.DeployIjazahContractRequest) (*pb.DeployIjazahContractResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}
// 	// Daftar field yang wajib diisi
// 	requiredFields := []string{"AccountType", "PrivateKey"}
// 	// Validasi request
// 	err := utils.ValidateFields(req, requiredFields)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// if req.AcountType == pb.AccountType_IMPORTED{

// 	// }

// 	// if req.GetUserId() == "\"\"" || req.GetPassword() == "\"\"" {
// 	// 	return nil, errors.New("user dan password tidak boleh kosong")
// 	// }

// 	contractAddress, txHash, err := s.client.DeployIjazahContract(ctx, req.PrivateKey)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &pb.DeployIjazahContractResponse{
// 		ContractAddress: contractAddress,
// 		TxHash:          txHash,
// 	}, nil

// }

// SendTransactionToContract: Mengirim data ke smart contract dengan memanggil fungsi tertentu
// func (s *BlockchainService) SendTransactionToContract(ctx context.Context, req *pb.SendTransactionToContractRequest) (*pb.SendTransactionToContractResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}

// 	// Kirim transaksi ke smart contract
// 	txHash, err := s.client.SendTransactionToContract(ctx, req.ContractAddress, req.Abi, req.Method, req.Params, req.PrivateKey, req.GasLimit)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.SendTransactionToContractResponse{
// 		TxHash: txHash,
// 	}, nil
// }

// GetConsensusAlgorithm: Mendapatkan algoritma konsensus (hanya untuk Quorum)
// func (s *BlockchainService) GetConsensusAlgorithm(ctx context.Context, _ *pb.Empty) (*pb.ConsensusAlgorithmResponse, error) {
// 	// Periksa apakah client adalah QuorumClient
// 	quorumClient, ok := s.client.(QuorumClient)
// 	if !ok {
// 		return nil, errors.New("fitur ini hanya tersedia untuk Quorum")
// 	}

// 	consensus, err := quorumClient.GetConsensusAlgorithm(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.ConsensusAlgorithmResponse{
// 		ConsensusAlgorithm: consensus,
// 	}, nil
// }

// ApproveToken: Memberikan izin kepada smart contract lain untuk menggunakan token ERC20
// func (s *BlockchainService) ApproveToken(ctx context.Context, req *pb.ApproveTokenRequest) (*pb.ApproveTokenResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}

// 	// Panggil fungsi "approve" dari kontrak ERC20
// 	txHash, err := s.client.SendTransactionToContract(ctx, req.TokenAddress, req.Abi, "approve", []string{req.Spender, req.Amount}, req.PrivateKey, req.GasLimit)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.ApproveTokenResponse{
// 		TxHash: txHash,
// 	}, nil
// }
// GetTokenAllowance: Mengecek jumlah token ERC20 yang telah diizinkan untuk digunakan oleh smart contract lain
// func (s *BlockchainService) GetTokenAllowance(ctx context.Context, req *pb.GetTokenAllowanceRequest) (*pb.GetTokenAllowanceResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}

// 	// Panggil fungsi "allowance" dari kontrak ERC20
// 	allowance, err := s.client.CallContractMethod(ctx, req.TokenAddress, req.Abi, "allowance", []string{req.Owner, req.Spender})
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.GetTokenAllowanceResponse{
// 		Allowance: allowance,
// 	}, nil
// }

// // GetGasPrice: Mendapatkan harga gas saat ini di jaringan blockchain
// func (s *BlockchainService) GetGasPrice(ctx context.Context, req *pb.GetGasPriceRequest) (*pb.GetGasPriceResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}

// 	// Ambil harga gas dari client
// 	gasPrice, err := s.client.GetGasPrice(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.GetGasPriceResponse{
// 		GasPrice: gasPrice,
// 	}, nil
// }

// // GetNonce: Mendapatkan nonce dari alamat tertentu
// func (s *BlockchainService) GetNonce(ctx context.Context, req *pb.GetNonceRequest) (*pb.GetNonceResponse, error) {
// 	if s.client.client == nil {
// 		return nil, errors.New("client belum dikonfigurasi")
// 	}

// 	nonce, err := s.GetNonce(ctx, req.Address)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.GetNonceResponse{
// 		Nonce: nonce,
// 	}, nil
// }

// func (s *BlockchainService) validateRequest(req any, requiredFields []string, checkEmptyFields map[string]func() string) error {
// 	if s.client.client == nil {
// 		return errors.New("client belum dikonfigurasi")
// 	}

// 	// Validasi apakah field-field wajib ada
// 	if err := utils.ValidateFields(req, requiredFields); err != nil {
// 		return err
// 	}

// 	// Validasi apakah field wajib kosong ("" atau nilai lain yang dianggap kosong)
// 	for field, getter := range checkEmptyFields {
// 		if getter() == "" || getter() == "\"\"" { // Sesuaikan dengan format yang mungkin terjadi
// 			return fmt.Errorf("%s tidak boleh kosong", field)
// 		}
// 	}

//		return nil
//	}
//
