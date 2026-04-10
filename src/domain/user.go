package domain

import "time"

type User struct {
	UserId    string    `db:"user_id" json:"user_id"`
	Name      string    `db:"name" json:"name"`
	Username  string    `db:"username" json:"username"`
	Password  string    `db:"password" json:"password"`
	Email     string    `db:"email" json:"email"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
