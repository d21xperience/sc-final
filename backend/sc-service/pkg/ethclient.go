package pkg

type ETHClientInfo struct {
}

type ETHClient struct {
	NetURL   string
	LocalURL string
	GasLimit uint64
}

// func NewETHClient(netURL, pvKeyHex string) *ETHClient {
// 	return &ETHClient{
// 		Client: ETHClientInfo{
// 			NetURL:        netURL,
// 			PvKeyHex:      pvKeyHex,
// 			PubAddressKey: publicAddress,
// 			ChainID:       chainID,
// 			GasPrice:      gasPrice,
// 			PvKey:         privateKey,
// 			GasLimit:      5000000,
// 			EthClient:     client,
// 		},
// 	}
// }

// Membuat transaksi dengan signed transactor
// func (e *ETHClient) AuthOfSC() *bind.TransactOpts {
// 	auth, err := bind.NewKeyedTransactorWithChainID(e.Client.PvKey, e.Client.ChainID)
// 	if err != nil {
// 		log.Fatalf("Failed to create transactor: %v", err)
// 	}
// 	auth.GasPrice = e.Client.GasPrice
// 	auth.GasLimit = e.Client.GasLimit
// 	auth.From = e.Client.PubAddressKey
// 	return auth
// }
