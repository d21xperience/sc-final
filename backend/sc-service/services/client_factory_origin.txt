package services

import (
	"errors"
)

// CreateClientFactory adalah factory untuk membuat Ethereum atau Quorum client
func CreateClientFactory(config *Config) (EthClient, error) {
	// switch config.BlockchainType {
	// case "ethereum":
	// 	client, err := NewDefaultEthClient(config.RPCURL)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return client, nil
	// case "quorum":
	// 	rpcURL := config.RPCURL
	// 	ETHclient, err := NewDefaultEthClient(rpcURL)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	client, err := NewDefaultQuorumClient(rpcURL, ETHclient)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	return client, nil
	// default:
	// 	return nil, errors.New("BlockchainType tidak dikenal: gunakan 'ethereum' atau 'quorum'")
	// }
	switch config.NetworkId {
	case 37:
		client, err := NewDefaultEthClient(config.RPCURL)
		if err != nil {
			return nil, err
		}
		return client, nil
	case 38:
		rpcURL := config.RPCURL
		ETHclient, err := NewDefaultEthClient(rpcURL)
		if err != nil {
			return nil, err
		}
		client, err := NewDefaultQuorumClient(rpcURL, ETHclient)
		if err != nil {
			return nil, err
		}
		return client, nil
	default:
		return nil, errors.New("BlockchainType tidak dikenal: gunakan 'ethereum' atau 'quorum'")
	}
}
