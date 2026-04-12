// Package usersservice implements service layer of feauture users
package usersservice

import (
	"context"

	"github.com/berezovskyivalerii/todo-app/internal/core/domain"
)

type UsersService struct {
	usersRepository UsersRepository
}

type UsersRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
	DeleteUser(ctx context.Context, id int) error
	PatchUser(ctx context.Context, id int, patch domain.User) (domain.User, error)
}

func NewUsersService(repo UsersRepository) *UsersService {
	return &UsersService{usersRepository: repo}
}
