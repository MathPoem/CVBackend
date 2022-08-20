package models

type User struct {
	ID       int    `json:"-" db:"id"`
	Username string `json:"username" db:"username" binding:"required"`
	Email    string `json:"email" db:"email_1" binding:"required"`
	Univ     string `json:"univ" db:"univ_1_id" binding:"required"`
	Program  string `json:"program" db:"program_1_id" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}
