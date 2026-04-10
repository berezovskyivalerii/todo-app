package coreutils

import (
	"fmt"
	"net/http"
	"strconv"

	coreerrors "github.com/berezovskyivalerii/todo-app/internal/core/errors"
)

func GetIntPathValue(r *http.Request, key string) (int, error) {
	pathValue := r.PathValue(key)
	if pathValue == "" {
		return 0, fmt.Errorf(
			"no key='%s' in path values: %w",
			key,
			coreerrors.ErrInvalidArgument,
		)
	}

	val, err := strconv.Atoi(pathValue)
	if err != nil {
		return 0, fmt.Errorf(
			"path value='%s' by key='%s' not a valid integer: %v: %w",
			val,
			key,
			err,
			coreerrors.ErrInvalidArgument,
		)
	}

	return val, nil
}
