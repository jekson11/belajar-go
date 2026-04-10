package domain

import "time"

type User struct {
	UserId    string    `db:"user_id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Username  string    `db:"username" json:"username"`
	Email     string    `db:"email" json:"email"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
