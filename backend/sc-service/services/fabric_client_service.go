package services

// import (
// 	"fmt"

// 	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
// )

// // FabricClient interface khusus untuk Hyperledger Fabric
// type FabricClient interface {
// 	SubmitTransaction(channel string, chaincode string, function string, args ...string) (string, error)
// }

// // HyperledgerClient struct untuk menyimpan instance Gateway Fabric
// type HyperledgerClient struct {
// 	Gateway *gateway.Gateway
// }

// // SubmitTransaction mengirim transaksi ke Fabric
// func (h *HyperledgerClient) SubmitTransaction(channel string, chaincode string, function string, args ...string) (string, error) {
// 	network, err := h.Gateway.GetNetwork(channel)
// 	if err != nil {
// 		return "", fmt.Errorf("gagal mendapatkan jaringan: %w", err)
// 	}

// 	contract := network.GetContract(chaincode)
// 	result, err := contract.SubmitTransaction(function, args...)
// 	if err != nil {
// 		return "", fmt.Errorf("gagal menjalankan transaksi: %w", err)
// 	}

// 	return string(result), nil
// }

// // NewHyperledgerFabricClient membuat klien untuk Hyperledger Fabric
// func NewHyperledgerFabricClient(config *Config) (FabricClient, error) {
// 	// Load wallet
// 	wallet, err := gateway.NewFileSystemWallet(config.FabricWallet)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal memuat wallet: %w", err)
// 	}

// 	// Buat koneksi ke gateway
// 	gw, err := gateway.Connect(
// 		gateway.WithConfigFromFile(config.FabricConfigPath), // Menggunakan file YAML sebagai konfigurasi
// 		gateway.WithIdentity(wallet, config.FabricIdentity),
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal membuat Hyperledger Fabric client: %w", err)
// 	}

// 	return &HyperledgerClient{Gateway: gw}, nil
// }
