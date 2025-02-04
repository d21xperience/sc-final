package config

import (
	"log"
	"sc-service/models"
	"time"

	"gorm.io/gorm"
)

func SeedBCNetworks() {
	networks := []models.BlockchainNetwork{
		{
			NetworkName:    "Hyperledger fabric",
			BlockchainType: "Private",
			ChainID:        1,
			RPCURL:         "localhot:6786",
			BlockExplorer:  "https://blockscout",
			Description:    "",
		},
		{

			NetworkName:    "Quorum",
			BlockchainType: "Private",
			ChainID:        12,
			RPCURL:         "",
			BlockExplorer:  "",
			Description:    "",
		},
		{

			NetworkName:    "Ethereum Mainnet",
			BlockchainType: "Public",
			ChainID:        1,
			RPCURL:         "https://mainnet.infura.io/v3/",
			BlockExplorer:  "https://etherscan.io",
		},
		{

			NetworkName:    "Binance Smart Chain (BSC)",
			BlockchainType: "Public",
			ChainID:        56,
			RPCURL:         "https://bsc-dataseed.binance.org",
			BlockExplorer:  "https://bscscan.com",
		},
		{

			NetworkName:    "Polygon (Matic) Mainnet",
			BlockchainType: "Public",
			ChainID:        137,
			RPCURL:         "https://polygon-rpc.com",
			BlockExplorer:  "https://polygonscan.com",
		},
		{

			NetworkName:    "Avalanche C-Chain",
			BlockchainType: "Public",
			ChainID:        43114,
			RPCURL:         "https://api.avax.network/ext/bc/C/rpc",
			BlockExplorer:  "https://snowtrace.io",
		},
		{

			NetworkName:    "Fantom Opera",
			BlockchainType: "Public",
			ChainID:        250,
			RPCURL:         "https://rpc.ftm.tools",
			BlockExplorer:  "https://ftmscan.com",
		},
		{

			NetworkName:    "Arbitrum One",
			BlockchainType: "Public",
			ChainID:        42161,
			RPCURL:         "https://arb1.arbitrum.io/rpc",
			BlockExplorer:  "https://arbiscan.io",
		},
		{

			NetworkName:    "Optimism",
			BlockchainType: "Public",
			ChainID:        10,
			RPCURL:         "https://mainnet.optimism.io",
			BlockExplorer:  "https://optimistic.etherscan.io",
		},
		{

			NetworkName:    "Ethereum Goerli Testnet",
			BlockchainType: "Public",
			ChainID:        5,
			RPCURL:         "https://goerli.infura.io/v3/",
			BlockExplorer:  "https://goerli.etherscan.io",
		},
		{

			NetworkName:    "Binance Smart Chain Testnet",
			BlockchainType: "Public",
			ChainID:        97,
			RPCURL:         "https://data-seed-prebsc-1-s1.binance.org:8545",
			BlockExplorer:  "https://testnet.bscscan.com",
		},
	}

	for _, network := range networks {
		if err := DB.Create(&network).Error; err != nil {
			log.Fatalf("Failed to seed blockchain networks: %v", err)
		}
	}

	log.Println("Seeding completed!")
}
// func SeedAccountBC() {
// 	accounts := []models.Account{
// 		{
// 			Address:      "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
// 			Name:         "account - 2",
// 			NetworkID:    72,
// 			Organization: "instansi", // Untuk Hyperledger Fabric
// 			Amount:       100,
// 			IsActive:     false,
// 		},
// 		{
// 			Address:      "0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
// 			Name:         "account - 3",
// 			NetworkID:    72,
// 			Organization: "instansi", // Untuk Hyperledger Fabric
// 			Amount:       100,
// 			IsActive:     false,
// 		},
// 		{
// 			Address:      "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC",
// 			Name:         "account - 4",
// 			NetworkID:    72,
// 			Organization: "instansi", // Untuk Hyperledger Fabric
// 			Amount:       100,
// 			IsActive:     false,
// 		},
// 		{
// 			Address:      "0x90F79bf6EB2c4f870365E785982E1f101E93b906",
// 			Name:         "account - 5",
// 			NetworkID:    72,
// 			Organization: "instansi", // Untuk Hyperledger Fabric
// 			Amount:       100,
// 			IsActive:     false,
// 		},
// 		{
// 			Address:      "0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65",
// 			Name:         "account - 6",
// 			NetworkID:    72,
// 			Organization: "instansi", // Untuk Hyperledger Fabric
// 			Amount:       100,
// 			IsActive:     false,
// 		},
// 	}
// 	for _, account := range accounts {
// 		if err := DB.Create(&account).Error; err != nil {
// 			log.Fatalf("Failed to seed blockchain networks: %v", err)
// 		}
// 	}
// }

