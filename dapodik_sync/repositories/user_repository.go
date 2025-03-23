package repositories

import "dapodik_sync/models"

type UserRepository interface {
	GetUserCredentials(username string) (*models.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetUserCredentials(username string) (*models.User, error) {
	// Simulasi data user dari database
	users := map[string]string{
		"user123": "password123",
		"admin":   "adminpass",
	}

	if password, exists := users[username]; exists {
		return &models.User{Username: username, Password: password}, nil
	}

	return nil, nil
}
