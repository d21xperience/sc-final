package services

import (
	"auth_service/models"
	"auth_service/repositories"
	"auth_service/utils"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type AuthService interface {
	IsAdminExists(schoolTenantID uint32) (bool, error)
	Register(user *models.User) error
	RegisterAdmin(user *models.User) error
	Login(username, password string) (*models.User, error)
	GenerateToken(userID int, role string) (string, error)
}

// AuthServiceImpl is the implementation of AuthService
type authServiceImpl struct {
	repo      repositories.UserRepository
	secretKey string
}

func NewAuthService(as repositories.UserRepository) AuthService {
	return &authServiceImpl{repo: as}
}

// IsAdminExists cek apakah admin sudah adah ada pada sekolah
func (s *authServiceImpl) IsAdminExists(schoolTenantID uint32) (bool, error) {
	admin, err := s.repo.FindUserByRoleAndSchoolID("admin", schoolTenantID)
	if err != nil {
		// Return false if no admin found or error is not nil
		if err == repositories.ErrUserNotFound {
			return false, nil
		}
		return false, err
	}
	return admin != nil, nil
}

func (s *authServiceImpl) Register(user *models.User) error {
	// Cek apakah username sudah ada
	existingUser, err := s.repo.FindByUsername(user.Username)
	var lanjutkan bool
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			lanjutkan = true
		} else {
			// Tangani error jika terjadi kesalahan dalam mencari user
			return fmt.Errorf("failed to check existing username: %w", err)
		}
	}

	if existingUser != nil {
		return errors.New("username already exists")
	}
	// Simpan user baru
	if !lanjutkan {
		return fmt.Errorf("failed to save user: %w", err)
	}
	// user.Password, _ = utils.EncryptPassword(user.Password) // Encrypt password
	// return s.repositories.Save(user)
	// Enkripsi password
	user.InitialPassword = &user.Password
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
		return s.repo.Save(user)
	case err = <-errorChan:
		return err
	}

}
func (s *authServiceImpl) RegisterAdmin(user *models.User) error {
	// Cek apakah email sudah ada dengan query efisien
	emailExists, err := s.repo.EmailExists(user.Email) // Hanya cek keberadaan email
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
		return s.repo.Save(user)
	case err = <-errorChan:
		return err
	}
}

func (s *authServiceImpl) Login(identifier, password string) (*models.User, error) {
	var user *models.User
	var err error

	// Deteksi apakah input adalah email
	if utils.IsEmail(identifier) {
		user, err = s.repo.FindByEmail(identifier)
	} else {
		user, err = s.repo.FindByUsername(identifier)
	}

	if err != nil || user == nil {
		return nil, errors.New("invalid credentials")
	}

	if !utils.VerifyPassword(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	err = s.repo.UpdateLastLogin(user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update last login: %w", err)
	}

	return user, nil
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
