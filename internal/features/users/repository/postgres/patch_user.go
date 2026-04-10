package usersrepository

import (
	"context"
	"errors"
	"fmt"

	"github.com/berezovskyivalerii/todo-app/internal/core/domain"
	coreerrors "github.com/berezovskyivalerii/todo-app/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *UsersRepository) PatchUser(ctx context.Context, id int, patch domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	UPDATE todoapp.users
	SET 
		full_name=$1,
		phone_number=$2,
		version=version+1
	WHERE id=$3 AND version=$4
	RETURNING 
		id, version, full_name, phone_number;
	`

	row := r.pool.QueryRow(ctx, query, patch.FullName, patch.PhoneNumber, id, patch.Version)

	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.User{}, fmt.Errorf("user with id=%d concurrently accessed: %w", id, coreerrors.ErrConflict)
		}

		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.FullName,
		userModel.PhoneNumber,
	)

	return userDomain, nil
}
