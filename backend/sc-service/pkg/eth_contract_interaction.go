package pkg

// Deploy smart contract
// func (e *ETHClient) DeployContract() (common.Address, string, error) {
// 	auth := e.AuthOfSC()
// 	contractAddress, tx, _, err := verval_ijazah.DeployVervalIjazah(auth, e.Client.EthClient)
// 	if err != nil {
// 		return common.Address{}, "", fmt.Errorf("failed to deploy contract: %w", err)
// 	}
// 	fmt.Printf("Contract deployed at: %s, tx hash: %s\n", contractAddress.Hex(), tx.Hash().Hex())
// 	return contractAddress, tx.Hash().Hex(), nil
// }

// // Interact with a smart contract
// func (e *ETHClient) InteractWithContract(contractAddress, dataID string) (string, string, error) {
// 	cAdd := common.HexToAddress(contractAddress)
// 	contract, err := verval_ijazah.NewVervalIjazah(cAdd, e.Client.EthClient)
// 	if err != nil {
// 		return "", "", fmt.Errorf("failed to initialize contract: %w", err)
// 	}

// 	_, nama, _, noIjazah, _, _, err := contract.Get(&bind.CallOpts{From: e.Client.PubAddressKey}, dataID)
// 	if err != nil {
// 		return "", "", fmt.Errorf("failed to fetch data from contract: %w", err)
// 	}
// 	return nama, noIjazah, nil
// }

// // Fungsi untuk menginisialisasi kontrak
// func (e *ETHClient) initializeContract(contractAddress string) (*verval_ijazah.VervalIjazah, error) {
// 	cAdd := common.HexToAddress(contractAddress)
// 	contract, err := verval_ijazah.NewVervalIjazah(cAdd, e.Client.EthClient)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal menginisialisasi kontrak: %w", err)
// 	}
// 	return contract, nil
// }

// // Fungsi untuk membuat transactor
// func (e *ETHClient) createTransactor() (*bind.TransactOpts, error) {
// 	transactor, err := bind.NewKeyedTransactorWithChainID(e.Client.PvKey, e.Client.ChainID)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal membuat transactor: %w", err)
// 	}
// 	transactor.GasLimit = 3000000
// 	transactor.GasPrice = e.Client.GasPrice
// 	return transactor, nil
// }

// // Fungsi untuk mengambil data dari kontrak
// func (e *ETHClient) fetchData(contract *verval_ijazah.VervalIjazah, id string) (string, string, error) {
// 	fromAddress := e.Client.PubAddressKey
// 	_, nama, _, noIjazah, _, _, err := contract.Get(&bind.CallOpts{
// 		From: fromAddress,
// 	}, id)
// 	if err != nil {
// 		return "", "", fmt.Errorf("gagal mengambil data dari kontrak: %w", err)
// 	}
// 	return nama, noIjazah, nil
// }
