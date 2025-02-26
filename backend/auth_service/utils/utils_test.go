package utils

import (
	"crypto/rand"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestEncryptPassword(t *testing.T) {
	password := "securepassword"

	// Call EncryptPassword
	hashedPassword, err := EncryptPassword(password)

	// Assertions
	assert.NoError(t, err)                       // Tidak ada error
	assert.NotEmpty(t, hashedPassword)           // Hasil enkripsi tidak kosong
	assert.NotEqual(t, password, hashedPassword) // Password tidak sama dengan hasil enkripsi
}

func TestEncryptPassword_Empty(t *testing.T) {
	hashedPassword, err := EncryptPassword("")
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
}

func TestVerifyPassword(t *testing.T) {
	password := "securepassword"
	wrongPassword := "wrongpassword"

	// Call EncryptPassword
	hashedPassword, err := EncryptPassword(password)
	assert.NoError(t, err)

	// Test VerifyPassword
	assert.True(t, VerifyPassword(password, hashedPassword))       // Password valid
	assert.False(t, VerifyPassword(wrongPassword, hashedPassword)) // Password tidak valid
}

func BenchmarkEncryptPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncryptPassword("securepassword")
	}
}

func BenchmarkBcryptCost(b *testing.B) {
	password := "securepassword"
	for cost := 10; cost <= 14; cost++ {
		b.Run("Cost="+string(rune(cost)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = bcrypt.GenerateFromPassword([]byte(password), cost)
			}
		})
	}
}

func TestGenerateJWT(t *testing.T) {
	// Call GenerateJWT
	token, err := GenerateJWT(nil)

	// Assertions
	assert.NoError(t, err)    // Tidak ada error
	assert.NotEmpty(t, token) // Token tidak kosong

	// Parse token kembali untuk memastikan validitasnya
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)

	// Validasi klaim
	claims := parsedToken.Claims.(jwt.MapClaims)
	assert.WithinDuration(t, time.Now(), time.Unix(int64(claims["iat"].(float64)), 0), time.Second)
	assert.WithinDuration(t, time.Now().Add(24*time.Hour), time.Unix(int64(claims["exp"].(float64)), 0), time.Second)
}

func TestJWTAuthMiddleware(t *testing.T) {
	// Dummy JWT key
	jwtKey = []byte("secret")

	// Generate a valid token for testing
	validToken, err := GenerateJWT(nil)
	assert.NoError(t, err)

	// Setup Gin
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Apply middleware
	router.Use(JWTAuthMiddleware())

	// Dummy endpoint
	router.GET("/protected", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Access granted"})
	})

	// Test cases
	tests := []struct {
		name       string
		header     string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "No Authorization Header",
			header:     "",
			wantStatus: http.StatusUnauthorized,
			wantBody:   `{"error":"Unauthorized"}`,
		},
		{
			name:       "Invalid Authorization Header",
			header:     "InvalidToken",
			wantStatus: http.StatusUnauthorized,
			wantBody:   `{"error":"Unauthorized"}`,
		},
		{
			name:       "Expired Token",
			header:     "Bearer expired.token.here",
			wantStatus: http.StatusUnauthorized,
			wantBody:   `{"error":"Invalid token"}`,
		},
		{
			name:       "Valid Token",
			header:     "Bearer " + validToken,
			wantStatus: http.StatusOK,
			wantBody:   `{"message":"Access granted"}`,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a request with header
			req := httptest.NewRequest("GET", "/protected", nil)
			if tt.header != "" {
				req.Header.Set("Authorization", tt.header)
			}

			// Record the response
			rec := httptest.NewRecorder()
			router.ServeHTTP(rec, req)

			// Validate response
			assert.Equal(t, tt.wantStatus, rec.Code)
			assert.JSONEq(t, tt.wantBody, rec.Body.String())
		})
	}
}

// func TestGenerateUsername(t *testing.T) {
// 	mockIsUsernameTaken := func(existingUsernames map[string]bool) func(string) bool {
// 		return func(username string) bool {
// 			return existingUsernames[username]
// 		}
// 	}

// 	tests := []struct {
// 		name              string
// 		inputName         string
// 		inputID           int
// 		existingUsernames map[string]bool
// 		expectedUsername  string
// 	}{
// 		{
// 			name:              "Simple unique username",
// 			inputName:         "JohnDoe",
// 			inputID:           123,
// 			existingUsernames: map[string]bool{},
// 			expectedUsername:  "johndoe123",
// 		},
// 		{
// 			name:              "Username with special characters",
// 			inputName:         "John.Doe@123!",
// 			inputID:           456,
// 			existingUsernames: map[string]bool{},
// 			expectedUsername:  "johndoe123456",
// 		},
// 		{
// 			name:              "Username already taken",
// 			inputName:         "JaneDoe",
// 			inputID:           789,
// 			existingUsernames: map[string]bool{"janedoe789": true},
// 			expectedUsername:  "janedoe790",
// 		},
// 		{
// 			name:              "Empty name input",
// 			inputName:         "",
// 			inputID:           101,
// 			existingUsernames: map[string]bool{},
// 			expectedUsername:  "user101",
// 		},
// 		{
// 			name:              "Multiple conflicts",
// 			inputName:         "Alice",
// 			inputID:           202,
// 			existingUsernames: map[string]bool{"alice202": true, "alice203": true},
// 			expectedUsername:  "alice204",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			isUsernameTaken := mockIsUsernameTaken(tt.existingUsernames)
// 			result := GenerateUsername(tt.inputName, tt.inputID, isUsernameTaken)

//				assert.Equal(t, tt.expectedUsername, result, fmt.Sprintf("Expected %s, but got %s", tt.expectedUsername, result))
//			})
//		}
//	}
func TestGenerateRefreshToken(t *testing.T) {
	t.Run("Successful token generation", func(t *testing.T) {
		token, err := GenerateRefreshToken(rand.Read)
		assert.NoError(t, err)          // Tidak ada error
		assert.NotEmpty(t, token)       // Token tidak kosong
		assert.Equal(t, 64, len(token)) // Panjang token (32 bytes * 2 karena hex encoding)
	})

	t.Run("Error during random bytes generation", func(t *testing.T) {
		// Mock rand.Read untuk mensimulasikan error
		mockRandRead := func(b []byte) (int, error) {
			return 0, errors.New("random generation error")
		}

		token, err := GenerateRefreshToken(mockRandRead)
		assert.Error(t, err)                                 // Pastikan ada error
		assert.Equal(t, "", token)                           // Token harus kosong
		assert.EqualError(t, err, "random generation error") // Pastikan error sesuai
	})
}
