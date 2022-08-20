package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	ID       int    `json:"-" db:"id"`
	Username string `json:"username" db:"username" binding:"required"`
	Email    string `json:"email" db:"email_1" binding:"required"`
	Univ     string `json:"univ" db:"univ_1_id" binding:"required"`
	Program  string `json:"program" db:"program_1_id" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, is.Email),
		validation.Field(&u.Password, validation.Length(6, 30)),
	)
}
