package services

import (
	"context"
	authpb "dapodik_sync/generated"
	"log"
	"time"

	"google.golang.org/grpc"
)

type AuthService interface {
	Login(username, password string) (string, error)
}

type authService struct {
	client authpb.AuthServiceClient
	// userRepository repositories.UserRepository
}

func NewAuthService(conn *grpc.ClientConn) AuthService {
	client := authpb.NewAuthServiceClient(conn)
	return &authService{client: client}
}

func (s *authService) Login(username, password string) (string, error) {
	// Panggil gRPC Login
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	req := &authpb.LoginRequest{Username: username, Password: password}
	res, err := s.client.Login(ctx, req)
	if err != nil {
		log.Printf("Login gagal: %v", err)
		return "", err
	}

	return res.Token, nil
}
