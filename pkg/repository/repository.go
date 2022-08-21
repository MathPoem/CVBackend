package repository

import (
	"CVBackend/models"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Default interface {
	InDefaultContent() error
}

type Private interface {
	CreateEstimate(estimate models.Estimate) (int, error)
	GetEstimate(userId int) ([]models.Estimate, error)
	CreatePerson(person models.Person, userId int) (int, error)
}

type Public interface {
	GetUniversity() ([]models.University, error)
	GetUniversityById(id int) (models.University, error)
	GetSchool() ([]models.School, error)
	GetSchoolById(id int) (models.School, error)
	GetDepartment() ([]models.Department, error)
	GetDepartmentById(id int) (models.Department, error)
}

type Repository struct {
	Authorization
	Default
	Private
	Public
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Default:       NewDefaultContent(db),
		Private:       NewPrivatePostgres(db),
		Public:        NewPublicPostgres(db),
	}
}
