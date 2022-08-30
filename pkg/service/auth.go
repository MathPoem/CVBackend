package service

import (
	"CVBackend/models"
	mymailer "CVBackend/pkg/mailer"
	"CVBackend/pkg/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/siruspen/logrus"
	"math/rand"
	"os"
	"time"
)

const (
	signInKey = "as6dsa8d76as(*&"
	salt      = "13h431k419asa@#!@$*d"
	tokenTTL  = time.Minute
)

type TokenClaims struct {
	jwt.RegisteredClaims
	UserID int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) GenerateToken(username, password string) (models.Tokens, error) {

	user, err := s.repo.GetUser(username, generateUserPasswordHash(password))
	if err != nil {
		return models.Tokens{}, err
	}

	return s.CreateSession(user.ID)
}

func (s *AuthService) Refresh(token string) (models.Tokens, error) {
	id, err := s.repo.GetUserByToken(token)
	if err != nil {
		return models.Tokens{}, err
	}

	return s.CreateSession(id)
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signInKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *token.Claims")
	}

	claims.ExpiresAt.Before(time.Now())

	return claims.UserID, nil
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	if err := user.Validate(); err != nil {
		return 0, err
	}
	user.Password = generateUserPasswordHash(user.Password)
	key, err := NewRefreshToken()

	if err != nil {
		return 0, err
	}
	user.ConfirmKey = key
	id, err := s.repo.CreateUser(user)
	if err != nil {
		return 0, err
	}
	mailer := mymailer.NewMailerService(os.Getenv("MAIL_PASSWORD"))
	if err := mailer.Mail.SendMail(user.Email, key); err != nil {
		logrus.Warning(err)
	}
	return id, nil
}

func generateUserPasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) CreateSession(id int) (models.Tokens, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID: id,
	})

	accessToken, err := token.SignedString([]byte(signInKey))
	if err != nil {
		return models.Tokens{}, err
	}

	refreshToken, err := NewRefreshToken()

	if err != nil {
		return models.Tokens{}, err
	}

	date := time.Now().Format("2006-01-02")

	err = s.repo.CreateUserSession(refreshToken, date, id)
	if err != nil {
		return models.Tokens{}, err
	}

	return models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

func (s *AuthService) GetUserByName(username, password string) (models.User, error) {
	return s.repo.GetUser(username, generateUserPasswordHash(password))
}

func (s *AuthService) Logout(id int, email string) error {
	return s.repo.DeleteUserSession(id, email)
}

func (s *AuthService) Activate(key string) error {
	return s.repo.ActivateUser(key)
}
