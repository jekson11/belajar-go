package user

import (
	"context"
	"database/sql"
	"fmt"

	"belajar-go/src/domain"
	"belajar-go/src/dto"

	"github.com/VauntDev/tqla"
)

func (d *userRepository) findAllUserFromSql(filter dto.UserFilter) ([]domain.User, int, error) {
	var results []domain.User
	var totalData int

	query, args, err := d.queryLoader.ExecuteTemplate("FindAllUserData", filter)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to load query template: %w", err)
	}

	queryC, argsC, err := d.queryLoader.ExecuteTemplate("CountDataUser", filter)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to load count template: %w", err)
	}

	if err = d.sql0.Select(&results, query, args...); err != nil {
		return nil, 0, fmt.Errorf("failed to execute query: %w", err)
	}

	if err = d.sql0.Get(&totalData, queryC, argsC...); err != nil {
		return nil, 0, fmt.Errorf("failed to execute count query: %w", err)
	}

	return results, totalData, nil
}

func (d *userRepository) createUserFromSql(ctx context.Context, users []domain.UserCreateDomain) ([]domain.UserCreateDomain, error) {
	begin, err := d.sql0.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	t, err := tqla.New(tqla.WithPlaceHolder(tqla.Dollar))
	if err != nil {
		return nil, fmt.Errorf("failed to create tqla: %w", err)
	}

	queryStr, ok := d.queryLoader.Get("CreateUser")
	if !ok {
		return nil, fmt.Errorf("query CreateUser not found")
	}

	query, args, err := t.Compile(queryStr, users)
	if err != nil {
		return nil, fmt.Errorf("failed to compile query: %w", err)
	}

	_, err = begin.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	if err = begin.Commit(); err != nil {
		if rollbackErr := begin.Rollback(); rollbackErr != nil {
			return nil, fmt.Errorf("rollback after commit: %w", rollbackErr)
		}
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}
	return users, nil
}
