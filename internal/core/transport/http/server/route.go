package coreserver

import (
	"net/http"

	coremiddleware "github.com/berezovskyivalerii/todo-app/internal/core/transport/http/middleware"
)

type Route struct {
	Method     string
	Path       string
	Handler    http.HandlerFunc
	Middleware []coremiddleware.Middleware
}

func (r *Route) WithMiddleware() http.Handler {
	return coremiddleware.ChainMiddleware(
		r.Handler,
		r.Middleware...,
	)
}
