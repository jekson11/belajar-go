package user

import (
	"belajar-go/src/domain"
	"belajar-go/src/dto"
	"belajar-go/src/repository/user"

	"context"
)

type UserServiceInterface interface {
	ListAllDataUser(filter dto.UserFilter) ([]domain.User, int, error)
	CreateDataUser(ctx context.Context, user []*domain.UserCreateDomain) ([]*domain.UserCreateDomain, error)
}

type userService struct {
	userRepository user.UserRepositoryInterface
}

func InitUserService(userRepository user.UserRepositoryInterface) UserServiceInterface {
	return &userService{
		userRepository: userRepository,
	}
}
