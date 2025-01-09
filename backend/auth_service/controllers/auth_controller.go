package controllers

import (
	"auth_service/models"
	"auth_service/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService    services.AuthService
	sekolahService services.SekolahService
	userProfile    services.UserProfileService
}

func NewAuthController(as services.AuthService, ss services.SekolahService, usPro services.UserProfileService) *AuthController {
	return &AuthController{
		authService:    as,
		sekolahService: ss,
		userProfile:    usPro,
	}
}

// Register handles user registration
func (c *AuthController) Register(ctx *gin.Context) {
	// Input data structure
	var input struct {
		// Npsn string      `json:"npsn" binding:"required"`
		Sekolah     models.Sekolah     `json:"sekolah" binding:"required"`
		User        models.User        `json:"user" binding:"required"`
		UserProfile models.UserProfile `json:"user_profile"`
	}

	// Bind input JSON to struct
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// // Validate NPSN format
	// if !isValidNpsn(input.Sekolah.NPSN) {
	// 	ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid NPSN format"})
	// 	return
	// }

	// Check if school exists
	sekolah, err := c.sekolahService.GetSekolahByNpsn(input.Sekolah.NPSN)
	if err != nil {
		if errors.Is(err, services.ErrNotFound) { // Use appropriate error type
			// ctx.JSON(http.StatusNotFound, gin.H{"error": "Sekolah tidak ditemukan"})
			// Jika sekolah tidak ditemukan, cek apakah input.role adalah admin
			// If user is admin, check if admin already exists for the school
			if input.User.Role == "admin" {
				// Buat sekolah
				// var newSekolah models.Sekolah
				var err1 error
				sekolah, err1 = c.sekolahService.CreateSekolah(&input.Sekolah)
				if err1 != nil {
					return
				}
			}
		} else {
			// Handle other errors
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Kacau, Server error"})
			return
		}
	}

	// Associate user with the school
	input.User.SchoolID = sekolah.ID

	if input.User.Role == "admin" {
		adminExists, err := c.authService.IsAdminExists(sekolah.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
			return
		}
		if adminExists {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Admin already exists for this school"})
			return
		} else {
			// Buat admin
			if err := c.authService.RegisterAdmin(&input.User); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}
	}
	// Register user
	if input.User.Role == "siswa" {
		if err := c.authService.Register(&input.User); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	}
	// Isi userID
	input.UserProfile.UserID = input.User.ID
	ups := c.userProfile.CreateUserProfile(&input.UserProfile)
	if ups != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	token, err := c.authService.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
