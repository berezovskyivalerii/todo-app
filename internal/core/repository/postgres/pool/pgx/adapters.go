package corepgxpool

import (
	"errors"
	"fmt"

	corepostgrespool "github.com/berezovskyivalerii/todo-app/internal/core/repository/postgres/pool"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type pgxRows struct {
	pgx.Rows
}

type pgxRow struct {
	pgx.Row
}

func (r pgxRow) Scan(dest ...any) error {
	err := r.Row.Scan(dest...)
	if err != nil {
		return mapErrors(err)
	}

	return nil
}

type pgxCommandTag struct {
	pgconn.CommandTag
}

func mapErrors(err error) error {
	const pgxViolatesForeignKeyErrorCode = "23503"

	if errors.Is(err, pgx.ErrNoRows) {
		return corepostgrespool.ErrNoRows
	}

	if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
		if pgErr.Code == pgxViolatesForeignKeyErrorCode {
			return fmt.Errorf(
				"%v: %w",
				pgErr, corepostgrespool.ErrViolatesForeignKey,
			)
		}
	}

	return fmt.Errorf(
		"%v: %w",
		err, corepostgrespool.ErrUnknown,
	)
}
