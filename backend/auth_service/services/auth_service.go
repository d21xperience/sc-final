package services

import (
	"auth_service/models"
	"auth_service/repository"
	"auth_service/utils"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	IsAdminExists(schoolID int) (bool, error)
	Register(user *models.User) error
	RegisterAdmin(user *models.User) error
	Login(email, password string) (string, error)
	GenerateToken(userID int, role string) (string, error)
}

// AuthServiceImpl is the implementation of AuthService
type authServiceImpl struct {
	repository repository.UserRepository
	secretKey  string
}

func NewAuthService(as repository.UserRepository) AuthService {
	return &authServiceImpl{repository: as}
}

// IsAdminExists cek apakah admin sudah adah ada pada sekolah
func (s *authServiceImpl) IsAdminExists(schoolID int) (bool, error) {
	admin, err := s.repository.FindUserByRoleAndSchoolID("admin", schoolID)
	if err != nil {
		// Return false if no admin found or error is not nil
		if err == repository.ErrUserNotFound {
			return false, nil
		}
		return false, err
	}
	return admin != nil, nil
}

func (s *authServiceImpl) Register(user *models.User) error {
	// Cek apakah username sudah ada
	existingUser, err := s.repository.FindByUsername(user.Username)
	if existingUser != nil {
		return errors.New("username already exists")
	}
	// Simpan user baru
	// user.Password, _ = utils.EncryptPassword(user.Password) // Encrypt password
	// return s.repository.Save(user)
	// Enkripsi password
	encryptedPasswordChan := make(chan string, 1)
	errorChan := make(chan error, 1)

	go func() {
		encryptedPassword, err := utils.EncryptPassword(user.Password)
		if err != nil {
			errorChan <- err
			return
		}
		encryptedPasswordChan <- encryptedPassword
	}()

	// Tunggu enkripsi selesai
	select {
	case user.Password = <-encryptedPasswordChan:
		// Simpan admin baru
		return s.repository.Save(user)
	case err = <-errorChan:
		return err
	}

}
func (s *authServiceImpl) RegisterAdmin(user *models.User) error {
	// Cek apakah email sudah ada dengan query efisien
	emailExists, err := s.repository.EmailExists(user.Email) // Hanya cek keberadaan email
	if err != nil {
		return err
	}
	if emailExists {
		return errors.New("email already exists")
	}

	// Enkripsi password
	encryptedPasswordChan := make(chan string, 1)
	errorChan := make(chan error, 1)

	go func() {
		encryptedPassword, err := utils.EncryptPassword(user.Password)
		if err != nil {
			errorChan <- err
			return
		}
		encryptedPasswordChan <- encryptedPassword
	}()

	// Tunggu enkripsi selesai
	select {
	case user.Password = <-encryptedPasswordChan:
		// Simpan admin baru
		return s.repository.Save(user)
	case err = <-errorChan:
		return err
	}
}


func (s *authServiceImpl) Login(email, password string) (string, error) {
	// Ambil user berdasarkan username
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Verifikasi password
	if !utils.VerifyPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT
	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
func (as *authServiceImpl) GenerateToken(userID int, role string) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"role":   role,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(as.secretKey))
}
