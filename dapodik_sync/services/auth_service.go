package services

import (
	"context"
	authpb "dapodik_sync/generated"
	"dapodik_sync/repositories"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

type AuthService interface {
	Login(username, password string) (string, error)
}

type authService struct {
	client         authpb.AuthServiceClient
	userRepository repositories.UserRepository
}

func NewAuthService(conn *grpc.ClientConn, userRepo repositories.UserRepository) AuthService {
	client := authpb.NewAuthServiceClient(conn)
	return &authService{client: client, userRepository: userRepo}
}

func (s *authService) Login(username, password string) (string, error) {
	// Simulasi validasi kredensial sebelum ke gRPC server
	user, err := s.userRepository.GetUserCredentials(username)
	if err != nil || user == nil {
		return "", fmt.Errorf("user tidak ditemukan")
	}

	if user.Password != password {
		return "", fmt.Errorf("password salah")
	}

	// Panggil gRPC Login
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &authpb.LoginRequest{Username: username, Password: password}
	res, err := s.client.Login(ctx, req)
	if err != nil {
		log.Printf("Login gagal: %v", err)
		return "", err
	}

	return res.Token, nil
}
