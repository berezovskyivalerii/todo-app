package coremiddleware

import (
	"fmt"
	"net/http"

	corelogger "github.com/berezovskyivalerii/todo-app/internal/core/logger"
)

func Dummy(s string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := corelogger.FromContext(ctx)

			log.Debug(fmt.Sprintf("-> before: %s", s))

			next.ServeHTTP(w, r)

			log.Debug(fmt.Sprintf("-> after: %s", s))
		})
	}
}
