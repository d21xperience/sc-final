package controller

import (
	"net/http"
	"sc-service/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// import (
// 	"fmt"
// 	"math/big"
// 	"net/http"
// 	"sc-service/config"
// 	"sc-service/pkg"

// 	"github.com/gin-gonic/gin"
// )

// func getBlockchainClient(platform string) (pkg.Blockchain, error) {
// 	switch platform {
// 	case "ethereum":
// 		client := &pkg.EthereumClient{ClientInfo: config.ETHClientInfo}
// 		client.Connect()
// 		return client, nil
// 	case "quorum":
// 		client := &pkg.QuorumClient{}
// 		client.Connect()
// 		return client, nil
// 	case "fabric":
// 		client := &pkg.FabricClient{}
// 		client.Connect()
// 		return client, nil
// 	default:
// 		return nil, fmt.Errorf("unsupported blockchain platform")
// 	}
// }

// func SendCrypto(c *gin.Context) {
// 	platform := c.Query("platform")
// 	client, err := getBlockchainClient(platform)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	address := c.Query("address")
// 	amountStr := c.Query("amount")
// 	amount, _ := new(big.Int).SetString(amountStr, 10)

// 	err = client.SendCrypto(address, amount)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Transfer failed"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"status": "Transfer successful"})
// }

type BlockchainController struct {
	DB *gorm.DB
}

func NewBlockchainController(db *gorm.DB) *BlockchainController {
	return &BlockchainController{DB: db}
}

func (bc *BlockchainController) GetBlockchainNetworks(c *gin.Context) {
	var networks []model.BlockchainNetwork
	bc.DB.Find(&networks)
	c.JSON(http.StatusOK, networks)
}

func (bc *BlockchainController) GetBlockchainNetwork(c *gin.Context) {
	id := c.Param("id")
	var network model.BlockchainNetwork
	if err := bc.DB.First(&network, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Network not found"})
		return
	}
	c.JSON(http.StatusOK, network)
}
func (bc *BlockchainController) CreateBlockchainNetwork(c *gin.Context) {
	var input model.BlockchainNetwork
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bc.DB.Create(&input)
	c.JSON(http.StatusCreated, input)
}
func (bc *BlockchainController) UpdateBlockchainNetwork(c *gin.Context) {
	id := c.Param("id")
	var network model.BlockchainNetwork
	if err := bc.DB.First(&network, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Network not found"})
		return
	}

	if err := c.ShouldBindJSON(&network); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bc.DB.Save(&network)
	c.JSON(http.StatusOK, network)
}
func (bc *BlockchainController) DeleteBlockchainNetwork(c *gin.Context) {
	id := c.Param("id")
	if err := bc.DB.Delete(&model.BlockchainNetwork{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete network"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Network deleted"})
}

// func (bc *BlockchainController) ActiveBlockchainNetwork(c *gin.Context) {
// 	id := c.Param("chainID")
// 	if err := bc.DB.Delete(&model.BlockchainNetwork{}, id).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengaktifkan jaringan!"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Jaringan diaktifkan"})
// }
