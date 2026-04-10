package user

import "belajar-go/src/domain"

func (s *userService) ListAllDataUser() ([]domain.User, error) {
	return s.userRepository.FindAll()
}
