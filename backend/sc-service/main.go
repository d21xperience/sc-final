package main

import "sc-service/server"

func main() {
	server.StartServer()
}


// // Load konfigurasi dari environment variables
	// config, err := services.LoadConfig()
	// if err != nil {
	// 	log.Fatalf("Gagal memuat konfigurasi: %v", err)
	// }

	// // Buat client berdasarkan konfigurasi
	// client, err := services.CreateClientFactory(config)
	// if err != nil {
	// 	log.Fatalf("Gagal membuat client: %v", err)
	// }

	// // Contoh penggunaan client
	// ctx := context.Background()

	// // Network ID
	// networkID, err := client.NetworkID(ctx)
	// if err != nil {
	// 	log.Fatalf("Gagal mendapatkan NetworkID: %v", err)
	// }
	// fmt.Printf("Network ID: %s\n", networkID.String())

	// // Jika Quorum, gunakan fitur tambahan
	// if quorumClient, ok := client.(services.QuorumClient); ok {
	// 	consensus, err := quorumClient.GetConsensusAlgorithm(ctx)
	// 	if err != nil {
	// 		log.Fatalf("Gagal mendapatkan algoritma konsensus: %v", err)
	// 	}
	// 	fmt.Printf("Konsensus Quorum: %s\n", consensus)
	// }
	// // Inisialisasi service Ethereum
	// ethClientService, err := services.NewRealEthClient("https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ethServiceServer := &server.EthServiceServer{
	// 	EthClientService: ethClientService,
	// }