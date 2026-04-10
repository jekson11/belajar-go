package user

import (
	"belajar-go/src/config/query"
	"belajar-go/src/domain"

	"github.com/jmoiron/sqlx"
)

type UserRepositoryInterface interface {
	FindAll() ([]domain.User, error)
}

type userRepository struct {
	sql0        *sqlx.DB
	queryLoader *query.LoadQuery
}

func InitUserRepository(sql0 *sqlx.DB, ql *query.LoadQuery) UserRepositoryInterface {
	return &userRepository{
		sql0:        sql0,
		queryLoader: ql,
	}
}
