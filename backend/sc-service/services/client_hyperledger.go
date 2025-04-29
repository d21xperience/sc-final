package services

import (
	"context"
	"crypto/x509"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/hash"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// HyperledgerFabricClient menggunakan Fabric Gateway
type HyperledgerFabricClient struct {
	gateway *client.Gateway
	network *client.Network
}

// NewHyperledgerFabricClient membuat koneksi ke Hyperledger Fabric menggunakan Fabric Gateway
func NewHyperledgerFabricClient(cfg *Config) (BlockchainClient, error) {
	// Load identitas pengguna dari file sertifikat
	identity, err := loadIdentity(cfg.CertPath, cfg.PeerHostOverride)
	if err != nil {
		return nil, fmt.Errorf("gagal memuat identitas: %v", err)
	}

	connection, err := grpc.NewClient(cfg.RPCURL, grpc.WithTransportCredentials(identity))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}
	// clientConnection := newGrpcConnection()
	// defer clientConnection.Close()
	// Buat koneksi ke gateway Fabric
	gw, err := client.Connect(
		newIdentity(cfg.CertPath),
		client.WithSign(newSign(cfg.KeyPath)),
		client.WithHash(hash.SHA256),
		client.WithClientConnection(connection),
	)
	if err != nil {
		return nil, fmt.Errorf("gagal terhubung ke jaringan Hyperledger Fabric: %v", err)
	}

	// Ambil jaringan (channel) dari gateway
	network := gw.GetNetwork(cfg.Channel)

	return &HyperledgerFabricClient{gateway: gw, network: network}, nil
}

// loadIdentity memuat sertifikat dan kunci dari file
func loadIdentity(certPath, peerHostOverride string) (credentials.TransportCredentials, error) {
	// Load sertifikat pengguna dari file
	certificatePEM, err := os.ReadFile(certPath)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca sertifikat: %v", err)
	}
	certificate, err := identity.CertificateFromPEM(certificatePEM)
	if err != nil {
		panic(err)
	}
	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, peerHostOverride)

	return transportCredentials, nil
}

// Connect menghubungkan ke jaringan Hyperledger Fabric
func (h *HyperledgerFabricClient) Connect() error {
	network := h.gateway.GetNetwork("mychannel")
	log.Println("terhubung ke Hyperledger Fabric di channel", network)
	return nil
}

func newIdentity(certPath string) *identity.X509Identity {
	certificatePEM, err := os.ReadFile(certPath)
	if err != nil {
		panic(fmt.Errorf("failed to read certificate file: %w", err))
	}

	certificate, err := identity.CertificateFromPEM(certificatePEM)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity("Org1MSP", certificate)
	if err != nil {
		panic(err)
	}

	return id
}

func newSign(keyPath string) identity.Sign {
	privateKeyPEM, err := readFirstFile(keyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}

func (h *HyperledgerFabricClient) NetworkID(ctx context.Context) (*big.Int, error) {
	return nil, fmt.Errorf("GetNetworkID tidak didukung di Hyperledger Fabric")
}

func (h *HyperledgerFabricClient) GenerateNewAccount(ctx context.Context, userId int32, password string) (map[string]interface{}, error) {
	return nil, errors.New("tes")
}

// func (h *HyperledgerFabricClient) IssueDegree(ctx context.Context, contractAddress string, degreeHash [32]byte, sekolah string, issueDate uint64, privateKey string, gasLimit uint64) (string, error) {
// 	network, := h.gateway.GetNetwork("mychannel")

// 	contract := network.GetContract("degree_contract") // Nama smart contract
// 	result, err := contract.SubmitTransaction("IssueDegree", string(degreeHash[:]), sekolah, fmt.Sprintf("%d", issueDate))
// 	if err != nil {
// 		return "", fmt.Errorf("gagal mengirim transaksi ke Hyperledger: %w", err)
// 	}

// 	return string(result), nil
// }

// readFirstFile membaca file pertama dalam folder atau satu file langsung
func readFirstFile(path string) ([]byte, error) {
	// Jika path adalah file, langsung baca file tersebut
	if fileInfo, err := os.Stat(path); err == nil && !fileInfo.IsDir() {
		return os.ReadFile(path)
	}

	// Jika path adalah folder, ambil file pertama di dalamnya
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca direktori: %v", err)
	}

	// Pastikan folder tidak kosong
	if len(files) == 0 {
		return nil, fmt.Errorf("tidak ada file dalam folder: %s", path)
	}

	// Ambil file pertama
	firstFile := filepath.Join(path, files[0].Name())
	return os.ReadFile(firstFile)
}
