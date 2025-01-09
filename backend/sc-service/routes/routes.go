package routes

import (
	"sc-service/controller"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Ac    *controller.AccountController
	IPFSc *controller.IPFSController
	Bc    *controller.BlockchainController
	Sc    *controller.StudentController
	Cc    *controller.CertificateController
}

// func InitRoutes(db *gorm.DB) {

// }

func RegisterRoutes(r *gin.Engine, rc Route) {
	api := r.Group("/api/v1")
	{
		// Account
		api.POST("/accounts", rc.Ac.CreateAccount)
		api.GET("/accounts", rc.Ac.GetAccounts)
		api.GET("/accounts/:id", rc.Ac.GetAccountByID)
		api.PUT("/accounts/:id", rc.Ac.UpdateAccount)
		api.DELETE("/accounts/:id", rc.Ac.DeleteAccount)

		// IPFS
		api.GET("/metadata", rc.IPFSc.GetMetadata)
		api.POST("/metadata", rc.IPFSc.PostMetadata)
		api.GET("/nodes", rc.IPFSc.GetNode)
		api.POST("/nodes", rc.IPFSc.PostNode)

		api.GET("/ipfs/status", rc.IPFSc.GetIPFSStatusHandler)
		api.GET("/ipfs/peers", rc.IPFSc.GetSwarmPeersHandler)
		api.GET("/ipfs/repo-stats", rc.IPFSc.GetRepoStatsHandler)

		// --------------------------------------------------
		// Blockchain
		// --------------------------------------------------

		// Blockchain network
		api.GET("/blockchain-networks", rc.Bc.GetBlockchainNetworks)
		api.GET("/blockchain-networks/:id", rc.Bc.GetBlockchainNetwork)
		api.POST("/blockchain-networks", rc.Bc.CreateBlockchainNetwork)
		api.PUT("/blockchain-networks/:id", rc.Bc.UpdateBlockchainNetwork)
		api.DELETE("/blockchain-networks/:id", rc.Bc.DeleteBlockchainNetwork)
		// Block

		// --------------------------------------------------
		// Students
		// --------------------------------------------------
		api.GET("/students", rc.Sc.GetStudents)
		api.GET("/student", rc.Sc.SearchStudent)
		api.GET("/students/:id", rc.Sc.GetStudent)
		// api.POST("/students", rc.SC.CreateStudent)
		// api.PUT("/students/:id", rc.SC.UpdateStudent)
		// api.DELETE("/students/:id", rc.SC.DeleteStudent)
		// --------------------------------------------------
		// Certificates
		// --------------------------------------------------
		api.GET("/certificates", rc.Cc.GetCertificates)
		api.GET("/certificates/:id", rc.Cc.GetCertificate)
		api.POST("/certificates", rc.Cc.CreateCertificate)
		api.PUT("/certificates/:id", rc.Cc.UpdateCertificate)
		api.DELETE("/certificates/:id", rc.Cc.DeleteCertificate)

	}
}

// func RegisterRoutes(r *gin.Engine, ac *controller.AccountController) {
// 	api := r.Group("/api/v1")
// 	{
// 		// Account
// 		api.POST("/accounts", ac.CreateAccount)
// 		api.GET("/accounts", ac.GetAccounts)
// 		api.GET("/accounts/:id", ac.GetAccountByID)
// 		api.PUT("/accounts/:id", ac.UpdateAccount)
// 		api.DELETE("/accounts/:id", ac.DeleteAccount)

// 	}
// }
