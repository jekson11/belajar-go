package repository

import (
	"github.com/jmoiron/sqlx"

	"belajar-go/src/config/query"
	"belajar-go/src/repository/user"
)

type Repository struct {
	User user.UserRepositoryInterface
}

func InitRepository(sql0 *sqlx.DB, ql *query.LoadQuery) *Repository {
	return &Repository{
		User: user.InitUserRepository(
			sql0,
			ql,
		),
	}
}
