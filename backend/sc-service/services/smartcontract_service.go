package services

// type SmartcontractService interface {
// 	LoginToContract(netURL, pvKeyHex string, chainID *big.Int, contractAddress string) error
// }

// type smartcontractServiceImpl struct {
// 	contract *verval_ijazah.VervalIjazah
// }

// func NewSmartcontractService(contract *verval_ijazah.VervalIjazah) SmartcontractService {
// 	return &smartcontractServiceImpl{
// 		contract: contract,
// 	}
// }

// func (s *ijazahServiceImpl) LoginToContract(netURL, pvKeyHex string, chainID *big.Int, contractAddress string) error {
// 	var err error
// 	// Inisialisasi kontrak
// 	cAdd := common.HexToAddress(contractAddress)
// 	s.contract, err = verval_ijazah.NewVervalIjazah(cAdd, client.Client.EthClient)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal menginisialisasi kontrak: %w", err)
// 	}
// }

// // FetchIjazahData mengambil data dari smart contract
// func (s *smartcontractServiceImpl) FetchIjazahData(id string) (string, string, error) {
// 	fromAddress := s.client.Client.PubAddressKey
// 	_, nama, _, noIjazah, _, _, err := s.contract.Get(&bind.CallOpts{
// 		From: fromAddress,
// 	}, id)
// 	if err != nil {
// 		return "", "", fmt.Errorf("gagal mengambil data dari kontrak: %w", err)
// 	}
// 	return nama, noIjazah, nil
// }
