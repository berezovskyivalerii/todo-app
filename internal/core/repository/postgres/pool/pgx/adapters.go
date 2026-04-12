package corepgxpool

import (
	"errors"

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
		if errors.Is(err, pgx.ErrNoRows) {
			return corepostgrespool.ErrNoRows
		}

		return err
	}

	return nil
}

type pgxCommandTag struct {
	pgconn.CommandTag
}
