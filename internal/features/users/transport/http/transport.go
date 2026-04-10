// Package userhttp contains users transport layer implementation
package userhttp

import (
	"context"
	"net/http"

	"github.com/berezovskyivalerii/todo-app/internal/core/domain"
	coreserver "github.com/berezovskyivalerii/todo-app/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	usersService UserService
}

type UserService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
	DeleteUser(ctx context.Context, id int) error
	PatchUser(ctx context.Context, id int, patch domain.UserPatch) (domain.User, error)
}

func NewUsersHTTPHandler(
	usersService UserService,
) *UsersHTTPHandler {
	return &UsersHTTPHandler{usersService: usersService}
}

func (h *UsersHTTPHandler) Routes() []coreserver.Route {
	return []coreserver.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: h.GetUsers,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: h.GetUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{id}",
			Handler: h.DeleteUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/{id}",
			Handler: h.PatchUser,
		},
	}
}
