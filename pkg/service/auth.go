package service

import (
	"CVBackend/models"
	"CVBackend/pkg/repository"
	"crypto/sha1"
	"fmt"
)

const (
	salt = "13h431k419asa@#!@$*d"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) GenerateToken(user models.User, password string) (string, error) {
	return "", nil
}

func (s *AuthService) ParseToken(token string) (int, error) {
	return 0, nil
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	user.Password = generateUserPasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generateUserPasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
