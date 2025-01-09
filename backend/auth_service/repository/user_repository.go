package repository

import (
	"auth_service/models"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByID(userID string) (*models.User, error)
	FindUserByRoleAndSchoolID(role string, schoolID int) (*models.User, error)
	Save(user *models.User) error
	EmailExists(email string) (bool, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

var ErrUserNotFound = errors.New("user not found")

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) FindByID(userID string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("ID = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) Save(user *models.User) error {
	return r.db.Create(user).Error
}

// FindUserByRoleAndSchoolID fetches a user with the given role and school ID
func (r *userRepositoryImpl) FindUserByRoleAndSchoolID(role string, schoolID int) (*models.User, error) {
	var user models.User
	err := r.db.Where("role = ? AND school_id = ?", role, schoolID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) EmailExists(email string) (bool, error) {
	var count int64
	if err := r.db.Model(&models.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
