// Package usersrepository implements repository layer for users
package usersrepository

import corepostgrespool "github.com/berezovskyivalerii/todo-app/internal/core/repository/postgres/pool"

type UsersRepository struct {
	pool corepostgrespool.Pool
}

func NewUsersRepository(pool corepostgrespool.Pool) *UsersRepository {
	return &UsersRepository{pool: pool}
}
