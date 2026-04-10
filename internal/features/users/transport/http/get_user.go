package userhttp

import (
	"net/http"

	corelogger "github.com/berezovskyivalerii/todo-app/internal/core/logger"
	coreresponse "github.com/berezovskyivalerii/todo-app/internal/core/transport/http/response"
	coreutils "github.com/berezovskyivalerii/todo-app/internal/core/transport/http/utils"
)

type GetUserResponse UserDTOResponse

func (h *UsersHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := corelogger.FromContext(ctx)
	responseHandler := coreresponse.NewHTTPResponseHandler(logger, rw)

	userID, err := coreutils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userID path value")
		return
	}

	user, err := h.usersService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user")
		return
	}

	response := GetUserResponse(userDTOFromDomain(user))

	responseHandler.JSONResponse(response, http.StatusOK)
}
