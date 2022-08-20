package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type DefaultContent interface {
}

type Repository struct {
	Authorization
	DefaultContent
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
