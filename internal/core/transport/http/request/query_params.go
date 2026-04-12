package corerequest

import (
	"fmt"
	"net/http"
	"strconv"

	coreerrors "github.com/berezovskyivalerii/todo-app/internal/core/errors"
)

func GetIntQueryParam(r *http.Request, key string) (*int, error) {
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	val, err := strconv.Atoi(param)
	if err != nil {
		return nil, fmt.Errorf(
			"param='%s' by key='%s' not a valid integer: %v: %w",
			param,
			key,
			err,
			coreerrors.ErrInvalidArgument,
		)
	}

	return &val, nil
}
