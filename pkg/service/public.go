package service

import (
	"CVBackend/models"
	"CVBackend/pkg/repository"
)

type PublicService struct {
	repo repository.Public
}

func NewPublicService(repo repository.Public) *PublicService {
	return &PublicService{repo: repo}
}

func (r *PublicService) GetUniversity() ([]models.University, error) {
	return r.repo.GetUniversity()
}

func (r *PublicService) GetUniversityById(id int) (models.University, error) {
	return r.repo.GetUniversityById(id)
}

func (r *PublicService) GetSchool() ([]models.School, error) {
	return r.repo.GetSchool()
}

func (r *PublicService) GetSchoolById(id int) (models.School, error) {
	return r.repo.GetSchoolById(id)
}

func (r *PublicService) GetDepartment() ([]models.Department, error) {
	return r.repo.GetDepartment()
}

func (r *PublicService) GetDepartmentById(id int) (models.Department, error) {
	return r.repo.GetDepartmentById(id)
}
