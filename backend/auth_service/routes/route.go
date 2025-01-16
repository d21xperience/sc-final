package routes

// type Route struct {
// 	Ac *controller.AuthController
// 	// IPFSc *controller.IPFSController
// 	// Bc    *controller.BlockchainController
// 	// Sc    *controller.StudentController
// 	// Cc    *controller.CertificateController
// }

// func RegisterRoutes(r *gin.Engine, rc Route) {
// 	// authMiddleware := utils.JWTAuthMiddleware()
// 	baseURL := "/api/v1"
// 	auth := r.Group(fmt.Sprintf("%s/%s", baseURL, "auth"))
// 	{
// 		auth.POST("/register", rc.Ac.Register)
// 		auth.POST("/login", rc.Ac.Login)
// 		// auth.GET("/users/:id", authMiddleware, rc.Ac.GetUserByID)  // Ambil user berdasarkan ID
// 		// auth.GET("/profile", authMiddleware, rc.Ac.GetUserProfile) // Ambil profil user dari JWT
// 		// api.POST("/upload/:id/profile_picture", authMiddleware, rc.Ac.UploadProfilePicture)
// 	}
// }
