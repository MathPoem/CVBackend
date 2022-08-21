package service

import (
	"CVBackend/models"
	"CVBackend/pkg/repository"
	"time"
)

type PrivateService struct {
	repo repository.Private
}

func NewPrivateService(repo repository.Private) *PrivateService {
	return &PrivateService{repo: repo}
}

func (s *PrivateService) Estimate(estimateInput models.Estimate, userId int) (int, error) {
	estimateInput.UserID = userId
	estimateInput.Date = time.Now().Format("2006-01-02")
	return s.repo.CreateEstimate(estimateInput)
}

func (s *PrivateService) GetEstimate(userId int) ([]models.Estimate, error) {
	return s.repo.GetEstimate(userId)
}

func (s *PrivateService) CreatePerson(person models.Person, userId int) (int, error) {
	return s.repo.CreatePerson(person, userId)
}
