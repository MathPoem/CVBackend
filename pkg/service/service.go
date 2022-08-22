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

type Public interface {
	GetUniversity() ([]models.University, error)
	GetUniversityById(id int) (models.University, error)
	GetSchool(id int) ([]models.School, error)
	GetSchoolById(id int) (models.School, error)
	GetDepartment(id int) ([]models.Department, error)
	GetDepartmentById(id int) (models.Department, error)
	GetPerson(id int) ([]models.Person, error)
	GetPersonById(id int) (models.Person, error)
	GetProgram(id int) ([]models.Program, error)
	GetProgramById(id int) (models.Program, error)
	GetCourse(id int) ([]models.Course, error)
	GetCourseById(id int) (models.Course, error)
	GetLecture(id int) ([]models.Lecture, error)
	GetLectureById(id int) (models.Lecture, error)
	GetSeminar(id int) ([]models.Seminar, error)
	GetSeminarById(id int) (models.Seminar, error)
}

type Service struct {
	Authorization
	Private
	Public
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		Private:       NewPrivateService(repo),
		Public:        NewPublicService(repo),
	}
}
