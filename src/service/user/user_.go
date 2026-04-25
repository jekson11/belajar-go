package user

import (
	"context"
	"strings"

	"belajar-go/src/domain"
	"belajar-go/src/dto"
)

func (s *userService) ListAllDataUser(filter dto.UserFilter) ([]domain.User, int, error) {
	filter.Page = (filter.Page - 1) * filter.Limit
	if filter.Name != "" {
		filter.Name = "%" + strings.ToLower(filter.Name) + "%"
	}
	return s.userRepository.FindAll(filter)
}

func (s *userService) CreateDataUser(ctx context.Context, request []domain.UserCreateDomain) ([]domain.UserCreateDomain, error) {
	if _, err := s.userRepository.Create(ctx, request); err != nil {
		return nil, err
	}

	return request, nil
}
