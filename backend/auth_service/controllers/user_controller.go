package controllers

// func (c *AuthController) UploadProfilePicture(ctx *gin.Context) {
// 	userID := ctx.Param("id") // Ambil ID user dari parameter URL

// 	// Ambil file dari form
// 	file, err := ctx.FormFile("profile_picture")
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
// 		return
// 	}

// 	// Simpan file ke direktori
// 	dir := "./uploads/profile_pictures"
// 	if _, err := os.Stat(dir); os.IsNotExist(err) {
// 		os.MkdirAll(dir, os.ModePerm) // Buat direktori jika belum ada
// 	}

// 	// Generate nama file unik
// 	filename := filepath.Join(dir, userID+"_"+file.Filename)
// 	if err := ctx.SaveUploadedFile(file, filename); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save file"})
// 		return
// 	}

// 	// Simpan path file ke database
// 	if err := c.authService.UpdateProfilePicture(userID, filename); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"message": "Profile picture uploaded successfully", "path": filename})
// }

// Get user by ID
// func (c *AuthController) GetUserByID(ctx *gin.Context) {
// 	userID := ctx.Param("id")
// 	user, err := c.authService.GetUserByID(userID)
// 	if err != nil {
// 		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, user)
// }

// func (c *AuthController) GetUserProfile(ctx *gin.Context) {
// 	claims, exists := ctx.Get("claims") // Claims berasal dari middleware JWT
// 	if !exists {
// 		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 		return
// 	}

// 	// Ambil user ID dari claims
// 	userID := claims.(map[string]interface{})["sub"].(string)
// 	user, err := c.authService.GetUserByID(userID)
// 	if err != nil {
// 		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, user)
// }
