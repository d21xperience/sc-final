package pkg

// import (
// 	"fmt"
// 	"net/rpc"
// )

// type EthereumClient struct {
// 	ClientIinfo *EthereumClientInfo
// 	Client      *rpc.Client
// }

// func (e *EthereumClient) Connect() error {
// 	rpcClient, err := rpc.Dial(e.ClientIinfo.NetURL)
// 	if err != nil {
// 		return fmt.Errorf("failed to connect to Ethereum: %v", err)
// 	}
// 	e.Client = rpcClient
// 	return nil
// }

// Implement SendCrypto, DeployContract, InteractWithContract sesuai dengan Ethereum
