package user

import (
	"belajar-go/src/domain"
	"belajar-go/src/dto"
)

func (d *userRepository) FindAll(filter dto.UserFilter) ([]domain.User, error) {
	result, err := d.findAllUserFromSql(filter)
	if err != nil {
		return nil, err
	}
	return result, nil
}
