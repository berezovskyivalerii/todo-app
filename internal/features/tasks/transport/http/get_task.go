package taskshttp

import (
	"net/http"

	corelogger "github.com/berezovskyivalerii/todo-app/internal/core/logger"
	corerequest "github.com/berezovskyivalerii/todo-app/internal/core/transport/http/request"
	coreresponse "github.com/berezovskyivalerii/todo-app/internal/core/transport/http/response"
)

type GetTaskResponse TaskDTOResponse

func (h *TasksHTTPHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := corelogger.FromContext(ctx)
	responseHandler := coreresponse.NewHTTPResponseHandler(log, w)

	taskID, err := corerequest.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get taskID path value")
		return
	}

	taskDomain, err := h.tasksService.GetTask(ctx, taskID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get task")
		return
	}

	task := GetTaskResponse(taskDTOFromDomain(taskDomain))

	responseHandler.JSONResponse(task, http.StatusOK)
}
