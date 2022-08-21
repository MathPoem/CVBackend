package service

import (
	"CVBackend/models"
	"CVBackend/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Private interface {
	Estimate(estimateInput models.Estimate, userId int) (int, error)
	GetEstimate(userId int) ([]models.Estimate, error)
	CreatePerson(person models.Person, userId int) (int, error)
}

type Service struct {
	Authorization
	Private
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		Private:       NewPrivateService(repo),
	}
}
