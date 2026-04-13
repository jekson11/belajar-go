package user

import (
	"fmt"

	"belajar-go/src/domain"
	"belajar-go/src/dto"
)

func (d *userRepository) findAllUserFromSql(filter dto.UserFilter) ([]domain.User, error) {
	var results []domain.User

	query, args, err := d.queryLoader.ExecuteTemplate("FindAllUserData", filter)

	if err != nil {
		return nil, fmt.Errorf("failed to load query template: %w", err)
	}

	err = d.sql0.Select(&results, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return results, nil
}
