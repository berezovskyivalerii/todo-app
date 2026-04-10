package userhttp

import (
	"fmt"
	"net/http"

	corelogger "github.com/berezovskyivalerii/todo-app/internal/core/logger"
	coreresponse "github.com/berezovskyivalerii/todo-app/internal/core/transport/http/response"
	coreutils "github.com/berezovskyivalerii/todo-app/internal/core/transport/http/utils"
)

type GetUsersResponse []UserDTOResponse

func (h *UsersHTTPHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := corelogger.FromContext(ctx)
	reponseHandler := coreresponse.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		reponseHandler.ErrorResponse(
			err,
			"failed to get 'limit'/'offset' query param",
		)

		return
	}

	userDomains, err := h.usersService.GetUsers(ctx, limit, offset)
	if err != nil {
		reponseHandler.ErrorResponse(
			err,
			"failed to get users",
		)

		return
	}

	response := GetUsersResponse(usersDTOFromDomains(userDomains))
	reponseHandler.JSONResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	limit, err := coreutils.GetIntQueryParam(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}

	offset, err := coreutils.GetIntQueryParam(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}

	return limit, offset, nil
}
