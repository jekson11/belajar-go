package service

import (
	"belajar-go/src/repository"
	"belajar-go/src/service/user"
)

type Service struct {
	User user.UserServiceInterface
}

func InitService(repository *repository.Repository) *Service {
	return &Service{
		User: user.InitUserService(
			repository.User,
		),
	}
}
