package services

import (
	"errors"
)

// CreateClientFactory adalah factory untuk membuat Ethereum atau Quorum client
func CreateClientFactory(config *Config) (EthClient, error) {
	switch config.BlockchainType {
	case "ethereum":
		client, err := NewDefaultEthClient(config.RPCURL)
		if err != nil {
			return nil, err
		}
		return client, nil
	case "quorum":
		client, err := NewDefaultQuorumClient(config.RPCURL)
		if err != nil {
			return nil, err
		}
		return client, nil
	default:
		return nil, errors.New("BlockchainType tidak dikenal: gunakan 'ethereum' atau 'quorum'")
	}
}
