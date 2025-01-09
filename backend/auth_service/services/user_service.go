package services

import (
	"auth_service/models"
	"auth_service/repository"
	"errors"
)

type UserService interface {
	GetUserByID(userID string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
}

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &UserServiceImpl{userRepo: ur}
}

// Get user by ID
func (s *UserServiceImpl) GetUserByID(userID string) (*models.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// Get user by username
func (s *UserServiceImpl) GetUserByUsername(username string) (*models.User, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
