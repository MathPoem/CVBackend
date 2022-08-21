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

type Repository struct {
	Authorization
	Default
	Private
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Default:       NewDefaultContent(db),
		Private:       NewPrivatePostgres(db),
	}
}
