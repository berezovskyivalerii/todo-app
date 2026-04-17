package taskshttp

import (
	"net/http"

	"github.com/berezovskyivalerii/todo-app/internal/core/domain"
	corelogger "github.com/berezovskyivalerii/todo-app/internal/core/logger"
	corerequest "github.com/berezovskyivalerii/todo-app/internal/core/transport/http/request"
	coreresponse "github.com/berezovskyivalerii/todo-app/internal/core/transport/http/response"
)

type CreateTaskRequest struct {
	Title        string  `json:"title"          validate:"required,min=1,max=100"`
	Description  *string `json:"description"    validate:"omitempty,min=1,max=1000"`
	AuthorUserID int     `json:"author_user_id" validate:"required"`
}

type CreateTaskResponse TaskDTOResponse

func (h *TasksHTTPHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := corelogger.FromContext(ctx)
	responseHandler := coreresponse.NewHTTPResponseHandler(log, w)

	var request CreateTaskRequest
	if err := corerequest.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	taskFromRequest := domain.NewTaskUninitialized(
		request.Title,
		request.Description,
		request.AuthorUserID,
	)

	taskDomain, err := h.tasksService.CreateTask(ctx, taskFromRequest)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create task")
		return
	}

	response := CreateTaskResponse(taskDTOFromDomain(taskDomain))

	responseHandler.JSONResponse(response, http.StatusCreated)
}
