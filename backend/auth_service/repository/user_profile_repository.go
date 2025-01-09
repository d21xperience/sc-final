package repository

import (
	"auth_service/models"

	"gorm.io/gorm"
)

type UserProfileRepository interface {
	Save(user *models.UserProfile) error
	FindByID(userID string) (*models.UserProfile, error)
	Update(user *models.UserProfile) error
	Delete(userID string) error
}

type userProfileRepositoryImpl struct {
	db *gorm.DB
}

func NewUserProfileRepository(db *gorm.DB) UserProfileRepository {
	return &userProfileRepositoryImpl{db: db}
}

// Create (Simpan User Profile)
func (r *userProfileRepositoryImpl) Save(userProfile *models.UserProfile) error {
	return r.db.Create(userProfile).Error
}

// Read (Cari User Profile berdasarkan ID)
func (r *userProfileRepositoryImpl) FindByID(userID string) (*models.UserProfile, error) {
	var userProfile models.UserProfile
	err := r.db.First(&userProfile, "id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	return &userProfile, nil
}

// Update (Perbarui User Profile)
func (r *userProfileRepositoryImpl) Update(userProfile *models.UserProfile) error {
	return r.db.Save(userProfile).Error
}

// Delete (Hapus User Profile berdasarkan ID)
func (r *userProfileRepositoryImpl) Delete(userID string) error {
	return r.db.Delete(&models.UserProfile{}, "id = ?", userID).Error
}
