package userhttp

import (
	"net/http"

	corelogger "github.com/berezovskyivalerii/todo-app/internal/core/logger"
	coreresponse "github.com/berezovskyivalerii/todo-app/internal/core/transport/http/response"
	coreutils "github.com/berezovskyivalerii/todo-app/internal/core/transport/http/utils"
)

type DeleteUserResponse UserDTOResponse

func (h *UsersHTTPHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := corelogger.FromContext(ctx)
	responseHandler := coreresponse.NewHTTPResponseHandler(logger, rw)

	userID, err := coreutils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userID path value")
		return
	}

	err = h.usersService.DeleteUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to delete user")
		return
	}

	responseHandler.NoContentResponse()
}
