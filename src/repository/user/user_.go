package user

import (
	"belajar-go/src/domain"
)

func (d *userRepository) FindAll() ([]domain.User, error) {
	result, err := d.findAllUserFromSql()
	if err != nil {
		return result, err
	}
	return result, nil
}
