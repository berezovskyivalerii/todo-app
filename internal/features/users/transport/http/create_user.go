package userhttp

import (
	"net/http"

	"github.com/berezovskyivalerii/todo-app/internal/core/domain"
	corelogger "github.com/berezovskyivalerii/todo-app/internal/core/logger"
	corerequest "github.com/berezovskyivalerii/todo-app/internal/core/transport/http/request"
	coreresponse "github.com/berezovskyivalerii/todo-app/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name"    validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,e164"`
}

type CreateUserResponse UserDTOResponse

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := corelogger.FromContext(ctx)
	responseHandler := coreresponse.NewHTTPResponseHandler(log, rw)

	var request CreateUserRequest
	if err := corerequest.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")

		return
	}

	userDomain, err := h.usersService.CreateUser(ctx, domainFromDTO(request))
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")

		return
	}

	response := CreateUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.FullName, dto.PhoneNumber)
}
