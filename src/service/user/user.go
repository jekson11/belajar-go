package user

import (
	"belajar-go/src/domain"
	"belajar-go/src/repository/user"
)

type UserServiceInterface interface {
	ListAllDataUser() ([]domain.User, error)
}

type userService struct {
	userRepository user.UserRepositoryInterface
}

func InitUserService(userRepository user.UserRepositoryInterface) UserServiceInterface {
	return &userService{
		userRepository: userRepository,
	}
}
