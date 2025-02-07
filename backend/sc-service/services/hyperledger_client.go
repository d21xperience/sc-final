package services

// import (
// 	"crypto/x509"
// 	"fmt"
// 	"os"

// 	"github.com/hyperledger/fabric-gateway/pkg/client"
// 	"github.com/hyperledger/fabric-gateway/pkg/hash"
// 	"github.com/hyperledger/fabric-gateway/pkg/identity"
// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials"
// )

// // HyperledgerFabricClient menggunakan Fabric Gateway
// type HyperledgerFabricClient struct {
// 	gateway *client.Gateway
// 	network *client.Network
// }

// // NewHyperledgerFabricClient membuat koneksi ke Hyperledger Fabric menggunakan Fabric Gateway
// func NewHyperledgerFabricClient(cfg *Config) (BlockchainClient, error) {
// 	// Load identitas pengguna dari file sertifikat
// 	identity, err := loadIdentity(cfg.CertPath, cfg.PeerHostOverride, cfg.MSPID)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal memuat identitas: %v", err)
// 	}

// 	connection, err := grpc.NewClient(cfg.RPCURL, grpc.WithTransportCredentials(identity))
// 	if err != nil {
// 		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
// 	}

// 	// Buat koneksi ke gateway Fabric
// 	gw, err := client.Connect(
// 		client.WithIdentity(identity),
// 		client.WithEndpoint(cfg.PeerEndpoint),
// 	)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal terhubung ke jaringan Hyperledger Fabric: %v", err)
// 	}

// 	// Ambil jaringan (channel) dari gateway
// 	network := gw.GetNetwork(cfg.Channel)

// 	return &HyperledgerFabricClient{gateway: gw, network: network}, nil
// }

// // loadIdentity memuat sertifikat dan kunci dari file
// func loadIdentity(certPath, peerHostOverride, mspID string) (credentials.TransportCredentials, error) {
// 	// Load sertifikat pengguna dari file
// 	certificatePEM, err := os.ReadFile(certPath)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal membaca sertifikat: %v", err)
// 	}
// 	certificate, err := identity.CertificateFromPEM(certificatePEM)
// 	if err != nil {
// 		panic(err)
// 	}
// 	certPool := x509.NewCertPool()
// 	certPool.AddCert(certificate)
// 	transportCredentials := credentials.NewClientTLSFromCert(certPool, peerHostOverride)

// 	return transportCredentials, nil
// }

// // Connect menghubungkan ke jaringan Hyperledger Fabric
// func (h *HyperledgerFabricClient) Connect() error {
// 	// network, err := h.gateway.GetNetwork("mychannel")
// 	// if err != nil {
// 	// 	return fmt.Errorf("gagal mendapatkan jaringan Hyperledger Fabric: %w", err)
// 	// }
// 	// log.Println("Terhubung ke Hyperledger Fabric di channel", network)
// 	// return nil
// 	clientConnection := newGrpcConnection()
// 	defer clientConnection.Close()
// 	gw, err := client.Connect(
// 		newIdentity(),
// 		client.WithSign(newSign()),
// 		client.WithHash(hash.SHA256),
// 		client.WithClientConnection(clientConnection),
// 	)
// 	// log.Println("Terhubung ke Hyperledger Fabric di channel", network)
// 	return nil
// }

// func newIdentity() *identity.X509Identity {
// 	certificatePEM, err := os.ReadFile(certPath)
// 	if err != nil {
// 		panic(fmt.Errorf("failed to read certificate file: %w", err))
// 	}

// 	certificate, err := identity.CertificateFromPEM(certificatePEM)
// 	if err != nil {
// 		panic(err)
// 	}

// 	id, err := identity.NewX509Identity("Org1MSP", certificate)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return id
// }

// func newSign() identity.Sign {
// 	privateKeyPEM, err := readFirstFile(keyPath)
// 	if err != nil {
// 		panic(fmt.Errorf("failed to read private key file: %w", err))
// 	}

// 	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
// 	if err != nil {
// 		panic(err)
// 	}

// 	sign, err := identity.NewPrivateKeySign(privateKey)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return sign
// }
