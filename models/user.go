package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type User struct {
	ID         int    `json:"-" db:"id"`
	Username   string `json:"username" db:"username" binding:"required"`
	Email      string `json:"email" db:"email_1" binding:"required"`
	Univ       int    `json:"univ" db:"univ_1_id" binding:"required"`
	Program    int    `json:"program" db:"program_1_id" binding:"required"`
	Password   string `json:"password" db:"password" binding:"required"`
	Refresh    string `json:"refresh" db:"refresh_token"`
	ExpiresAt  string `json:"expiresAt" db:"expires_at"`
	Confirmed  bool   `json:"confirmed" db:"confirmed"`
	ConfirmKey string `json:"confirmKey" db:"confirm_key"`
}

type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, is.Email),
		validation.Field(&u.Password, validation.Length(6, 30)),
	)
}
