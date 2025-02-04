package services

// import (
// 	"context"
// 	"fmt"
// 	"math/big"
// 	"path/filepath"
// 	"sc-service/config"
// 	"sc-service/models"
// 	"sc-service/repositories"
// 	"sc-service/utils"

// 	verifikasiIjazah "sc-service/smartcontract/gen"

// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/accounts/keystore"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"github.com/ethereum/go-ethereum/rpc"
// )

// // Default implementasi EthClient menggunakan go-ethereum dan RPC
// type Account struct {
// 	rpcClient *rpc.Client
// 	client    *ethclient.Client
// 	repo      *repositories.GenericRepository[models.Account]
// }

// func NewAccount(rawUrl string) (*Account, error) {
// 	// Buat koneksi RPC
// 	rpcClient, err := rpc.Dial(rawUrl)
// 	if err != nil {
// 		return nil, fmt.Errorf("gagal menghubungkan ke RPC: %v", err)
// 	}
// 	repo := repositories.NewAccountRepository(config.DB)
// 	// Gunakan ethclient sebagai wrapper untuk RPC
// 	client := ethclient.NewClient(rpcClient)

// 	return &Account{
// 		rpcClient: rpcClient,
// 		client:    client, // Sekarang client diinisialisasi
// 		repo:      repo,
// 	}, nil
// }

// // =============================
// // =============Akun============
// func (c *Account) CreateAccount(ctx context.Context, userId int32, accountType, password string) (string, error) {
// 	key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)

// 	a, err := key.NewAccount(password)
// 	if err != nil {
// 		return "", err
// 	}
// 	// simpan ke database
// 	pass, err := utils.EncryptPassword(password)
// 	if err != nil {
// 		return "", err
// 	}
// 	simpan := c.repo.Save(ctx, &models.Account{
// 		UserID:            userId,
// 		Address:           a.Address.Hex(),
// 		Type:              models.AccountType(accountType),
// 		Password:          pass,
// 		KeystrokeFilename: filepath.Base(a.URL.Path),
// 	}, "public")
// 	if simpan != nil {
// 		return "", simpan
// 	}

// 	return string(a.Address.Hex()), nil
// }
// func (c *Account) DeployIjazahContract(ctx context.Context, privateKeyHex string) (contracAddress string, txHash string, err error) {

// 	pvKey, pubKey, err := ConvertStringPrivateKey(privateKeyHex)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	// akun, err := c.GetAccounts(ctx, userId)
// 	// if err != nil {
// 	// 	return "","", err
// 	// }
// 	// // baca file utc
// 	// wd, err := os.Getwd()
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to get working directory: %v", err)
// 	// }

// 	// // Menggunakan filepath.Join agar sesuai dengan OS
// 	// path := filepath.Join(wd, "wallet", akun[0].WalletFilename)
// 	// b, err := os.ReadFile(path)
// 	// if err != nil {
// 	// 	return "","", err
// 	// }

// 	// pas := utils.VerifyPassword(password, akun[0].Password)
// 	// if pas {
// 	// key, err := keystore.DecryptKey(b, password)
// 	// if err != nil {
// 	// 	return "","", err
// 	// }
// 	// add := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)

// 	nonce, err := c.client.PendingNonceAt(ctx, pubKey)
// 	if err != nil {
// 		return "", "", err
// 	}
// 	gasPrice, err := c.client.SuggestGasPrice(ctx)
// 	if err != nil {
// 		return "", "", err
// 	}
// 	chainId, err := c.client.ChainID(ctx)
// 	if err != nil {
// 		return "", "", err
// 	}
// 	gasLimit := uint64(3000000)
// 	// auth := TransactOptsAuth(key, chainId, gasPrice, nonce, gasLimit)
// 	auth, err := bind.NewKeyedTransactorWithChainID(pvKey, chainId)
// 	if err != nil {
// 		return "", "", err
// 	}
// 	auth.Nonce = big.NewInt(int64(nonce))
// 	auth.Value = big.NewInt(0)
// 	auth.GasLimit = gasLimit
// 	auth.GasPrice = gasPrice
// 	a, tx, _, err := verifikasiIjazah.DeployVerifikasiIjazah(auth, c.client)
// 	if err != nil {
// 		return "", "", err
// 	}
// 	fmt.Println(tx.Hash().Hex())
// 	return a.Hex(), tx.Hash().Hex(), nil

// 	// }
// 	// return "", "",nil
// }

// func (c *Account) GetAccounts(ctx context.Context, userId string) ([]*models.WalletTable, error) {
// 	var modelTableWallet []*models.WalletTable
// 	var err error
// 	if userId == "" {
// 		modelTableWallet, err = c.repo.FindAll(ctx, "public", 100, 0)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return modelTableWallet, nil
// 	}

// 	var condition = map[string]interface{}{
// 		"user_id": userId,
// 	}
// 	modelTableWallet, err = c.repo.FindAllByConditions(ctx, "public", condition, 100, 0)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return modelTableWallet, nil
// }

// // =============================
