package services

// import (
// 	"fmt"

// )

// // HyperledgerClient struct menyimpan instance Fabric Gateway
// type HyperledgerClient struct {
// 	Gateway *gateway.Gateway
// }

// // Connect membuat koneksi ke Hyperledger Fabric
// func (h *HyperledgerClient) Connect() error {
// 	_, err := h.Gateway.GetNetwork("mychannel")
// 	if err != nil {
// 		return fmt.Errorf("gagal mendapatkan jaringan Fabric: %w", err)
// 	}
// 	return nil
// }

// // NewHyperledgerFabricClient membuat klien Hyperledger Fabric
// func NewHyperledgerFabricClient(cfg *Config) (BlockchainClient, error) {
// 	// Load wallet
// 	wallet, err := gateway.NewFileSystemWallet(cfg.FabricWallet)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal memuat wallet: %w", err)
// 	}

// 	// Buat koneksi ke gateway
// 	gw, err := gateway.Connect(
// 		gateway.WithConfigFromFile(cfg.FabricConfigPath),
// 		gateway.WithIdentity(wallet, cfg.FabricIdentity),
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal membuat Hyperledger Fabric client: %w", err)
// 	}

// 	return &HyperledgerClient{Gateway: gw}, nil
// }
