package config

import (
	"log"
	"sc-service/models"
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
