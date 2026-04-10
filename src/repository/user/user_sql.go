package user

import (
	"belajar-go/src/domain"
	"fmt"
)

func (d *userRepository) findAllUserFromSql() ([]domain.User, error) {
	var results []domain.User

	query, args, err := d.queryLoader.ExecuteTemplate("FindAllUserData", nil)

	if err != nil {
		return nil, fmt.Errorf("failed to load query template: %w", err)
	}
	err = d.sql0.Select(&results, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	return results, nil
}
