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

type Repository struct {
	Authorization
	Default
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Default:       NewDefaultContent(db),
	}
}
