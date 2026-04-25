package user

import (
	"belajar-go/src/domain"
	"belajar-go/src/dto"

	"context"
)

func (d *userRepository) FindAll(filter dto.UserFilter) ([]domain.User, int, error) {
	result, total, err := d.findAllUserFromSql(filter)
	if err != nil {
		return nil, 0, err
	}
	return result, total, nil
}

func (d *userRepository) Create(ctx context.Context, user []*domain.UserCreateDomain) ([]*domain.UserCreateDomain, error) {
	return d.createUserFromSql(ctx, user)
}
