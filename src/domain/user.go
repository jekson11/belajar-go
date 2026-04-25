package domain

import "time"

type User struct {
	UserId    string    `db:"user_id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Username  string    `db:"username" json:"username"`
	Email     string    `db:"email" json:"email"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type UserCreateDomain struct {
	Name     string `db:"name" json:"name" binding:"required,min=2,max=100"`
	Username string `db:"username" json:"username" binding:"required,min=2,max=100"`
	Password string `db:"password" json:"password" binding:"required,min=2,max=100"`
	Email    string `db:"email" json:"email" binding:"required,min=2,max=100"`
}