func SeedStudent() {
	// Students Data
	students := []models.Student{
		{
			StudentID:   1,
			Nama:        "John Doe",
			NIS:         "12345678",
			NISN:        "87654321",
			NIK:         "87654323",
			TptLahir:    time.Date(2005, 10, 12, 0, 0, 0, 0, time.UTC),
			TglLahir:    time.Date(2005, 10, 12, 0, 0, 0, 0, time.UTC),
			AsalSekolah: "SMA Negeri 1",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			StudentID:   2,
			Nama:        "Jane Smith",
			NIS:         "23456789",
			NISN:        "98765432",
			NIK:         "98765434",
			TptLahir:    time.Date(2006, 5, 21, 0, 0, 0, 0, time.UTC),
			TglLahir:    time.Date(2006, 5, 21, 0, 0, 0, 0, time.UTC),
			AsalSekolah: "SMA Negeri 2",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	// Certificates Data
	certificates := []models.Certificate{
		{
			CertificateID:   1,
			StudentID:       1,
			CertificateHash: "hash123...",
			IPFSCID:         "Qm12345...",
			IssueDate:       time.Now(),
			IssuerName:      "Ministry of Education",
			NilaiRata:       "85.5",
			BlockchainType:  "Ethereum",
			TransactionHash: "tx456...",
			ContractAddress: "0xcontract123...",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
		{
			CertificateID:   2,
			StudentID:       2,
			CertificateHash: "hash456...",
			IPFSCID:         "Qm67890...",
			IssueDate:       time.Now(),
			IssuerName:      "Ministry of Education",
			NilaiRata:       "90.0",
			BlockchainType:  "Hyperledger Fabric",
			TransactionHash: "tx789...",
			ContractAddress: "0xcontract456...",
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
		},
	}

	// Insert data into database
	err := DB.Transaction(func(tx *gorm.DB) error {
		// Seed Students
		if err := tx.Create(&students).Error; err != nil {
			return err
		}

		// Seed Certificates
		if err := tx.Create(&certificates).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Seeding failed: %v", err)
	}

	log.Println("Database seeded successfully.")
}

//func dbSeed() {
// 	// Dummy data (tambahkan jika database kosong)
// 	DB.Create(&models.BlockchainNetwork{
// 		NetworkName: "Ethereum-Local",
// 		Unit:        "ETH",
// 		Account: []models.Account{
// 			{Name: "Account 1", Address: "39TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTp", Amount: 100, NetworkID: 1},
// 			{Name: "Account 2", Address: "9TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTpU", Amount: 200, NetworkID: 1},
// 		},
// 	})

// 	DB.Create(&models.BlockchainNetwork{
// 		NetworkName: "Quorum",
// 		Unit:        "ETH",
// 		Account: []models.Account{
// 			{Name: "quo 1", Address: "quuro9TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTp", Amount: 500, NetworkID: 2},
// 			{Name: "quo 2", Address: "quuro2TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTpU", Amount: 120, NetworkID: 2},
// 		},
// 	})

// 	DB.Create(&models.BlockchainNetwork{
// 		NetworkName: "Hyperledger Fabric",
// 		Unit:        "No limit",
// 		Account: []models.Account{
// 			{Name: "HF-01", Address: "39TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTp", Amount: 350, NetworkID: 3},
// 			{Name: "HF-02", Address: "10TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTpU", Amount: 130, NetworkID: 3},
// 		},
// 	})
// }
