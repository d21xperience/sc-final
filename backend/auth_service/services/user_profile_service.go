package services

import (
	"auth_service/models"
	"auth_service/repository"
	"errors"
)

type UserProfileService interface {
	CreateUserProfile(up *models.UserProfile) error
	GetUserProfileByID(userID string) (*models.UserProfile, error)
	UpdateUserProfile(userID string, updatedData *models.UserProfile) error
	DeleteUserProfile(userID string) error
}

type userProfileServiceImpl struct {
	// userRepo        repository.UserRepository
	userProfileRepo repository.UserProfileRepository
}

func NewUserProfileService(upRepo repository.UserProfileRepository) UserProfileService {
	return &userProfileServiceImpl{
		userProfileRepo: upRepo,
	}
}

func (uPS userProfileServiceImpl) CreateUserProfile(up *models.UserProfile) error {
	err := uPS.userProfileRepo.Save(up)
	if err != nil {
		return errors.New("failed to create user profile: " + err.Error())
	}
	return nil
}

func (uPS userProfileServiceImpl) GetUserProfileByID(userID string) (*models.UserProfile, error) {
	userProfile, err := uPS.userProfileRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("failed to retrieve user profile: " + err.Error())
	}
	if userProfile == nil {
		return nil, errors.New("user profile not found")
	}
	return userProfile, nil
}

func (uPS userProfileServiceImpl) UpdateUserProfile(userID string, updatedData *models.UserProfile) error {
	existingProfile, err := uPS.userProfileRepo.FindByID(userID)
	if err != nil {
		return errors.New("failed to find user profile: " + err.Error())
	}
	if existingProfile == nil {
		return errors.New("user profile not found")
	}

	updatedData.ID = existingProfile.ID // Pastikan ID tidak berubah
	err = uPS.userProfileRepo.Update(updatedData)
	if err != nil {
		return errors.New("failed to update user profile: " + err.Error())
	}
	return nil
}

func (uPS userProfileServiceImpl) DeleteUserProfile(userID string) error {
	err := uPS.userProfileRepo.Delete(userID)
	if err != nil {
		return errors.New("failed to delete user profile: " + err.Error())
	}
	return nil
}

// func (s *userProfileServiceImpl) UpdateProfilePicture(userID string, filepath string) error {
// 	user, err := s.userProfileRepo.FindByID(userID)
// 	if err != nil {
// 		return errors.New("user not found")
// 	}

// 	s.userProfileRepo.ProfilePicture = filepath
// 	return s.userRepo.Save(user)
// }
