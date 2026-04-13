package user

import (
	"strings"

	"belajar-go/src/domain"
	"belajar-go/src/dto"
)

func (s *userService) ListAllDataUser(filter dto.UserFilter) ([]domain.User, error) {
	filter.Page = (filter.Page - 1) * filter.Limit
	if filter.Name != "" {
		filter.Name = "%" + strings.ToLower(filter.Name) + "%"
	}
	return s.userRepository.FindAll(filter)
}
