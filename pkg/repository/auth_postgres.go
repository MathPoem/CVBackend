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
	query := fmt.Sprintf("INSERT INTO %s (username, password, email_1, univ_1_id, program_1_id, confirm_key) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		userTable)
	row := r.db.QueryRow(query, user.Username, user.Password, user.Email, user.Univ, user.Program, user.ConfirmKey)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
func (r *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id, username, email_1, univ_1_id, program_1_id, confirmed FROM %s WHERE password=$1 AND username=$2", userTable)
	err := r.db.Get(&user, query, password, username)
	return user, err
}

func (r *AuthPostgres) CreateUserSession(token string, date string, userId int) error {
	query := fmt.Sprintf("UPDATE %s SET refresh_token=$1, expires_at=$2 WHERE id=$3", userTable)
	_, err := r.db.Exec(query, token, date, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthPostgres) GetUserByToken(token string) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE refresh_token=$1", userTable)
	row := r.db.QueryRow(query, token)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) DeleteUserSession(id int, email string) error {
	query := fmt.Sprintf("UPDATE %s SET refresh_token=$1 WHERE id=$2 AND email_1=$3", userTable)
	_, err := r.db.Exec(query, "invalid", id, email)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthPostgres) ActivateUser(key string) error {
	query := fmt.Sprintf("UPDATE %s SET confirmed=$1 WHERE confirm_key=$2", userTable)
	_, err := r.db.Exec(query, true, key)
	if err != nil {
		return err
	}
	return nil
}
