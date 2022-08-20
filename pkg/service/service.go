package service

import "CVBackend/pkg/repository"

type Service struct {
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
