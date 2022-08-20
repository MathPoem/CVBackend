package repository

import (
	"CVBackend/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password, email_1, univ_1_id, program_1_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		userTable)
	row := r.db.QueryRow(query, user.Username, user.Password, user.Email, user.Univ, user.Program)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (r *AuthPostgres) GetUser(username, password string) (models.User, error) {
	return models.User{}, nil
}
