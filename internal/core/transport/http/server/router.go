package coreserver

import (
	"fmt"
	"net/http"

	coremiddleware "github.com/berezovskyivalerii/todo-app/internal/core/transport/http/middleware"
)

type APIVersion string

const (
	APIVersion1 = APIVersion("v1")
	APIVersion2 = APIVersion("v2")
	APIVersion3 = APIVersion("v3")
)

type APIVersionRouter struct {
	*http.ServeMux
	apiVersion APIVersion
	middleware []coremiddleware.Middleware
}

func NewAPIVersionRouter(apiVersion APIVersion, middleware ...coremiddleware.Middleware) *APIVersionRouter {
	return &APIVersionRouter{
		ServeMux:   http.NewServeMux(),
		apiVersion: apiVersion,
		middleware: middleware,
	}
}

func (r *APIVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		r.Handle(pattern, route.WithMiddleware())
	}
}

func (r *APIVersionRouter) WithMiddleware() http.Handler {
	return coremiddleware.ChainMiddleware(
		r,
		r.middleware...,
	)
}
