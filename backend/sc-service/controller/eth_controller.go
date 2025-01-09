package controller

// import (
// 	"math/big"
// 	"net/http"
// 	"sc-service/config"

// 	"github.com/gin-gonic/gin"
// )

// // Initialisasi BC
// func InitBC(c *gin.Context) {

// }

// // Endpoint untuk Transfer ETH
// func SendCrypto(c *gin.Context) {
// 	address := c.Query("address")
// 	amountStr := c.Query("amount")
// 	amount, _ := new(big.Int).SetString(amountStr, 10)

// 	config.ETHClient.SendCrypto(address, amount)
// 	// if err != nil {
// 	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Transfer failed"})
// 	// 	return
// 	// }
// 	c.JSON(http.StatusOK, gin.H{"status": "Transfer successful"})
// }

// // Endpoint untuk Deploy Smart Contract
// func DeployContract(c *gin.Context) {
// 	address, txHash, err := config.ETHClient.DeployContract()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"contract_address": address.Hex(),
// 		"transaction_hash": txHash,
// 	})
// }

// // Endpoint untuk Interaksi Kontrak
// func InteractContract(c *gin.Context) {
// 	dataID := c.Query("data_id")
// 	name, noIjazah, err := config.ETHClient.InteractWithContract(dataID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"nama":      name,
// 		"no_ijazah": noIjazah,
// 	})
// }
