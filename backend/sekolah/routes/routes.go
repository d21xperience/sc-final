package routes

import (
	"sekolah/controllers"

	"github.com/gin-gonic/gin"
)

// func DapodikRoutes(router *gin.Engine) {
// 	// Setup Dapodik

// 	dapodikRoutes := router.Group("/WebService")
// 	{
// 		dapodikRoutes.GET("/getSekolah", controllers.GetPesertaDidik())
// 		dapodikRoutes.GET("/getRombonganBelajar")
// 		dapodikRoutes.GET("/getPesertaDidik")
// 	}
// }

func SetupDapodik(router *gin.Engine, dapoController *controllers.DapodikClient) {
	dapodikRouter := router.Group("/api/v1/dapodik")
	{
		dapodikRouter.GET("", dapoController.GetDapodikApp)
		dapodikRouter.POST("/login", dapoController.LoginToDapodik)
		dapodikRouter.GET("/GeW", dapoController.GetSekolah)
		dapodikRouter.GET("/GetSekolah", dapoController.GetSekolah)
		dapodikRouter.GET("/GetGtk", dapoController.GetGTK)
		dapodikRouter.GET("/GetPesertaDidik", dapoController.GetPesertaDidik)
		dapodikRouter.GET("/GetPengguna", dapoController.GetPengguna)
	}

	dapodikRouterInternal := router.Group("/rest")
	{
		dapodikRouterInternal.GET("/WsAplikasi", dapoController.GetWSDapodik)
	}

	// dapodikInternalService := router.Group("/rest")
	// {
	// 	// Web Service
	// 	// dapodikInternalService.GET("/WsAplikasi",dapoController.GetWSDapodik)
	// 	// dapodikInternalService.POST("/WsAplikasi")
	// 	// dapodikInternalService.DELETE("/WsAplikasi")
	// }
}
